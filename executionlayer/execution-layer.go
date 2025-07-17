package executionlayer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer/stakewise"
	"github.com/Rocket-Rescue-Node/rescue-proxy/metrics"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rocket-pool/rocketpool-go/dao/trustednode"
	"github.com/rocket-pool/rocketpool-go/minipool"
	"github.com/rocket-pool/rocketpool-go/node"
	"github.com/rocket-pool/rocketpool-go/rocketpool"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type ForEachNodeClosure func(common.Address) bool

const reconnectRetries = 10
const maxCacheAgeBlocks = 64

type nodeInfo struct {
	inSmoothingPool bool
	feeDistributor  common.Address
}

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

// CachingExecutionLayer is a bespoke execution layer client for the rescue proxy.
// It abstracts away all the work to cache in-memory the data needed to enforce
// that fee recipients are 'correct'.
type CachingExecutionLayer struct {
	// Fields passed in by the constructor which are later referenced
	Logger               *zap.Logger
	ECURL                *url.URL
	RocketStorageAddr    string
	SWVaultsRegistryAddr string

	// The rocketpool-go client and its ethclient instance

	rp     *rocketpool.RocketPool
	client *ethclient.Client

	// Smart contracts we either read from or need the address of

	rocketNodeManager           *rocketpool.Contract
	rocketMinipoolManager       *rocketpool.Contract
	smoothingPool               *rocketpool.Contract
	rEth                        *rocketpool.Contract
	rocketDaoNodeTrustedActions *rocketpool.Contract

	// The "topics" of the events we subscribe to

	nodeRegisteredTopic             common.Hash
	smoothingPoolStatusChangedTopic common.Hash
	minipoolLaunchedTopic           common.Hash
	odaoJoinedTopic                 common.Hash
	odaoLeftTopic                   common.Hash
	odaoKickedTopic                 common.Hash

	// The "topics" and contract filter for the events we subscribe to
	query ethereum.FilterQuery

	// Channels for those subscriptions
	events     chan types.Log
	newHeaders chan *types.Header

	// Somewhere to store chain data we care about
	CachePath string
	cache     Cache

	// Checkers for vaults and mev escrow
	vaultsChecker *stakewise.VaultsChecker

	// ethclient subscription needs to be manually closed on shutdown
	ethclientShutdownCb func()

	// A context that can be canceled in order to gracefully stop the EL abstraction
	ctx context.Context
	// A function that cancels that context
	shutdown func()

	m *metrics.MetricsRegistry

	connected chan bool
}

func (e *CachingExecutionLayer) setECShutdownCb(cb func()) {
	if cb == nil {
		e.ethclientShutdownCb = nil
		return
	}

	e.ethclientShutdownCb = func() {
		cb()
		e.Logger.Info("Unsubscribed from EL events")
	}
}

func (e *CachingExecutionLayer) handleNodeEvent(event types.Log) {

	// Check if it's a node registration
	if bytes.Equal(event.Topics[0].Bytes(), e.nodeRegisteredTopic.Bytes()) {
		var err error

		addr := common.BytesToAddress(event.Topics[1].Bytes())
		// When we see new nodes register, assume they aren't in the SP and add to index
		nodeInfo := &nodeInfo{}
		// Get their fee distributor address
		nodeInfo.feeDistributor, err = node.GetDistributorAddress(e.rp, addr, nil)
		if err != nil {
			e.Logger.Warn("Couldn't get fee distributor address for newly registered node", zap.String("node", addr.String()))
		}
		err = e.cache.addNodeInfo(addr, nodeInfo)
		if err != nil {
			e.Logger.Error("Failed to add nodeInfo to cache", zap.Error(err))
		}

		e.m.Counter("node_registration_added").Inc()
		e.Logger.Info("New node registered", zap.String("addr", addr.String()))
		return
	}

	// Otherwise it should be a smoothing pool update
	if bytes.Equal(event.Topics[0].Bytes(), e.smoothingPoolStatusChangedTopic.Bytes()) {
		var n *nodeInfo
		// When we see a SP status change, update the pointer in the index
		nodeAddr := common.BytesToAddress(event.Topics[1].Bytes())
		status := big.NewInt(0).SetBytes(event.Data)

		// Attempt to load the node
		n, err := e.cache.getNodeInfo(nodeAddr)
		if err != nil {
			_, ok := err.(*NotFoundError)

			if !ok {
				e.Logger.Panic("Got an error from the cache while looking up a node",
					zap.String("addr", nodeAddr.String()), zap.Error(err))
			}

			// Odd that we don't have this node already, but add it and carry on
			e.Logger.Warn("Unknown node updated its smoothing pool status", zap.String("addr", nodeAddr.String()))
			n = &nodeInfo{}
			// Get their fee distributor address
			n.feeDistributor, err = node.GetDistributorAddress(e.rp, nodeAddr, nil)
			if err != nil {
				e.Logger.Warn("Couldn't compute fee distributor address for unknown node", zap.String("node", nodeAddr.String()))
			}

		}

		e.Logger.Info("Node SP status changed", zap.String("addr", nodeAddr.String()), zap.Bool("in_sp", status.Cmp(big.NewInt(1)) == 0))
		n.inSmoothingPool = status.Cmp(big.NewInt(1)) == 0
		err = e.cache.addNodeInfo(nodeAddr, n)
		if err != nil {
			e.Logger.Error("Failed to add nodeInfo to cache", zap.Error(err))
		}

		e.m.Counter("smoothing_pool_status_changed").Inc()
		return
	}

	e.Logger.Warn("Event with unknown topic received", zap.String("string", event.Topics[0].String()))
}

func (e *CachingExecutionLayer) handleMinipoolEvent(event types.Log) {

	// Make sure it's an event for the only topic we subscribed to, minipool launches
	if !bytes.Equal(event.Topics[0].Bytes(), e.minipoolLaunchedTopic.Bytes()) {
		e.Logger.Warn("Event with unknown topic received", zap.String("string", event.Topics[0].String()))
		return
	}

	// When a new minipool launches, grab its node address.
	nodeAddr := common.BytesToAddress(event.Topics[2].Bytes())

	// Grab its minipool (contract) address and use that to find its public key
	minipoolAddr := common.BytesToAddress(event.Topics[1].Bytes())
	pubkey, err := minipool.GetMinipoolPubkey(e.rp, minipoolAddr, nil)
	if err != nil {
		e.Logger.Warn("Error fetching minipool pubkey for new minipool", zap.String("minipool", minipoolAddr.String()), zap.Error(err))
		return
	}

	// Finally, update the minipool index
	err = e.cache.addMinipoolNode(pubkey, nodeAddr)
	if err != nil {
		e.Logger.Warn("Error updating minipool cache", zap.Error(err))
	}
	e.m.Counter("minipool_launch_received").Inc()
	e.Logger.Info("Added new minipool", zap.String("pubkey", pubkey.String()), zap.String("node", nodeAddr.String()))
}

func (e *CachingExecutionLayer) handleOdaoEvent(event types.Log) {

	if bytes.Equal(event.Topics[0].Bytes(), e.odaoJoinedTopic.Bytes()) {
		addr := common.BytesToAddress(event.Topics[1].Bytes())

		err := e.cache.addOdaoNode(addr)
		if err != nil {
			e.Logger.Warn("Error updating odao cache", zap.Error(err))
		}
		return
	}

	if bytes.Equal(event.Topics[0].Bytes(), e.odaoLeftTopic.Bytes()) ||
		bytes.Equal(event.Topics[0].Bytes(), e.odaoKickedTopic.Bytes()) {

		addr := common.BytesToAddress(event.Topics[1].Bytes())

		err := e.cache.removeOdaoNode(addr)
		if err != nil {
			e.Logger.Warn("Error updating odao cache", zap.Error(err))
		}
		return
	}

	e.Logger.Warn("Event with unknown topic received", zap.String("string", event.Topics[0].String()))
}

func (e *CachingExecutionLayer) handleEvent(event types.Log) {
	// events from the rocketNodeManager contract
	e.m.Counter("subscription_event_received").Inc()
	if bytes.Equal(e.rocketNodeManager.Address[:], event.Address[:]) {
		e.handleNodeEvent(event)
		goto out
	}

	// events from the rocketMinipoolManager contract
	if bytes.Equal(e.rocketMinipoolManager.Address[:], event.Address[:]) {
		e.handleMinipoolEvent(event)
		goto out
	}

	// events from the rocketDAONodeTrustedActions contract
	if bytes.Equal(e.rocketDaoNodeTrustedActions.Address[:], event.Address[:]) {
		e.handleOdaoEvent(event)
		goto out
	}

	// Shouldn't ever happen, barring a bug in ethclient
	e.Logger.Warn("Received event for unknown contract", zap.String("address", event.Address.String()))
out:
	// We should always update highestBlock when we receive any event
	e.cache.setHighestBlock(big.NewInt(int64(event.BlockNumber)))
}

// Gets the current block and loads any events we missed between highestBlock and the current one
// If we get disconnected from the EC, we may need to backfill.
// Additionally, we do some slow work on startup which takes 8-9 blocks, so we do that work,
// subscribe for future events, and then backfill the events in between before processing future events.
//
// All of this is only necessary because SubscribeFilterLogs doesn't seem to send old events, no matter
// what FromBlock is set to.
func (e *CachingExecutionLayer) backfillEvents() error {
	// Since highestBlock was the highest processed block, start one block after
	start := big.NewInt(0).Add(e.cache.getHighestBlock(), big.NewInt(1))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get current block
	header, err := e.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	stop := header.Number

	e.Logger.Info("Checking if backfill neeeded",
		zap.Uint64("cache height", start.Uint64()),
		zap.Uint64("current height", stop.Uint64()),
		zap.Int("cmp", start.Cmp(stop)))

	// Make sure there is actually a gap before backfilling
	if stop.Cmp(start) < 0 {
		e.Logger.Info("No blocks to backfill events from")
		return nil
	}

	missedEvents, err := e.client.FilterLogs(ctx, ethereum.FilterQuery{
		// We only want events for 2 contracts
		Addresses: []common.Address{*e.rocketMinipoolManager.Address, *e.rocketNodeManager.Address},
		FromBlock: start,
		// The current block is actually the last block processed by the EC, so play any events from it as well
		// The range is inclusive
		ToBlock: stop,
		Topics: [][]common.Hash{{
			e.nodeRegisteredTopic,
			e.smoothingPoolStatusChangedTopic,
			e.minipoolLaunchedTopic,
			e.odaoJoinedTopic,
			e.odaoLeftTopic,
			e.odaoKickedTopic,
		}},
	})

	if err != nil {
		return err
	}

	for _, event := range missedEvents {
		e.handleEvent(event)
		e.m.Counter("backfill_events").Inc()
	}

	// Force the highest block to update, as we may not have received any events in it, which would have updated it
	e.cache.setHighestBlock(stop)

	delta := big.NewInt(0).Sub(stop, start)

	// If start == stop we actually fill that one block, so add one to delta
	delta = delta.Add(delta, big.NewInt(1))
	e.m.Counter("backfill_blocks").Add(float64(delta.Uint64()))

	e.Logger.Info("Backfilled events", zap.Int("events", len(missedEvents)),
		zap.Uint64("blocks", delta.Uint64()),
		zap.Int64("start", start.Int64()), zap.Int64("stop", stop.Int64()))
	return nil
}

// Will likely attempt to reconnect, and will overwrite the pointers passed with the new subscription objects
func (e *CachingExecutionLayer) handleSubscriptionError(err error, logEventSub **ethereum.Subscription, headerSub **ethereum.Subscription) {
	if e.ctx.Err() != nil {
		// We're shutting down, so return quietly
		return
	}

	e.m.Counter("subscription_disconnected").Inc()
	e.Logger.Warn("Error received from eth client subscription", zap.Error(err))
	// Attempt to reconnect `reconnectRetries` times with steadily increasing waits
	for i := 0; i < reconnectRetries; i++ {

		e.Logger.Warn("Attempting to reconnect", zap.Int("attempt", i+1))
		e.m.Counter("reconnection_attempt").Inc()
		s, err := e.client.SubscribeFilterLogs(context.Background(), e.query, e.events)
		if err == nil {
			e.Logger.Warn("Reconnected", zap.Int("attempt", i+1))

			// Resubscribe to new headers - no retries
			h, err := e.client.SubscribeNewHead(context.Background(), e.newHeaders)
			if err != nil {
				e.Logger.Warn("Couldn't resubscribe to block headers after reconnecting")
				break
			}

			e.setECShutdownCb(func() {
				s.Unsubscribe()
				h.Unsubscribe()
			})

			// Now that we've reconnected, we need to backfill
			err = e.backfillEvents()
			if err != nil {
				// Failed to backfill
				e.Logger.Panic("Couldn't backfill blocks after reconnecting to execution client")
			}

			*logEventSub = &s
			*headerSub = &h
			return
		}

		e.Logger.Warn("Error trying to reconnect to execution client", zap.Error(err))
		select {
		case <-e.ctx.Done():
			// We're shutting down, so exit now
			e.Logger.Info("Terminating while re-establishing the connection to the EL")
			return
		case <-time.After(time.Duration(i) * (5 * time.Second)):
			// Loop again
		}
	}

	// Failed to reconnect after 10 tries
	e.Logger.Panic("Couldn't re-establish eth client connection")
}

// Registers to receive the events we care about
func (e *CachingExecutionLayer) ecEventsConnect(_ *bind.CallOpts) error {
	var err error

	e.nodeRegisteredTopic = crypto.Keccak256Hash([]byte("NodeRegistered(address,uint256)"))
	e.smoothingPoolStatusChangedTopic = crypto.Keccak256Hash([]byte("NodeSmoothingPoolStateChanged(address,bool)"))
	e.minipoolLaunchedTopic = crypto.Keccak256Hash([]byte("MinipoolCreated(address,address,uint256)"))
	e.odaoJoinedTopic = crypto.Keccak256Hash([]byte("ActionJoined(address,uint256,uint256)"))
	e.odaoLeftTopic = crypto.Keccak256Hash([]byte("ActionLeave(address,uint256,uint256)"))
	e.odaoKickedTopic = crypto.Keccak256Hash([]byte("ActionKick(address,uint256,uint256)"))

	// Subscribe to events from rocketNodeManager and rocketMinipoolManager
	e.query = ethereum.FilterQuery{
		Addresses: []common.Address{*e.rocketMinipoolManager.Address, *e.rocketNodeManager.Address},
		Topics: [][]common.Hash{{
			e.nodeRegisteredTopic,
			e.smoothingPoolStatusChangedTopic,
			e.minipoolLaunchedTopic,
			e.odaoJoinedTopic,
			e.odaoLeftTopic,
			e.odaoKickedTopic,
		}},
	}

	e.events = make(chan types.Log, 32)
	sub, err := e.client.SubscribeFilterLogs(context.Background(), e.query, e.events)
	if err != nil {
		return err
	}

	e.newHeaders = make(chan *types.Header, 32)
	newHeadSub, err := e.client.SubscribeNewHead(context.Background(), e.newHeaders)
	if err != nil {
		return err
	}

	e.Logger.Info("Subscribed to EL events")

	// After subscribing, we need to grab the current block and replay events between highestBlock and the current one.
	// While we were building the cache from cold, we may have missed some events.
	err = e.backfillEvents()
	if err != nil {
		return err
	}

	// Make sure we can unsubscribe on shutdown
	e.setECShutdownCb(func() {
		sub.Unsubscribe()
		newHeadSub.Unsubscribe()
	})

	e.connected <- true
	{
		var noMoreEvents bool
		var noMoreHeaders bool

		logSubscription := &sub
		newHeadSubscription := &newHeadSub
		for {

			select {
			case err := <-(*logSubscription).Err():
				(*newHeadSubscription).Unsubscribe()
				e.handleSubscriptionError(err, &logSubscription, &newHeadSubscription)
			case err := <-(*newHeadSubscription).Err():
				(*logSubscription).Unsubscribe()
				e.handleSubscriptionError(err, &logSubscription, &newHeadSubscription)
			case event, ok := <-e.events:
				noMoreEvents = !ok
				if !noMoreEvents {
					e.handleEvent(event)
				}
			case newHeader, ok := <-e.newHeaders:
				noMoreHeaders = !ok
				if !noMoreHeaders {
					// Just advance highest block
					e.m.Counter("block_header_received").Inc()
					e.Logger.Debug("New block received",
						zap.Int64("new height", newHeader.Number.Int64()),
						zap.Int64("old height", e.cache.getHighestBlock().Int64()))
					e.cache.setHighestBlock(newHeader.Number)

					// Continue here to check for new events
					continue
				}
			}

			// If we didn't process any events in the select and the channels are closed,
			// no new events will come, so break the loop
			if noMoreEvents && noMoreHeaders {
				e.Logger.Debug("Finished processing events", zap.Int64("height", e.cache.getHighestBlock().Int64()))
				break
			}

		}
	}

	return nil
}

// Init creates and warms up the ExecutionLayer cache.
func (e *CachingExecutionLayer) Init() error {
	var err error

	e.m = metrics.NewMetricsRegistry("execution_layer")
	e.connected = make(chan bool, 1)

	// Pick a cache
	if e.CachePath == "" {
		e.cache = &MapsCache{}
	} else {
		e.cache = &SqliteCache{
			Path: e.CachePath,
		}
	}

	e.ctx, e.shutdown = context.WithCancel(context.Background())

	if err := e.cache.init(); err != nil {
		return err
	}
	cacheBlock := e.cache.getHighestBlock()

	e.client, err = ethclient.Dial(e.ECURL.String())
	if err != nil {
		return err
	}
	e.rp, err = rocketpool.NewRocketPool(e.client, common.HexToAddress(e.RocketStorageAddr))
	if err != nil {
		return err
	}

	// Set up the vaults checker
	if e.SWVaultsRegistryAddr != "" {
		e.vaultsChecker = stakewise.NewVaultsChecker(e.client, common.HexToAddress(e.SWVaultsRegistryAddr))
	}

	// First, get the current block
	ctx, cancel := context.WithTimeout(e.ctx, 5*time.Second)
	defer cancel()
	header, err := e.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	// Subtract the cache's highest block from the current block
	delta := big.NewInt(0)
	delta.Sub(header.Number, cacheBlock)
	if delta.Int64() < 0 || delta.Int64() > maxCacheAgeBlocks {
		// Reset caches from the future and the distance past
		e.Logger.Info("Cache is stale or from the future, resetting...",
			zap.Int64("cache block", cacheBlock.Int64()),
			zap.Int64("current block", header.Number.Int64()),
			zap.Int64("delta", delta.Int64()))
		err = e.cache.reset()
		if err != nil {
			return err
		}

		cacheBlock = big.NewInt(0)
	}

	// Create opts to query state at the latest block
	opts := &bind.CallOpts{BlockNumber: header.Number}

	// Load contracts
	e.rocketNodeManager, err = e.rp.GetContract("rocketNodeManager", opts)
	if err != nil {
		return err
	}

	e.rocketMinipoolManager, err = e.rp.GetContract("rocketMinipoolManager", opts)
	if err != nil {
		return err
	}

	e.smoothingPool, err = e.rp.GetContract("rocketSmoothingPool", opts)
	if err != nil {
		return err
	}

	e.rEth, err = e.rp.GetContract("rocketTokenRETH", opts)
	if err != nil {
		return err
	}

	e.rocketDaoNodeTrustedActions, err = e.rp.GetContract("rocketDAONodeTrustedActions", opts)
	if err != nil {
		return err
	}

	// If the cache is warm, skip the slow path
	if cacheBlock.Cmp(big.NewInt(0)) != 0 {
		return nil
	}
	e.Logger.Info("Warming up the cache")

	// Get all nodes at the given block
	nodes, err := node.GetNodeAddresses(e.rp, opts)
	if err != nil {
		return err
	}
	e.Logger.Info("Found nodes to preload", zap.Int("count", len(nodes)), zap.Int64("block", opts.BlockNumber.Int64()))

	minipoolCount := 0
	for _, addr := range nodes {
		// Allocate a pointer for this node
		nodeInfo := &nodeInfo{}
		// Determine their smoothing pool status
		nodeInfo.inSmoothingPool, err = node.GetSmoothingPoolRegistrationState(e.rp, addr, opts)
		if err != nil {
			return err
		}

		// Get their fee distributor address
		nodeInfo.feeDistributor, err = node.GetDistributorAddress(e.rp, addr, opts)
		if err != nil {
			return err
		}

		// Store the smoothing pool state / fee distributor in the node index
		err = e.cache.addNodeInfo(addr, nodeInfo)
		if err != nil {
			return err
		}

		// Also grab their minipools
		minipoolAddresses, err := minipool.GetNodeMinipoolAddresses(e.rp, addr, opts)
		if err != nil {
			return err
		}

		minipoolCount += len(minipoolAddresses)
		var wg errgroup.Group
		wg.SetLimit(64)
		for _, m := range minipoolAddresses {
			m := m
			wg.Go(func() error {
				pubkey, err := minipool.GetMinipoolPubkey(e.rp, m, opts)
				if err != nil {
					return err
				}
				err = e.cache.addMinipoolNode(pubkey, addr)
				if err != nil {
					return err
				}
				return nil
			})
		}
		err = wg.Wait()
		if err != nil {
			return err
		}
	}

	// Get all odao nodes at the given block
	odaoNodes, err := trustednode.GetMemberAddresses(e.rp, opts)
	if err != nil {
		return err
	}

	for _, member := range odaoNodes {

		err = e.cache.addOdaoNode(member)
		if err != nil {
			return err
		}
	}
	e.Logger.Info("Found odao nodes to preload", zap.Int("count", len(odaoNodes)), zap.Int64("block", opts.BlockNumber.Int64()))

	// Set highestBlock to the cache's highestBlock, since it was just warmed up
	e.cache.setHighestBlock(opts.BlockNumber)

	e.Logger.Info("Pre-loaded nodes and minipools",
		zap.Int("nodes", len(nodes)),
		zap.Int("minipools", minipoolCount),
		zap.Int("odao nodes", len(odaoNodes)))

	return nil
}

func (e *CachingExecutionLayer) Start() error {
	// First, get the current block
	header, err := e.client.HeaderByNumber(e.ctx, nil)
	if err != nil {
		return err
	}

	// Create opts to query state at the latest block
	opts := &bind.CallOpts{BlockNumber: header.Number}

	return e.ecEventsConnect(opts)
}

// Stop shuts down this ExecutionLayer
func (e *CachingExecutionLayer) Stop() {
	e.Logger.Info("Stopping ethclient")
	e.shutdown()
	if e.ethclientShutdownCb != nil {
		e.ethclientShutdownCb()
	}
	close(e.events)
	close(e.newHeaders)
	e.Logger.Info("Stopping EL cache")
	err := e.cache.deinit()
	if err != nil {
		e.Logger.Error("error while stopping the cache", zap.Error(err))
	}
	close(e.connected)
}

// ForEachNode calls the provided closure with the address of every rocket pool node the ExecutionLayer has observed
func (e *CachingExecutionLayer) ForEachNode(closure ForEachNodeClosure) error {
	return e.cache.forEachNode(closure)
}

// ForEachOdaoNode calls the provided closure with the address of every odao node the ExecutionLayer has observed
func (e *CachingExecutionLayer) ForEachOdaoNode(closure ForEachNodeClosure) error {
	return e.cache.forEachOdaoNode(closure)
}

// GetRPInfo returns the expected fee recipient and node address for a validator, or nil if the validator is not a minipool
func (e *CachingExecutionLayer) GetRPInfo(pubkey rptypes.ValidatorPubkey) (*RPInfo, error) {

	nodeAddr, err := e.cache.getMinipoolNode(pubkey)
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

	nodeInfo, err := e.cache.getNodeInfo(nodeAddr)
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

	if nodeInfo.inSmoothingPool {
		return &RPInfo{
			ExpectedFeeRecipient: e.smoothingPool.Address,
			NodeAddress:          nodeAddr,
		}, nil
	}

	return &RPInfo{
		ExpectedFeeRecipient: &nodeInfo.feeDistributor,
		NodeAddress:          nodeAddr,
	}, nil
}

// REthAddress is a convenience function to get the rEth contract address
func (e *CachingExecutionLayer) REthAddress() *common.Address {
	return e.rEth.Address
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
