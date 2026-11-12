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

// Multicall3Call is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call struct {
	Target   common.Address
	CallData []byte
}

// Multicall3Call3 is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call3 struct {
	Target       common.Address
	AllowFailure bool
	CallData     []byte
}

// Multicall3Call3Value is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call3Value struct {
	Target       common.Address
	AllowFailure bool
	Value        *big.Int
	CallData     []byte
}

// Multicall3Result is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Result struct {
	Success    bool
	ReturnData []byte
}

// Multicall3MetaData contains all meta data concerning the Multicall3 contract.
var Multicall3MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call3[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call3Value[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3Value\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasefee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"basefee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "Multicall3",
}

// Multicall3 is an auto generated Go binding around an Ethereum contract.
type Multicall3 struct {
	abi abi.ABI
}

// NewMulticall3 creates a new instance of Multicall3.
func NewMulticall3() *Multicall3 {
	parsed, err := Multicall3MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Multicall3{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Multicall3) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAggregate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x252dba42.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (multicall3 *Multicall3) PackAggregate(calls []Multicall3Call) []byte {
	enc, err := multicall3.abi.Pack("aggregate", calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAggregate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x252dba42.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (multicall3 *Multicall3) TryPackAggregate(calls []Multicall3Call) ([]byte, error) {
	return multicall3.abi.Pack("aggregate", calls)
}

// AggregateOutput serves as a container for the return parameters of contract
// method Aggregate.
type AggregateOutput struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}

// UnpackAggregate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (multicall3 *Multicall3) UnpackAggregate(data []byte) (AggregateOutput, error) {
	out, err := multicall3.abi.Unpack("aggregate", data)
	outstruct := new(AggregateOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.BlockNumber = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.ReturnData = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)
	return *outstruct, nil
}

// PackAggregate3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82ad56cb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (multicall3 *Multicall3) PackAggregate3(calls []Multicall3Call3) []byte {
	enc, err := multicall3.abi.Pack("aggregate3", calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAggregate3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82ad56cb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (multicall3 *Multicall3) TryPackAggregate3(calls []Multicall3Call3) ([]byte, error) {
	return multicall3.abi.Pack("aggregate3", calls)
}

// UnpackAggregate3 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (multicall3 *Multicall3) UnpackAggregate3(data []byte) ([]Multicall3Result, error) {
	out, err := multicall3.abi.Unpack("aggregate3", data)
	if err != nil {
		return *new([]Multicall3Result), err
	}
	out0 := *abi.ConvertType(out[0], new([]Multicall3Result)).(*[]Multicall3Result)
	return out0, nil
}

// PackAggregate3Value is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x174dea71.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (multicall3 *Multicall3) PackAggregate3Value(calls []Multicall3Call3Value) []byte {
	enc, err := multicall3.abi.Pack("aggregate3Value", calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAggregate3Value is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x174dea71.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (multicall3 *Multicall3) TryPackAggregate3Value(calls []Multicall3Call3Value) ([]byte, error) {
	return multicall3.abi.Pack("aggregate3Value", calls)
}

// UnpackAggregate3Value is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (multicall3 *Multicall3) UnpackAggregate3Value(data []byte) ([]Multicall3Result, error) {
	out, err := multicall3.abi.Unpack("aggregate3Value", data)
	if err != nil {
		return *new([]Multicall3Result), err
	}
	out0 := *abi.ConvertType(out[0], new([]Multicall3Result)).(*[]Multicall3Result)
	return out0, nil
}

// PackBlockAndAggregate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3077fa9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (multicall3 *Multicall3) PackBlockAndAggregate(calls []Multicall3Call) []byte {
	enc, err := multicall3.abi.Pack("blockAndAggregate", calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBlockAndAggregate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3077fa9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (multicall3 *Multicall3) TryPackBlockAndAggregate(calls []Multicall3Call) ([]byte, error) {
	return multicall3.abi.Pack("blockAndAggregate", calls)
}

// BlockAndAggregateOutput serves as a container for the return parameters of contract
// method BlockAndAggregate.
type BlockAndAggregateOutput struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall3Result
}

// UnpackBlockAndAggregate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (multicall3 *Multicall3) UnpackBlockAndAggregate(data []byte) (BlockAndAggregateOutput, error) {
	out, err := multicall3.abi.Unpack("blockAndAggregate", data)
	outstruct := new(BlockAndAggregateOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.BlockNumber = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ReturnData = *abi.ConvertType(out[2], new([]Multicall3Result)).(*[]Multicall3Result)
	return *outstruct, nil
}

// PackGetBasefee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e64a696.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (multicall3 *Multicall3) PackGetBasefee() []byte {
	enc, err := multicall3.abi.Pack("getBasefee")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBasefee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e64a696.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (multicall3 *Multicall3) TryPackGetBasefee() ([]byte, error) {
	return multicall3.abi.Pack("getBasefee")
}

// UnpackGetBasefee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (multicall3 *Multicall3) UnpackGetBasefee(data []byte) (*big.Int, error) {
	out, err := multicall3.abi.Unpack("getBasefee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetBlockHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xee82ac5e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (multicall3 *Multicall3) PackGetBlockHash(blockNumber *big.Int) []byte {
	enc, err := multicall3.abi.Pack("getBlockHash", blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBlockHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xee82ac5e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (multicall3 *Multicall3) TryPackGetBlockHash(blockNumber *big.Int) ([]byte, error) {
	return multicall3.abi.Pack("getBlockHash", blockNumber)
}

// UnpackGetBlockHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (multicall3 *Multicall3) UnpackGetBlockHash(data []byte) ([32]byte, error) {
	out, err := multicall3.abi.Unpack("getBlockHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetBlockNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42cbb15c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (multicall3 *Multicall3) PackGetBlockNumber() []byte {
	enc, err := multicall3.abi.Pack("getBlockNumber")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBlockNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42cbb15c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (multicall3 *Multicall3) TryPackGetBlockNumber() ([]byte, error) {
	return multicall3.abi.Pack("getBlockNumber")
}

// UnpackGetBlockNumber is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (multicall3 *Multicall3) UnpackGetBlockNumber(data []byte) (*big.Int, error) {
	out, err := multicall3.abi.Unpack("getBlockNumber", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetChainId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3408e470.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (multicall3 *Multicall3) PackGetChainId() []byte {
	enc, err := multicall3.abi.Pack("getChainId")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetChainId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3408e470.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (multicall3 *Multicall3) TryPackGetChainId() ([]byte, error) {
	return multicall3.abi.Pack("getChainId")
}

// UnpackGetChainId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (multicall3 *Multicall3) UnpackGetChainId(data []byte) (*big.Int, error) {
	out, err := multicall3.abi.Unpack("getChainId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetCurrentBlockCoinbase is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8b0574e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (multicall3 *Multicall3) PackGetCurrentBlockCoinbase() []byte {
	enc, err := multicall3.abi.Pack("getCurrentBlockCoinbase")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCurrentBlockCoinbase is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8b0574e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (multicall3 *Multicall3) TryPackGetCurrentBlockCoinbase() ([]byte, error) {
	return multicall3.abi.Pack("getCurrentBlockCoinbase")
}

// UnpackGetCurrentBlockCoinbase is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (multicall3 *Multicall3) UnpackGetCurrentBlockCoinbase(data []byte) (common.Address, error) {
	out, err := multicall3.abi.Unpack("getCurrentBlockCoinbase", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetCurrentBlockDifficulty is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x72425d9d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (multicall3 *Multicall3) PackGetCurrentBlockDifficulty() []byte {
	enc, err := multicall3.abi.Pack("getCurrentBlockDifficulty")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCurrentBlockDifficulty is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x72425d9d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (multicall3 *Multicall3) TryPackGetCurrentBlockDifficulty() ([]byte, error) {
	return multicall3.abi.Pack("getCurrentBlockDifficulty")
}

// UnpackGetCurrentBlockDifficulty is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (multicall3 *Multicall3) UnpackGetCurrentBlockDifficulty(data []byte) (*big.Int, error) {
	out, err := multicall3.abi.Unpack("getCurrentBlockDifficulty", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetCurrentBlockGasLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x86d516e8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (multicall3 *Multicall3) PackGetCurrentBlockGasLimit() []byte {
	enc, err := multicall3.abi.Pack("getCurrentBlockGasLimit")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCurrentBlockGasLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x86d516e8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (multicall3 *Multicall3) TryPackGetCurrentBlockGasLimit() ([]byte, error) {
	return multicall3.abi.Pack("getCurrentBlockGasLimit")
}

// UnpackGetCurrentBlockGasLimit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (multicall3 *Multicall3) UnpackGetCurrentBlockGasLimit(data []byte) (*big.Int, error) {
	out, err := multicall3.abi.Unpack("getCurrentBlockGasLimit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetCurrentBlockTimestamp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f28c97d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (multicall3 *Multicall3) PackGetCurrentBlockTimestamp() []byte {
	enc, err := multicall3.abi.Pack("getCurrentBlockTimestamp")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCurrentBlockTimestamp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f28c97d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (multicall3 *Multicall3) TryPackGetCurrentBlockTimestamp() ([]byte, error) {
	return multicall3.abi.Pack("getCurrentBlockTimestamp")
}

// UnpackGetCurrentBlockTimestamp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (multicall3 *Multicall3) UnpackGetCurrentBlockTimestamp(data []byte) (*big.Int, error) {
	out, err := multicall3.abi.Unpack("getCurrentBlockTimestamp", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetEthBalance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4d2301cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (multicall3 *Multicall3) PackGetEthBalance(addr common.Address) []byte {
	enc, err := multicall3.abi.Pack("getEthBalance", addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetEthBalance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4d2301cc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (multicall3 *Multicall3) TryPackGetEthBalance(addr common.Address) ([]byte, error) {
	return multicall3.abi.Pack("getEthBalance", addr)
}

// UnpackGetEthBalance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (multicall3 *Multicall3) UnpackGetEthBalance(data []byte) (*big.Int, error) {
	out, err := multicall3.abi.Unpack("getEthBalance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetLastBlockHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x27e86d6e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (multicall3 *Multicall3) PackGetLastBlockHash() []byte {
	enc, err := multicall3.abi.Pack("getLastBlockHash")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLastBlockHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x27e86d6e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (multicall3 *Multicall3) TryPackGetLastBlockHash() ([]byte, error) {
	return multicall3.abi.Pack("getLastBlockHash")
}

// UnpackGetLastBlockHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (multicall3 *Multicall3) UnpackGetLastBlockHash(data []byte) ([32]byte, error) {
	out, err := multicall3.abi.Unpack("getLastBlockHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackTryAggregate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbce38bd7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (multicall3 *Multicall3) PackTryAggregate(requireSuccess bool, calls []Multicall3Call) []byte {
	enc, err := multicall3.abi.Pack("tryAggregate", requireSuccess, calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTryAggregate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbce38bd7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (multicall3 *Multicall3) TryPackTryAggregate(requireSuccess bool, calls []Multicall3Call) ([]byte, error) {
	return multicall3.abi.Pack("tryAggregate", requireSuccess, calls)
}

// UnpackTryAggregate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (multicall3 *Multicall3) UnpackTryAggregate(data []byte) ([]Multicall3Result, error) {
	out, err := multicall3.abi.Unpack("tryAggregate", data)
	if err != nil {
		return *new([]Multicall3Result), err
	}
	out0 := *abi.ConvertType(out[0], new([]Multicall3Result)).(*[]Multicall3Result)
	return out0, nil
}

// PackTryBlockAndAggregate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x399542e9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (multicall3 *Multicall3) PackTryBlockAndAggregate(requireSuccess bool, calls []Multicall3Call) []byte {
	enc, err := multicall3.abi.Pack("tryBlockAndAggregate", requireSuccess, calls)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTryBlockAndAggregate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x399542e9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (multicall3 *Multicall3) TryPackTryBlockAndAggregate(requireSuccess bool, calls []Multicall3Call) ([]byte, error) {
	return multicall3.abi.Pack("tryBlockAndAggregate", requireSuccess, calls)
}

// TryBlockAndAggregateOutput serves as a container for the return parameters of contract
// method TryBlockAndAggregate.
type TryBlockAndAggregateOutput struct {
	BlockNumber *big.Int
	BlockHash   [32]byte
	ReturnData  []Multicall3Result
}

// UnpackTryBlockAndAggregate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (multicall3 *Multicall3) UnpackTryBlockAndAggregate(data []byte) (TryBlockAndAggregateOutput, error) {
	out, err := multicall3.abi.Unpack("tryBlockAndAggregate", data)
	outstruct := new(TryBlockAndAggregateOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.BlockNumber = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.BlockHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ReturnData = *abi.ConvertType(out[2], new([]Multicall3Result)).(*[]Multicall3Result)
	return *outstruct, nil
}
