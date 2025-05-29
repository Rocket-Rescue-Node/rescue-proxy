// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakewise

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

// EthPrivVaultMetaData contains all meta data concerning the EthPrivVault contract.
var EthPrivVaultMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"mevEscrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "EthPrivVault",
}

// EthPrivVault is an auto generated Go binding around an Ethereum contract.
type EthPrivVault struct {
	abi abi.ABI
}

// NewEthPrivVault creates a new instance of EthPrivVault.
func NewEthPrivVault() *EthPrivVault {
	parsed, err := EthPrivVaultMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &EthPrivVault{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *EthPrivVault) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackMevEscrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3229fa95.
//
// Solidity: function mevEscrow() view returns(address)
func (ethPrivVault *EthPrivVault) PackMevEscrow() []byte {
	enc, err := ethPrivVault.abi.Pack("mevEscrow")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMevEscrow is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3229fa95.
//
// Solidity: function mevEscrow() view returns(address)
func (ethPrivVault *EthPrivVault) UnpackMevEscrow(data []byte) (common.Address, error) {
	out, err := ethPrivVault.abi.Unpack("mevEscrow", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}
