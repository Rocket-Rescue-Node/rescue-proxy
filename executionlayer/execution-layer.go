package executionlayer

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer/dataprovider"
	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer/stakewise"
	"github.com/Rocket-Rescue-Node/rescue-proxy/metrics"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	rptypes "github.com/rocket-pool/smartnode/bindings/types"
	"go.uber.org/zap"
)

type ForEachNodeClosure func(common.Address) bool

const multicall3Addr = "0xcA11bde05977b3631167028862bE2a173976CA11"

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
	Logger               *zap.Logger
	ECURL                *url.URL
	RocketStorageAddr    string
	SWVaultsRegistryAddr string
	Context              context.Context
	RefreshInterval      time.Duration

	client *ethclient.Client
	cache  atomic.Pointer[cacheRCU]

	// Checkers for vaults and mev escrow
	vaultsChecker *stakewise.VaultsChecker

	m *metrics.MetricsRegistry

	// A ticker to refresh the cache
	ticker *time.Ticker
}

func (e *CachingExecutionLayer) newCache(loggerFunc func(fmt string, fields ...zap.Field)) error {
	out := MapsCache{}

	if err := out.init(); err != nil {
		return fmt.Errorf("unable to init cache: %w", err)
	}

	rocketStorageAddr := common.HexToAddress(e.RocketStorageAddr)

	// First, get the current block
	ctx, cancel := context.WithTimeout(e.Context, 5*time.Second)
	defer cancel()
	header, err := e.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	// Create a new context for multicall
	// It's pretty snappy, so timeout after just 2 minutes.
	mcCtx, cancel := context.WithTimeout(e.Context, 2*time.Minute)
	defer cancel()

	// Create opts to query state at the latest block
	opts := &bind.CallOpts{BlockNumber: header.Number, Context: mcCtx}

	loggerFunc("Warming up the cache")

	// Get all nodes at the given block
	mc, err := dataprovider.NewMulticall(e.Context, e.client, rocketStorageAddr, common.HexToAddress(multicall3Addr))
	if err != nil {
		return fmt.Errorf("error initializing multicall3 dataprovider: %w", err)
	}
	nodes, err := mc.GetAllNodes(opts)
	if err != nil {
		return fmt.Errorf("error getting all nodes: %w", err)
	}
	loggerFunc("Loaded nodes", zap.Int("count", len(nodes)), zap.Int64("block", opts.BlockNumber.Int64()))

	// Get all minipools at the given block
	minipools, err := mc.GetAllMinipools(nodes, opts)
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
	odaoNodes, err := mc.GetAllOdaoNodes(opts)
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

	smoothingPoolAddress := mc.GetSmoothingPoolAddress()
	rethAddress := mc.GetREthAddress()

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
	var err error

	e.m = metrics.NewMetricsRegistry("execution_layer")
	// Make sure RocketStorageAddr is a valid address
	decoded, err := hex.DecodeString(e.RocketStorageAddr[2:])
	if err != nil {
		return fmt.Errorf("invalid rocket storage address: %w", err)
	}
	if len(decoded) != len(common.Address{}) {
		return fmt.Errorf("invalid rocket storage address: %w", err)
	}

	if e.RefreshInterval == 0 {
		return fmt.Errorf("must specify refresh interval")
	}

	e.client, err = ethclient.Dial(e.ECURL.String())
	if err != nil {
		return err
	}

	// Set up the vaults checker
	if e.SWVaultsRegistryAddr != "" {
		e.vaultsChecker = stakewise.NewVaultsChecker(e.client, common.HexToAddress(e.SWVaultsRegistryAddr))
	}

	if e.Context == nil {
		e.Context = context.Background()
	}

	if err := e.newCache(e.Logger.Info); err != nil {
		return fmt.Errorf("error warming up the cache: %w", err)
	}

	// Once the cache is warm, start a background process to refresh it
	e.ticker = time.NewTicker(e.RefreshInterval)
	go func() {
		e.Logger.Info("Starting cache refresh", zap.Duration("interval", e.RefreshInterval))
		for {
			select {
			case <-e.Context.Done():
				e.Logger.Info("Stopping cache refresh")
				e.ticker.Stop()
				return
			case <-e.ticker.C:
				e.Logger.Info("Refreshing cache")
				e.newCache(e.Logger.Debug)
			}
		}
	}()

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

// EIP1271ABI is the ABI for the EIP-1271 isValidSignature function
var eip1271ABI *abi.ABI

// getEIP1271ABI returns the EIP1271ABI
func getEIP1271ABI() *abi.ABI {
	if eip1271ABI != nil {
		return eip1271ABI
	}

	const abiJSON = `[{"inputs":[{"name":"_hash","type":"bytes32"},{"name":"_signature","type":"bytes"}],"name":"isValidSignature","outputs":[{"type":"bytes4"}],"stateMutability":"view","type":"function"}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		panic(fmt.Sprintf("failed to parse EIP1271 ABI: %v", err))
	}
	eip1271ABI = &parsedABI
	return eip1271ABI
}

var ErrNoData = errors.New("no data were returned from the EVM, did you pass the correct smart contract wallet address?")
var ErrBadData = errors.New("the evm returned data with an unexpected length, did you pass the correct smart contract wallet address?")
var ErrInternal = errors.New("an internal error occurred, please contact the maintainers")

// ValidateEIP1271 validates an EIP-1271 signature
func (e *CachingExecutionLayer) ValidateEIP1271(ctx context.Context, dataHash common.Hash, signature []byte, address common.Address) (bool, error) {
	parsedABI := getEIP1271ABI()

	// Encode the function call
	encodedData, err := parsedABI.Pack("isValidSignature", dataHash, signature)
	if err != nil {
		e.Logger.Warn("error packing isValidSignature call", zap.Error(err))
		return false, ErrInternal
	}

	// Make the contract call
	data, err := e.client.CallContract(ctx, ethereum.CallMsg{
		To:   &address,
		Data: encodedData,
	}, nil)
	if err != nil {
		e.Logger.Warn("error querying the execution client to validate an EIP1271 signature", zap.Error(err))
		return false, ErrInternal
	}

	if len(data) == 0 {
		return false, ErrNoData
	}

	if len(data) < 4 {
		return false, ErrBadData
	}

	// Trim the trailing bytes from the evm
	data = data[:4]

	// Check the return value, it should be exactly 4 bytes long
	if len(data) != 4 {
		return false, ErrBadData
	}

	// The expected return value for a valid signature is 0x1626ba7e
	// bytes4(keccak256("isValidSignature(bytes32,bytes)")
	// invalid signatures return 4 bytes that do not match the magic
	expectedReturnValue := [4]byte{0x16, 0x26, 0xba, 0x7e}
	return bytes.Equal(data, expectedReturnValue[:]), nil
}

func (e *CachingExecutionLayer) StakewiseFeeRecipient(ctx context.Context, address common.Address) (*common.Address, error) {
	if e.vaultsChecker != nil {
		return e.vaultsChecker.IsVault(ctx, address)
	}
	return nil, nil
}
