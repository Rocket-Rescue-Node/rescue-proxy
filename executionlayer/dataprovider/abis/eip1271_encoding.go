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

// EIP1271MetaData contains all meta data concerning the EIP1271 contract.
var EIP1271MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "EIP1271",
}

// EIP1271 is an auto generated Go binding around an Ethereum contract.
type EIP1271 struct {
	abi abi.ABI
}

// NewEIP1271 creates a new instance of EIP1271.
func NewEIP1271() *EIP1271 {
	parsed, err := EIP1271MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &EIP1271{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *EIP1271) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackIsValidSignature is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1626ba7e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (eIP1271 *EIP1271) PackIsValidSignature(hash [32]byte, signature []byte) []byte {
	enc, err := eIP1271.abi.Pack("isValidSignature", hash, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsValidSignature is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1626ba7e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (eIP1271 *EIP1271) TryPackIsValidSignature(hash [32]byte, signature []byte) ([]byte, error) {
	return eIP1271.abi.Pack("isValidSignature", hash, signature)
}

// UnpackIsValidSignature is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (eIP1271 *EIP1271) UnpackIsValidSignature(data []byte) ([4]byte, error) {
	out, err := eIP1271.abi.Unpack("isValidSignature", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}
