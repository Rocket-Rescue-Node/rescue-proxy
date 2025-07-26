// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// NodeDetails is an auto generated low-level Go binding around an user-defined struct.
type NodeDetails struct {
	Exists                           bool
	RegistrationTime                 *big.Int
	TimezoneLocation                 string
	FeeDistributorInitialised        bool
	FeeDistributorAddress            common.Address
	RewardNetwork                    *big.Int
	RplStake                         *big.Int
	EffectiveRPLStake                *big.Int
	MinimumRPLStake                  *big.Int
	MaximumRPLStake                  *big.Int
	EthMatched                       *big.Int
	EthMatchedLimit                  *big.Int
	MinipoolCount                    *big.Int
	BalanceETH                       *big.Int
	BalanceRETH                      *big.Int
	BalanceRPL                       *big.Int
	BalanceOldRPL                    *big.Int
	DepositCreditBalance             *big.Int
	DistributorBalanceUserETH        *big.Int
	DistributorBalanceNodeETH        *big.Int
	WithdrawalAddress                common.Address
	PendingWithdrawalAddress         common.Address
	SmoothingPoolRegistrationState   bool
	SmoothingPoolRegistrationChanged *big.Int
	NodeAddress                      common.Address
}

// RocketNodeManagerInterfaceTimezoneCount is an auto generated low-level Go binding around an user-defined struct.
type RocketNodeManagerInterfaceTimezoneCount struct {
	Timezone string
	Count    *big.Int
}

// RocketNodeManagerMetaData contains all meta data concerning the RocketNodeManager contract.
var RocketNodeManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRocketStorageInterface\",\"name\":\"_rocketStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawalAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"NodeRPLWithdrawalAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"NodeRPLWithdrawalAddressUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"NodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"}],\"name\":\"NodeRewardNetworkChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"NodeSmoothingPoolStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"NodeTimezoneLocationSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"confirmRPLWithdrawalAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getAverageNodeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getFeeDistributorInitialised\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getNodeAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNodeAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getNodeCountPerTimezone\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"timezone\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"internalType\":\"structRocketNodeManagerInterface.TimezoneCount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"registrationTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"timezoneLocation\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"feeDistributorInitialised\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"feeDistributorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardNetwork\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rplStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effectiveRPLStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumRPLStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumRPLStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethMatched\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethMatchedLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minipoolCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceRETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceRPL\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceOldRPL\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositCreditBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributorBalanceUserETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributorBalanceNodeETH\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"withdrawalAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendingWithdrawalAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"smoothingPoolRegistrationState\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"smoothingPoolRegistrationChanged\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"internalType\":\"structNodeDetails\",\"name\":\"nodeDetails\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodePendingRPLWithdrawalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodePendingWithdrawalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeRPLWithdrawalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeRPLWithdrawalAddressIsSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeRegistrationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeTimezoneLocation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeWithdrawalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getRewardNetwork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getSmoothingPoolRegisteredNodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getSmoothingPoolRegistrationChanged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getSmoothingPoolRegistrationState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialiseFeeDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_timezoneLocation\",\"type\":\"string\"}],\"name\":\"registerNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newRPLWithdrawalAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_confirm\",\"type\":\"bool\"}],\"name\":\"setRPLWithdrawalAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_network\",\"type\":\"uint256\"}],\"name\":\"setRewardNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setSmoothingPoolRegistrationState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_timezoneLocation\",\"type\":\"string\"}],\"name\":\"setTimezoneLocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"unsetRPLWithdrawalAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "RocketNodeManager",
}

// RocketNodeManager is an auto generated Go binding around an Ethereum contract.
type RocketNodeManager struct {
	abi abi.ABI
}

// NewRocketNodeManager creates a new instance of RocketNodeManager.
func NewRocketNodeManager() *RocketNodeManager {
	parsed, err := RocketNodeManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RocketNodeManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RocketNodeManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _rocketStorageAddress) returns()
func (rocketNodeManager *RocketNodeManager) PackConstructor(_rocketStorageAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("", _rocketStorageAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackConfirmRPLWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a643648.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmRPLWithdrawalAddress(address _nodeAddress) returns()
func (rocketNodeManager *RocketNodeManager) PackConfirmRPLWithdrawalAddress(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("confirmRPLWithdrawalAddress", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmRPLWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a643648.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmRPLWithdrawalAddress(address _nodeAddress) returns()
func (rocketNodeManager *RocketNodeManager) TryPackConfirmRPLWithdrawalAddress(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("confirmRPLWithdrawalAddress", nodeAddress)
}

// PackGetAverageNodeFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x414dd1d2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAverageNodeFee(address _nodeAddress) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) PackGetAverageNodeFee(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getAverageNodeFee", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAverageNodeFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x414dd1d2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAverageNodeFee(address _nodeAddress) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) TryPackGetAverageNodeFee(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getAverageNodeFee", nodeAddress)
}

// UnpackGetAverageNodeFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x414dd1d2.
//
// Solidity: function getAverageNodeFee(address _nodeAddress) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) UnpackGetAverageNodeFee(data []byte) (*big.Int, error) {
	out, err := rocketNodeManager.abi.Unpack("getAverageNodeFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetFeeDistributorInitialised is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x927ece4f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getFeeDistributorInitialised(address _nodeAddress) view returns(bool)
func (rocketNodeManager *RocketNodeManager) PackGetFeeDistributorInitialised(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getFeeDistributorInitialised", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetFeeDistributorInitialised is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x927ece4f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getFeeDistributorInitialised(address _nodeAddress) view returns(bool)
func (rocketNodeManager *RocketNodeManager) TryPackGetFeeDistributorInitialised(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getFeeDistributorInitialised", nodeAddress)
}

// UnpackGetFeeDistributorInitialised is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x927ece4f.
//
// Solidity: function getFeeDistributorInitialised(address _nodeAddress) view returns(bool)
func (rocketNodeManager *RocketNodeManager) UnpackGetFeeDistributorInitialised(data []byte) (bool, error) {
	out, err := rocketNodeManager.abi.Unpack("getFeeDistributorInitialised", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGetNodeAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d7f21d0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeAddresses(uint256 _offset, uint256 _limit) view returns(address[])
func (rocketNodeManager *RocketNodeManager) PackGetNodeAddresses(offset *big.Int, limit *big.Int) []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodeAddresses", offset, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d7f21d0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeAddresses(uint256 _offset, uint256 _limit) view returns(address[])
func (rocketNodeManager *RocketNodeManager) TryPackGetNodeAddresses(offset *big.Int, limit *big.Int) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodeAddresses", offset, limit)
}

// UnpackGetNodeAddresses is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2d7f21d0.
//
// Solidity: function getNodeAddresses(uint256 _offset, uint256 _limit) view returns(address[])
func (rocketNodeManager *RocketNodeManager) UnpackGetNodeAddresses(data []byte) ([]common.Address, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodeAddresses", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetNodeAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba75d806.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeAt(uint256 _index) view returns(address)
func (rocketNodeManager *RocketNodeManager) PackGetNodeAt(index *big.Int) []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodeAt", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba75d806.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeAt(uint256 _index) view returns(address)
func (rocketNodeManager *RocketNodeManager) TryPackGetNodeAt(index *big.Int) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodeAt", index)
}

// UnpackGetNodeAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xba75d806.
//
// Solidity: function getNodeAt(uint256 _index) view returns(address)
func (rocketNodeManager *RocketNodeManager) UnpackGetNodeAt(data []byte) (common.Address, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodeAt", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNodeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x39bf397e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeCount() view returns(uint256)
func (rocketNodeManager *RocketNodeManager) PackGetNodeCount() []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x39bf397e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeCount() view returns(uint256)
func (rocketNodeManager *RocketNodeManager) TryPackGetNodeCount() ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodeCount")
}

// UnpackGetNodeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (rocketNodeManager *RocketNodeManager) UnpackGetNodeCount(data []byte) (*big.Int, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodeCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetNodeCountPerTimezone is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x29554540.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeCountPerTimezone(uint256 _offset, uint256 _limit) view returns((string,uint256)[])
func (rocketNodeManager *RocketNodeManager) PackGetNodeCountPerTimezone(offset *big.Int, limit *big.Int) []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodeCountPerTimezone", offset, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeCountPerTimezone is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x29554540.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeCountPerTimezone(uint256 _offset, uint256 _limit) view returns((string,uint256)[])
func (rocketNodeManager *RocketNodeManager) TryPackGetNodeCountPerTimezone(offset *big.Int, limit *big.Int) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodeCountPerTimezone", offset, limit)
}

// UnpackGetNodeCountPerTimezone is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x29554540.
//
// Solidity: function getNodeCountPerTimezone(uint256 _offset, uint256 _limit) view returns((string,uint256)[])
func (rocketNodeManager *RocketNodeManager) UnpackGetNodeCountPerTimezone(data []byte) ([]RocketNodeManagerInterfaceTimezoneCount, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodeCountPerTimezone", data)
	if err != nil {
		return *new([]RocketNodeManagerInterfaceTimezoneCount), err
	}
	out0 := *abi.ConvertType(out[0], new([]RocketNodeManagerInterfaceTimezoneCount)).(*[]RocketNodeManagerInterfaceTimezoneCount)
	return out0, nil
}

// PackGetNodeDetails is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbafb3581.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeDetails(address _nodeAddress) view returns((bool,uint256,string,bool,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,bool,uint256,address) nodeDetails)
func (rocketNodeManager *RocketNodeManager) PackGetNodeDetails(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodeDetails", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeDetails is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbafb3581.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeDetails(address _nodeAddress) view returns((bool,uint256,string,bool,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,bool,uint256,address) nodeDetails)
func (rocketNodeManager *RocketNodeManager) TryPackGetNodeDetails(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodeDetails", nodeAddress)
}

// UnpackGetNodeDetails is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbafb3581.
//
// Solidity: function getNodeDetails(address _nodeAddress) view returns((bool,uint256,string,bool,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,bool,uint256,address) nodeDetails)
func (rocketNodeManager *RocketNodeManager) UnpackGetNodeDetails(data []byte) (NodeDetails, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodeDetails", data)
	if err != nil {
		return *new(NodeDetails), err
	}
	out0 := *abi.ConvertType(out[0], new(NodeDetails)).(*NodeDetails)
	return out0, nil
}

// PackGetNodeExists is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x65d4176f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeExists(address _nodeAddress) view returns(bool)
func (rocketNodeManager *RocketNodeManager) PackGetNodeExists(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodeExists", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeExists is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x65d4176f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeExists(address _nodeAddress) view returns(bool)
func (rocketNodeManager *RocketNodeManager) TryPackGetNodeExists(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodeExists", nodeAddress)
}

// UnpackGetNodeExists is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x65d4176f.
//
// Solidity: function getNodeExists(address _nodeAddress) view returns(bool)
func (rocketNodeManager *RocketNodeManager) UnpackGetNodeExists(data []byte) (bool, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodeExists", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGetNodePendingRPLWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1ac3c0a8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodePendingRPLWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketNodeManager *RocketNodeManager) PackGetNodePendingRPLWithdrawalAddress(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodePendingRPLWithdrawalAddress", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodePendingRPLWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1ac3c0a8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodePendingRPLWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketNodeManager *RocketNodeManager) TryPackGetNodePendingRPLWithdrawalAddress(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodePendingRPLWithdrawalAddress", nodeAddress)
}

// UnpackGetNodePendingRPLWithdrawalAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1ac3c0a8.
//
// Solidity: function getNodePendingRPLWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketNodeManager *RocketNodeManager) UnpackGetNodePendingRPLWithdrawalAddress(data []byte) (common.Address, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodePendingRPLWithdrawalAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNodePendingWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfd412513.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodePendingWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketNodeManager *RocketNodeManager) PackGetNodePendingWithdrawalAddress(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodePendingWithdrawalAddress", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodePendingWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfd412513.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodePendingWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketNodeManager *RocketNodeManager) TryPackGetNodePendingWithdrawalAddress(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodePendingWithdrawalAddress", nodeAddress)
}

// UnpackGetNodePendingWithdrawalAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfd412513.
//
// Solidity: function getNodePendingWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketNodeManager *RocketNodeManager) UnpackGetNodePendingWithdrawalAddress(data []byte) (common.Address, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodePendingWithdrawalAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNodeRPLWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb71f0c7c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeRPLWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketNodeManager *RocketNodeManager) PackGetNodeRPLWithdrawalAddress(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodeRPLWithdrawalAddress", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeRPLWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb71f0c7c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeRPLWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketNodeManager *RocketNodeManager) TryPackGetNodeRPLWithdrawalAddress(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodeRPLWithdrawalAddress", nodeAddress)
}

// UnpackGetNodeRPLWithdrawalAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb71f0c7c.
//
// Solidity: function getNodeRPLWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketNodeManager *RocketNodeManager) UnpackGetNodeRPLWithdrawalAddress(data []byte) (common.Address, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodeRPLWithdrawalAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNodeRPLWithdrawalAddressIsSet is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe667d828.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeRPLWithdrawalAddressIsSet(address _nodeAddress) view returns(bool)
func (rocketNodeManager *RocketNodeManager) PackGetNodeRPLWithdrawalAddressIsSet(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodeRPLWithdrawalAddressIsSet", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeRPLWithdrawalAddressIsSet is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe667d828.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeRPLWithdrawalAddressIsSet(address _nodeAddress) view returns(bool)
func (rocketNodeManager *RocketNodeManager) TryPackGetNodeRPLWithdrawalAddressIsSet(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodeRPLWithdrawalAddressIsSet", nodeAddress)
}

// UnpackGetNodeRPLWithdrawalAddressIsSet is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe667d828.
//
// Solidity: function getNodeRPLWithdrawalAddressIsSet(address _nodeAddress) view returns(bool)
func (rocketNodeManager *RocketNodeManager) UnpackGetNodeRPLWithdrawalAddressIsSet(data []byte) (bool, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodeRPLWithdrawalAddressIsSet", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGetNodeRegistrationTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02d8a732.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeRegistrationTime(address _nodeAddress) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) PackGetNodeRegistrationTime(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodeRegistrationTime", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeRegistrationTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02d8a732.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeRegistrationTime(address _nodeAddress) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) TryPackGetNodeRegistrationTime(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodeRegistrationTime", nodeAddress)
}

// UnpackGetNodeRegistrationTime is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x02d8a732.
//
// Solidity: function getNodeRegistrationTime(address _nodeAddress) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) UnpackGetNodeRegistrationTime(data []byte) (*big.Int, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodeRegistrationTime", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetNodeTimezoneLocation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb018f026.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeTimezoneLocation(address _nodeAddress) view returns(string)
func (rocketNodeManager *RocketNodeManager) PackGetNodeTimezoneLocation(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodeTimezoneLocation", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeTimezoneLocation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb018f026.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeTimezoneLocation(address _nodeAddress) view returns(string)
func (rocketNodeManager *RocketNodeManager) TryPackGetNodeTimezoneLocation(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodeTimezoneLocation", nodeAddress)
}

// UnpackGetNodeTimezoneLocation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb018f026.
//
// Solidity: function getNodeTimezoneLocation(address _nodeAddress) view returns(string)
func (rocketNodeManager *RocketNodeManager) UnpackGetNodeTimezoneLocation(data []byte) (string, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodeTimezoneLocation", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackGetNodeWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b49ff62.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketNodeManager *RocketNodeManager) PackGetNodeWithdrawalAddress(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getNodeWithdrawalAddress", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b49ff62.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketNodeManager *RocketNodeManager) TryPackGetNodeWithdrawalAddress(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getNodeWithdrawalAddress", nodeAddress)
}

// UnpackGetNodeWithdrawalAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5b49ff62.
//
// Solidity: function getNodeWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketNodeManager *RocketNodeManager) UnpackGetNodeWithdrawalAddress(data []byte) (common.Address, error) {
	out, err := rocketNodeManager.abi.Unpack("getNodeWithdrawalAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetRewardNetwork is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43f88981.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRewardNetwork(address _nodeAddress) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) PackGetRewardNetwork(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getRewardNetwork", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRewardNetwork is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43f88981.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRewardNetwork(address _nodeAddress) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) TryPackGetRewardNetwork(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getRewardNetwork", nodeAddress)
}

// UnpackGetRewardNetwork is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x43f88981.
//
// Solidity: function getRewardNetwork(address _nodeAddress) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) UnpackGetRewardNetwork(data []byte) (*big.Int, error) {
	out, err := rocketNodeManager.abi.Unpack("getRewardNetwork", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetSmoothingPoolRegisteredNodeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb715a1aa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSmoothingPoolRegisteredNodeCount(uint256 _offset, uint256 _limit) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) PackGetSmoothingPoolRegisteredNodeCount(offset *big.Int, limit *big.Int) []byte {
	enc, err := rocketNodeManager.abi.Pack("getSmoothingPoolRegisteredNodeCount", offset, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSmoothingPoolRegisteredNodeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb715a1aa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSmoothingPoolRegisteredNodeCount(uint256 _offset, uint256 _limit) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) TryPackGetSmoothingPoolRegisteredNodeCount(offset *big.Int, limit *big.Int) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getSmoothingPoolRegisteredNodeCount", offset, limit)
}

// UnpackGetSmoothingPoolRegisteredNodeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb715a1aa.
//
// Solidity: function getSmoothingPoolRegisteredNodeCount(uint256 _offset, uint256 _limit) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) UnpackGetSmoothingPoolRegisteredNodeCount(data []byte) (*big.Int, error) {
	out, err := rocketNodeManager.abi.Unpack("getSmoothingPoolRegisteredNodeCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetSmoothingPoolRegistrationChanged is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4d99f633.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSmoothingPoolRegistrationChanged(address _nodeAddress) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) PackGetSmoothingPoolRegistrationChanged(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getSmoothingPoolRegistrationChanged", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSmoothingPoolRegistrationChanged is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4d99f633.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSmoothingPoolRegistrationChanged(address _nodeAddress) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) TryPackGetSmoothingPoolRegistrationChanged(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getSmoothingPoolRegistrationChanged", nodeAddress)
}

// UnpackGetSmoothingPoolRegistrationChanged is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4d99f633.
//
// Solidity: function getSmoothingPoolRegistrationChanged(address _nodeAddress) view returns(uint256)
func (rocketNodeManager *RocketNodeManager) UnpackGetSmoothingPoolRegistrationChanged(data []byte) (*big.Int, error) {
	out, err := rocketNodeManager.abi.Unpack("getSmoothingPoolRegistrationChanged", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetSmoothingPoolRegistrationState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4cef9dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSmoothingPoolRegistrationState(address _nodeAddress) view returns(bool)
func (rocketNodeManager *RocketNodeManager) PackGetSmoothingPoolRegistrationState(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("getSmoothingPoolRegistrationState", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSmoothingPoolRegistrationState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4cef9dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSmoothingPoolRegistrationState(address _nodeAddress) view returns(bool)
func (rocketNodeManager *RocketNodeManager) TryPackGetSmoothingPoolRegistrationState(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("getSmoothingPoolRegistrationState", nodeAddress)
}

// UnpackGetSmoothingPoolRegistrationState is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa4cef9dd.
//
// Solidity: function getSmoothingPoolRegistrationState(address _nodeAddress) view returns(bool)
func (rocketNodeManager *RocketNodeManager) UnpackGetSmoothingPoolRegistrationState(data []byte) (bool, error) {
	out, err := rocketNodeManager.abi.Unpack("getSmoothingPoolRegistrationState", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackInitialiseFeeDistributor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x64908a86.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialiseFeeDistributor() returns()
func (rocketNodeManager *RocketNodeManager) PackInitialiseFeeDistributor() []byte {
	enc, err := rocketNodeManager.abi.Pack("initialiseFeeDistributor")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialiseFeeDistributor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x64908a86.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialiseFeeDistributor() returns()
func (rocketNodeManager *RocketNodeManager) TryPackInitialiseFeeDistributor() ([]byte, error) {
	return rocketNodeManager.abi.Pack("initialiseFeeDistributor")
}

// PackRegisterNode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x27c6f43e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registerNode(string _timezoneLocation) returns()
func (rocketNodeManager *RocketNodeManager) PackRegisterNode(timezoneLocation string) []byte {
	enc, err := rocketNodeManager.abi.Pack("registerNode", timezoneLocation)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegisterNode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x27c6f43e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registerNode(string _timezoneLocation) returns()
func (rocketNodeManager *RocketNodeManager) TryPackRegisterNode(timezoneLocation string) ([]byte, error) {
	return rocketNodeManager.abi.Pack("registerNode", timezoneLocation)
}

// PackSetRPLWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5b17b42.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setRPLWithdrawalAddress(address _nodeAddress, address _newRPLWithdrawalAddress, bool _confirm) returns()
func (rocketNodeManager *RocketNodeManager) PackSetRPLWithdrawalAddress(nodeAddress common.Address, newRPLWithdrawalAddress common.Address, confirm bool) []byte {
	enc, err := rocketNodeManager.abi.Pack("setRPLWithdrawalAddress", nodeAddress, newRPLWithdrawalAddress, confirm)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetRPLWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5b17b42.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setRPLWithdrawalAddress(address _nodeAddress, address _newRPLWithdrawalAddress, bool _confirm) returns()
func (rocketNodeManager *RocketNodeManager) TryPackSetRPLWithdrawalAddress(nodeAddress common.Address, newRPLWithdrawalAddress common.Address, confirm bool) ([]byte, error) {
	return rocketNodeManager.abi.Pack("setRPLWithdrawalAddress", nodeAddress, newRPLWithdrawalAddress, confirm)
}

// PackSetRewardNetwork is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd565f276.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setRewardNetwork(address _nodeAddress, uint256 _network) returns()
func (rocketNodeManager *RocketNodeManager) PackSetRewardNetwork(nodeAddress common.Address, network *big.Int) []byte {
	enc, err := rocketNodeManager.abi.Pack("setRewardNetwork", nodeAddress, network)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetRewardNetwork is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd565f276.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setRewardNetwork(address _nodeAddress, uint256 _network) returns()
func (rocketNodeManager *RocketNodeManager) TryPackSetRewardNetwork(nodeAddress common.Address, network *big.Int) ([]byte, error) {
	return rocketNodeManager.abi.Pack("setRewardNetwork", nodeAddress, network)
}

// PackSetSmoothingPoolRegistrationState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x99283f8b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSmoothingPoolRegistrationState(bool _state) returns()
func (rocketNodeManager *RocketNodeManager) PackSetSmoothingPoolRegistrationState(state bool) []byte {
	enc, err := rocketNodeManager.abi.Pack("setSmoothingPoolRegistrationState", state)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSmoothingPoolRegistrationState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x99283f8b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSmoothingPoolRegistrationState(bool _state) returns()
func (rocketNodeManager *RocketNodeManager) TryPackSetSmoothingPoolRegistrationState(state bool) ([]byte, error) {
	return rocketNodeManager.abi.Pack("setSmoothingPoolRegistrationState", state)
}

// PackSetTimezoneLocation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7e6e8b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setTimezoneLocation(string _timezoneLocation) returns()
func (rocketNodeManager *RocketNodeManager) PackSetTimezoneLocation(timezoneLocation string) []byte {
	enc, err := rocketNodeManager.abi.Pack("setTimezoneLocation", timezoneLocation)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetTimezoneLocation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7e6e8b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setTimezoneLocation(string _timezoneLocation) returns()
func (rocketNodeManager *RocketNodeManager) TryPackSetTimezoneLocation(timezoneLocation string) ([]byte, error) {
	return rocketNodeManager.abi.Pack("setTimezoneLocation", timezoneLocation)
}

// PackUnsetRPLWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2a7968eb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unsetRPLWithdrawalAddress(address _nodeAddress) returns()
func (rocketNodeManager *RocketNodeManager) PackUnsetRPLWithdrawalAddress(nodeAddress common.Address) []byte {
	enc, err := rocketNodeManager.abi.Pack("unsetRPLWithdrawalAddress", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnsetRPLWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2a7968eb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unsetRPLWithdrawalAddress(address _nodeAddress) returns()
func (rocketNodeManager *RocketNodeManager) TryPackUnsetRPLWithdrawalAddress(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeManager.abi.Pack("unsetRPLWithdrawalAddress", nodeAddress)
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(uint8)
func (rocketNodeManager *RocketNodeManager) PackVersion() []byte {
	enc, err := rocketNodeManager.abi.Pack("version")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function version() view returns(uint8)
func (rocketNodeManager *RocketNodeManager) TryPackVersion() ([]byte, error) {
	return rocketNodeManager.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (rocketNodeManager *RocketNodeManager) UnpackVersion(data []byte) (uint8, error) {
	out, err := rocketNodeManager.abi.Unpack("version", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// RocketNodeManagerNodeRPLWithdrawalAddressSet represents a NodeRPLWithdrawalAddressSet event raised by the RocketNodeManager contract.
type RocketNodeManagerNodeRPLWithdrawalAddressSet struct {
	Node              common.Address
	WithdrawalAddress common.Address
	Time              *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const RocketNodeManagerNodeRPLWithdrawalAddressSetEventName = "NodeRPLWithdrawalAddressSet"

// ContractEventName returns the user-defined event name.
func (RocketNodeManagerNodeRPLWithdrawalAddressSet) ContractEventName() string {
	return RocketNodeManagerNodeRPLWithdrawalAddressSetEventName
}

// UnpackNodeRPLWithdrawalAddressSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NodeRPLWithdrawalAddressSet(address indexed node, address indexed withdrawalAddress, uint256 time)
func (rocketNodeManager *RocketNodeManager) UnpackNodeRPLWithdrawalAddressSetEvent(log *types.Log) (*RocketNodeManagerNodeRPLWithdrawalAddressSet, error) {
	event := "NodeRPLWithdrawalAddressSet"
	if log.Topics[0] != rocketNodeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketNodeManagerNodeRPLWithdrawalAddressSet)
	if len(log.Data) > 0 {
		if err := rocketNodeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketNodeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// RocketNodeManagerNodeRPLWithdrawalAddressUnset represents a NodeRPLWithdrawalAddressUnset event raised by the RocketNodeManager contract.
type RocketNodeManagerNodeRPLWithdrawalAddressUnset struct {
	Node common.Address
	Time *big.Int
	Raw  *types.Log // Blockchain specific contextual infos
}

const RocketNodeManagerNodeRPLWithdrawalAddressUnsetEventName = "NodeRPLWithdrawalAddressUnset"

// ContractEventName returns the user-defined event name.
func (RocketNodeManagerNodeRPLWithdrawalAddressUnset) ContractEventName() string {
	return RocketNodeManagerNodeRPLWithdrawalAddressUnsetEventName
}

// UnpackNodeRPLWithdrawalAddressUnsetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NodeRPLWithdrawalAddressUnset(address indexed node, uint256 time)
func (rocketNodeManager *RocketNodeManager) UnpackNodeRPLWithdrawalAddressUnsetEvent(log *types.Log) (*RocketNodeManagerNodeRPLWithdrawalAddressUnset, error) {
	event := "NodeRPLWithdrawalAddressUnset"
	if log.Topics[0] != rocketNodeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketNodeManagerNodeRPLWithdrawalAddressUnset)
	if len(log.Data) > 0 {
		if err := rocketNodeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketNodeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// RocketNodeManagerNodeRegistered represents a NodeRegistered event raised by the RocketNodeManager contract.
type RocketNodeManagerNodeRegistered struct {
	Node common.Address
	Time *big.Int
	Raw  *types.Log // Blockchain specific contextual infos
}

const RocketNodeManagerNodeRegisteredEventName = "NodeRegistered"

// ContractEventName returns the user-defined event name.
func (RocketNodeManagerNodeRegistered) ContractEventName() string {
	return RocketNodeManagerNodeRegisteredEventName
}

// UnpackNodeRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NodeRegistered(address indexed node, uint256 time)
func (rocketNodeManager *RocketNodeManager) UnpackNodeRegisteredEvent(log *types.Log) (*RocketNodeManagerNodeRegistered, error) {
	event := "NodeRegistered"
	if log.Topics[0] != rocketNodeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketNodeManagerNodeRegistered)
	if len(log.Data) > 0 {
		if err := rocketNodeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketNodeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// RocketNodeManagerNodeRewardNetworkChanged represents a NodeRewardNetworkChanged event raised by the RocketNodeManager contract.
type RocketNodeManagerNodeRewardNetworkChanged struct {
	Node    common.Address
	Network *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const RocketNodeManagerNodeRewardNetworkChangedEventName = "NodeRewardNetworkChanged"

// ContractEventName returns the user-defined event name.
func (RocketNodeManagerNodeRewardNetworkChanged) ContractEventName() string {
	return RocketNodeManagerNodeRewardNetworkChangedEventName
}

// UnpackNodeRewardNetworkChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NodeRewardNetworkChanged(address indexed node, uint256 network)
func (rocketNodeManager *RocketNodeManager) UnpackNodeRewardNetworkChangedEvent(log *types.Log) (*RocketNodeManagerNodeRewardNetworkChanged, error) {
	event := "NodeRewardNetworkChanged"
	if log.Topics[0] != rocketNodeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketNodeManagerNodeRewardNetworkChanged)
	if len(log.Data) > 0 {
		if err := rocketNodeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketNodeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// RocketNodeManagerNodeSmoothingPoolStateChanged represents a NodeSmoothingPoolStateChanged event raised by the RocketNodeManager contract.
type RocketNodeManagerNodeSmoothingPoolStateChanged struct {
	Node  common.Address
	State bool
	Raw   *types.Log // Blockchain specific contextual infos
}

const RocketNodeManagerNodeSmoothingPoolStateChangedEventName = "NodeSmoothingPoolStateChanged"

// ContractEventName returns the user-defined event name.
func (RocketNodeManagerNodeSmoothingPoolStateChanged) ContractEventName() string {
	return RocketNodeManagerNodeSmoothingPoolStateChangedEventName
}

// UnpackNodeSmoothingPoolStateChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NodeSmoothingPoolStateChanged(address indexed node, bool state)
func (rocketNodeManager *RocketNodeManager) UnpackNodeSmoothingPoolStateChangedEvent(log *types.Log) (*RocketNodeManagerNodeSmoothingPoolStateChanged, error) {
	event := "NodeSmoothingPoolStateChanged"
	if log.Topics[0] != rocketNodeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketNodeManagerNodeSmoothingPoolStateChanged)
	if len(log.Data) > 0 {
		if err := rocketNodeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketNodeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// RocketNodeManagerNodeTimezoneLocationSet represents a NodeTimezoneLocationSet event raised by the RocketNodeManager contract.
type RocketNodeManagerNodeTimezoneLocationSet struct {
	Node common.Address
	Time *big.Int
	Raw  *types.Log // Blockchain specific contextual infos
}

const RocketNodeManagerNodeTimezoneLocationSetEventName = "NodeTimezoneLocationSet"

// ContractEventName returns the user-defined event name.
func (RocketNodeManagerNodeTimezoneLocationSet) ContractEventName() string {
	return RocketNodeManagerNodeTimezoneLocationSetEventName
}

// UnpackNodeTimezoneLocationSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NodeTimezoneLocationSet(address indexed node, uint256 time)
func (rocketNodeManager *RocketNodeManager) UnpackNodeTimezoneLocationSetEvent(log *types.Log) (*RocketNodeManagerNodeTimezoneLocationSet, error) {
	event := "NodeTimezoneLocationSet"
	if log.Topics[0] != rocketNodeManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketNodeManagerNodeTimezoneLocationSet)
	if len(log.Data) > 0 {
		if err := rocketNodeManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketNodeManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}
