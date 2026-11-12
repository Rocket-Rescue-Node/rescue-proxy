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

// RocketMinipoolManagerMetaData contains all meta data concerning the RocketMinipoolManager contract.
var RocketMinipoolManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRocketStorageInterface\",\"name\":\"_rocketStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minipool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BeginBondReduction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minipool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"CancelReductionVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minipool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"MinipoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minipool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"MinipoolDestroyed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minipool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ReductionCancelled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"createMinipool\",\"outputs\":[{\"internalType\":\"contractRocketMinipoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_bondAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_currentBalance\",\"type\":\"uint256\"}],\"name\":\"createVacantMinipool\",\"outputs\":[{\"internalType\":\"contractRocketMinipoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"decrementNodeStakingMinipoolCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroyMinipool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalisedMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getMinipoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"getMinipoolByPubkey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getMinipoolCountPerStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialisedCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prelaunchCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawableCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dissolvedCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolDepositType\",\"outputs\":[{\"internalType\":\"enumMinipoolDeposit\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolDestroyed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolRPLSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolWithdrawalCredentials\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeActiveMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeFinalisedMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNodeMinipoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeStakingMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_depositSize\",\"type\":\"uint256\"}],\"name\":\"getNodeStakingMinipoolCountBySize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNodeValidatingMinipoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeValidatingMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getPrelaunchMinipools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getVacantMinipoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVacantMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"incrementNodeFinalisedMinipoolCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"incrementNodeStakingMinipoolCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeVacantMinipool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"setMinipoolPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"tryDistribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_previousBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_previousFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"updateNodeStakingMinipoolCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "RocketMinipoolManager",
}

// RocketMinipoolManager is an auto generated Go binding around an Ethereum contract.
type RocketMinipoolManager struct {
	abi abi.ABI
}

// NewRocketMinipoolManager creates a new instance of RocketMinipoolManager.
func NewRocketMinipoolManager() *RocketMinipoolManager {
	parsed, err := RocketMinipoolManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RocketMinipoolManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RocketMinipoolManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _rocketStorageAddress) returns()
func (rocketMinipoolManager *RocketMinipoolManager) PackConstructor(_rocketStorageAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("", _rocketStorageAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreateMinipool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc64372bb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createMinipool(address _nodeAddress, uint256 _salt) returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) PackCreateMinipool(nodeAddress common.Address, salt *big.Int) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("createMinipool", nodeAddress, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreateMinipool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc64372bb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createMinipool(address _nodeAddress, uint256 _salt) returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackCreateMinipool(nodeAddress common.Address, salt *big.Int) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("createMinipool", nodeAddress, salt)
}

// UnpackCreateMinipool is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc64372bb.
//
// Solidity: function createMinipool(address _nodeAddress, uint256 _salt) returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackCreateMinipool(data []byte) (common.Address, error) {
	out, err := rocketMinipoolManager.abi.Unpack("createMinipool", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCreateVacantMinipool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa179778b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createVacantMinipool(address _nodeAddress, uint256 _salt, bytes _validatorPubkey, uint256 _bondAmount, uint256 _currentBalance) returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) PackCreateVacantMinipool(nodeAddress common.Address, salt *big.Int, validatorPubkey []byte, bondAmount *big.Int, currentBalance *big.Int) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("createVacantMinipool", nodeAddress, salt, validatorPubkey, bondAmount, currentBalance)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreateVacantMinipool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa179778b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createVacantMinipool(address _nodeAddress, uint256 _salt, bytes _validatorPubkey, uint256 _bondAmount, uint256 _currentBalance) returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackCreateVacantMinipool(nodeAddress common.Address, salt *big.Int, validatorPubkey []byte, bondAmount *big.Int, currentBalance *big.Int) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("createVacantMinipool", nodeAddress, salt, validatorPubkey, bondAmount, currentBalance)
}

// UnpackCreateVacantMinipool is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa179778b.
//
// Solidity: function createVacantMinipool(address _nodeAddress, uint256 _salt, bytes _validatorPubkey, uint256 _bondAmount, uint256 _currentBalance) returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackCreateVacantMinipool(data []byte) (common.Address, error) {
	out, err := rocketMinipoolManager.abi.Unpack("createVacantMinipool", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDecrementNodeStakingMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75b59c7f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function decrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (rocketMinipoolManager *RocketMinipoolManager) PackDecrementNodeStakingMinipoolCount(nodeAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("decrementNodeStakingMinipoolCount", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDecrementNodeStakingMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75b59c7f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function decrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (rocketMinipoolManager *RocketMinipoolManager) TryPackDecrementNodeStakingMinipoolCount(nodeAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("decrementNodeStakingMinipoolCount", nodeAddress)
}

// PackDestroyMinipool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7bb40aaf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function destroyMinipool() returns()
func (rocketMinipoolManager *RocketMinipoolManager) PackDestroyMinipool() []byte {
	enc, err := rocketMinipoolManager.abi.Pack("destroyMinipool")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDestroyMinipool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7bb40aaf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function destroyMinipool() returns()
func (rocketMinipoolManager *RocketMinipoolManager) TryPackDestroyMinipool() ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("destroyMinipool")
}

// PackGetActiveMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce9b79ad.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getActiveMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetActiveMinipoolCount() []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getActiveMinipoolCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetActiveMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce9b79ad.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getActiveMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetActiveMinipoolCount() ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getActiveMinipoolCount")
}

// UnpackGetActiveMinipoolCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xce9b79ad.
//
// Solidity: function getActiveMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetActiveMinipoolCount(data []byte) (*big.Int, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getActiveMinipoolCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetFinalisedMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1ea6ce0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getFinalisedMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetFinalisedMinipoolCount() []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getFinalisedMinipoolCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetFinalisedMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1ea6ce0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getFinalisedMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetFinalisedMinipoolCount() ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getFinalisedMinipoolCount")
}

// UnpackGetFinalisedMinipoolCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd1ea6ce0.
//
// Solidity: function getFinalisedMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetFinalisedMinipoolCount(data []byte) (*big.Int, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getFinalisedMinipoolCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetMinipoolAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeff7319f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinipoolAt(uint256 _index) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetMinipoolAt(index *big.Int) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getMinipoolAt", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinipoolAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeff7319f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinipoolAt(uint256 _index) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetMinipoolAt(index *big.Int) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getMinipoolAt", index)
}

// UnpackGetMinipoolAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xeff7319f.
//
// Solidity: function getMinipoolAt(uint256 _index) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetMinipoolAt(data []byte) (common.Address, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getMinipoolAt", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetMinipoolByPubkey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcf6a4763.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinipoolByPubkey(bytes _pubkey) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetMinipoolByPubkey(pubkey []byte) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getMinipoolByPubkey", pubkey)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinipoolByPubkey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcf6a4763.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinipoolByPubkey(bytes _pubkey) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetMinipoolByPubkey(pubkey []byte) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getMinipoolByPubkey", pubkey)
}

// UnpackGetMinipoolByPubkey is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6a4763.
//
// Solidity: function getMinipoolByPubkey(bytes _pubkey) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetMinipoolByPubkey(data []byte) (common.Address, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getMinipoolByPubkey", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xae4d0bed.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetMinipoolCount() []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getMinipoolCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xae4d0bed.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetMinipoolCount() ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getMinipoolCount")
}

// UnpackGetMinipoolCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xae4d0bed.
//
// Solidity: function getMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetMinipoolCount(data []byte) (*big.Int, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getMinipoolCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetMinipoolCountPerStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3b5ecefa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinipoolCountPerStatus(uint256 _offset, uint256 _limit) view returns(uint256 initialisedCount, uint256 prelaunchCount, uint256 stakingCount, uint256 withdrawableCount, uint256 dissolvedCount)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetMinipoolCountPerStatus(offset *big.Int, limit *big.Int) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getMinipoolCountPerStatus", offset, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinipoolCountPerStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3b5ecefa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinipoolCountPerStatus(uint256 _offset, uint256 _limit) view returns(uint256 initialisedCount, uint256 prelaunchCount, uint256 stakingCount, uint256 withdrawableCount, uint256 dissolvedCount)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetMinipoolCountPerStatus(offset *big.Int, limit *big.Int) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getMinipoolCountPerStatus", offset, limit)
}

// GetMinipoolCountPerStatusOutput serves as a container for the return parameters of contract
// method GetMinipoolCountPerStatus.
type GetMinipoolCountPerStatusOutput struct {
	InitialisedCount  *big.Int
	PrelaunchCount    *big.Int
	StakingCount      *big.Int
	WithdrawableCount *big.Int
	DissolvedCount    *big.Int
}

// UnpackGetMinipoolCountPerStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3b5ecefa.
//
// Solidity: function getMinipoolCountPerStatus(uint256 _offset, uint256 _limit) view returns(uint256 initialisedCount, uint256 prelaunchCount, uint256 stakingCount, uint256 withdrawableCount, uint256 dissolvedCount)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetMinipoolCountPerStatus(data []byte) (GetMinipoolCountPerStatusOutput, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getMinipoolCountPerStatus", data)
	outstruct := new(GetMinipoolCountPerStatusOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.InitialisedCount = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.PrelaunchCount = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.StakingCount = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	outstruct.WithdrawableCount = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.DissolvedCount = abi.ConvertType(out[4], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGetMinipoolDepositType is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ea1a6e2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinipoolDepositType(address _minipoolAddress) view returns(uint8)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetMinipoolDepositType(minipoolAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getMinipoolDepositType", minipoolAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinipoolDepositType is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ea1a6e2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinipoolDepositType(address _minipoolAddress) view returns(uint8)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetMinipoolDepositType(minipoolAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getMinipoolDepositType", minipoolAddress)
}

// UnpackGetMinipoolDepositType is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5ea1a6e2.
//
// Solidity: function getMinipoolDepositType(address _minipoolAddress) view returns(uint8)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetMinipoolDepositType(data []byte) (uint8, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getMinipoolDepositType", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackGetMinipoolDestroyed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa757987a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinipoolDestroyed(address _minipoolAddress) view returns(bool)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetMinipoolDestroyed(minipoolAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getMinipoolDestroyed", minipoolAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinipoolDestroyed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa757987a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinipoolDestroyed(address _minipoolAddress) view returns(bool)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetMinipoolDestroyed(minipoolAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getMinipoolDestroyed", minipoolAddress)
}

// UnpackGetMinipoolDestroyed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa757987a.
//
// Solidity: function getMinipoolDestroyed(address _minipoolAddress) view returns(bool)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetMinipoolDestroyed(data []byte) (bool, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getMinipoolDestroyed", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGetMinipoolExists is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x606bb62e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinipoolExists(address _minipoolAddress) view returns(bool)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetMinipoolExists(minipoolAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getMinipoolExists", minipoolAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinipoolExists is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x606bb62e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinipoolExists(address _minipoolAddress) view returns(bool)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetMinipoolExists(minipoolAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getMinipoolExists", minipoolAddress)
}

// UnpackGetMinipoolExists is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x606bb62e.
//
// Solidity: function getMinipoolExists(address _minipoolAddress) view returns(bool)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetMinipoolExists(data []byte) (bool, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getMinipoolExists", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGetMinipoolPubkey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3eb535e9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinipoolPubkey(address _minipoolAddress) view returns(bytes)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetMinipoolPubkey(minipoolAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getMinipoolPubkey", minipoolAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinipoolPubkey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3eb535e9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinipoolPubkey(address _minipoolAddress) view returns(bytes)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetMinipoolPubkey(minipoolAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getMinipoolPubkey", minipoolAddress)
}

// UnpackGetMinipoolPubkey is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3eb535e9.
//
// Solidity: function getMinipoolPubkey(address _minipoolAddress) view returns(bytes)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetMinipoolPubkey(data []byte) ([]byte, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getMinipoolPubkey", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackGetMinipoolRPLSlashed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c21b8a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinipoolRPLSlashed(address _minipoolAddress) view returns(bool)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetMinipoolRPLSlashed(minipoolAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getMinipoolRPLSlashed", minipoolAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinipoolRPLSlashed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c21b8a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinipoolRPLSlashed(address _minipoolAddress) view returns(bool)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetMinipoolRPLSlashed(minipoolAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getMinipoolRPLSlashed", minipoolAddress)
}

// UnpackGetMinipoolRPLSlashed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0c21b8a7.
//
// Solidity: function getMinipoolRPLSlashed(address _minipoolAddress) view returns(bool)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetMinipoolRPLSlashed(data []byte) (bool, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getMinipoolRPLSlashed", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGetMinipoolWithdrawalCredentials is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2cb76c37.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinipoolWithdrawalCredentials(address _minipoolAddress) pure returns(bytes)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetMinipoolWithdrawalCredentials(minipoolAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getMinipoolWithdrawalCredentials", minipoolAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinipoolWithdrawalCredentials is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2cb76c37.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinipoolWithdrawalCredentials(address _minipoolAddress) pure returns(bytes)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetMinipoolWithdrawalCredentials(minipoolAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getMinipoolWithdrawalCredentials", minipoolAddress)
}

// UnpackGetMinipoolWithdrawalCredentials is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2cb76c37.
//
// Solidity: function getMinipoolWithdrawalCredentials(address _minipoolAddress) pure returns(bytes)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetMinipoolWithdrawalCredentials(data []byte) ([]byte, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getMinipoolWithdrawalCredentials", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackGetNodeActiveMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1844ec01.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeActiveMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetNodeActiveMinipoolCount(nodeAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getNodeActiveMinipoolCount", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeActiveMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1844ec01.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeActiveMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetNodeActiveMinipoolCount(nodeAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getNodeActiveMinipoolCount", nodeAddress)
}

// UnpackGetNodeActiveMinipoolCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1844ec01.
//
// Solidity: function getNodeActiveMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetNodeActiveMinipoolCount(data []byte) (*big.Int, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getNodeActiveMinipoolCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetNodeFinalisedMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88a89f7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeFinalisedMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetNodeFinalisedMinipoolCount(nodeAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getNodeFinalisedMinipoolCount", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeFinalisedMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88a89f7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeFinalisedMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetNodeFinalisedMinipoolCount(nodeAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getNodeFinalisedMinipoolCount", nodeAddress)
}

// UnpackGetNodeFinalisedMinipoolCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb88a89f7.
//
// Solidity: function getNodeFinalisedMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetNodeFinalisedMinipoolCount(data []byte) (*big.Int, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getNodeFinalisedMinipoolCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetNodeMinipoolAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8b300029.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetNodeMinipoolAt(nodeAddress common.Address, index *big.Int) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getNodeMinipoolAt", nodeAddress, index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeMinipoolAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8b300029.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetNodeMinipoolAt(nodeAddress common.Address, index *big.Int) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getNodeMinipoolAt", nodeAddress, index)
}

// UnpackGetNodeMinipoolAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8b300029.
//
// Solidity: function getNodeMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetNodeMinipoolAt(data []byte) (common.Address, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getNodeMinipoolAt", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNodeMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1ce9ec33.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetNodeMinipoolCount(nodeAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getNodeMinipoolCount", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1ce9ec33.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetNodeMinipoolCount(nodeAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getNodeMinipoolCount", nodeAddress)
}

// UnpackGetNodeMinipoolCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1ce9ec33.
//
// Solidity: function getNodeMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetNodeMinipoolCount(data []byte) (*big.Int, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getNodeMinipoolCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetNodeStakingMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57b4ef6b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeStakingMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetNodeStakingMinipoolCount(nodeAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getNodeStakingMinipoolCount", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeStakingMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57b4ef6b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeStakingMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetNodeStakingMinipoolCount(nodeAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getNodeStakingMinipoolCount", nodeAddress)
}

// UnpackGetNodeStakingMinipoolCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x57b4ef6b.
//
// Solidity: function getNodeStakingMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetNodeStakingMinipoolCount(data []byte) (*big.Int, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getNodeStakingMinipoolCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetNodeStakingMinipoolCountBySize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x240eb330.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeStakingMinipoolCountBySize(address _nodeAddress, uint256 _depositSize) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetNodeStakingMinipoolCountBySize(nodeAddress common.Address, depositSize *big.Int) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getNodeStakingMinipoolCountBySize", nodeAddress, depositSize)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeStakingMinipoolCountBySize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x240eb330.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeStakingMinipoolCountBySize(address _nodeAddress, uint256 _depositSize) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetNodeStakingMinipoolCountBySize(nodeAddress common.Address, depositSize *big.Int) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getNodeStakingMinipoolCountBySize", nodeAddress, depositSize)
}

// UnpackGetNodeStakingMinipoolCountBySize is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x240eb330.
//
// Solidity: function getNodeStakingMinipoolCountBySize(address _nodeAddress, uint256 _depositSize) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetNodeStakingMinipoolCountBySize(data []byte) (*big.Int, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getNodeStakingMinipoolCountBySize", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetNodeValidatingMinipoolAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9da0700f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeValidatingMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetNodeValidatingMinipoolAt(nodeAddress common.Address, index *big.Int) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getNodeValidatingMinipoolAt", nodeAddress, index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeValidatingMinipoolAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9da0700f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeValidatingMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetNodeValidatingMinipoolAt(nodeAddress common.Address, index *big.Int) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getNodeValidatingMinipoolAt", nodeAddress, index)
}

// UnpackGetNodeValidatingMinipoolAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9da0700f.
//
// Solidity: function getNodeValidatingMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetNodeValidatingMinipoolAt(data []byte) (common.Address, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getNodeValidatingMinipoolAt", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNodeValidatingMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf90267c4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeValidatingMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetNodeValidatingMinipoolCount(nodeAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getNodeValidatingMinipoolCount", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNodeValidatingMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf90267c4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNodeValidatingMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetNodeValidatingMinipoolCount(nodeAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getNodeValidatingMinipoolCount", nodeAddress)
}

// UnpackGetNodeValidatingMinipoolCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf90267c4.
//
// Solidity: function getNodeValidatingMinipoolCount(address _nodeAddress) view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetNodeValidatingMinipoolCount(data []byte) (*big.Int, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getNodeValidatingMinipoolCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetPrelaunchMinipools is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5dfef965.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPrelaunchMinipools(uint256 _offset, uint256 _limit) view returns(address[])
func (rocketMinipoolManager *RocketMinipoolManager) PackGetPrelaunchMinipools(offset *big.Int, limit *big.Int) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getPrelaunchMinipools", offset, limit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPrelaunchMinipools is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5dfef965.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPrelaunchMinipools(uint256 _offset, uint256 _limit) view returns(address[])
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetPrelaunchMinipools(offset *big.Int, limit *big.Int) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getPrelaunchMinipools", offset, limit)
}

// UnpackGetPrelaunchMinipools is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5dfef965.
//
// Solidity: function getPrelaunchMinipools(uint256 _offset, uint256 _limit) view returns(address[])
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetPrelaunchMinipools(data []byte) ([]common.Address, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getPrelaunchMinipools", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetStakingMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67bca235.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getStakingMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetStakingMinipoolCount() []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getStakingMinipoolCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetStakingMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67bca235.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getStakingMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetStakingMinipoolCount() ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getStakingMinipoolCount")
}

// UnpackGetStakingMinipoolCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x67bca235.
//
// Solidity: function getStakingMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetStakingMinipoolCount(data []byte) (*big.Int, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getStakingMinipoolCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetVacantMinipoolAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1401991.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVacantMinipoolAt(uint256 _index) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetVacantMinipoolAt(index *big.Int) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getVacantMinipoolAt", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVacantMinipoolAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1401991.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVacantMinipoolAt(uint256 _index) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetVacantMinipoolAt(index *big.Int) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getVacantMinipoolAt", index)
}

// UnpackGetVacantMinipoolAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd1401991.
//
// Solidity: function getVacantMinipoolAt(uint256 _index) view returns(address)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetVacantMinipoolAt(data []byte) (common.Address, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getVacantMinipoolAt", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetVacantMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1286377e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVacantMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) PackGetVacantMinipoolCount() []byte {
	enc, err := rocketMinipoolManager.abi.Pack("getVacantMinipoolCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVacantMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1286377e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVacantMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) TryPackGetVacantMinipoolCount() ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("getVacantMinipoolCount")
}

// UnpackGetVacantMinipoolCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1286377e.
//
// Solidity: function getVacantMinipoolCount() view returns(uint256)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackGetVacantMinipoolCount(data []byte) (*big.Int, error) {
	out, err := rocketMinipoolManager.abi.Unpack("getVacantMinipoolCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackIncrementNodeFinalisedMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb04e8868.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function incrementNodeFinalisedMinipoolCount(address _nodeAddress) returns()
func (rocketMinipoolManager *RocketMinipoolManager) PackIncrementNodeFinalisedMinipoolCount(nodeAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("incrementNodeFinalisedMinipoolCount", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIncrementNodeFinalisedMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb04e8868.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function incrementNodeFinalisedMinipoolCount(address _nodeAddress) returns()
func (rocketMinipoolManager *RocketMinipoolManager) TryPackIncrementNodeFinalisedMinipoolCount(nodeAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("incrementNodeFinalisedMinipoolCount", nodeAddress)
}

// PackIncrementNodeStakingMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9907288c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function incrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (rocketMinipoolManager *RocketMinipoolManager) PackIncrementNodeStakingMinipoolCount(nodeAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("incrementNodeStakingMinipoolCount", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIncrementNodeStakingMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9907288c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function incrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (rocketMinipoolManager *RocketMinipoolManager) TryPackIncrementNodeStakingMinipoolCount(nodeAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("incrementNodeStakingMinipoolCount", nodeAddress)
}

// PackRemoveVacantMinipool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44e51a03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeVacantMinipool() returns()
func (rocketMinipoolManager *RocketMinipoolManager) PackRemoveVacantMinipool() []byte {
	enc, err := rocketMinipoolManager.abi.Pack("removeVacantMinipool")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveVacantMinipool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44e51a03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeVacantMinipool() returns()
func (rocketMinipoolManager *RocketMinipoolManager) TryPackRemoveVacantMinipool() ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("removeVacantMinipool")
}

// PackSetMinipoolPubkey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c7f64d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMinipoolPubkey(bytes _pubkey) returns()
func (rocketMinipoolManager *RocketMinipoolManager) PackSetMinipoolPubkey(pubkey []byte) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("setMinipoolPubkey", pubkey)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMinipoolPubkey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c7f64d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMinipoolPubkey(bytes _pubkey) returns()
func (rocketMinipoolManager *RocketMinipoolManager) TryPackSetMinipoolPubkey(pubkey []byte) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("setMinipoolPubkey", pubkey)
}

// PackTryDistribute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1afe958.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function tryDistribute(address _nodeAddress) returns()
func (rocketMinipoolManager *RocketMinipoolManager) PackTryDistribute(nodeAddress common.Address) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("tryDistribute", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTryDistribute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1afe958.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function tryDistribute(address _nodeAddress) returns()
func (rocketMinipoolManager *RocketMinipoolManager) TryPackTryDistribute(nodeAddress common.Address) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("tryDistribute", nodeAddress)
}

// PackUpdateNodeStakingMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0fcc8178.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateNodeStakingMinipoolCount(uint256 _previousBond, uint256 _newBond, uint256 _previousFee, uint256 _newFee) returns()
func (rocketMinipoolManager *RocketMinipoolManager) PackUpdateNodeStakingMinipoolCount(previousBond *big.Int, newBond *big.Int, previousFee *big.Int, newFee *big.Int) []byte {
	enc, err := rocketMinipoolManager.abi.Pack("updateNodeStakingMinipoolCount", previousBond, newBond, previousFee, newFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateNodeStakingMinipoolCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0fcc8178.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateNodeStakingMinipoolCount(uint256 _previousBond, uint256 _newBond, uint256 _previousFee, uint256 _newFee) returns()
func (rocketMinipoolManager *RocketMinipoolManager) TryPackUpdateNodeStakingMinipoolCount(previousBond *big.Int, newBond *big.Int, previousFee *big.Int, newFee *big.Int) ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("updateNodeStakingMinipoolCount", previousBond, newBond, previousFee, newFee)
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(uint8)
func (rocketMinipoolManager *RocketMinipoolManager) PackVersion() []byte {
	enc, err := rocketMinipoolManager.abi.Pack("version")
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
func (rocketMinipoolManager *RocketMinipoolManager) TryPackVersion() ([]byte, error) {
	return rocketMinipoolManager.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackVersion(data []byte) (uint8, error) {
	out, err := rocketMinipoolManager.abi.Unpack("version", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// RocketMinipoolManagerBeginBondReduction represents a BeginBondReduction event raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerBeginBondReduction struct {
	Minipool common.Address
	Time     *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const RocketMinipoolManagerBeginBondReductionEventName = "BeginBondReduction"

// ContractEventName returns the user-defined event name.
func (RocketMinipoolManagerBeginBondReduction) ContractEventName() string {
	return RocketMinipoolManagerBeginBondReductionEventName
}

// UnpackBeginBondReductionEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BeginBondReduction(address indexed minipool, uint256 time)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackBeginBondReductionEvent(log *types.Log) (*RocketMinipoolManagerBeginBondReduction, error) {
	event := "BeginBondReduction"
	if log.Topics[0] != rocketMinipoolManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketMinipoolManagerBeginBondReduction)
	if len(log.Data) > 0 {
		if err := rocketMinipoolManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketMinipoolManager.abi.Events[event].Inputs {
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

// RocketMinipoolManagerCancelReductionVoted represents a CancelReductionVoted event raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerCancelReductionVoted struct {
	Minipool common.Address
	Member   common.Address
	Time     *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const RocketMinipoolManagerCancelReductionVotedEventName = "CancelReductionVoted"

// ContractEventName returns the user-defined event name.
func (RocketMinipoolManagerCancelReductionVoted) ContractEventName() string {
	return RocketMinipoolManagerCancelReductionVotedEventName
}

// UnpackCancelReductionVotedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CancelReductionVoted(address indexed minipool, address indexed member, uint256 time)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackCancelReductionVotedEvent(log *types.Log) (*RocketMinipoolManagerCancelReductionVoted, error) {
	event := "CancelReductionVoted"
	if log.Topics[0] != rocketMinipoolManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketMinipoolManagerCancelReductionVoted)
	if len(log.Data) > 0 {
		if err := rocketMinipoolManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketMinipoolManager.abi.Events[event].Inputs {
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

// RocketMinipoolManagerMinipoolCreated represents a MinipoolCreated event raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerMinipoolCreated struct {
	Minipool common.Address
	Node     common.Address
	Time     *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const RocketMinipoolManagerMinipoolCreatedEventName = "MinipoolCreated"

// ContractEventName returns the user-defined event name.
func (RocketMinipoolManagerMinipoolCreated) ContractEventName() string {
	return RocketMinipoolManagerMinipoolCreatedEventName
}

// UnpackMinipoolCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinipoolCreated(address indexed minipool, address indexed node, uint256 time)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackMinipoolCreatedEvent(log *types.Log) (*RocketMinipoolManagerMinipoolCreated, error) {
	event := "MinipoolCreated"
	if log.Topics[0] != rocketMinipoolManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketMinipoolManagerMinipoolCreated)
	if len(log.Data) > 0 {
		if err := rocketMinipoolManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketMinipoolManager.abi.Events[event].Inputs {
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

// RocketMinipoolManagerMinipoolDestroyed represents a MinipoolDestroyed event raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerMinipoolDestroyed struct {
	Minipool common.Address
	Node     common.Address
	Time     *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const RocketMinipoolManagerMinipoolDestroyedEventName = "MinipoolDestroyed"

// ContractEventName returns the user-defined event name.
func (RocketMinipoolManagerMinipoolDestroyed) ContractEventName() string {
	return RocketMinipoolManagerMinipoolDestroyedEventName
}

// UnpackMinipoolDestroyedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinipoolDestroyed(address indexed minipool, address indexed node, uint256 time)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackMinipoolDestroyedEvent(log *types.Log) (*RocketMinipoolManagerMinipoolDestroyed, error) {
	event := "MinipoolDestroyed"
	if log.Topics[0] != rocketMinipoolManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketMinipoolManagerMinipoolDestroyed)
	if len(log.Data) > 0 {
		if err := rocketMinipoolManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketMinipoolManager.abi.Events[event].Inputs {
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

// RocketMinipoolManagerReductionCancelled represents a ReductionCancelled event raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerReductionCancelled struct {
	Minipool common.Address
	Time     *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const RocketMinipoolManagerReductionCancelledEventName = "ReductionCancelled"

// ContractEventName returns the user-defined event name.
func (RocketMinipoolManagerReductionCancelled) ContractEventName() string {
	return RocketMinipoolManagerReductionCancelledEventName
}

// UnpackReductionCancelledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ReductionCancelled(address indexed minipool, uint256 time)
func (rocketMinipoolManager *RocketMinipoolManager) UnpackReductionCancelledEvent(log *types.Log) (*RocketMinipoolManagerReductionCancelled, error) {
	event := "ReductionCancelled"
	if log.Topics[0] != rocketMinipoolManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketMinipoolManagerReductionCancelled)
	if len(log.Data) > 0 {
		if err := rocketMinipoolManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketMinipoolManager.abi.Events[event].Inputs {
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
