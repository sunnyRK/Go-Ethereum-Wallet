// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CryptoContractABI is the input ABI used to generate the binding from.
const CryptoContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipeintAddress\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_recipeintAddress\",\"type\":\"address\"}],\"name\":\"getBalanceOfAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipeintAddress\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"moneyTrnasfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenName\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// CryptoContractBin is the compiled bytecode used for deploying new contracts.
const CryptoContractBin = `0x6060604052341561000f57600080fd5b60405161021e38038061021e833981016040528080515050506101e7806100376000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312065fe0811461006857806369ca02dd1461008d57806370a08231146100a4578063936710bf146100c3578063c5fb7b76146100e257600080fd5b341561007357600080fd5b61007b610106565b60405190815260200160405180910390f35b61007b600160a060020a0360043516602435610114565b34156100af57600080fd5b61007b600160a060020a036004351661013a565b34156100ce57600080fd5b61007b600160a060020a036004351661014c565b34156100ed57600080fd5b610104600160a060020a0360043516602435610167565b005b600160a060020a0330163190565b50600160a060020a03908116600090815260208190526040902080543401905533163190565b60006020819052908152604090205481565b600160a060020a031660009081526020819052604090205490565b600160a060020a03821681156108fc0282604051600060405180830381858888f19350505050151561019857600080fd5b600160a060020a03909116600090815260208190526040902080549190910390555600a165627a7a72305820659bce97dfde6b6df470ea3c7d6fe9264eec9b2614ca72e608fcd9073a3ac96b0029`

// DeployCryptoContract deploys a new Ethereum contract, binding an instance of CryptoContract to it.
func DeployCryptoContract(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenName string) (common.Address, *types.Transaction, *CryptoContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CryptoContractBin), backend, _tokenName)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CryptoContract{CryptoContractCaller: CryptoContractCaller{contract: contract}, CryptoContractTransactor: CryptoContractTransactor{contract: contract}, CryptoContractFilterer: CryptoContractFilterer{contract: contract}}, nil
}

// CryptoContract is an auto generated Go binding around an Ethereum contract.
type CryptoContract struct {
	CryptoContractCaller     // Read-only binding to the contract
	CryptoContractTransactor // Write-only binding to the contract
	CryptoContractFilterer   // Log filterer for contract events
}

// CryptoContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CryptoContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CryptoContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CryptoContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CryptoContractSession struct {
	Contract     *CryptoContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CryptoContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CryptoContractCallerSession struct {
	Contract *CryptoContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CryptoContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CryptoContractTransactorSession struct {
	Contract     *CryptoContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CryptoContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CryptoContractRaw struct {
	Contract *CryptoContract // Generic contract binding to access the raw methods on
}

// CryptoContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CryptoContractCallerRaw struct {
	Contract *CryptoContractCaller // Generic read-only contract binding to access the raw methods on
}

// CryptoContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CryptoContractTransactorRaw struct {
	Contract *CryptoContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCryptoContract creates a new instance of CryptoContract, bound to a specific deployed contract.
func NewCryptoContract(address common.Address, backend bind.ContractBackend) (*CryptoContract, error) {
	contract, err := bindCryptoContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CryptoContract{CryptoContractCaller: CryptoContractCaller{contract: contract}, CryptoContractTransactor: CryptoContractTransactor{contract: contract}, CryptoContractFilterer: CryptoContractFilterer{contract: contract}}, nil
}

// NewCryptoContractCaller creates a new read-only instance of CryptoContract, bound to a specific deployed contract.
func NewCryptoContractCaller(address common.Address, caller bind.ContractCaller) (*CryptoContractCaller, error) {
	contract, err := bindCryptoContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoContractCaller{contract: contract}, nil
}

// NewCryptoContractTransactor creates a new write-only instance of CryptoContract, bound to a specific deployed contract.
func NewCryptoContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CryptoContractTransactor, error) {
	contract, err := bindCryptoContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoContractTransactor{contract: contract}, nil
}

// NewCryptoContractFilterer creates a new log filterer instance of CryptoContract, bound to a specific deployed contract.
func NewCryptoContractFilterer(address common.Address, filterer bind.ContractFilterer) (*CryptoContractFilterer, error) {
	contract, err := bindCryptoContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CryptoContractFilterer{contract: contract}, nil
}

// bindCryptoContract binds a generic wrapper to an already deployed contract.
func bindCryptoContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoContract *CryptoContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CryptoContract.Contract.CryptoContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoContract *CryptoContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoContract.Contract.CryptoContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoContract *CryptoContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoContract.Contract.CryptoContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoContract *CryptoContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CryptoContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoContract *CryptoContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoContract *CryptoContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_CryptoContract *CryptoContractCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CryptoContract.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_CryptoContract *CryptoContractSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CryptoContract.Contract.BalanceOf(&_CryptoContract.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_CryptoContract *CryptoContractCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CryptoContract.Contract.BalanceOf(&_CryptoContract.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_CryptoContract *CryptoContractCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CryptoContract.contract.Call(opts, out, "getBalance")
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_CryptoContract *CryptoContractSession) GetBalance() (*big.Int, error) {
	return _CryptoContract.Contract.GetBalance(&_CryptoContract.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_CryptoContract *CryptoContractCallerSession) GetBalance() (*big.Int, error) {
	return _CryptoContract.Contract.GetBalance(&_CryptoContract.CallOpts)
}

// GetBalanceOfAccount is a free data retrieval call binding the contract method 0x936710bf.
//
// Solidity: function getBalanceOfAccount(address _recipeintAddress) constant returns(uint256)
func (_CryptoContract *CryptoContractCaller) GetBalanceOfAccount(opts *bind.CallOpts, _recipeintAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CryptoContract.contract.Call(opts, out, "getBalanceOfAccount", _recipeintAddress)
	return *ret0, err
}

// GetBalanceOfAccount is a free data retrieval call binding the contract method 0x936710bf.
//
// Solidity: function getBalanceOfAccount(address _recipeintAddress) constant returns(uint256)
func (_CryptoContract *CryptoContractSession) GetBalanceOfAccount(_recipeintAddress common.Address) (*big.Int, error) {
	return _CryptoContract.Contract.GetBalanceOfAccount(&_CryptoContract.CallOpts, _recipeintAddress)
}

// GetBalanceOfAccount is a free data retrieval call binding the contract method 0x936710bf.
//
// Solidity: function getBalanceOfAccount(address _recipeintAddress) constant returns(uint256)
func (_CryptoContract *CryptoContractCallerSession) GetBalanceOfAccount(_recipeintAddress common.Address) (*big.Int, error) {
	return _CryptoContract.Contract.GetBalanceOfAccount(&_CryptoContract.CallOpts, _recipeintAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address _recipeintAddress, uint256 value) returns(uint256)
func (_CryptoContract *CryptoContractTransactor) Transfer(opts *bind.TransactOpts, _recipeintAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _CryptoContract.contract.Transact(opts, "Transfer", _recipeintAddress, value)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address _recipeintAddress, uint256 value) returns(uint256)
func (_CryptoContract *CryptoContractSession) Transfer(_recipeintAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _CryptoContract.Contract.Transfer(&_CryptoContract.TransactOpts, _recipeintAddress, value)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address _recipeintAddress, uint256 value) returns(uint256)
func (_CryptoContract *CryptoContractTransactorSession) Transfer(_recipeintAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _CryptoContract.Contract.Transfer(&_CryptoContract.TransactOpts, _recipeintAddress, value)
}

// MoneyTrnasfer is a paid mutator transaction binding the contract method 0xc5fb7b76.
//
// Solidity: function moneyTrnasfer(address _recipeintAddress, uint256 value) returns()
func (_CryptoContract *CryptoContractTransactor) MoneyTrnasfer(opts *bind.TransactOpts, _recipeintAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _CryptoContract.contract.Transact(opts, "moneyTrnasfer", _recipeintAddress, value)
}

// MoneyTrnasfer is a paid mutator transaction binding the contract method 0xc5fb7b76.
//
// Solidity: function moneyTrnasfer(address _recipeintAddress, uint256 value) returns()
func (_CryptoContract *CryptoContractSession) MoneyTrnasfer(_recipeintAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _CryptoContract.Contract.MoneyTrnasfer(&_CryptoContract.TransactOpts, _recipeintAddress, value)
}

// MoneyTrnasfer is a paid mutator transaction binding the contract method 0xc5fb7b76.
//
// Solidity: function moneyTrnasfer(address _recipeintAddress, uint256 value) returns()
func (_CryptoContract *CryptoContractTransactorSession) MoneyTrnasfer(_recipeintAddress common.Address, value *big.Int) (*types.Transaction, error) {
	return _CryptoContract.Contract.MoneyTrnasfer(&_CryptoContract.TransactOpts, _recipeintAddress, value)
}
