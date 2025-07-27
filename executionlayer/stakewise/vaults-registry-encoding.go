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

// VaultsRegistryMetaData contains all meta data concerning the VaultsRegistry contract.
var VaultsRegistryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "VaultsRegistry",
}

// VaultsRegistry is an auto generated Go binding around an Ethereum contract.
type VaultsRegistry struct {
	abi abi.ABI
}

// NewVaultsRegistry creates a new instance of VaultsRegistry.
func NewVaultsRegistry() *VaultsRegistry {
	parsed, err := VaultsRegistryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &VaultsRegistry{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *VaultsRegistry) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackVaults is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa622ee7c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function vaults(address ) view returns(bool)
func (vaultsRegistry *VaultsRegistry) PackVaults(arg0 common.Address) []byte {
	enc, err := vaultsRegistry.abi.Pack("vaults", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVaults is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa622ee7c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function vaults(address ) view returns(bool)
func (vaultsRegistry *VaultsRegistry) TryPackVaults(arg0 common.Address) ([]byte, error) {
	return vaultsRegistry.abi.Pack("vaults", arg0)
}

// UnpackVaults is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa622ee7c.
//
// Solidity: function vaults(address ) view returns(bool)
func (vaultsRegistry *VaultsRegistry) UnpackVaults(data []byte) (bool, error) {
	out, err := vaultsRegistry.abi.Unpack("vaults", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}
