// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package erc

import (
	"math/big"
	"strings"

	"github.com/ShyftNetwork/go-empyrean/accounts/abi"
	"github.com/ShyftNetwork/go-empyrean/accounts/abi/bind"
	"github.com/ShyftNetwork/go-empyrean/common"
	"github.com/ShyftNetwork/go-empyrean/core/types"
)

// ErcABI is the input ABI used to generate the binding from.
const ErcABI = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"tokens","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"tokens","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"tokenOwner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"tokens","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"tokenOwner","type":"address"},{"name":"spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"tokens","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"tokenOwner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"tokens","type":"uint256"}],"name":"Approval","type":"event"}]`

// Erc is an auto generated Go binding around an Ethereum contract.
type Erc struct {
	ErcCaller     // Read-only binding to the contract
	ErcTransactor // Write-only binding to the contract
}

// ErcCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErcSession struct {
	Contract     *Erc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErcCallerSession struct {
	Contract *ErcCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ErcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErcTransactorSession struct {
	Contract     *ErcTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErcRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErcRaw struct {
	Contract *Erc // Generic contract binding to access the raw methods on
}

// ErcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErcCallerRaw struct {
	Contract *ErcCaller // Generic read-only contract binding to access the raw methods on
}

// ErcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErcTransactorRaw struct {
	Contract *ErcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc creates a new instance of Erc, bound to a specific deployed contract.
func NewErc(address common.Address, backend bind.ContractBackend) (*Erc, error) {
	contract, err := bindErc(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &Erc{ErcCaller: ErcCaller{contract: contract}, ErcTransactor: ErcTransactor{contract: contract}}, nil
}

// NewErcCaller creates a new read-only instance of Erc, bound to a specific deployed contract.
func NewErcCaller(address common.Address, caller bind.ContractCaller) (*ErcCaller, error) {
	contract, err := bindErc(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ErcCaller{contract: contract}, nil
}

// NewErcTransactor creates a new write-only instance of Erc, bound to a specific deployed contract.
func NewErcTransactor(address common.Address, transactor bind.ContractTransactor) (*ErcTransactor, error) {
	contract, err := bindErc(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ErcTransactor{contract: contract}, nil
}

// bindErc binds a generic wrapper to an already deployed contract.
func bindErc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ErcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, nil), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc *ErcRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc.Contract.ErcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc *ErcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc.Contract.ErcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc *ErcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc.Contract.ErcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc *ErcCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc *ErcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc *ErcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_Erc *ErcCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Erc.contract.Call(opts, out, "allowance", tokenOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_Erc *ErcSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc.Contract.Allowance(&_Erc.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(tokenOwner address, spender address) constant returns(remaining uint256)
func (_Erc *ErcCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc.Contract.Allowance(&_Erc.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_Erc *ErcCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Erc.contract.Call(opts, out, "balanceOf", tokenOwner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_Erc *ErcSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Erc.Contract.BalanceOf(&_Erc.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(tokenOwner address) constant returns(balance uint256)
func (_Erc *ErcCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Erc.Contract.BalanceOf(&_Erc.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Erc *ErcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Erc.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Erc *ErcSession) Decimals() (uint8, error) {
	return _Erc.Contract.Decimals(&_Erc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Erc *ErcCallerSession) Decimals() (uint8, error) {
	return _Erc.Contract.Decimals(&_Erc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Erc *ErcCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Erc.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Erc *ErcSession) Name() (string, error) {
	return _Erc.Contract.Name(&_Erc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Erc *ErcCallerSession) Name() (string, error) {
	return _Erc.Contract.Name(&_Erc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Erc *ErcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Erc.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Erc *ErcSession) Symbol() (string, error) {
	return _Erc.Contract.Symbol(&_Erc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Erc *ErcCallerSession) Symbol() (string, error) {
	return _Erc.Contract.Symbol(&_Erc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Erc *ErcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Erc.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Erc *ErcSession) TotalSupply() (*big.Int, error) {
	return _Erc.Contract.TotalSupply(&_Erc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Erc *ErcCallerSession) TotalSupply() (*big.Int, error) {
	return _Erc.Contract.TotalSupply(&_Erc.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_Erc *ErcTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_Erc *ErcSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc.Contract.Approve(&_Erc.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(success bool)
func (_Erc *ErcTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc.Contract.Approve(&_Erc.TransactOpts, spender, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_Erc *ErcTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_Erc *ErcSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc.Contract.Transfer(&_Erc.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(success bool)
func (_Erc *ErcTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc.Contract.Transfer(&_Erc.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_Erc *ErcTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_Erc *ErcSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc.Contract.TransferFrom(&_Erc.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(success bool)
func (_Erc *ErcTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Erc.Contract.TransferFrom(&_Erc.TransactOpts, from, to, tokens)
}
