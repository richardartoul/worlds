// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SingleMessage

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SingleMessageABI is the input ABI used to generate the binding from.
const SingleMessageABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"priceInWei\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMessage\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"message\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"amountInWei\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialMessage\",\"type\":\"string\"},{\"name\":\"initialPriceInWei\",\"type\":\"uint256\"},{\"name\":\"maxLengthArg\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPriceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"payer\",\"type\":\"address\"}],\"name\":\"MessageSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// SingleMessage is an auto generated Go binding around an Ethereum contract.
type SingleMessage struct {
	SingleMessageCaller     // Read-only binding to the contract
	SingleMessageTransactor // Write-only binding to the contract
}

// SingleMessageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SingleMessageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleMessageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SingleMessageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleMessageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SingleMessageSession struct {
	Contract     *SingleMessage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SingleMessageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SingleMessageCallerSession struct {
	Contract *SingleMessageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SingleMessageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SingleMessageTransactorSession struct {
	Contract     *SingleMessageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SingleMessageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SingleMessageRaw struct {
	Contract *SingleMessage // Generic contract binding to access the raw methods on
}

// SingleMessageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SingleMessageCallerRaw struct {
	Contract *SingleMessageCaller // Generic read-only contract binding to access the raw methods on
}

// SingleMessageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SingleMessageTransactorRaw struct {
	Contract *SingleMessageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSingleMessage creates a new instance of SingleMessage, bound to a specific deployed contract.
func NewSingleMessage(address common.Address, backend bind.ContractBackend) (*SingleMessage, error) {
	contract, err := bindSingleMessage(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SingleMessage{SingleMessageCaller: SingleMessageCaller{contract: contract}, SingleMessageTransactor: SingleMessageTransactor{contract: contract}}, nil
}

// NewSingleMessageCaller creates a new read-only instance of SingleMessage, bound to a specific deployed contract.
func NewSingleMessageCaller(address common.Address, caller bind.ContractCaller) (*SingleMessageCaller, error) {
	contract, err := bindSingleMessage(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SingleMessageCaller{contract: contract}, nil
}

// NewSingleMessageTransactor creates a new write-only instance of SingleMessage, bound to a specific deployed contract.
func NewSingleMessageTransactor(address common.Address, transactor bind.ContractTransactor) (*SingleMessageTransactor, error) {
	contract, err := bindSingleMessage(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SingleMessageTransactor{contract: contract}, nil
}

// bindSingleMessage binds a generic wrapper to an already deployed contract.
func bindSingleMessage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SingleMessageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SingleMessage *SingleMessageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SingleMessage.Contract.SingleMessageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SingleMessage *SingleMessageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleMessage.Contract.SingleMessageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SingleMessage *SingleMessageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingleMessage.Contract.SingleMessageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SingleMessage *SingleMessageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SingleMessage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SingleMessage *SingleMessageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleMessage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SingleMessage *SingleMessageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingleMessage.Contract.contract.Transact(opts, method, params...)
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_SingleMessage *SingleMessageCaller) MaxLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SingleMessage.contract.Call(opts, out, "maxLength")
	return *ret0, err
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_SingleMessage *SingleMessageSession) MaxLength() (*big.Int, error) {
	return _SingleMessage.Contract.MaxLength(&_SingleMessage.CallOpts)
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_SingleMessage *SingleMessageCallerSession) MaxLength() (*big.Int, error) {
	return _SingleMessage.Contract.MaxLength(&_SingleMessage.CallOpts)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() constant returns(string)
func (_SingleMessage *SingleMessageCaller) Message(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SingleMessage.contract.Call(opts, out, "message")
	return *ret0, err
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() constant returns(string)
func (_SingleMessage *SingleMessageSession) Message() (string, error) {
	return _SingleMessage.Contract.Message(&_SingleMessage.CallOpts)
}

// Message is a free data retrieval call binding the contract method 0xe21f37ce.
//
// Solidity: function message() constant returns(string)
func (_SingleMessage *SingleMessageCallerSession) Message() (string, error) {
	return _SingleMessage.Contract.Message(&_SingleMessage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SingleMessage *SingleMessageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SingleMessage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SingleMessage *SingleMessageSession) Owner() (common.Address, error) {
	return _SingleMessage.Contract.Owner(&_SingleMessage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SingleMessage *SingleMessageCallerSession) Owner() (common.Address, error) {
	return _SingleMessage.Contract.Owner(&_SingleMessage.CallOpts)
}

// PriceInWei is a free data retrieval call binding the contract method 0x3c8da588.
//
// Solidity: function priceInWei() constant returns(uint256)
func (_SingleMessage *SingleMessageCaller) PriceInWei(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SingleMessage.contract.Call(opts, out, "priceInWei")
	return *ret0, err
}

// PriceInWei is a free data retrieval call binding the contract method 0x3c8da588.
//
// Solidity: function priceInWei() constant returns(uint256)
func (_SingleMessage *SingleMessageSession) PriceInWei() (*big.Int, error) {
	return _SingleMessage.Contract.PriceInWei(&_SingleMessage.CallOpts)
}

// PriceInWei is a free data retrieval call binding the contract method 0x3c8da588.
//
// Solidity: function priceInWei() constant returns(uint256)
func (_SingleMessage *SingleMessageCallerSession) PriceInWei() (*big.Int, error) {
	return _SingleMessage.Contract.PriceInWei(&_SingleMessage.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(newMessage string) returns()
func (_SingleMessage *SingleMessageTransactor) Set(opts *bind.TransactOpts, newMessage string) (*types.Transaction, error) {
	return _SingleMessage.contract.Transact(opts, "set", newMessage)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(newMessage string) returns()
func (_SingleMessage *SingleMessageSession) Set(newMessage string) (*types.Transaction, error) {
	return _SingleMessage.Contract.Set(&_SingleMessage.TransactOpts, newMessage)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(newMessage string) returns()
func (_SingleMessage *SingleMessageTransactorSession) Set(newMessage string) (*types.Transaction, error) {
	return _SingleMessage.Contract.Set(&_SingleMessage.TransactOpts, newMessage)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SingleMessage *SingleMessageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SingleMessage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SingleMessage *SingleMessageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SingleMessage.Contract.TransferOwnership(&_SingleMessage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SingleMessage *SingleMessageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SingleMessage.Contract.TransferOwnership(&_SingleMessage.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(destination address, amountInWei uint256) returns()
func (_SingleMessage *SingleMessageTransactor) Withdraw(opts *bind.TransactOpts, destination common.Address, amountInWei *big.Int) (*types.Transaction, error) {
	return _SingleMessage.contract.Transact(opts, "withdraw", destination, amountInWei)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(destination address, amountInWei uint256) returns()
func (_SingleMessage *SingleMessageSession) Withdraw(destination common.Address, amountInWei *big.Int) (*types.Transaction, error) {
	return _SingleMessage.Contract.Withdraw(&_SingleMessage.TransactOpts, destination, amountInWei)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(destination address, amountInWei uint256) returns()
func (_SingleMessage *SingleMessageTransactorSession) Withdraw(destination common.Address, amountInWei *big.Int) (*types.Transaction, error) {
	return _SingleMessage.Contract.Withdraw(&_SingleMessage.TransactOpts, destination, amountInWei)
}
