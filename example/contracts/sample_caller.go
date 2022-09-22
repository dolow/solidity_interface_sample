// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sample_caller

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SampleCallerMetaData contains all meta data concerning the SampleCaller contract.
var SampleCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setSampleAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"setSampleValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSampleValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SampleCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use SampleCallerMetaData.ABI instead.
var SampleCallerABI = SampleCallerMetaData.ABI

// SampleCaller is an auto generated Go binding around an Ethereum contract.
type SampleCaller struct {
	SampleCallerCaller     // Read-only binding to the contract
	SampleCallerTransactor // Write-only binding to the contract
	SampleCallerFilterer   // Log filterer for contract events
}

// SampleCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SampleCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleCallerSession struct {
	Contract     *SampleCaller     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SampleCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleCallerCallerSession struct {
	Contract *SampleCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SampleCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleCallerTransactorSession struct {
	Contract     *SampleCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SampleCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleCallerRaw struct {
	Contract *SampleCaller // Generic contract binding to access the raw methods on
}

// SampleCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleCallerCallerRaw struct {
	Contract *SampleCallerCaller // Generic read-only contract binding to access the raw methods on
}

// SampleCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleCallerTransactorRaw struct {
	Contract *SampleCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampleCaller creates a new instance of SampleCaller, bound to a specific deployed contract.
func NewSampleCaller(address common.Address, backend bind.ContractBackend) (*SampleCaller, error) {
	contract, err := bindSampleCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SampleCaller{SampleCallerCaller: SampleCallerCaller{contract: contract}, SampleCallerTransactor: SampleCallerTransactor{contract: contract}, SampleCallerFilterer: SampleCallerFilterer{contract: contract}}, nil
}

// NewSampleCallerCaller creates a new read-only instance of SampleCaller, bound to a specific deployed contract.
func NewSampleCallerCaller(address common.Address, caller bind.ContractCaller) (*SampleCallerCaller, error) {
	contract, err := bindSampleCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SampleCallerCaller{contract: contract}, nil
}

// NewSampleCallerTransactor creates a new write-only instance of SampleCaller, bound to a specific deployed contract.
func NewSampleCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleCallerTransactor, error) {
	contract, err := bindSampleCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SampleCallerTransactor{contract: contract}, nil
}

// NewSampleCallerFilterer creates a new log filterer instance of SampleCaller, bound to a specific deployed contract.
func NewSampleCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*SampleCallerFilterer, error) {
	contract, err := bindSampleCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SampleCallerFilterer{contract: contract}, nil
}

// bindSampleCaller binds a generic wrapper to an already deployed contract.
func bindSampleCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleCallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleCaller *SampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SampleCaller.Contract.SampleCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleCaller *SampleCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCaller.Contract.SampleCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleCaller *SampleCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleCaller.Contract.SampleCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleCaller *SampleCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SampleCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleCaller *SampleCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleCaller *SampleCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleCaller.Contract.contract.Transact(opts, method, params...)
}

// GetSampleValue is a free data retrieval call binding the contract method 0xd29cd7ac.
//
// Solidity: function getSampleValue() view returns(uint256)
func (_SampleCaller *SampleCallerCaller) GetSampleValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SampleCaller.contract.Call(opts, &out, "getSampleValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSampleValue is a free data retrieval call binding the contract method 0xd29cd7ac.
//
// Solidity: function getSampleValue() view returns(uint256)
func (_SampleCaller *SampleCallerSession) GetSampleValue() (*big.Int, error) {
	return _SampleCaller.Contract.GetSampleValue(&_SampleCaller.CallOpts)
}

// GetSampleValue is a free data retrieval call binding the contract method 0xd29cd7ac.
//
// Solidity: function getSampleValue() view returns(uint256)
func (_SampleCaller *SampleCallerCallerSession) GetSampleValue() (*big.Int, error) {
	return _SampleCaller.Contract.GetSampleValue(&_SampleCaller.CallOpts)
}

// SetSampleAddress is a paid mutator transaction binding the contract method 0xbdbad548.
//
// Solidity: function setSampleAddress(address a) returns()
func (_SampleCaller *SampleCallerTransactor) SetSampleAddress(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _SampleCaller.contract.Transact(opts, "setSampleAddress", a)
}

// SetSampleAddress is a paid mutator transaction binding the contract method 0xbdbad548.
//
// Solidity: function setSampleAddress(address a) returns()
func (_SampleCaller *SampleCallerSession) SetSampleAddress(a common.Address) (*types.Transaction, error) {
	return _SampleCaller.Contract.SetSampleAddress(&_SampleCaller.TransactOpts, a)
}

// SetSampleAddress is a paid mutator transaction binding the contract method 0xbdbad548.
//
// Solidity: function setSampleAddress(address a) returns()
func (_SampleCaller *SampleCallerTransactorSession) SetSampleAddress(a common.Address) (*types.Transaction, error) {
	return _SampleCaller.Contract.SetSampleAddress(&_SampleCaller.TransactOpts, a)
}

// SetSampleValue is a paid mutator transaction binding the contract method 0xacf97b56.
//
// Solidity: function setSampleValue(uint256 i) returns()
func (_SampleCaller *SampleCallerTransactor) SetSampleValue(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _SampleCaller.contract.Transact(opts, "setSampleValue", i)
}

// SetSampleValue is a paid mutator transaction binding the contract method 0xacf97b56.
//
// Solidity: function setSampleValue(uint256 i) returns()
func (_SampleCaller *SampleCallerSession) SetSampleValue(i *big.Int) (*types.Transaction, error) {
	return _SampleCaller.Contract.SetSampleValue(&_SampleCaller.TransactOpts, i)
}

// SetSampleValue is a paid mutator transaction binding the contract method 0xacf97b56.
//
// Solidity: function setSampleValue(uint256 i) returns()
func (_SampleCaller *SampleCallerTransactorSession) SetSampleValue(i *big.Int) (*types.Transaction, error) {
	return _SampleCaller.Contract.SetSampleValue(&_SampleCaller.TransactOpts, i)
}
