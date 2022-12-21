package consensuslayer

import (
	"context"
	"encoding/hex"
	"net/url"
	"strconv"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
	"github.com/allegro/bigcache/v3"
	"github.com/attestantio/go-eth2-client/http"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
)

const cacheTTL time.Duration = 10 * time.Hour
const cacheShards int = 32
const cacheGC time.Duration = 30 * time.Second
const cacheHardMaxMB int = 512

// ConsensusLayer provides an abstraction for the rescue proxy over the consensus layer
// It's specifically needed to map validator indices to pubkeys prior to EL validation
type ConsensusLayer struct {
	bnURL  *url.URL
	logger *zap.Logger

	// Client for the BN
	client *http.Service

	// Caches index->pubkey for prepare_beacon_proposer
	pubkeyCache *bigcache.BigCache

	// Disconnects from the bn
	disconnect func()

	m *metrics.MetricsRegistry
}

// NewConsensusLayer creates a new consensus layer client using the provided url and logger
func NewConsensusLayer(bnURL *url.URL, logger *zap.Logger) *ConsensusLayer {
	out := &ConsensusLayer{}
	out.bnURL = bnURL
	out.logger = logger
	out.m = metrics.NewMetricsRegistry("consensus_layer")

	return out
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
		http.WithLogLevel(zerolog.WarnLevel))
	if err != nil {
		return err
	}
	c.client = client.(*http.Service)

	c.logger.Debug("Connected to Beacon Node", zap.String("url", c.bnURL.String()))

	cacheConfig := bigcache.DefaultConfig(cacheTTL)
	cacheConfig.CleanWindow = cacheGC
	cacheConfig.Shards = cacheShards
	cacheConfig.HardMaxCacheSize = cacheHardMaxMB

	c.pubkeyCache, err = bigcache.New(ctx, cacheConfig)
	if err != nil {
		return err
	}

	c.logger.Debug("Initialized pubkey cache")

	return nil
}

const pubkeyBytes = 48

// GetValidatorPubkey maps a validator index to a pubkey.
// It caches responses from the beacon client in memory for an arbitrary amount of time to save resources.
func (c *ConsensusLayer) GetValidatorPubkey(validatorIndices []string) (map[string]rptypes.ValidatorPubkey, error) {

	// Pre-allocate the retval based on the argument length
	out := make(map[string]rptypes.ValidatorPubkey, len(validatorIndices))
	missing := make([]phase0.ValidatorIndex, 0, len(validatorIndices))

	for _, validatorIndex := range validatorIndices {
		// Check the cache first
		pubkey, err := c.pubkeyCache.Get(validatorIndex)
		if err == nil {
			if len(pubkey) != pubkeyBytes {
				c.logger.Warn("Invalid pubkey from beacon node", zap.String("key", hex.EncodeToString(pubkey)))
				continue
			}
			// Add the pubkey to the output. We have to cast it to an array, but the length is correct (see above)
			out[validatorIndex] = *(*rptypes.ValidatorPubkey)(pubkey)
			c.logger.Debug("Cache hit", zap.String("validator", validatorIndex))
			c.m.Counter("cache_hit").Inc()
		} else {
			// An error means the record wasn't in the cache
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
		out[strIndex] = pubkey

		// Add it to the cache. Ignore errors, we can always look the key up later
		_ = c.pubkeyCache.Set(strIndex, pubkey[:])
		c.m.Counter("cache_add").Inc()
	}

	return out, nil
}

// Deinit shuts down the consensus layer client
func (c *ConsensusLayer) Deinit() {
	c.pubkeyCache.Close()
	c.disconnect()
	c.logger.Debug("HTTP Client Disconnected from the BN")
}
