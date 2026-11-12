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

// RocketMegapoolMetaData contains all meta data concerning the RocketMegapool contract.
var RocketMegapoolMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getValidatorCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_validatorId\",\"type\":\"uint32\"}],\"name\":\"getValidatorPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "RocketMegapool",
}

// RocketMegapool is an auto generated Go binding around an Ethereum contract.
type RocketMegapool struct {
	abi abi.ABI
}

// NewRocketMegapool creates a new instance of RocketMegapool.
func NewRocketMegapool() *RocketMegapool {
	parsed, err := RocketMegapoolMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RocketMegapool{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RocketMegapool) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetValidatorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7071688a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getValidatorCount() view returns(uint32)
func (rocketMegapool *RocketMegapool) PackGetValidatorCount() []byte {
	enc, err := rocketMegapool.abi.Pack("getValidatorCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetValidatorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7071688a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getValidatorCount() view returns(uint32)
func (rocketMegapool *RocketMegapool) TryPackGetValidatorCount() ([]byte, error) {
	return rocketMegapool.abi.Pack("getValidatorCount")
}

// UnpackGetValidatorCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7071688a.
//
// Solidity: function getValidatorCount() view returns(uint32)
func (rocketMegapool *RocketMegapool) UnpackGetValidatorCount(data []byte) (uint32, error) {
	out, err := rocketMegapool.abi.Unpack("getValidatorCount", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackGetValidatorPubkey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e335206.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getValidatorPubkey(uint32 _validatorId) view returns(bytes)
func (rocketMegapool *RocketMegapool) PackGetValidatorPubkey(validatorId uint32) []byte {
	enc, err := rocketMegapool.abi.Pack("getValidatorPubkey", validatorId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetValidatorPubkey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e335206.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getValidatorPubkey(uint32 _validatorId) view returns(bytes)
func (rocketMegapool *RocketMegapool) TryPackGetValidatorPubkey(validatorId uint32) ([]byte, error) {
	return rocketMegapool.abi.Pack("getValidatorPubkey", validatorId)
}

// UnpackGetValidatorPubkey is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0e335206.
//
// Solidity: function getValidatorPubkey(uint32 _validatorId) view returns(bytes)
func (rocketMegapool *RocketMegapool) UnpackGetValidatorPubkey(data []byte) ([]byte, error) {
	out, err := rocketMegapool.abi.Unpack("getValidatorPubkey", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}
