package consensuslayer

import (
	"bytes"
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
	"github.com/allegro/bigcache/v3"
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
type ConsensusLayer struct {
	bnURL  *url.URL
	logger *zap.Logger

	// Client for the BN
	client *http.Service

	// Caches index->validatorInfo for prepare_beacon_proposer
	validatorCache *validatorCache

	// Disconnects from the bn
	disconnect func()

	m             *metrics.MetricsRegistry
	slotsPerEpoch uint64
}

type ValidatorInfo struct {
	Pubkey            rptypes.ValidatorPubkey
	WithdrawalAddress common.Address
}

// NewConsensusLayer creates a new consensus layer client using the provided url and logger
func NewConsensusLayer(bnURL *url.URL, logger *zap.Logger) *ConsensusLayer {
	out := &ConsensusLayer{}
	out.bnURL = bnURL
	out.logger = logger
	out.m = metrics.NewMetricsRegistry("consensus_layer")

	return out
}

func (c *ConsensusLayer) onHeadUpdate(e *apiv1.Event) {
	headEvent, ok := e.Data.(*apiv1.HeadEvent)
	if !ok {
		c.logger.Warn("Couldn't convert event to headEvent", zap.Any("event", e))
		return
	}

	c.logger.Debug("Observed consensus slot", zap.Uint64("slot", uint64(headEvent.Slot)), zap.Bool("new_epoch", headEvent.EpochTransition))

	// The CL doesn't report events very reliably, probably an issue with the attestantio client.
	// So, every single slot, we will check to see if the epoch advanced.
	epoch := uint64(headEvent.Slot) / c.slotsPerEpoch

	metrics.OnHead(epoch)
}

// Init connects to the consensus layer and initializes the cache
func (c *ConsensusLayer) Init() error {
	var err error
	var ctx context.Context

	// Connect to BN
	ctx, c.disconnect = context.WithCancel(context.Background())
	client, err := http.New(ctx,
		http.WithAddress(c.bnURL.String()),
		// It's very chatty if we don't quiet it down
		http.WithLogLevel(zerolog.WarnLevel),
		// Set a sensible timeout. This is used as a maximum. Requests can set their own via ctx.
		http.WithTimeout(1*time.Minute))
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

	c.logger.Debug("Connected to Beacon Node", zap.String("url", c.bnURL.String()))

	// Listen for head updates
	err = c.client.Events(context.Background(), []string{"head"}, c.onHeadUpdate)
	if err != nil {
		c.logger.Warn("Clouldn't subscribe to CL events. Metrics will be inaccurate", zap.Error(err))
	}

	validatorCacheConfig := bigcache.DefaultConfig(10 * time.Hour)
	validatorCacheConfig.CleanWindow = 30 * time.Second
	validatorCacheConfig.Shards = 32
	validatorCacheConfig.HardMaxCacheSize = 512

	c.validatorCache, err = newValidatorCache(ctx, validatorCacheConfig)
	if err != nil {
		return err
	}

	c.logger.Debug("Initialized pubkey cache")

	return nil
}

// GetValidatorIfno maps a validator index to a pubkey and withdrawal credential.
// It caches responses from the beacon client in memory for an arbitrary amount of time to save resources.
func (c *ConsensusLayer) GetValidatorInfo(validatorIndices []string) (map[string]*ValidatorInfo, error) {

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

	// Grab the index->validator map from the client if missing from the cache
	resp, err := c.client.Validators(context.Background(), "head", missing)
	if err != nil {
		return nil, err
	}
	for index, validator := range resp {
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
		}

		// Add it to the cache. Ignore errors, we can always look the key up later
		c.validatorCache.Set(strIndex, out[strIndex])
		c.m.Counter("cache_add").Inc()
	}

	return out, nil
}

// GetValidators gets the list of all validators for the finalized state
// It does no caching- the response is large, so caching should be done downstream, for the data the caller cares about.
func (c *ConsensusLayer) GetValidators() ([]*apiv1.Validator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	vmap, err := c.client.Validators(ctx, "finalized", nil)
	if err != nil {
		return nil, err
	}

	out := make([]*apiv1.Validator, 0, len(vmap))
	for _, v := range vmap {
		out = append(out, v)
	}

	return out, nil
}

// Deinit shuts down the consensus layer client
func (c *ConsensusLayer) Deinit() {
	c.validatorCache.Close()
	c.disconnect()
	c.logger.Debug("HTTP Client Disconnected from the BN")
}
