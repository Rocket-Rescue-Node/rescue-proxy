package dataprovider

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer/dataprovider/abis"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/smartnode/bindings/types"
)

type Multicall struct {
	NodeBatchSize int

	client                       bind.ContractBackend
	multicallInstance            *bind.BoundContract
	rocketDaoNodeTrustedInstance *bind.BoundContract

	rocketNodeManagerAddress            common.Address
	rocketNodeDistributorFactoryAddress common.Address
	rocketMinipoolManagerAddress        common.Address
	rocketTokenREthAddress              common.Address
	rocketSmoothingPoolAddress          common.Address
}

var multicall3 *abis.Multicall3
var rocketStorage *abis.RocketStorage
var rocketNodeManager *abis.RocketNodeManager
var rocketNodeDistributorFactory *abis.RocketNodeDistributorFactory
var rocketMinipoolManager *abis.RocketMinipoolManager
var rocketDaoNodeTrusted *abis.RocketDaoNodeTrusted

func init() {
	multicall3 = abis.NewMulticall3()
	rocketStorage = abis.NewRocketStorage()
	rocketNodeManager = abis.NewRocketNodeManager()
	rocketNodeDistributorFactory = abis.NewRocketNodeDistributorFactory()
	rocketMinipoolManager = abis.NewRocketMinipoolManager()
	rocketDaoNodeTrusted = abis.NewRocketDaoNodeTrusted()
}

func NewMulticall(ctx context.Context, client bind.ContractBackend,
	rocketStorageAddress common.Address,
	contractAddress common.Address,
) (DataProvider, error) {

	storageInstance := rocketStorage.Instance(client, rocketStorageAddress)

	// get the head block height
	headBlock, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get head block: %w", err)
	}

	headBlockHeight := headBlock.Number
	opts := bind.CallOpts{
		Context:     ctx,
		BlockNumber: headBlockHeight,
	}

	// hashes precomputed from https://emn178.github.io/online-tools/keccak_256.html

	// nodeManagerKey is a keccak256 of "contract.addressrocketNodeManager"
	nodeManagerKey := common.HexToHash("0xaf00be55c9fb8f543c04e0aa0d70351b880c1bfafffd15b60065a4a50c85ec94")
	// get nodeManagerAddress from storage
	nodeManagerAddressCallPacked := rocketStorage.PackGetAddress(nodeManagerKey)
	nodeManagerAddressResponsePacked, err := storageInstance.CallRaw(&opts, nodeManagerAddressCallPacked)
	if err != nil {
		return nil, fmt.Errorf("failed to get node manager address: %w", err)
	}
	nodeManagerAddress, err := rocketStorage.UnpackGetAddress(nodeManagerAddressResponsePacked)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack node manager address: %w", err)
	}

	// nodeDistributorFactoryKey is a keccak256 of "contract.addressrocketNodeDistributorFactory"
	nodeDistributorFactoryKey := common.HexToHash("0xea051094896ef3b09ab1b794ad5ea695a5ff3906f74a9328e2c16db69d0f3123")
	nodeDistributorFactoryAddressCallPacked := rocketStorage.PackGetAddress(nodeDistributorFactoryKey)
	nodeDistributorFactoryAddressResponsePacked, err := storageInstance.CallRaw(&opts, nodeDistributorFactoryAddressCallPacked)
	if err != nil {
		return nil, fmt.Errorf("failed to get node distributor factory address: %w", err)
	}
	nodeDistributorFactoryAddress, err := rocketStorage.UnpackGetAddress(nodeDistributorFactoryAddressResponsePacked)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack node distributor factory address: %w", err)
	}

	// minipoolManagerKey is a keccak256 of "contract.addressrocketMinipoolManager"
	minipoolManagerKey := common.HexToHash("0xe9dfec9339b94a131861a58f1bb4ac4c1ce55c7ffe8550e0b6ebcfde87bb012f")
	minipoolManagerAddressCallPacked := rocketStorage.PackGetAddress(minipoolManagerKey)
	minipoolManagerAddressResponsePacked, err := storageInstance.CallRaw(&opts, minipoolManagerAddressCallPacked)
	if err != nil {
		return nil, fmt.Errorf("failed to get minipool manager address: %w", err)
	}
	minipoolManagerAddress, err := rocketStorage.UnpackGetAddress(minipoolManagerAddressResponsePacked)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack minipool manager address: %w", err)
	}

	// rocketDAONodeTrustedKey is a keccak256 of "contract.addressrocketDAONodeTrusted"
	rocketDAONodeTrustedKey := common.HexToHash("0x9a354e1bb2e38ca826db7a8d061cfb0ed7dbd83d241a2cbe4fd5218f9bb4333f")
	rocketDAONodeTrustedAddressCallPacked := rocketStorage.PackGetAddress(rocketDAONodeTrustedKey)
	rocketDAONodeTrustedAddressResponsePacked, err := storageInstance.CallRaw(&opts, rocketDAONodeTrustedAddressCallPacked)
	if err != nil {
		return nil, fmt.Errorf("failed to get rocket dao node trusted address: %w", err)
	}
	rocketDAONodeTrustedAddress, err := rocketStorage.UnpackGetAddress(rocketDAONodeTrustedAddressResponsePacked)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack rocket dao node trusted address: %w", err)
	}

	// rethAddressKey is a keccak256 of "contract.addressrocketTokenRETH"
	rethAddressKey := common.HexToHash("0xe3744443225bff7cc22028be036b80de58057d65a3fdca0a3df329f525e31ccc")
	rethAddressCallPacked := rocketStorage.PackGetAddress(rethAddressKey)
	rethAddressResponsePacked, err := storageInstance.CallRaw(&opts, rethAddressCallPacked)
	if err != nil {
		return nil, fmt.Errorf("failed to get reth address: %w", err)
	}
	rethAddress, err := rocketStorage.UnpackGetAddress(rethAddressResponsePacked)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack reth address: %w", err)
	}

	// smoothingPoolAddressKey is a keccak256 of "contract.addressrocketSmoothingPool"
	smoothingPoolAddressKey := common.HexToHash("0x822231720aef9b264db1d9ca053137498f759c28b243f45c44db1d39d6bce46e")
	smoothingPoolAddressCallPacked := rocketStorage.PackGetAddress(smoothingPoolAddressKey)
	smoothingPoolAddressResponsePacked, err := storageInstance.CallRaw(&opts, smoothingPoolAddressCallPacked)
	if err != nil {
		return nil, fmt.Errorf("failed to get smoothing pool address: %w", err)
	}
	smoothingPoolAddress, err := rocketStorage.UnpackGetAddress(smoothingPoolAddressResponsePacked)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack smoothing pool address: %w", err)
	}

	return &Multicall{
		client:                              client,
		multicallInstance:                   multicall3.Instance(client, contractAddress),
		rocketNodeManagerAddress:            nodeManagerAddress,
		rocketNodeDistributorFactoryAddress: nodeDistributorFactoryAddress,
		rocketMinipoolManagerAddress:        minipoolManagerAddress,
		rocketDaoNodeTrustedInstance:        rocketDaoNodeTrusted.Instance(client, rocketDAONodeTrustedAddress),
		rocketTokenREthAddress:              rethAddress,
		rocketSmoothingPoolAddress:          smoothingPoolAddress,
		NodeBatchSize:                       100,
	}, nil
}

func (m *Multicall) GetAllNodes(opts *bind.CallOpts) (map[common.Address]*NodeInfo, error) {
	// Create the rocketNodeManagerInstance
	rocketNodeManagerInstance := rocketNodeManager.Instance(m.client, m.rocketNodeManagerAddress)
	// Get the total number of registered nodes
	getNodeCountPacked := rocketNodeManager.PackGetNodeCount()
	getNodeCountResponsePacked, err := rocketNodeManagerInstance.CallRaw(opts, getNodeCountPacked)
	if err != nil {
		return nil, fmt.Errorf("failed to get node count: %w", err)
	}
	nodeCount, err := rocketNodeManager.UnpackGetNodeCount(getNodeCountResponsePacked)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack node count: %w", err)
	}

	// Grab up to m.NodeBatchSize at a time, and return the total number of nodes
	nodeMap := make(map[common.Address]*NodeInfo)
	nodeAddresses := make([]common.Address, 0, nodeCount.Int64())
	for i := int64(0); i < nodeCount.Int64(); i += int64(m.NodeBatchSize) {
		getNodeAddressesPacked := rocketNodeManager.PackGetNodeAddresses(big.NewInt(i), big.NewInt(int64(m.NodeBatchSize)))
		getNodesResponsePacked, err := rocketNodeManagerInstance.CallRaw(opts, getNodeAddressesPacked)
		if err != nil {
			return nil, fmt.Errorf("failed to get nodes: %w", err)
		}
		nodes, err := rocketNodeManager.UnpackGetNodeAddresses(getNodesResponsePacked)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack nodes: %w", err)
		}
		nodeAddresses = append(nodeAddresses, nodes...)
	}

	// In batches of 500, populate FeeDistributor
	for i := 0; i < len(nodeAddresses); i += m.NodeBatchSize {
		batch := nodeAddresses[i:min(i+m.NodeBatchSize, len(nodeAddresses))]

		calls := make([]abis.Multicall3Call3, 0, len(batch))
		for _, node := range batch {
			calls = append(calls, abis.Multicall3Call3{
				Target:       m.rocketNodeDistributorFactoryAddress,
				AllowFailure: false,
				CallData:     rocketNodeDistributorFactory.PackGetProxyAddress(node),
			})
		}

		// Run the multicall call
		mcPacked := multicall3.PackAggregate3(calls)
		mcResponsePacked, err := m.multicallInstance.CallRaw(opts, mcPacked)
		if err != nil {
			return nil, fmt.Errorf("failed to call multicall: %w", err)
		}
		mcResponse, err := multicall3.UnpackAggregate3(mcResponsePacked)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack multicall response: %w", err)
		}

		// Unpack the inner responses.
		for j, call := range mcResponse {
			if !call.Success {
				return nil, errors.New("multicall element failed")
			}
			rawResult := call.ReturnData
			distributorAddress, err := rocketNodeDistributorFactory.UnpackGetProxyAddress(rawResult)
			if err != nil {
				return nil, fmt.Errorf("failed to unpack distributor address: %w", err)
			}
			// batch is 0-index, but the response will be in the same order as the calls
			// Therefor, batch[j] is the address of the node for each result
			// Since this is the first datum for each node, create a new NodeInfo struct.
			nodeMap[batch[j]] = &NodeInfo{
				FeeDistributor: distributorAddress,
			}
		}
	}

	// In batches of 500, populate InSmoothingPool
	for i := 0; i < len(nodeAddresses); i += m.NodeBatchSize {
		batch := nodeAddresses[i:min(i+m.NodeBatchSize, len(nodeAddresses))]

		calls := make([]abis.Multicall3Call3, 0, len(batch))
		for _, node := range batch {
			calls = append(calls, abis.Multicall3Call3{
				Target:       m.rocketNodeManagerAddress,
				AllowFailure: false,
				CallData:     rocketNodeManager.PackGetSmoothingPoolRegistrationState(node),
			})
		}

		// Run the multicall call
		mcPacked := multicall3.PackAggregate3(calls)
		mcResponsePacked, err := m.multicallInstance.CallRaw(opts, mcPacked)
		if err != nil {
			return nil, fmt.Errorf("failed to call multicall: %w", err)
		}
		mcResponse, err := multicall3.UnpackAggregate3(mcResponsePacked)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack multicall response: %w", err)
		}

		// Unpack the inner responses.
		for j, call := range mcResponse {
			if !call.Success {
				return nil, errors.New("multicall element failed")
			}
			rawResult := call.ReturnData
			smoothingPoolRegistrationState, err := rocketNodeManager.UnpackGetSmoothingPoolRegistrationState(rawResult)
			if err != nil {
				return nil, fmt.Errorf("failed to unpack smoothing pool registration state: %w", err)
			}
			nodeMap[batch[j]].InSmoothingPool = smoothingPoolRegistrationState
		}
	}
	return nodeMap, nil
}

func (m *Multicall) GetAllMinipools(nodes map[common.Address]*NodeInfo, opts *bind.CallOpts) (map[common.Address][]rptypes.ValidatorPubkey, error) {
	type nodeInfo struct {
		address       common.Address
		minipoolCount *big.Int
		minipools     []common.Address
		validators    []rptypes.ValidatorPubkey
	}

	// Create a nodeInfo slice from the nodeMap
	nodeInfos := make([]*nodeInfo, 0, len(nodes))
	for addr := range nodes {
		nodeInfos = append(nodeInfos, &nodeInfo{
			address: addr,
		})
	}

	// Working 500 nodes at a time, get the minipool count
	for i := 0; i < len(nodeInfos); i += m.NodeBatchSize {
		batch := nodeInfos[i:min(i+m.NodeBatchSize, len(nodeInfos))]

		calls := make([]abis.Multicall3Call3, 0, len(batch))
		for _, node := range batch {
			calls = append(calls, abis.Multicall3Call3{
				Target:       m.rocketMinipoolManagerAddress,
				AllowFailure: false,
				CallData:     rocketMinipoolManager.PackGetNodeMinipoolCount(node.address),
			})
		}

		mcPacked := multicall3.PackAggregate3(calls)
		mcResponsePacked, err := m.multicallInstance.CallRaw(opts, mcPacked)
		if err != nil {
			return nil, fmt.Errorf("failed to call multicall: %w", err)
		}
		mcResponse, err := multicall3.UnpackAggregate3(mcResponsePacked)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack multicall response: %w", err)
		}

		for j, call := range mcResponse {
			if !call.Success {
				return nil, errors.New("multicall element failed")
			}
			rawResult := call.ReturnData
			minipoolCount, err := rocketMinipoolManager.UnpackGetNodeMinipoolCount(rawResult)
			if err != nil {
				return nil, fmt.Errorf("failed to unpack minipool count: %w", err)
			}
			batch[j].minipoolCount = minipoolCount
		}
	}

	// Working 500 nodes at a time, get the minipool addresses
	for i := 0; i < len(nodeInfos); i += m.NodeBatchSize {
		batch := nodeInfos[i:min(i+m.NodeBatchSize, len(nodeInfos))]

		calls := make([]abis.Multicall3Call3, 0, len(batch))
		for _, node := range batch {
			for j := big.NewInt(0); j.Cmp(node.minipoolCount) < 0; j.Add(j, big.NewInt(1)) {
				calls = append(calls, abis.Multicall3Call3{
					Target:       m.rocketMinipoolManagerAddress,
					AllowFailure: false,
					CallData:     rocketMinipoolManager.PackGetNodeMinipoolAt(node.address, big.NewInt(0).Set(j)),
				})
			}
		}

		mcPacked := multicall3.PackAggregate3(calls)
		mcResponsePacked, err := m.multicallInstance.CallRaw(opts, mcPacked)
		if err != nil {
			return nil, fmt.Errorf("failed to call multicall: %w", err)
		}
		mcResponse, err := multicall3.UnpackAggregate3(mcResponsePacked)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack multicall response: %w", err)
		}

		resultIndex := 0
		for j := range batch {
			for range batch[j].minipoolCount.Int64() {
				rawResult := mcResponse[resultIndex].ReturnData
				minipoolAddress, err := rocketMinipoolManager.UnpackGetNodeMinipoolAt(rawResult)
				if err != nil {
					return nil, fmt.Errorf("failed to unpack minipool address: %w", err)
				}
				batch[j].minipools = append(batch[j].minipools, minipoolAddress)
				resultIndex++
			}
		}
	}

	// Working 500 nodes at a time, get the minipool pubkeys
	for i := 0; i < len(nodeInfos); i += m.NodeBatchSize {
		batch := nodeInfos[i:min(i+m.NodeBatchSize, len(nodeInfos))]

		calls := make([]abis.Multicall3Call3, 0, len(batch))
		for _, node := range batch {
			for _, minipool := range node.minipools {
				calls = append(calls, abis.Multicall3Call3{
					Target:       m.rocketMinipoolManagerAddress,
					AllowFailure: false,
					CallData:     rocketMinipoolManager.PackGetMinipoolPubkey(minipool),
				})
			}
		}

		mcPacked := multicall3.PackAggregate3(calls)
		mcResponsePacked, err := m.multicallInstance.CallRaw(opts, mcPacked)
		if err != nil {
			return nil, fmt.Errorf("failed to call multicall: %w", err)
		}
		mcResponse, err := multicall3.UnpackAggregate3(mcResponsePacked)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack multicall response: %w", err)
		}

		resultIndex := 0
		for j := range batch {
			for range batch[j].minipoolCount.Int64() {
				rawResult := mcResponse[resultIndex].ReturnData
				minipoolPubkey, err := rocketMinipoolManager.UnpackGetMinipoolPubkey(rawResult)
				if err != nil {
					return nil, fmt.Errorf("failed to unpack minipool pubkey: %w", err)
				}
				batch[j].validators = append(batch[j].validators, rptypes.ValidatorPubkey(minipoolPubkey))
				resultIndex++
			}
		}
	}

	// Finally, create an address->[]pubkey map to return
	out := make(map[common.Address][]rptypes.ValidatorPubkey, len(nodeInfos))
	for _, nodeInfo := range nodeInfos {
		out[nodeInfo.address] = nodeInfo.validators
	}

	return out, nil
}

func (m *Multicall) GetAllOdaoNodes(opts *bind.CallOpts) ([]common.Address, error) {
	// There aren't enough odao nodes to justify custom multicall
	// Just use RP's exisitng functionality
	memberCountCallPacked := rocketDaoNodeTrusted.PackGetMemberCount()
	memberCountResponsePacked, err := m.rocketDaoNodeTrustedInstance.CallRaw(opts, memberCountCallPacked)
	if err != nil {
		return nil, fmt.Errorf("failed to call multicall: %w", err)
	}
	memberCount, err := rocketDaoNodeTrusted.UnpackGetMemberCount(memberCountResponsePacked)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack member count: %w", err)
	}
	out := make([]common.Address, 0, memberCount.Int64())
	for i := big.NewInt(0); i.Cmp(memberCount) < 0; i.Add(i, big.NewInt(1)) {
		getMemberCallPacked := rocketDaoNodeTrusted.PackGetMemberAt(i)
		getMemberResponsePacked, err := m.rocketDaoNodeTrustedInstance.CallRaw(opts, getMemberCallPacked)
		if err != nil {
			return nil, fmt.Errorf("failed to call getMemberAt: %w", err)
		}
		memberAddress, err := rocketDaoNodeTrusted.UnpackGetMemberAt(getMemberResponsePacked)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack member address: %w", err)
		}
		out = append(out, memberAddress)
	}

	return out, nil
}

func (m *Multicall) GetREthAddress() common.Address {
	return m.rocketTokenREthAddress
}

func (m *Multicall) GetSmoothingPoolAddress() common.Address {
	return m.rocketSmoothingPoolAddress
}
