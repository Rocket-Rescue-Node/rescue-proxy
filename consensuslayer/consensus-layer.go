package consensuslayer

import (
	"bytes"
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/Rocket-Rescue-Node/rescue-proxy/metrics"
	"github.com/allegro/bigcache/v3"
	"github.com/attestantio/go-eth2-client/api"
	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/http"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
)

// ConsensusLayer provides an abstraction for the rescue proxy over the consensus layer
// It's specifically needed to map validator indices to pubkeys prior to EL validation
type ConsensusLayer interface {
	GetValidatorInfo([]string) (map[string]*ValidatorInfo, error)
	GetValidators() ([]*apiv1.Validator, error)
}

// CachingConsensusLayer provides rescue-proxy with the consensus layer information needed
// to enforce fee recipients and caches it for performance.
type CachingConsensusLayer struct {
	bnURL  *url.URL
	logger *zap.Logger

	// Client for the BN
	client *http.Service

	// Caches index->validatorInfo for prepare_beacon_proposer
	validatorCache *validatorCache

	// Disconnects from the bn
	disconnect func()

	// Force attestantio client to use json
	forceJSON bool

	m             *metrics.MetricsRegistry
	slotsPerEpoch uint64
}

type ValidatorInfo struct {
	Pubkey            rptypes.ValidatorPubkey
	WithdrawalAddress common.Address
	Is0x01            bool
}

// NewConsensusLayer creates a new consensus layer client using the provided url and logger
func NewCachingConsensusLayer(bnURL *url.URL, logger *zap.Logger, forceJSON bool) *CachingConsensusLayer {
	out := &CachingConsensusLayer{}
	out.bnURL = bnURL
	out.logger = logger
	out.forceJSON = forceJSON
	out.m = metrics.NewMetricsRegistry("consensus_layer")

	return out
}

func (c *CachingConsensusLayer) onHeadUpdate(slot uint64) {

	c.logger.Debug("Observed consensus slot", zap.Uint64("slot", slot))
	metrics.OnHead(slot / c.slotsPerEpoch)
}

// Init connects to the consensus layer and initializes the cache
func (c *CachingConsensusLayer) Init(ctx context.Context) error {

	// Connect to BN
	ctx, c.disconnect = context.WithCancel(ctx)
	client, err := http.New(ctx,
		http.WithAddress(c.bnURL.String()),
		// It's very chatty if we don't quiet it down
		http.WithLogLevel(zerolog.WarnLevel),
		// Set a sensible timeout. This is used as a maximum. Requests can set their own via ctx.
		http.WithTimeout(5*time.Minute),
		http.WithEnforceJSON(c.forceJSON))
	if err != nil {
		return err
	}
	c.client = client.(*http.Service)

	speCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	c.slotsPerEpoch, err = c.client.SlotsPerEpoch(speCtx)
	if err != nil {
		c.logger.Warn("Couldn't get slots per epoch, defaulting to 32", zap.Error(err))
		c.slotsPerEpoch = 32
	} else {
		c.logger.Debug("Fetched slots per epoch", zap.Uint64("slots", c.slotsPerEpoch))
	}

	c.logger.Info("Connected to Beacon Node", zap.String("url", c.bnURL.String()))

	// Poll for head updates
	tickerCtx, tickerCtxCancel := context.WithCancel(ctx)
	ticker := time.NewTicker(12 * time.Second)
	go func() {
		for {
			select {
			case <-tickerCtx.Done():
				// The parent context was canceled, so exit now
				c.logger.Debug("ConsensusLayer context canceled, exiting head update ticker")
				tickerCtxCancel()
				return
			case <-ticker.C:
				// Poll for head updates
				syncingCtx, syncingCtxCancel := context.WithTimeout(tickerCtx, 2*time.Second)
				nodeSyncing, err := c.client.NodeSyncing(syncingCtx, &api.NodeSyncingOpts{
					Common: api.CommonOpts{
						Timeout: 2 * time.Second,
					},
				})
				// Cancel the context before looping again
				syncingCtxCancel()
				if err != nil {
					c.logger.Warn("Error polling for node syncing", zap.Error(err))
					break
				}
				c.onHeadUpdate(uint64(nodeSyncing.Data.HeadSlot))
			}
		}
	}()

	validatorCacheConfig := bigcache.DefaultConfig(10 * time.Hour)
	validatorCacheConfig.CleanWindow = 30 * time.Second
	validatorCacheConfig.Shards = 32
	validatorCacheConfig.HardMaxCacheSize = 512

	c.validatorCache, err = newValidatorCache(ctx, validatorCacheConfig)
	if err != nil {
		return err
	}

	c.logger.Info("Initialized pubkey cache")

	return nil
}

// GetValidatorIfno maps a validator index to a pubkey and withdrawal credential.
// It caches responses from the beacon client in memory for an arbitrary amount of time to save resources.
func (c *CachingConsensusLayer) GetValidatorInfo(validatorIndices []string) (map[string]*ValidatorInfo, error) {

	// Pre-allocate the retval based on the argument length
	out := make(map[string]*ValidatorInfo, len(validatorIndices))
	missing := make([]phase0.ValidatorIndex, 0, len(validatorIndices))

	for _, validatorIndex := range validatorIndices {
		// Check the cache first
		validatorInfo := c.validatorCache.Get(validatorIndex)
		if validatorInfo != nil {
			// Add the pubkey to the output. We have to cast it to an array, but the length is correct (see above)
			out[validatorIndex] = validatorInfo
			c.logger.Debug("Cache hit", zap.String("validator", validatorIndex))
			c.m.Counter("cache_hit").Inc()
		} else {
			// A nil value means the record wasn't in the cache or there was an error
			// Add the index to the list to be queried against the BN
			index, err := strconv.ParseUint(validatorIndex, 10, 64)
			if err != nil {
				c.logger.Warn("Invalid validator index", zap.String("index", validatorIndex))
			}
			missing = append(missing, phase0.ValidatorIndex(index))
			c.m.Counter("cache_miss").Inc()
			c.logger.Debug("Cache miss", zap.String("validator", validatorIndex))
		}
	}

	if len(missing) == 0 {
		// All pubkeys were cached
		c.m.Counter("all_keys_cache_hit").Inc()
		return out, nil
	}

	// Create a context to enforce a timeout
	vCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// Grab the index->validator map from the client if missing from the cache
	resp, err := c.client.Validators(vCtx, &api.ValidatorsOpts{
		State:   "head",
		Indices: missing,
	})
	if err != nil {
		return nil, err
	}
	for index, validator := range resp.Data {
		strIndex := strconv.FormatUint(uint64(index), 10)
		pubkey := rptypes.ValidatorPubkey(validator.Validator.PublicKey)
		withdrawalCredentials := validator.Validator.WithdrawalCredentials

		out[strIndex] = &ValidatorInfo{
			Pubkey: pubkey,
		}

		if !bytes.HasPrefix(withdrawalCredentials, []byte{0x01}) {
			c.logger.Warn("0x00 Validator seen", zap.Binary("pubkey", pubkey.Bytes()))
		} else {
			// BytesToAddress will cut off all but the last 20 bytes
			out[strIndex].WithdrawalAddress = common.BytesToAddress(withdrawalCredentials)
			out[strIndex].Is0x01 = true
		}

		// Add it to the cache. Ignore errors, we can always look the key up later
		err = c.validatorCache.Set(strIndex, out[strIndex])
		if err != nil {
			c.logger.Warn("Error encountered while saving blob to cache", zap.Error(err))
			c.m.Counter("cache_add_failed")
		} else {
			c.m.Counter("cache_add").Inc()
		}
	}

	return out, nil
}

// GetValidators gets the list of all validators for the finalized state
// It does no caching- the response is large, so caching should be done downstream, for the data the caller cares about.
func (c *CachingConsensusLayer) GetValidators() ([]*apiv1.Validator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	vmap, err := c.client.Validators(ctx, &api.ValidatorsOpts{State: "finalized"})
	if err != nil {
		return nil, err
	}

	out := make([]*apiv1.Validator, 0, len(vmap.Data))
	for _, v := range vmap.Data {
		out = append(out, v)
	}

	return out, nil
}

// Deinit shuts down the consensus layer client
func (c *CachingConsensusLayer) Deinit() {
	err := c.validatorCache.Close()
	if err != nil {
		c.logger.Info("Error closing validator cache", zap.Error(err))
	}
	c.disconnect()
	c.logger.Info("HTTP Client Disconnected from the BN")
}
