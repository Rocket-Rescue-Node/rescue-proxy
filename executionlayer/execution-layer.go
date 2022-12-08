package executionlayer

import (
	"bytes"
	"context"
	"math/big"
	"net/url"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rocket-pool/rocketpool-go/minipool"
	"github.com/rocket-pool/rocketpool-go/node"
	"github.com/rocket-pool/rocketpool-go/rocketpool"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"go.uber.org/zap"
)

type nodeInfo struct {
	inSmoothingPool bool
	feeDistributor  common.Address
}

// ExecutionLayer is a bespoke execution layer client for the rescue proxy.
// It abstracts away all the work to cache in-memory the data needed to enforce
// that fee recipients are 'correct'.
type ExecutionLayer struct {
	// Fields passed in by the constructor which are later referenced

	logger            *zap.Logger
	ecURL             *url.URL
	rocketStorageAddr string

	// The rocketpool-go client and its ethclient instance

	rp     *rocketpool.RocketPool
	client *ethclient.Client

	// Smart contracts we either read from or need the address of

	rocketNodeManager     *rocketpool.Contract
	rocketMinipoolManager *rocketpool.Contract
	smoothingPool         *rocketpool.Contract

	// The "topics" of the events we subscribe to

	nodeRegisteredTopic             common.Hash
	smoothingPoolStatusChangedTopic common.Hash
	minipoolLaunchedTopic           common.Hash

	// A long-lived index of pubkey->node address
	//
	// If a guarded query contains a pubkey we've seen before, and the fee recipient is
	// the smoothing pool, no further validation is needed, so we can exit early based
	// on membership in this map.
	//
	// Since this index is expected to strictly grow, we can use sync.Map to deal with
	// concurrent access. Elements are only inserted, never deleted.
	minipoolIndex *sync.Map

	// We need to store each node's smoothing pool status and fee recipient address.
	// We will subscribe to rocketNodeManager's events stream, which will notify us of
	// changes- to keep map contention down, we will use pointers as elements.
	// Ergo, this is a map of node address -> *Node
	nodeIndex *sync.Map

	// We need to detect gaps in the event stream when there are connection issues, and
	// backfill missing data, so we keep track of the highest block for which we received
	// an event here.
	highestBlock *big.Int

	// ethclient subscription needs to be manually closed on shutdown
	ethclientShutdownCb func()
}

// NewExecutionLayer creates an ExecutionLayer with the provided ec URL, rocketStorage address, and logger
func NewExecutionLayer(ecURL *url.URL, rocketStorageAddr string, logger *zap.Logger) *ExecutionLayer {
	out := &ExecutionLayer{}
	out.logger = logger
	out.minipoolIndex = &sync.Map{}
	out.nodeIndex = &sync.Map{}
	out.nodeRegisteredTopic = crypto.Keccak256Hash([]byte("NodeRegistered(address,uint256)"))
	out.smoothingPoolStatusChangedTopic = crypto.Keccak256Hash([]byte("NodeSmoothingPoolStateChanged(address,bool)"))
	out.minipoolLaunchedTopic = crypto.Keccak256Hash([]byte("MinipoolCreated(address,address,uint256)"))
	out.rocketStorageAddr = rocketStorageAddr
	out.ecURL = ecURL

	return out
}

func (e *ExecutionLayer) setECShutdownCb(cb func()) {
	if cb == nil {
		e.ethclientShutdownCb = nil
		return
	}

	e.ethclientShutdownCb = func() {
		cb()
		e.logger.Debug("Unsubscribed from EL events")
	}
}

func (e *ExecutionLayer) handleNodeEvent(event types.Log) {

	// Check if it's a node registration
	if bytes.Equal(event.Topics[0].Bytes(), e.nodeRegisteredTopic.Bytes()) {
		// When we see new nodes register, assume they aren't in the SP and add to index
		nodeInfo := &nodeInfo{}
		addr := common.BytesToAddress(event.Topics[1].Bytes())
		e.nodeIndex.Store(addr, nodeInfo)
		e.logger.Debug("New node registered", zap.String("addr", addr.String()))
		return
	}

	// Otherwise it should be a smoothing pool update
	if bytes.Equal(event.Topics[0].Bytes(), e.smoothingPoolStatusChangedTopic.Bytes()) {
		var n *nodeInfo
		// When we see a SP status change, update the pointer in the index
		nodeAddr := common.BytesToAddress(event.Topics[1].Bytes())
		status := big.NewInt(0).SetBytes(event.Data)

		// Attempt to load the node
		ptr, ok := e.nodeIndex.Load(nodeAddr)
		if ok {
			n = ptr.(*nodeInfo)
		} else {
			var err error

			// Odd that we don't have this node already, but add it and carry on
			e.logger.Warn("Unknown node updated its smoothing pool status", zap.String("addr", nodeAddr.String()))
			n = &nodeInfo{}
			// Get their fee distributor address
			n.feeDistributor, err = node.GetDistributorAddress(e.rp, nodeAddr, nil)
			if err != nil {
				e.logger.Warn("Couldn't compute fee distributor address for unknown node", zap.String("node", nodeAddr.String()))
			}
		}

		e.logger.Debug("Node SP status changed", zap.String("addr", nodeAddr.String()), zap.Bool("in_sp", status.Cmp(big.NewInt(1)) == 0))
		n.inSmoothingPool = status.Cmp(big.NewInt(1)) == 0
		return
	}

	e.logger.Warn("Event with unknown topic received", zap.String("string", event.Topics[0].String()))
}

func (e *ExecutionLayer) handleMinipoolEvent(event types.Log) {

	// Make sure it's an event for the only topic we subscribed to, minipool launches
	if !bytes.Equal(event.Topics[0].Bytes(), e.minipoolLaunchedTopic.Bytes()) {
		e.logger.Warn("Event with unknown topic received", zap.String("string", event.Topics[0].String()))
		return
	}

	// When a new minipool launches, grab its node address.
	nodeAddr := common.BytesToAddress(event.Topics[2].Bytes())

	// Grab its minipool (contract) address and use that to find its public key
	minipoolAddr := common.BytesToAddress(event.Topics[1].Bytes())
	minipoolDetails, err := minipool.GetMinipoolDetails(e.rp, minipoolAddr, nil)
	if err != nil {
		e.logger.Warn("Error fetching minipool details for new minipools", zap.String("minipool", minipoolAddr.String()), zap.Error(err))
		return
	}

	// Finally, update the minipool index
	e.minipoolIndex.Store(minipoolDetails.Pubkey, nodeAddr)
	e.logger.Debug("Added new minipool", zap.String("pubkey", minipoolDetails.Pubkey.String()), zap.String("node", nodeAddr.String()))
}

func (e *ExecutionLayer) handleEvent(event types.Log) {
	// events from the rocketNodeManager contract
	if bytes.Equal(e.rocketNodeManager.Address[:], event.Address[:]) {
		e.handleNodeEvent(event)
		goto out
	}

	// events from the rocketMinipoolManager contract
	if bytes.Equal(e.rocketMinipoolManager.Address[:], event.Address[:]) {
		e.handleMinipoolEvent(event)
		goto out
	}

	// Shouldn't ever happen, barring a bug in ethclient
	e.logger.Warn("Received event for unknown contract", zap.String("address", event.Address.String()))
out:
	// We should always update highestBlock when we receive any event
	e.highestBlock = big.NewInt(int64(event.BlockNumber))
}

// Gets the current block and loads any events we missed between highestBlock and the current one
// If we get disconnected from the EC, we may need to backfill.
// Additionally, we do some slow work on startup which takes 8-9 blocks, so we do that work,
// subscribe for future events, and then backfill the events in between before processing future events.
//
// All of this is only necessary because SubscribeFilterLogs doesn't seem to send old events, no matter
// what FromBlock is set to.
func (e *ExecutionLayer) backfillEvents() error {
	start := big.NewInt(0).Set(e.highestBlock)

	// Get current block
	header, err := e.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}
	stop := header.Number

	// Make sure there is actually a gap before backfilling
	if stop.Cmp(start) < 1 {
		e.logger.Debug("No blocks to backfill events from")
		return nil
	}

	missedEvents, err := e.client.FilterLogs(context.Background(), ethereum.FilterQuery{
		// We only want events for 2 contracts
		Addresses: []common.Address{*e.rocketMinipoolManager.Address, *e.rocketNodeManager.Address},
		// Since highestBlock was the highest processed block, start one block after
		FromBlock: big.NewInt(0).Add(start, big.NewInt(1)),
		// The current block is actually the last processed block, so play any events from it as well
		// The range is inclusive
		ToBlock: stop,
		// And we only care about 3 event types
		Topics: [][]common.Hash{[]common.Hash{e.nodeRegisteredTopic, e.smoothingPoolStatusChangedTopic, e.minipoolLaunchedTopic}},
	})

	if err != nil {
		return err
	}

	for _, event := range missedEvents {
		e.handleEvent(event)
	}

	e.logger.Debug("Backfilled events", zap.Int("events", len(missedEvents)), zap.Uint64("blocks", stop.Sub(stop, start).Uint64()))
	return nil
}

// Registers to receive the events we care about
func (e *ExecutionLayer) ecEventsConnect(opts *bind.CallOpts) error {
	var err error

	// Subscribe to events from rocketNodeManager and rocketMinipoolManager
	query := ethereum.FilterQuery{
		Addresses: []common.Address{*e.rocketMinipoolManager.Address, *e.rocketNodeManager.Address},
		Topics:    [][]common.Hash{[]common.Hash{e.nodeRegisteredTopic, e.smoothingPoolStatusChangedTopic, e.minipoolLaunchedTopic}},
	}

	// Set highestBlock to the same block that we used to build the cache from cold
	// TODO: If we add snapshots, save the highest block of the snapshot and start from there
	e.highestBlock = opts.BlockNumber

	events := make(chan types.Log)
	sub, err := e.client.SubscribeFilterLogs(context.Background(), query, events)
	if err != nil {
		return err
	}

	e.logger.Debug("Subscribed to EL events")

	// After subscribing, we need to grab the current block and replay events between highestBlock and the current one.
	// While we were building the cache from cold, we may have missed some events.
	err = e.backfillEvents()
	if err != nil {
		return err
	}

	// Make sure we can unsubscribe on shutdown
	e.setECShutdownCb(sub.Unsubscribe)

	// Start listening for events in a separate routine
	go func(subscription *ethereum.Subscription) {
		for {
			select {
			case err := <-(*subscription).Err():
				e.logger.Warn("Error received from eth client subscription", zap.Error(err))
				// Attempt to reconnect 10 times with steadily increasing waits
				reconnected := false
				for i := 0; i < 10; i++ {
					e.logger.Warn("Attempting to reconnect", zap.Int("attempt", i+1))
					s, err := e.client.SubscribeFilterLogs(context.Background(), query, events)
					if err == nil {
						e.logger.Warn("Reconnected", zap.Int("attempt", i+1))
						e.setECShutdownCb(s.Unsubscribe)
						reconnected = true
						subscription = &s
						break
					}

					e.logger.Warn("Error trying to reconnect to execution client", zap.Error(err))
					time.Sleep(time.Duration(i) * (5 * time.Second))
				}

				if !reconnected {
					// Failed to reconnect after 10 tries
					e.logger.Panic("Couldn't re-establish eth client connection")
				}

				// Now that we've reconnected, we need to backfill
				err = e.backfillEvents()
				if err != nil {
					// Failed to backfill
					e.logger.Panic("Couldn't backfill blocks after reconnecting to execution client")
				}

			case event := <-events:
				e.handleEvent(event)
			}
		}
	}(&sub)

	return nil
}

// Init creates and warms up the ExecutionLayer cache.
func (e *ExecutionLayer) Init() error {
	var err error

	e.client, err = ethclient.Dial(e.ecURL.String())
	if err != nil {
		return err
	}
	e.rp, err = rocketpool.NewRocketPool(e.client, common.HexToAddress(e.rocketStorageAddr))
	if err != nil {
		return err
	}
	// First, get the current block
	header, err := e.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	// Create opts to query state at that block
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

	// Get all nodes at the given block
	nodes, err := node.GetNodes(e.rp, opts)
	if err != nil {
		return err
	}
	e.logger.Debug("Found nodes to preload", zap.Int("count", len(nodes)))

	minipoolCount := 0
	for _, n := range nodes {
		// Allocate a pointer for this node
		nodeInfo := &nodeInfo{}
		// Determine their smoothing pool status
		nodeInfo.inSmoothingPool, err = node.GetSmoothingPoolRegistrationState(e.rp, n.Address, opts)
		if err != nil {
			return err
		}

		// Get their fee distributor address
		nodeInfo.feeDistributor, err = node.GetDistributorAddress(e.rp, n.Address, opts)
		if err != nil {
			return err
		}

		// Store the smoothing pool state / fee distributor in the node index
		e.nodeIndex.Store(n.Address, nodeInfo)

		// Also grab their minipools
		minipools, err := minipool.GetNodeMinipools(e.rp, n.Address, opts)
		if err != nil {
			return err
		}

		minipoolCount += len(minipools)
		for _, minipool := range minipools {
			e.minipoolIndex.Store(minipool.Pubkey, n.Address)
		}
	}
	e.logger.Debug("Pre-loaded nodes and minipools", zap.Int("nodes", len(nodes)), zap.Int("minipools", minipoolCount))

	// Listen for updates
	e.ecEventsConnect(opts)

	return nil
}

// Deinit shuts down this ExecutionLayer
func (e *ExecutionLayer) Deinit() {
	if e.ethclientShutdownCb == nil {
		return
	}
	e.ethclientShutdownCb()
}

type ForEachNodeClosure func(common.Address) bool

// ForEachNode calls the provided closure with the address of every rocket pool node the ExecutionLayer has observed
func (e *ExecutionLayer) ForEachNode(closure ForEachNodeClosure) {
	e.nodeIndex.Range(func (k any, value any) bool {
		return closure(k.(common.Address))
	})
}

// ValidatorFeeRecipient returns the expected fee recipient for a validator, or nil if the validator is "unknown"
// If the queryNodeAddr is not nil and the validator is a minipool but isn't owned by that node, (nil, true) is returned
func (e *ExecutionLayer) ValidatorFeeRecipient(pubkey rptypes.ValidatorPubkey, queryNodeAddr *common.Address) (*common.Address, bool) {

	void, ok := e.minipoolIndex.Load(pubkey)
	if !ok {
		// Validator (hopefully) isn't a minipool
		return nil, false
	}

	nodeAddr := void.(common.Address)

	if queryNodeAddr != nil && !bytes.Equal(queryNodeAddr.Bytes(), nodeAddr.Bytes()) {
		return nil, true
	}

	ptr, ok := e.nodeIndex.Load(nodeAddr)
	if !ok {
		// Validator was a minipool, but we don't have a node record for it. This is bad.
		e.logger.Error("Validator was in the minipool index, but not the node index",
			zap.String("pubkey", pubkey.String()),
			zap.String("node", nodeAddr.String()))
		return nil, false
	}

	nodeInfo := ptr.(*nodeInfo)
	if nodeInfo.inSmoothingPool {
		return e.smoothingPool.Address, false
	}

	return &nodeInfo.feeDistributor, false
}
