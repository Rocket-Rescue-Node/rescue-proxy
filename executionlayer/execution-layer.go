package executionlayer

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer/dataprovider"
	"github.com/Rocket-Rescue-Node/rescue-proxy/metrics"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/smartnode/bindings/types"
	"go.uber.org/zap"
)

type ForEachNodeClosure func(common.Address) bool

type RPInfo struct {
	ExpectedFeeRecipient *common.Address
	NodeAddress          common.Address
}

// ExecutionLayer is the abstract interface which provides the rescue proxy
// with all the data needed to enforce fee recipients are 'correct'.
type ExecutionLayer interface {
	ForEachNode(ForEachNodeClosure) error
	ForEachOdaoNode(ForEachNodeClosure) error
	GetRPInfo(rptypes.ValidatorPubkey) (*RPInfo, error)
	REthAddress() *common.Address
	StakewiseFeeRecipient(ctx context.Context, address common.Address) (*common.Address, error)
	ValidateEIP1271(ctx context.Context, dataHash common.Hash, signature []byte, address common.Address) (bool, error)
}

type cacheRCU struct {
	Cache

	smoothingPoolAddress common.Address
	rethAddress          common.Address
}

// CachingExecutionLayer is a bespoke execution layer client for the rescue proxy.
// It abstracts away all the work to cache in-memory the data needed to enforce
// that fee recipients are 'correct'.
type CachingExecutionLayer struct {
	// Fields passed in by the constructor which are later referenced
	Logger          *zap.Logger
	Context         context.Context
	RefreshInterval time.Duration
	DataProvider    dataprovider.DataProvider

	cache atomic.Pointer[cacheRCU]

	m *metrics.MetricsRegistry

	// A ticker to refresh the cache
	ticker *time.Ticker
}

func (e *CachingExecutionLayer) newCache(loggerFunc func(fmt string, fields ...zap.Field)) error {
	out := MapsCache{}

	if err := out.init(); err != nil {
		return fmt.Errorf("unable to init cache: %w", err)
	}

	// First, get the current block
	ctx, cancel := context.WithTimeout(e.Context, 5*time.Second)
	defer cancel()
	header, err := e.DataProvider.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	// Create a new context for multicall
	// It's pretty snappy, so timeout after just 2 minutes.
	mcCtx, cancel := context.WithTimeout(e.Context, 2*time.Minute)
	defer cancel()

	// Create opts to query state at the latest block
	opts := &bind.CallOpts{BlockNumber: header.Number, Context: mcCtx}

	// Refresh the addresses from rocketStorage in case there has been a protocol upgrade
	e.DataProvider.RefreshAddresses(opts)

	loggerFunc("Warming up the cache")

	// Get all nodes at the given block
	nodes, err := e.DataProvider.GetAllNodes(opts)
	if err != nil {
		return fmt.Errorf("error getting all nodes: %w", err)
	}
	loggerFunc("Loaded nodes", zap.Int("count", len(nodes)), zap.Int64("block", opts.BlockNumber.Int64()))

	// Get all minipools at the given block
	minipools, err := e.DataProvider.GetAllMinipools(nodes, opts)
	if err != nil {
		return fmt.Errorf("error getting all minipools: %w", err)
	}
	loggerFunc("Loaded minipools", zap.Int("count", len(minipools)), zap.Int64("block", opts.BlockNumber.Int64()))

	minipoolCount := 0
	for addr, node := range nodes {
		// Allocate a pointer for this node
		nodeInfo := &dataprovider.NodeInfo{}

		// Determine their smoothing pool status
		nodeInfo.InSmoothingPool = node.InSmoothingPool
		nodeInfo.FeeDistributor = node.FeeDistributor

		// Store the smoothing pool state / fee distributor in the node index
		err = out.addNodeInfo(addr, nodeInfo)
		if err != nil {
			return fmt.Errorf("unable to add node info: %w", err)
		}

		// Also grab their minipools
		minipools, ok := minipools[addr]
		if !ok {
			continue
		}

		minipoolCount += len(minipools)
		for _, minipool := range minipools {
			err = out.addMinipoolNode(minipool, addr)
			if err != nil {
				return fmt.Errorf("unable to add minipool node: %w", err)
			}
		}
	}

	// Get all odao nodes at the given block
	odaoNodes, err := e.DataProvider.GetAllOdaoNodes(opts)
	if err != nil {
		return fmt.Errorf("error getting all odao nodes: %w", err)
	}

	for _, member := range odaoNodes {

		err = out.addOdaoNode(member)
		if err != nil {
			return fmt.Errorf("unable to add odao node: %w", err)
		}
	}
	loggerFunc("Loaded odao nodes", zap.Int("count", len(odaoNodes)), zap.Int64("block", opts.BlockNumber.Int64()))

	smoothingPoolAddress := e.DataProvider.GetSmoothingPoolAddress()
	rethAddress := e.DataProvider.GetREthAddress()

	// Set highestBlock to the cache's highestBlock, since it was just warmed up
	out.setHighestBlock(opts.BlockNumber)

	loggerFunc("Loaded nodes and minipools snapshot",
		zap.Int("nodes", len(nodes)),
		zap.Int("minipools", minipoolCount),
		zap.Int("odao nodes", len(odaoNodes)))

	e.cache.Store(&cacheRCU{Cache: &out, smoothingPoolAddress: smoothingPoolAddress, rethAddress: rethAddress})
	return nil
}

// Init creates and warms up the ExecutionLayer cache.
func (e *CachingExecutionLayer) Init() error {

	e.m = metrics.NewMetricsRegistry("execution_layer")

	if e.Context == nil {
		e.Context = context.Background()
	}

	if err := e.newCache(e.Logger.Info); err != nil {
		return fmt.Errorf("error warming up the cache: %w", err)
	}

	// Once the cache is warm, start a background process to refresh it
	if e.RefreshInterval > 0 {
		e.ticker = time.NewTicker(e.RefreshInterval)
		go func() {
			e.Logger.Info("Starting cache refresh", zap.Duration("interval", e.RefreshInterval))
			for {
				select {
				case <-e.Context.Done():
					e.ticker.Stop()
					return
				case <-e.ticker.C:
					e.Logger.Info("Refreshing cache")
					if err := e.newCache(e.Logger.Debug); err != nil {
						e.Logger.Error("error refreshing cache", zap.Error(err))
					}
				}
			}
		}()
	}

	return nil
}

// ForEachNode calls the provided closure with the address of every rocket pool node the ExecutionLayer has observed
func (e *CachingExecutionLayer) ForEachNode(closure ForEachNodeClosure) error {
	c := e.cache.Load()
	return c.forEachNode(closure)
}

// ForEachOdaoNode calls the provided closure with the address of every odao node the ExecutionLayer has observed
func (e *CachingExecutionLayer) ForEachOdaoNode(closure ForEachNodeClosure) error {
	c := e.cache.Load()
	return c.forEachOdaoNode(closure)
}

// GetRPInfo returns the expected fee recipient and node address for a validator, or nil if the validator is not a minipool
func (e *CachingExecutionLayer) GetRPInfo(pubkey rptypes.ValidatorPubkey) (*RPInfo, error) {
	c := e.cache.Load()

	nodeAddr, err := c.getMinipoolNode(pubkey)
	if err != nil {
		_, ok := err.(*NotFoundError)
		if !ok {
			e.Logger.Panic("error querying cache for minipool", zap.String("pubkey", pubkey.String()), zap.Error(err))
			return nil, err
		}

		// Validator (hopefully) isn't a minipool
		e.m.Counter("non_minipool_detected").Inc()
		return nil, nil
	}

	nodeInfo, err := c.getNodeInfo(nodeAddr)
	if err != nil {
		_, ok := err.(*NotFoundError)
		if !ok {
			e.Logger.Panic("error querying cache for node",
				zap.String("pubkey", pubkey.String()),
				zap.String("node", nodeAddr.String()),
				zap.Error(err))
			return nil, err
		}

		// Validator was a minipool, but we don't have a node record for it. This is bad.
		e.m.Counter("cache_inconsistent").Inc()
		e.Logger.Error("Validator was in the minipool index, but not the node index",
			zap.String("pubkey", pubkey.String()),
			zap.String("node", nodeAddr.String()))
		return nil, fmt.Errorf("node %s not found in cache despite pubkey %s being present", nodeAddr.String(), pubkey.String())
	}

	if nodeInfo.InSmoothingPool {
		return &RPInfo{
			ExpectedFeeRecipient: &c.smoothingPoolAddress,
			NodeAddress:          nodeAddr,
		}, nil
	}

	return &RPInfo{
		ExpectedFeeRecipient: &nodeInfo.FeeDistributor,
		NodeAddress:          nodeAddr,
	}, nil
}

// REthAddress is a convenience function to get the rEth contract address
func (e *CachingExecutionLayer) REthAddress() *common.Address {
	c := e.cache.Load()
	return &c.rethAddress
}

// ValidateEIP1271 validates an EIP-1271 signature
func (e *CachingExecutionLayer) ValidateEIP1271(ctx context.Context, dataHash common.Hash, signature []byte, address common.Address) (bool, error) {
	return e.DataProvider.ValidateEIP1271(&bind.CallOpts{Context: ctx}, dataHash, signature, address)
}

func (e *CachingExecutionLayer) StakewiseFeeRecipient(ctx context.Context, address common.Address) (*common.Address, error) {
	return e.DataProvider.StakewiseFeeRecipient(&bind.CallOpts{Context: ctx}, address)
}
