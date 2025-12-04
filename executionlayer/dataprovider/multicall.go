package dataprovider

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"slices"

	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer/dataprovider/abis"
	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer/stakewise"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	rptypes "github.com/rocket-pool/smartnode/bindings/types"
)

type Multicall struct {
	NodeBatchSize        int
	SWVaultsRegistryAddr string

	client                       bind.ContractBackend
	rocketStorageAddress         common.Address
	multicallInstance            *bind.BoundContract
	rocketDaoNodeTrustedInstance *bind.BoundContract

	rocketNodeManagerAddress            common.Address
	rocketNodeDistributorFactoryAddress common.Address
	rocketMinipoolManagerAddress        common.Address
	rocketTokenREthAddress              common.Address
	rocketSmoothingPoolAddress          common.Address

	// Checkers for vaults and mev escrow
	vaultsChecker *stakewise.VaultsChecker
}

// Call is the generic type that represents a unresolved multicall3
// query. It can be de-genericified with getCallRecord().
type call[T any] struct {
	ContractAddress common.Address
	PackedCall      []byte
	Unpacker        func([]byte) (T, error)
	Destination     *T
}

func (c call[T]) getCallRecord() callRecord {
	return callRecord{
		multicall3Call3: abis.Multicall3Call3{
			Target:       c.ContractAddress,
			AllowFailure: false,
			CallData:     c.PackedCall,
		},
		unpack: func(rawResult []byte) error {
			var err error
			*c.Destination, err = c.Unpacker(rawResult)
			return err
		},
	}
}

// callRecord is a wrapper around a Multicall3Call3 that has an unpacker
// function. Executing a callRecord will query the client and unpack the result.
type callRecord struct {
	multicall3Call3 abis.Multicall3Call3
	unpack          func([]byte) error
}

type callRecords []callRecord

func (calls callRecords) execute(multicallInstance *bind.BoundContract, opts *bind.CallOpts) error {
	mcCalls := make([]abis.Multicall3Call3, 0, len(calls))
	for _, call := range calls {
		mcCalls = append(mcCalls, call.multicall3Call3)
	}
	mcPacked := multicall3.PackAggregate3(mcCalls)
	mcResponsePacked, err := multicallInstance.CallRaw(opts, mcPacked)
	if err != nil {
		return fmt.Errorf("failed to call multicall: %w", err)
	}
	mcResponse, err := multicall3.UnpackAggregate3(mcResponsePacked)
	if err != nil {
		return fmt.Errorf("failed to unpack multicall response: %w", err)
	}

	if len(mcResponse) != len(calls) {
		return fmt.Errorf("expected %d responses, got %d", len(calls), len(mcResponse))
	}

	for i := range calls {
		if !mcResponse[i].Success {
			return fmt.Errorf("multicall element failed")
		}
	}

	for i := range calls {
		rawResult := mcResponse[i].ReturnData
		err := calls[i].unpack(rawResult)
		if err != nil {
			return fmt.Errorf("failed to unpack result: %w", err)
		}
	}

	return nil
}

func (calls callRecords) executeBatched(multicallInstance *bind.BoundContract, opts *bind.CallOpts, batchSize int) error {
	for batch := range slices.Chunk(calls, batchSize) {
		if err := batch.execute(multicallInstance, opts); err != nil {
			return fmt.Errorf("failed to execute batch: %w", err)
		}
	}
	return nil
}

var multicall3 *abis.Multicall3
var rocketStorage *abis.RocketStorage
var rocketNodeManager *abis.RocketNodeManager
var rocketNodeDistributorFactory *abis.RocketNodeDistributorFactory
var rocketMinipoolManager *abis.RocketMinipoolManager
var rocketDaoNodeTrusted *abis.RocketDaoNodeTrusted
var eip1271 *abis.EIP1271

func init() {
	multicall3 = abis.NewMulticall3()
	rocketStorage = abis.NewRocketStorage()
	rocketNodeManager = abis.NewRocketNodeManager()
	rocketNodeDistributorFactory = abis.NewRocketNodeDistributorFactory()
	rocketMinipoolManager = abis.NewRocketMinipoolManager()
	rocketDaoNodeTrusted = abis.NewRocketDaoNodeTrusted()
	eip1271 = abis.NewEIP1271()
}

func NewMulticall(ctx context.Context, client bind.ContractBackend,
	rocketStorageAddress common.Address,
	contractAddress common.Address,
) (*Multicall, error) {

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

	out := &Multicall{
		client:               client,
		rocketStorageAddress: rocketStorageAddress,
		multicallInstance:    multicall3.Instance(client, contractAddress),
		NodeBatchSize:        500,
	}

	err = out.RefreshAddresses(&opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (m *Multicall) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return m.client.HeaderByNumber(ctx, number)
}

func (m *Multicall) RefreshAddresses(opts *bind.CallOpts) error {
	// hashes precomputed from https://emn178.github.io/online-tools/keccak_256.html

	var rocketDaoNodeTrustedAddress common.Address

	getAddressCalls := make(callRecords, 0)

	for _, contract := range []struct {
		Key         string
		Destination *common.Address
	}{
		{
			// keccak256 of "contract.addressrocketNodeManager"
			Key:         "0xaf00be55c9fb8f543c04e0aa0d70351b880c1bfafffd15b60065a4a50c85ec94",
			Destination: &(m.rocketNodeManagerAddress),
		},
		{
			// keccak256 of "contract.addressrocketNodeDistributorFactory"
			Key:         "0xea051094896ef3b09ab1b794ad5ea695a5ff3906f74a9328e2c16db69d0f3123",
			Destination: &(m.rocketNodeDistributorFactoryAddress),
		},
		{
			// keccak256 of "contract.addressrocketMinipoolManager"
			Key:         "0xe9dfec9339b94a131861a58f1bb4ac4c1ce55c7ffe8550e0b6ebcfde87bb012f",
			Destination: &(m.rocketMinipoolManagerAddress),
		},
		{
			// keccak256 of "contract.addressrocketDAONodeTrusted"
			Key:         "0x9a354e1bb2e38ca826db7a8d061cfb0ed7dbd83d241a2cbe4fd5218f9bb4333f",
			Destination: &rocketDaoNodeTrustedAddress,
		},
		{
			// keccak256 of "contract.addressrocketTokenRETH"
			Key:         "0xe3744443225bff7cc22028be036b80de58057d65a3fdca0a3df329f525e31ccc",
			Destination: &(m.rocketTokenREthAddress),
		},
		{
			// keccak256 of "contract.addressrocketSmoothingPool"
			Key:         "0x822231720aef9b264db1d9ca053137498f759c28b243f45c44db1d39d6bce46e",
			Destination: &(m.rocketSmoothingPoolAddress),
		},
	} {
		getAddressCalls = append(getAddressCalls, call[common.Address]{
			ContractAddress: m.rocketStorageAddress,
			PackedCall:      rocketStorage.PackGetAddress(common.HexToHash(contract.Key)),
			Unpacker:        rocketStorage.UnpackGetAddress,
			Destination:     contract.Destination,
		}.getCallRecord())
	}

	if err := getAddressCalls.execute(m.multicallInstance, opts); err != nil {
		return fmt.Errorf("failed to get address calls: %w", err)
	}

	m.rocketDaoNodeTrustedInstance = rocketDaoNodeTrusted.Instance(m.client, rocketDaoNodeTrustedAddress)

	return nil
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

	// populate FeeDistributor and InSmoothingPool
	records := make(callRecords, 0, len(nodeAddresses)*2)
	for _, nodeAddress := range nodeAddresses {
		nodeInfo := &NodeInfo{}
		nodeMap[nodeAddress] = nodeInfo
		records = append(records, call[common.Address]{
			ContractAddress: m.rocketNodeDistributorFactoryAddress,
			PackedCall:      rocketNodeDistributorFactory.PackGetProxyAddress(nodeAddress),
			Unpacker:        rocketNodeDistributorFactory.UnpackGetProxyAddress,
			Destination:     &(nodeInfo.FeeDistributor),
		}.getCallRecord())
		records = append(records, call[bool]{
			ContractAddress: m.rocketNodeManagerAddress,
			PackedCall:      rocketNodeManager.PackGetSmoothingPoolRegistrationState(nodeAddress),
			Unpacker:        rocketNodeManager.UnpackGetSmoothingPoolRegistrationState,
			Destination:     &(nodeInfo.InSmoothingPool),
		}.getCallRecord())
	}

	if err := records.executeBatched(m.multicallInstance, opts, m.NodeBatchSize); err != nil {
		return nil, fmt.Errorf("failed to execute fee distributor and smoothing pool batch: %w", err)
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

	// Get the minipool counts
	records := make(callRecords, 0, len(nodeInfos))
	for _, node := range nodeInfos {
		records = append(records, call[*big.Int]{
			ContractAddress: m.rocketMinipoolManagerAddress,
			PackedCall:      rocketMinipoolManager.PackGetNodeMinipoolCount(node.address),
			Unpacker:        rocketMinipoolManager.UnpackGetNodeMinipoolCount,
			Destination:     &(node.minipoolCount),
		}.getCallRecord())
	}
	if err := records.executeBatched(m.multicallInstance, opts, m.NodeBatchSize); err != nil {
		return nil, fmt.Errorf("failed to execute minipool count batch: %w", err)
	}
	records = records[:0]

	// Get the minipool addresses
	for _, node := range nodeInfos {
		node.minipools = make([]common.Address, node.minipoolCount.Int64())
		for j := int64(0); j < node.minipoolCount.Int64(); j++ {
			records = append(records, call[common.Address]{
				ContractAddress: m.rocketMinipoolManagerAddress,
				PackedCall:      rocketMinipoolManager.PackGetNodeMinipoolAt(node.address, big.NewInt(j)),
				Unpacker:        rocketMinipoolManager.UnpackGetNodeMinipoolAt,
				Destination:     &(node.minipools[j]),
			}.getCallRecord())
		}
	}
	if err := records.executeBatched(m.multicallInstance, opts, m.NodeBatchSize); err != nil {
		return nil, fmt.Errorf("failed to execute minipool address batch: %w", err)
	}
	records = records[:0]

	// Get the minipool pubkeys
	for _, node := range nodeInfos {
		node.validators = make([]rptypes.ValidatorPubkey, node.minipoolCount.Int64())
		for j := int64(0); j < node.minipoolCount.Int64(); j++ {
			records = append(records, call[rptypes.ValidatorPubkey]{
				ContractAddress: m.rocketMinipoolManagerAddress,
				PackedCall:      rocketMinipoolManager.PackGetMinipoolPubkey(node.minipools[j]),
				Unpacker: func(rawResult []byte) (rptypes.ValidatorPubkey, error) {
					unpacked, err := rocketMinipoolManager.UnpackGetMinipoolPubkey(rawResult)
					if err != nil {
						return rptypes.ValidatorPubkey{}, fmt.Errorf("failed to unpack minipool pubkey: %w", err)
					}
					return rptypes.BytesToValidatorPubkey(unpacked), nil
				},
				Destination: &(node.validators[j]),
			}.getCallRecord())
		}
	}

	if err := records.executeBatched(m.multicallInstance, opts, m.NodeBatchSize); err != nil {
		return nil, fmt.Errorf("failed to execute minipool pubkey batch: %w", err)
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

func (m *Multicall) StakewiseFeeRecipient(opts *bind.CallOpts, address common.Address) (*common.Address, error) {
	if m.SWVaultsRegistryAddr == "" {
		return nil, errors.New("SWVaultsRegistryAddr is not set")
	}
	if m.vaultsChecker == nil {
		m.vaultsChecker = stakewise.NewVaultsChecker(m.client, common.HexToAddress(m.SWVaultsRegistryAddr))
	}

	return m.vaultsChecker.IsVault(opts, address)
}

var ErrNoData = errors.New("no data were returned from the EVM, did you pass the correct smart contract wallet address?")
var ErrBadData = errors.New("the evm returned data with an unexpected length, did you pass the correct smart contract wallet address?")
var ErrInternal = errors.New("an internal error occurred, please contact the maintainers")

func (m *Multicall) ValidateEIP1271(opts *bind.CallOpts, dataHash common.Hash, signature []byte, address common.Address) (bool, error) {
	// Encode the function call
	encodedData := eip1271.PackIsValidSignature(dataHash, signature)

	// Make the contract call
	data, err := m.client.CallContract(opts.Context, ethereum.CallMsg{
		To:   &address,
		Data: encodedData,
	}, nil)
	if err != nil {
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
