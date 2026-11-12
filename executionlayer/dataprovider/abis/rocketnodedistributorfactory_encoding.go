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

// RocketNodeDistributorFactoryMetaData contains all meta data concerning the RocketNodeDistributorFactory contract.
var RocketNodeDistributorFactoryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRocketStorageInterface\",\"name\":\"_rocketStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"createProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyBytecode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "RocketNodeDistributorFactory",
}

// RocketNodeDistributorFactory is an auto generated Go binding around an Ethereum contract.
type RocketNodeDistributorFactory struct {
	abi abi.ABI
}

// NewRocketNodeDistributorFactory creates a new instance of RocketNodeDistributorFactory.
func NewRocketNodeDistributorFactory() *RocketNodeDistributorFactory {
	parsed, err := RocketNodeDistributorFactoryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RocketNodeDistributorFactory{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RocketNodeDistributorFactory) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _rocketStorageAddress) returns()
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) PackConstructor(_rocketStorageAddress common.Address) []byte {
	enc, err := rocketNodeDistributorFactory.abi.Pack("", _rocketStorageAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreateProxy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6140c54c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createProxy(address _nodeAddress) returns()
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) PackCreateProxy(nodeAddress common.Address) []byte {
	enc, err := rocketNodeDistributorFactory.abi.Pack("createProxy", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreateProxy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6140c54c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createProxy(address _nodeAddress) returns()
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) TryPackCreateProxy(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeDistributorFactory.abi.Pack("createProxy", nodeAddress)
}

// PackGetProxyAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa2a5b01.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getProxyAddress(address _nodeAddress) view returns(address)
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) PackGetProxyAddress(nodeAddress common.Address) []byte {
	enc, err := rocketNodeDistributorFactory.abi.Pack("getProxyAddress", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetProxyAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa2a5b01.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getProxyAddress(address _nodeAddress) view returns(address)
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) TryPackGetProxyAddress(nodeAddress common.Address) ([]byte, error) {
	return rocketNodeDistributorFactory.abi.Pack("getProxyAddress", nodeAddress)
}

// UnpackGetProxyAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfa2a5b01.
//
// Solidity: function getProxyAddress(address _nodeAddress) view returns(address)
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) UnpackGetProxyAddress(data []byte) (common.Address, error) {
	out, err := rocketNodeDistributorFactory.abi.Unpack("getProxyAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetProxyBytecode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb416663e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getProxyBytecode() pure returns(bytes)
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) PackGetProxyBytecode() []byte {
	enc, err := rocketNodeDistributorFactory.abi.Pack("getProxyBytecode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetProxyBytecode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb416663e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getProxyBytecode() pure returns(bytes)
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) TryPackGetProxyBytecode() ([]byte, error) {
	return rocketNodeDistributorFactory.abi.Pack("getProxyBytecode")
}

// UnpackGetProxyBytecode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb416663e.
//
// Solidity: function getProxyBytecode() pure returns(bytes)
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) UnpackGetProxyBytecode(data []byte) ([]byte, error) {
	out, err := rocketNodeDistributorFactory.abi.Unpack("getProxyBytecode", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(uint8)
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) PackVersion() []byte {
	enc, err := rocketNodeDistributorFactory.abi.Pack("version")
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
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) TryPackVersion() ([]byte, error) {
	return rocketNodeDistributorFactory.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) UnpackVersion(data []byte) (uint8, error) {
	out, err := rocketNodeDistributorFactory.abi.Unpack("version", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// RocketNodeDistributorFactoryProxyCreated represents a ProxyCreated event raised by the RocketNodeDistributorFactory contract.
type RocketNodeDistributorFactoryProxyCreated struct {
	Address common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const RocketNodeDistributorFactoryProxyCreatedEventName = "ProxyCreated"

// ContractEventName returns the user-defined event name.
func (RocketNodeDistributorFactoryProxyCreated) ContractEventName() string {
	return RocketNodeDistributorFactoryProxyCreatedEventName
}

// UnpackProxyCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProxyCreated(address _address)
func (rocketNodeDistributorFactory *RocketNodeDistributorFactory) UnpackProxyCreatedEvent(log *types.Log) (*RocketNodeDistributorFactoryProxyCreated, error) {
	event := "ProxyCreated"
	if log.Topics[0] != rocketNodeDistributorFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RocketNodeDistributorFactoryProxyCreated)
	if len(log.Data) > 0 {
		if err := rocketNodeDistributorFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rocketNodeDistributorFactory.abi.Events[event].Inputs {
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
