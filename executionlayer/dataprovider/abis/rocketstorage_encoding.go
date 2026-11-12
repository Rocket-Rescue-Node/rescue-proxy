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

// RocketStorageMetaData contains all meta data concerning the RocketStorage contract.
var RocketStorageMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"GuardianChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawalAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"NodeWithdrawalAddressSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addUint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"confirmWithdrawalAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteBool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteBytes32\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteInt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteString\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"deleteUint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeployedStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"r\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodePendingWithdrawalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeWithdrawalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setBool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"setBytes32\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setDeployedStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"setGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"setInt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"setString\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setUint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newWithdrawalAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_confirm\",\"type\":\"bool\"}],\"name\":\"setWithdrawalAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"subUint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "RocketStorage",
}

// RocketStorage is an auto generated Go binding around an Ethereum contract.
type RocketStorage struct {
	abi abi.ABI
}

// NewRocketStorage creates a new instance of RocketStorage.
func NewRocketStorage() *RocketStorage {
	parsed, err := RocketStorageMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RocketStorage{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RocketStorage) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddUint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xadb353dc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addUint(bytes32 _key, uint256 _amount) returns()
func (rocketStorage *RocketStorage) PackAddUint(key [32]byte, amount *big.Int) []byte {
	enc, err := rocketStorage.abi.Pack("addUint", key, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddUint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xadb353dc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addUint(bytes32 _key, uint256 _amount) returns()
func (rocketStorage *RocketStorage) TryPackAddUint(key [32]byte, amount *big.Int) ([]byte, error) {
	return rocketStorage.abi.Pack("addUint", key, amount)
}

// PackConfirmGuardian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e0ea61e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmGuardian() returns()
func (rocketStorage *RocketStorage) PackConfirmGuardian() []byte {
	enc, err := rocketStorage.abi.Pack("confirmGuardian")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmGuardian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e0ea61e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmGuardian() returns()
func (rocketStorage *RocketStorage) TryPackConfirmGuardian() ([]byte, error) {
	return rocketStorage.abi.Pack("confirmGuardian")
}

// PackConfirmWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd439126.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmWithdrawalAddress(address _nodeAddress) returns()
func (rocketStorage *RocketStorage) PackConfirmWithdrawalAddress(nodeAddress common.Address) []byte {
	enc, err := rocketStorage.abi.Pack("confirmWithdrawalAddress", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd439126.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmWithdrawalAddress(address _nodeAddress) returns()
func (rocketStorage *RocketStorage) TryPackConfirmWithdrawalAddress(nodeAddress common.Address) ([]byte, error) {
	return rocketStorage.abi.Pack("confirmWithdrawalAddress", nodeAddress)
}

// PackDeleteAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e14a376.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deleteAddress(bytes32 _key) returns()
func (rocketStorage *RocketStorage) PackDeleteAddress(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("deleteAddress", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeleteAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e14a376.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deleteAddress(bytes32 _key) returns()
func (rocketStorage *RocketStorage) TryPackDeleteAddress(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("deleteAddress", key)
}

// PackDeleteBool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c62ff2d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deleteBool(bytes32 _key) returns()
func (rocketStorage *RocketStorage) PackDeleteBool(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("deleteBool", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeleteBool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c62ff2d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deleteBool(bytes32 _key) returns()
func (rocketStorage *RocketStorage) TryPackDeleteBool(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("deleteBool", key)
}

// PackDeleteBytes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x616b59f6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deleteBytes(bytes32 _key) returns()
func (rocketStorage *RocketStorage) PackDeleteBytes(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("deleteBytes", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeleteBytes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x616b59f6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deleteBytes(bytes32 _key) returns()
func (rocketStorage *RocketStorage) TryPackDeleteBytes(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("deleteBytes", key)
}

// PackDeleteBytes32 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0b9adc57.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deleteBytes32(bytes32 _key) returns()
func (rocketStorage *RocketStorage) PackDeleteBytes32(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("deleteBytes32", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeleteBytes32 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0b9adc57.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deleteBytes32(bytes32 _key) returns()
func (rocketStorage *RocketStorage) TryPackDeleteBytes32(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("deleteBytes32", key)
}

// PackDeleteInt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8c160095.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deleteInt(bytes32 _key) returns()
func (rocketStorage *RocketStorage) PackDeleteInt(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("deleteInt", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeleteInt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8c160095.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deleteInt(bytes32 _key) returns()
func (rocketStorage *RocketStorage) TryPackDeleteInt(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("deleteInt", key)
}

// PackDeleteString is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf6bb3cc4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deleteString(bytes32 _key) returns()
func (rocketStorage *RocketStorage) PackDeleteString(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("deleteString", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeleteString is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf6bb3cc4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deleteString(bytes32 _key) returns()
func (rocketStorage *RocketStorage) TryPackDeleteString(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("deleteString", key)
}

// PackDeleteUint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2b202bf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deleteUint(bytes32 _key) returns()
func (rocketStorage *RocketStorage) PackDeleteUint(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("deleteUint", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeleteUint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2b202bf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deleteUint(bytes32 _key) returns()
func (rocketStorage *RocketStorage) TryPackDeleteUint(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("deleteUint", key)
}

// PackGetAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x21f8a721.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAddress(bytes32 _key) view returns(address r)
func (rocketStorage *RocketStorage) PackGetAddress(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("getAddress", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x21f8a721.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAddress(bytes32 _key) view returns(address r)
func (rocketStorage *RocketStorage) TryPackGetAddress(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("getAddress", key)
}

// UnpackGetAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address r)
func (rocketStorage *RocketStorage) UnpackGetAddress(data []byte) (common.Address, error) {
	out, err := rocketStorage.abi.Unpack("getAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetBool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ae1cfca.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBool(bytes32 _key) view returns(bool r)
func (rocketStorage *RocketStorage) PackGetBool(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("getBool", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ae1cfca.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBool(bytes32 _key) view returns(bool r)
func (rocketStorage *RocketStorage) TryPackGetBool(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("getBool", key)
}

// UnpackGetBool is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ae1cfca.
//
// Solidity: function getBool(bytes32 _key) view returns(bool r)
func (rocketStorage *RocketStorage) UnpackGetBool(data []byte) (bool, error) {
	out, err := rocketStorage.abi.Unpack("getBool", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGetBytes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc031a180.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBytes(bytes32 _key) view returns(bytes)
func (rocketStorage *RocketStorage) PackGetBytes(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("getBytes", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBytes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc031a180.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBytes(bytes32 _key) view returns(bytes)
func (rocketStorage *RocketStorage) TryPackGetBytes(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("getBytes", key)
}

// UnpackGetBytes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc031a180.
//
// Solidity: function getBytes(bytes32 _key) view returns(bytes)
func (rocketStorage *RocketStorage) UnpackGetBytes(data []byte) ([]byte, error) {
	out, err := rocketStorage.abi.Unpack("getBytes", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackGetBytes32 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa6ed563e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBytes32(bytes32 _key) view returns(bytes32 r)
func (rocketStorage *RocketStorage) PackGetBytes32(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("getBytes32", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBytes32 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa6ed563e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBytes32(bytes32 _key) view returns(bytes32 r)
func (rocketStorage *RocketStorage) TryPackGetBytes32(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("getBytes32", key)
}

// UnpackGetBytes32 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa6ed563e.
//
// Solidity: function getBytes32(bytes32 _key) view returns(bytes32 r)
func (rocketStorage *RocketStorage) UnpackGetBytes32(data []byte) ([32]byte, error) {
	out, err := rocketStorage.abi.Unpack("getBytes32", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetDeployedStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1bed5241.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getDeployedStatus() view returns(bool)
func (rocketStorage *RocketStorage) PackGetDeployedStatus() []byte {
	enc, err := rocketStorage.abi.Pack("getDeployedStatus")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetDeployedStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1bed5241.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getDeployedStatus() view returns(bool)
func (rocketStorage *RocketStorage) TryPackGetDeployedStatus() ([]byte, error) {
	return rocketStorage.abi.Pack("getDeployedStatus")
}

// UnpackGetDeployedStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1bed5241.
//
// Solidity: function getDeployedStatus() view returns(bool)
func (rocketStorage *RocketStorage) UnpackGetDeployedStatus(data []byte) (bool, error) {
	out, err := rocketStorage.abi.Unpack("getDeployedStatus", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGetGuardian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa75b87d2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getGuardian() view returns(address)
func (rocketStorage *RocketStorage) PackGetGuardian() []byte {
	enc, err := rocketStorage.abi.Pack("getGuardian")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetGuardian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa75b87d2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getGuardian() view returns(address)
func (rocketStorage *RocketStorage) TryPackGetGuardian() ([]byte, error) {
	return rocketStorage.abi.Pack("getGuardian")
}

// UnpackGetGuardian is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa75b87d2.
//
// Solidity: function getGuardian() view returns(address)
func (rocketStorage *RocketStorage) UnpackGetGuardian(data []byte) (common.Address, error) {
	out, err := rocketStorage.abi.Unpack("getGuardian", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetInt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdc97d962.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getInt(bytes32 _key) view returns(int256 r)
func (rocketStorage *RocketStorage) PackGetInt(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("getInt", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetInt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdc97d962.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getInt(bytes32 _key) view returns(int256 r)
func (rocketStorage *RocketStorage) TryPackGetInt(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("getInt", key)
}

// UnpackGetInt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdc97d962.
//
// Solidity: function getInt(bytes32 _key) view returns(int256 r)
func (rocketStorage *RocketStorage) UnpackGetInt(data []byte) (*big.Int, error) {
	out, err := rocketStorage.abi.Unpack("getInt", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetNodePendingWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfd412513.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodePendingWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketStorage *RocketStorage) PackGetNodePendingWithdrawalAddress(nodeAddress common.Address) []byte {
	enc, err := rocketStorage.abi.Pack("getNodePendingWithdrawalAddress", nodeAddress)
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
func (rocketStorage *RocketStorage) TryPackGetNodePendingWithdrawalAddress(nodeAddress common.Address) ([]byte, error) {
	return rocketStorage.abi.Pack("getNodePendingWithdrawalAddress", nodeAddress)
}

// UnpackGetNodePendingWithdrawalAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfd412513.
//
// Solidity: function getNodePendingWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketStorage *RocketStorage) UnpackGetNodePendingWithdrawalAddress(data []byte) (common.Address, error) {
	out, err := rocketStorage.abi.Unpack("getNodePendingWithdrawalAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNodeWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b49ff62.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNodeWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketStorage *RocketStorage) PackGetNodeWithdrawalAddress(nodeAddress common.Address) []byte {
	enc, err := rocketStorage.abi.Pack("getNodeWithdrawalAddress", nodeAddress)
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
func (rocketStorage *RocketStorage) TryPackGetNodeWithdrawalAddress(nodeAddress common.Address) ([]byte, error) {
	return rocketStorage.abi.Pack("getNodeWithdrawalAddress", nodeAddress)
}

// UnpackGetNodeWithdrawalAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5b49ff62.
//
// Solidity: function getNodeWithdrawalAddress(address _nodeAddress) view returns(address)
func (rocketStorage *RocketStorage) UnpackGetNodeWithdrawalAddress(data []byte) (common.Address, error) {
	out, err := rocketStorage.abi.Unpack("getNodeWithdrawalAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetString is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x986e791a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getString(bytes32 _key) view returns(string)
func (rocketStorage *RocketStorage) PackGetString(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("getString", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetString is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x986e791a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getString(bytes32 _key) view returns(string)
func (rocketStorage *RocketStorage) TryPackGetString(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("getString", key)
}

// UnpackGetString is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x986e791a.
//
// Solidity: function getString(bytes32 _key) view returns(string)
func (rocketStorage *RocketStorage) UnpackGetString(data []byte) (string, error) {
	out, err := rocketStorage.abi.Unpack("getString", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackGetUint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd02d0f5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getUint(bytes32 _key) view returns(uint256 r)
func (rocketStorage *RocketStorage) PackGetUint(key [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("getUint", key)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetUint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd02d0f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getUint(bytes32 _key) view returns(uint256 r)
func (rocketStorage *RocketStorage) TryPackGetUint(key [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("getUint", key)
}

// UnpackGetUint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbd02d0f5.
//
// Solidity: function getUint(bytes32 _key) view returns(uint256 r)
func (rocketStorage *RocketStorage) UnpackGetUint(data []byte) (*big.Int, error) {
	out, err := rocketStorage.abi.Unpack("getUint", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSetAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca446dd9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setAddress(bytes32 _key, address _value) returns()
func (rocketStorage *RocketStorage) PackSetAddress(key [32]byte, value common.Address) []byte {
	enc, err := rocketStorage.abi.Pack("setAddress", key, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca446dd9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setAddress(bytes32 _key, address _value) returns()
func (rocketStorage *RocketStorage) TryPackSetAddress(key [32]byte, value common.Address) ([]byte, error) {
	return rocketStorage.abi.Pack("setAddress", key, value)
}

// PackSetBool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xabfdcced.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setBool(bytes32 _key, bool _value) returns()
func (rocketStorage *RocketStorage) PackSetBool(key [32]byte, value bool) []byte {
	enc, err := rocketStorage.abi.Pack("setBool", key, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetBool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xabfdcced.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setBool(bytes32 _key, bool _value) returns()
func (rocketStorage *RocketStorage) TryPackSetBool(key [32]byte, value bool) ([]byte, error) {
	return rocketStorage.abi.Pack("setBool", key, value)
}

// PackSetBytes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2e28d084.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setBytes(bytes32 _key, bytes _value) returns()
func (rocketStorage *RocketStorage) PackSetBytes(key [32]byte, value []byte) []byte {
	enc, err := rocketStorage.abi.Pack("setBytes", key, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetBytes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2e28d084.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setBytes(bytes32 _key, bytes _value) returns()
func (rocketStorage *RocketStorage) TryPackSetBytes(key [32]byte, value []byte) ([]byte, error) {
	return rocketStorage.abi.Pack("setBytes", key, value)
}

// PackSetBytes32 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4e91db08.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setBytes32(bytes32 _key, bytes32 _value) returns()
func (rocketStorage *RocketStorage) PackSetBytes32(key [32]byte, value [32]byte) []byte {
	enc, err := rocketStorage.abi.Pack("setBytes32", key, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetBytes32 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4e91db08.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setBytes32(bytes32 _key, bytes32 _value) returns()
func (rocketStorage *RocketStorage) TryPackSetBytes32(key [32]byte, value [32]byte) ([]byte, error) {
	return rocketStorage.abi.Pack("setBytes32", key, value)
}

// PackSetDeployedStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfebffd99.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setDeployedStatus() returns()
func (rocketStorage *RocketStorage) PackSetDeployedStatus() []byte {
	enc, err := rocketStorage.abi.Pack("setDeployedStatus")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetDeployedStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfebffd99.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setDeployedStatus() returns()
func (rocketStorage *RocketStorage) TryPackSetDeployedStatus() ([]byte, error) {
	return rocketStorage.abi.Pack("setDeployedStatus")
}

// PackSetGuardian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a0dac4a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setGuardian(address _newAddress) returns()
func (rocketStorage *RocketStorage) PackSetGuardian(newAddress common.Address) []byte {
	enc, err := rocketStorage.abi.Pack("setGuardian", newAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetGuardian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a0dac4a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setGuardian(address _newAddress) returns()
func (rocketStorage *RocketStorage) TryPackSetGuardian(newAddress common.Address) ([]byte, error) {
	return rocketStorage.abi.Pack("setGuardian", newAddress)
}

// PackSetInt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e49bed0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setInt(bytes32 _key, int256 _value) returns()
func (rocketStorage *RocketStorage) PackSetInt(key [32]byte, value *big.Int) []byte {
	enc, err := rocketStorage.abi.Pack("setInt", key, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetInt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e49bed0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setInt(bytes32 _key, int256 _value) returns()
func (rocketStorage *RocketStorage) TryPackSetInt(key [32]byte, value *big.Int) ([]byte, error) {
	return rocketStorage.abi.Pack("setInt", key, value)
}

// PackSetString is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e899550.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setString(bytes32 _key, string _value) returns()
func (rocketStorage *RocketStorage) PackSetString(key [32]byte, value string) []byte {
	enc, err := rocketStorage.abi.Pack("setString", key, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetString is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e899550.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setString(bytes32 _key, string _value) returns()
func (rocketStorage *RocketStorage) TryPackSetString(key [32]byte, value string) ([]byte, error) {
	return rocketStorage.abi.Pack("setString", key, value)
}

// PackSetUint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2a4853a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setUint(bytes32 _key, uint256 _value) returns()
func (rocketStorage *RocketStorage) PackSetUint(key [32]byte, value *big.Int) []byte {
	enc, err := rocketStorage.abi.Pack("setUint", key, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetUint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2a4853a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setUint(bytes32 _key, uint256 _value) returns()
func (rocketStorage *RocketStorage) TryPackSetUint(key [32]byte, value *big.Int) ([]byte, error) {
	return rocketStorage.abi.Pack("setUint", key, value)
}

// PackSetWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa543ccea.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setWithdrawalAddress(address _nodeAddress, address _newWithdrawalAddress, bool _confirm) returns()
func (rocketStorage *RocketStorage) PackSetWithdrawalAddress(nodeAddress common.Address, newWithdrawalAddress common.Address, confirm bool) []byte {
	enc, err := rocketStorage.abi.Pack("setWithdrawalAddress", nodeAddress, newWithdrawalAddress, confirm)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetWithdrawalAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa543ccea.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setWithdrawalAddress(address _nodeAddress, address _newWithdrawalAddress, bool _confirm) returns()
func (rocketStorage *RocketStorage) TryPackSetWithdrawalAddress(nodeAddress common.Address, newWithdrawalAddress common.Address, confirm bool) ([]byte, error) {
	return rocketStorage.abi.Pack("setWithdrawalAddress", nodeAddress, newWithdrawalAddress, confirm)
}

// PackSubUint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xebb9d8c9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function subUint(bytes32 _key, uint256 _amount) returns()
func (rocketStorage *RocketStorage) PackSubUint(key [32]byte, amount *big.Int) []byte {
	enc, err := rocketStorage.abi.Pack("subUint", key, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubUint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xebb9d8c9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function subUint(bytes32 _key, uint256 _amount) returns()
func (rocketStorage *RocketStorage) TryPackSubUint(key [32]byte, amount *big.Int) ([]byte, error) {
	return rocketStorage.abi.Pack("subUint", key, amount)
}

// RocketStorageGuardianChanged represents a GuardianChanged event raised by the RocketStorage contract.
type RocketStorageGuardianChanged struct {
	OldGuardian common.Address
	NewGuardian common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const RocketStorageGuardianChangedEventName = "GuardianChanged"

// ContractEventName returns the user-defined event name.
func (RocketStorageGuardianChanged) ContractEventName() string {
	return RocketStorageGuardianChangedEventName
}

// UnpackGuardianChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GuardianChanged(address oldGuardian, address newGuardian)
func (rocketStorage *RocketStorage) UnpackGuardianChangedEvent(log *types.Log) (*RocketStorageGuardianChanged, error) {
	event := "GuardianChanged"
	if log.Topics[0] != rocketStorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketStorageGuardianChanged)
	if len(log.Data) > 0 {
		if err := rocketStorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketStorage.abi.Events[event].Inputs {
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

// RocketStorageNodeWithdrawalAddressSet represents a NodeWithdrawalAddressSet event raised by the RocketStorage contract.
type RocketStorageNodeWithdrawalAddressSet struct {
	Node              common.Address
	WithdrawalAddress common.Address
	Time              *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const RocketStorageNodeWithdrawalAddressSetEventName = "NodeWithdrawalAddressSet"

// ContractEventName returns the user-defined event name.
func (RocketStorageNodeWithdrawalAddressSet) ContractEventName() string {
	return RocketStorageNodeWithdrawalAddressSetEventName
}

// UnpackNodeWithdrawalAddressSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NodeWithdrawalAddressSet(address indexed node, address indexed withdrawalAddress, uint256 time)
func (rocketStorage *RocketStorage) UnpackNodeWithdrawalAddressSetEvent(log *types.Log) (*RocketStorageNodeWithdrawalAddressSet, error) {
	event := "NodeWithdrawalAddressSet"
	if log.Topics[0] != rocketStorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketStorageNodeWithdrawalAddressSet)
	if len(log.Data) > 0 {
		if err := rocketStorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketStorage.abi.Events[event].Inputs {
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
