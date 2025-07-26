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

// RocketDaoNodeTrustedMetaData contains all meta data concerning the RocketDaoNodeTrusted contract.
var RocketDaoNodeTrustedMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRocketStorageInterface\",\"name\":\"_rocketStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_confirmDisableBootstrapMode\",\"type\":\"bool\"}],\"name\":\"bootstrapDisable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"bootstrapMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_settingContractName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_settingPath\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"bootstrapSettingBool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_settingContractName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_settingPath\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"bootstrapSettingUint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"decrementMemberUnbondedValidatorCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBootstrapModeDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getMemberAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getMemberID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getMemberIsChallenged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getMemberIsValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getMemberJoinedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getMemberLastProposalTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMemberMinRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_proposalType\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getMemberProposalExecutedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMemberQuorumVotesRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getMemberRPLBondAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getMemberUnbondedValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getMemberUrl\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"incrementMemberUnbondedValidatorCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"memberJoinRequired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "RocketDaoNodeTrusted",
}

// RocketDaoNodeTrusted is an auto generated Go binding around an Ethereum contract.
type RocketDaoNodeTrusted struct {
	abi abi.ABI
}

// NewRocketDaoNodeTrusted creates a new instance of RocketDaoNodeTrusted.
func NewRocketDaoNodeTrusted() *RocketDaoNodeTrusted {
	parsed, err := RocketDaoNodeTrustedMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RocketDaoNodeTrusted{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RocketDaoNodeTrusted) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _rocketStorageAddress) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackConstructor(_rocketStorageAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("", _rocketStorageAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBootstrapDisable is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1503944.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bootstrapDisable(bool _confirmDisableBootstrapMode) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackBootstrapDisable(confirmDisableBootstrapMode bool) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("bootstrapDisable", confirmDisableBootstrapMode)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBootstrapDisable is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1503944.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function bootstrapDisable(bool _confirmDisableBootstrapMode) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackBootstrapDisable(confirmDisableBootstrapMode bool) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("bootstrapDisable", confirmDisableBootstrapMode)
}

// PackBootstrapMember is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48795904.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bootstrapMember(string _id, string _url, address _nodeAddress) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackBootstrapMember(id string, url string, nodeAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("bootstrapMember", id, url, nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBootstrapMember is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48795904.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function bootstrapMember(string _id, string _url, address _nodeAddress) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackBootstrapMember(id string, url string, nodeAddress common.Address) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("bootstrapMember", id, url, nodeAddress)
}

// PackBootstrapSettingBool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3edad14.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bootstrapSettingBool(string _settingContractName, string _settingPath, bool _value) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackBootstrapSettingBool(settingContractName string, settingPath string, value bool) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("bootstrapSettingBool", settingContractName, settingPath, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBootstrapSettingBool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3edad14.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function bootstrapSettingBool(string _settingContractName, string _settingPath, bool _value) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackBootstrapSettingBool(settingContractName string, settingPath string, value bool) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("bootstrapSettingBool", settingContractName, settingPath, value)
}

// PackBootstrapSettingUint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb3b0db22.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bootstrapSettingUint(string _settingContractName, string _settingPath, uint256 _value) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackBootstrapSettingUint(settingContractName string, settingPath string, value *big.Int) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("bootstrapSettingUint", settingContractName, settingPath, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBootstrapSettingUint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb3b0db22.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function bootstrapSettingUint(string _settingContractName, string _settingPath, uint256 _value) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackBootstrapSettingUint(settingContractName string, settingPath string, value *big.Int) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("bootstrapSettingUint", settingContractName, settingPath, value)
}

// PackDecrementMemberUnbondedValidatorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54d28878.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function decrementMemberUnbondedValidatorCount(address _nodeAddress) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackDecrementMemberUnbondedValidatorCount(nodeAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("decrementMemberUnbondedValidatorCount", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDecrementMemberUnbondedValidatorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54d28878.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function decrementMemberUnbondedValidatorCount(address _nodeAddress) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackDecrementMemberUnbondedValidatorCount(nodeAddress common.Address) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("decrementMemberUnbondedValidatorCount", nodeAddress)
}

// PackGetBootstrapModeDisabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf54746e4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBootstrapModeDisabled() view returns(bool)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetBootstrapModeDisabled() []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getBootstrapModeDisabled")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBootstrapModeDisabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf54746e4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBootstrapModeDisabled() view returns(bool)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetBootstrapModeDisabled() ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getBootstrapModeDisabled")
}

// UnpackGetBootstrapModeDisabled is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf54746e4.
//
// Solidity: function getBootstrapModeDisabled() view returns(bool)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetBootstrapModeDisabled(data []byte) (bool, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getBootstrapModeDisabled", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGetMemberAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe992c817.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberAt(uint256 _index) view returns(address)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberAt(index *big.Int) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberAt", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe992c817.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberAt(uint256 _index) view returns(address)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberAt(index *big.Int) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberAt", index)
}

// UnpackGetMemberAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe992c817.
//
// Solidity: function getMemberAt(uint256 _index) view returns(address)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberAt(data []byte) (common.Address, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberAt", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetMemberCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x997072f7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberCount() view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberCount() []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x997072f7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberCount() view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberCount() ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberCount")
}

// UnpackGetMemberCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x997072f7.
//
// Solidity: function getMemberCount() view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberCount(data []byte) (*big.Int, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetMemberID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e2d45d1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberID(address _nodeAddress) view returns(string)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberID(nodeAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberID", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e2d45d1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberID(address _nodeAddress) view returns(string)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberID(nodeAddress common.Address) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberID", nodeAddress)
}

// UnpackGetMemberID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e2d45d1.
//
// Solidity: function getMemberID(address _nodeAddress) view returns(string)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberID(data []byte) (string, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberID", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackGetMemberIsChallenged is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a1b2327.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberIsChallenged(address _nodeAddress) view returns(bool)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberIsChallenged(nodeAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberIsChallenged", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberIsChallenged is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a1b2327.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberIsChallenged(address _nodeAddress) view returns(bool)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberIsChallenged(nodeAddress common.Address) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberIsChallenged", nodeAddress)
}

// UnpackGetMemberIsChallenged is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7a1b2327.
//
// Solidity: function getMemberIsChallenged(address _nodeAddress) view returns(bool)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberIsChallenged(data []byte) (bool, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberIsChallenged", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGetMemberIsValid is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5dc33bdd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberIsValid(address _nodeAddress) view returns(bool)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberIsValid(nodeAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberIsValid", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberIsValid is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5dc33bdd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberIsValid(address _nodeAddress) view returns(bool)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberIsValid(nodeAddress common.Address) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberIsValid", nodeAddress)
}

// UnpackGetMemberIsValid is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5dc33bdd.
//
// Solidity: function getMemberIsValid(address _nodeAddress) view returns(bool)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberIsValid(data []byte) (bool, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberIsValid", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackGetMemberJoinedTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5987956e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberJoinedTime(address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberJoinedTime(nodeAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberJoinedTime", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberJoinedTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5987956e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberJoinedTime(address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberJoinedTime(nodeAddress common.Address) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberJoinedTime", nodeAddress)
}

// UnpackGetMemberJoinedTime is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5987956e.
//
// Solidity: function getMemberJoinedTime(address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberJoinedTime(data []byte) (*big.Int, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberJoinedTime", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetMemberLastProposalTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x51553095.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberLastProposalTime(address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberLastProposalTime(nodeAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberLastProposalTime", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberLastProposalTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x51553095.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberLastProposalTime(address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberLastProposalTime(nodeAddress common.Address) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberLastProposalTime", nodeAddress)
}

// UnpackGetMemberLastProposalTime is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x51553095.
//
// Solidity: function getMemberLastProposalTime(address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberLastProposalTime(data []byte) (*big.Int, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberLastProposalTime", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetMemberMinRequired is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc1eb7b2a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberMinRequired() pure returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberMinRequired() []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberMinRequired")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberMinRequired is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc1eb7b2a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberMinRequired() pure returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberMinRequired() ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberMinRequired")
}

// UnpackGetMemberMinRequired is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc1eb7b2a.
//
// Solidity: function getMemberMinRequired() pure returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberMinRequired(data []byte) (*big.Int, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberMinRequired", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetMemberProposalExecutedTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x803f94e3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberProposalExecutedTime(string _proposalType, address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberProposalExecutedTime(proposalType string, nodeAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberProposalExecutedTime", proposalType, nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberProposalExecutedTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x803f94e3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberProposalExecutedTime(string _proposalType, address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberProposalExecutedTime(proposalType string, nodeAddress common.Address) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberProposalExecutedTime", proposalType, nodeAddress)
}

// UnpackGetMemberProposalExecutedTime is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x803f94e3.
//
// Solidity: function getMemberProposalExecutedTime(string _proposalType, address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberProposalExecutedTime(data []byte) (*big.Int, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberProposalExecutedTime", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetMemberQuorumVotesRequired is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43906fea.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberQuorumVotesRequired() view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberQuorumVotesRequired() []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberQuorumVotesRequired")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberQuorumVotesRequired is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43906fea.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberQuorumVotesRequired() view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberQuorumVotesRequired() ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberQuorumVotesRequired")
}

// UnpackGetMemberQuorumVotesRequired is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x43906fea.
//
// Solidity: function getMemberQuorumVotesRequired() view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberQuorumVotesRequired(data []byte) (*big.Int, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberQuorumVotesRequired", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetMemberRPLBondAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x03c86bbd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberRPLBondAmount(address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberRPLBondAmount(nodeAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberRPLBondAmount", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberRPLBondAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x03c86bbd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberRPLBondAmount(address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberRPLBondAmount(nodeAddress common.Address) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberRPLBondAmount", nodeAddress)
}

// UnpackGetMemberRPLBondAmount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x03c86bbd.
//
// Solidity: function getMemberRPLBondAmount(address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberRPLBondAmount(data []byte) (*big.Int, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberRPLBondAmount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetMemberUnbondedValidatorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d89846e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberUnbondedValidatorCount(address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberUnbondedValidatorCount(nodeAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberUnbondedValidatorCount", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberUnbondedValidatorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d89846e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberUnbondedValidatorCount(address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberUnbondedValidatorCount(nodeAddress common.Address) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberUnbondedValidatorCount", nodeAddress)
}

// UnpackGetMemberUnbondedValidatorCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7d89846e.
//
// Solidity: function getMemberUnbondedValidatorCount(address _nodeAddress) view returns(uint256)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberUnbondedValidatorCount(data []byte) (*big.Int, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberUnbondedValidatorCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetMemberUrl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8840fe0c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMemberUrl(address _nodeAddress) view returns(string)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackGetMemberUrl(nodeAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("getMemberUrl", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMemberUrl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8840fe0c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMemberUrl(address _nodeAddress) view returns(string)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackGetMemberUrl(nodeAddress common.Address) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("getMemberUrl", nodeAddress)
}

// UnpackGetMemberUrl is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8840fe0c.
//
// Solidity: function getMemberUrl(address _nodeAddress) view returns(string)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackGetMemberUrl(data []byte) (string, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("getMemberUrl", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackIncrementMemberUnbondedValidatorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x72043ec4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function incrementMemberUnbondedValidatorCount(address _nodeAddress) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackIncrementMemberUnbondedValidatorCount(nodeAddress common.Address) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("incrementMemberUnbondedValidatorCount", nodeAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIncrementMemberUnbondedValidatorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x72043ec4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function incrementMemberUnbondedValidatorCount(address _nodeAddress) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackIncrementMemberUnbondedValidatorCount(nodeAddress common.Address) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("incrementMemberUnbondedValidatorCount", nodeAddress)
}

// PackMemberJoinRequired is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x636e3e41.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function memberJoinRequired(string _id, string _url) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackMemberJoinRequired(id string, url string) []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("memberJoinRequired", id, url)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMemberJoinRequired is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x636e3e41.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function memberJoinRequired(string _id, string _url) returns()
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackMemberJoinRequired(id string, url string) ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("memberJoinRequired", id, url)
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(uint8)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) PackVersion() []byte {
	enc, err := rocketDaoNodeTrusted.abi.Pack("version")
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
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) TryPackVersion() ([]byte, error) {
	return rocketDaoNodeTrusted.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (rocketDaoNodeTrusted *RocketDaoNodeTrusted) UnpackVersion(data []byte) (uint8, error) {
	out, err := rocketDaoNodeTrusted.abi.Unpack("version", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}
