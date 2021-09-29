// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package authority

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

// AuthorityMetaData contains all meta data concerning the Authority contract.
var AuthorityMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"authorizeRecordWriter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_taxID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"createDB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_taxID\",\"type\":\"string\"}],\"name\":\"getDB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isAllowedToCreateRecords\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recordWriters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"unauthorizeRecordWriter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6bca1725": "authorizeRecordWriter(address)",
		"75badbe3": "createDB(string,address)",
		"3f43ef6f": "getDB(string)",
		"6182dbe3": "isAllowedToCreateRecords(address)",
		"fbcb50e4": "recordWriters(address)",
		"ebf0c717": "root()",
		"a1df1ccb": "unauthorizeRecordWriter(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50336100725760405162461bcd60e51b815260206004820152602760248201527f526f6f7420616464726573732063616e6e6f7420626520746865207a65726f206044820152666164647265737360c81b606482015260840160405180910390fd5b600080546001600160a01b03191633179055610376806100936000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806375badbe31161005b57806375badbe314610109578063a1df1ccb1461011c578063ebf0c7171461014b578063fbcb50e41461015e5761007d565b80633f43ef6f146100825780636182dbe3146100b25780636bca1725146100d5575b600080fd5b61009561009036600461029e565b610181565b6040516001600160a01b0390911681526020015b60405180910390f35b6100c56100c036600461027d565b6101b6565b60405190151581526020016100a9565b6101076100e336600461027d565b6001600160a01b03166000908152600260205260409020805460ff19166001179055565b005b6101076101173660046102de565b6101d8565b61010761012a36600461027d565b6001600160a01b03166000908152600260205260409020805460ff19169055565b600054610095906001600160a01b031681565b6100c561016c36600461027d565b60026020526000908152604090205460ff1681565b600060018383604051610195929190610330565b908152604051908190036020019020546001600160a01b0316905092915050565b6001600160a01b03811660009081526002602052604090205460ff165b919050565b80600184846040516101eb929190610330565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b0319909216919091179055505050565b80356001600160a01b03811681146101d357600080fd5b60008083601f840112610247578182fd5b50813567ffffffffffffffff81111561025e578182fd5b60208301915083602082850101111561027657600080fd5b9250929050565b60006020828403121561028e578081fd5b6102978261021f565b9392505050565b600080602083850312156102b0578081fd5b823567ffffffffffffffff8111156102c6578182fd5b6102d285828601610236565b90969095509350505050565b6000806000604084860312156102f2578081fd5b833567ffffffffffffffff811115610308578182fd5b61031486828701610236565b909450925061032790506020850161021f565b90509250925092565b600082848337910190815291905056fea264697066735822122001ce259d3fe3cdd4865394283963f8530fb16cc75686c509066cdf844e37060064736f6c63430008020033",
}

// AuthorityABI is the input ABI used to generate the binding from.
// Deprecated: Use AuthorityMetaData.ABI instead.
var AuthorityABI = AuthorityMetaData.ABI

// Deprecated: Use AuthorityMetaData.Sigs instead.
// AuthorityFuncSigs maps the 4-byte function signature to its string representation.
var AuthorityFuncSigs = AuthorityMetaData.Sigs

// AuthorityBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuthorityMetaData.Bin instead.
var AuthorityBin = AuthorityMetaData.Bin

// DeployAuthority deploys a new Ethereum contract, binding an instance of Authority to it.
func DeployAuthority(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Authority, error) {
	parsed, err := AuthorityMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuthorityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Authority{AuthorityCaller: AuthorityCaller{contract: contract}, AuthorityTransactor: AuthorityTransactor{contract: contract}, AuthorityFilterer: AuthorityFilterer{contract: contract}}, nil
}

// Authority is an auto generated Go binding around an Ethereum contract.
type Authority struct {
	AuthorityCaller     // Read-only binding to the contract
	AuthorityTransactor // Write-only binding to the contract
	AuthorityFilterer   // Log filterer for contract events
}

// AuthorityCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthorityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthorityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthorityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthoritySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthoritySession struct {
	Contract     *Authority        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthorityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthorityCallerSession struct {
	Contract *AuthorityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AuthorityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthorityTransactorSession struct {
	Contract     *AuthorityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AuthorityRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthorityRaw struct {
	Contract *Authority // Generic contract binding to access the raw methods on
}

// AuthorityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthorityCallerRaw struct {
	Contract *AuthorityCaller // Generic read-only contract binding to access the raw methods on
}

// AuthorityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthorityTransactorRaw struct {
	Contract *AuthorityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthority creates a new instance of Authority, bound to a specific deployed contract.
func NewAuthority(address common.Address, backend bind.ContractBackend) (*Authority, error) {
	contract, err := bindAuthority(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Authority{AuthorityCaller: AuthorityCaller{contract: contract}, AuthorityTransactor: AuthorityTransactor{contract: contract}, AuthorityFilterer: AuthorityFilterer{contract: contract}}, nil
}

// NewAuthorityCaller creates a new read-only instance of Authority, bound to a specific deployed contract.
func NewAuthorityCaller(address common.Address, caller bind.ContractCaller) (*AuthorityCaller, error) {
	contract, err := bindAuthority(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorityCaller{contract: contract}, nil
}

// NewAuthorityTransactor creates a new write-only instance of Authority, bound to a specific deployed contract.
func NewAuthorityTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthorityTransactor, error) {
	contract, err := bindAuthority(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorityTransactor{contract: contract}, nil
}

// NewAuthorityFilterer creates a new log filterer instance of Authority, bound to a specific deployed contract.
func NewAuthorityFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthorityFilterer, error) {
	contract, err := bindAuthority(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthorityFilterer{contract: contract}, nil
}

// bindAuthority binds a generic wrapper to an already deployed contract.
func bindAuthority(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authority *AuthorityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Authority.Contract.AuthorityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authority *AuthorityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authority.Contract.AuthorityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authority *AuthorityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authority.Contract.AuthorityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authority *AuthorityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Authority.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authority *AuthorityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authority.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authority *AuthorityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authority.Contract.contract.Transact(opts, method, params...)
}

// GetDB is a free data retrieval call binding the contract method 0x3f43ef6f.
//
// Solidity: function getDB(string _taxID) view returns(address)
func (_Authority *AuthorityCaller) GetDB(opts *bind.CallOpts, _taxID string) (common.Address, error) {
	var out []interface{}
	err := _Authority.contract.Call(opts, &out, "getDB", _taxID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDB is a free data retrieval call binding the contract method 0x3f43ef6f.
//
// Solidity: function getDB(string _taxID) view returns(address)
func (_Authority *AuthoritySession) GetDB(_taxID string) (common.Address, error) {
	return _Authority.Contract.GetDB(&_Authority.CallOpts, _taxID)
}

// GetDB is a free data retrieval call binding the contract method 0x3f43ef6f.
//
// Solidity: function getDB(string _taxID) view returns(address)
func (_Authority *AuthorityCallerSession) GetDB(_taxID string) (common.Address, error) {
	return _Authority.Contract.GetDB(&_Authority.CallOpts, _taxID)
}

// IsAllowedToCreateRecords is a free data retrieval call binding the contract method 0x6182dbe3.
//
// Solidity: function isAllowedToCreateRecords(address _addr) view returns(bool)
func (_Authority *AuthorityCaller) IsAllowedToCreateRecords(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Authority.contract.Call(opts, &out, "isAllowedToCreateRecords", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedToCreateRecords is a free data retrieval call binding the contract method 0x6182dbe3.
//
// Solidity: function isAllowedToCreateRecords(address _addr) view returns(bool)
func (_Authority *AuthoritySession) IsAllowedToCreateRecords(_addr common.Address) (bool, error) {
	return _Authority.Contract.IsAllowedToCreateRecords(&_Authority.CallOpts, _addr)
}

// IsAllowedToCreateRecords is a free data retrieval call binding the contract method 0x6182dbe3.
//
// Solidity: function isAllowedToCreateRecords(address _addr) view returns(bool)
func (_Authority *AuthorityCallerSession) IsAllowedToCreateRecords(_addr common.Address) (bool, error) {
	return _Authority.Contract.IsAllowedToCreateRecords(&_Authority.CallOpts, _addr)
}

// RecordWriters is a free data retrieval call binding the contract method 0xfbcb50e4.
//
// Solidity: function recordWriters(address ) view returns(bool)
func (_Authority *AuthorityCaller) RecordWriters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Authority.contract.Call(opts, &out, "recordWriters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RecordWriters is a free data retrieval call binding the contract method 0xfbcb50e4.
//
// Solidity: function recordWriters(address ) view returns(bool)
func (_Authority *AuthoritySession) RecordWriters(arg0 common.Address) (bool, error) {
	return _Authority.Contract.RecordWriters(&_Authority.CallOpts, arg0)
}

// RecordWriters is a free data retrieval call binding the contract method 0xfbcb50e4.
//
// Solidity: function recordWriters(address ) view returns(bool)
func (_Authority *AuthorityCallerSession) RecordWriters(arg0 common.Address) (bool, error) {
	return _Authority.Contract.RecordWriters(&_Authority.CallOpts, arg0)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_Authority *AuthorityCaller) Root(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Authority.contract.Call(opts, &out, "root")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_Authority *AuthoritySession) Root() (common.Address, error) {
	return _Authority.Contract.Root(&_Authority.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_Authority *AuthorityCallerSession) Root() (common.Address, error) {
	return _Authority.Contract.Root(&_Authority.CallOpts)
}

// AuthorizeRecordWriter is a paid mutator transaction binding the contract method 0x6bca1725.
//
// Solidity: function authorizeRecordWriter(address _addr) returns()
func (_Authority *AuthorityTransactor) AuthorizeRecordWriter(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "authorizeRecordWriter", _addr)
}

// AuthorizeRecordWriter is a paid mutator transaction binding the contract method 0x6bca1725.
//
// Solidity: function authorizeRecordWriter(address _addr) returns()
func (_Authority *AuthoritySession) AuthorizeRecordWriter(_addr common.Address) (*types.Transaction, error) {
	return _Authority.Contract.AuthorizeRecordWriter(&_Authority.TransactOpts, _addr)
}

// AuthorizeRecordWriter is a paid mutator transaction binding the contract method 0x6bca1725.
//
// Solidity: function authorizeRecordWriter(address _addr) returns()
func (_Authority *AuthorityTransactorSession) AuthorizeRecordWriter(_addr common.Address) (*types.Transaction, error) {
	return _Authority.Contract.AuthorizeRecordWriter(&_Authority.TransactOpts, _addr)
}

// CreateDB is a paid mutator transaction binding the contract method 0x75badbe3.
//
// Solidity: function createDB(string _taxID, address _addr) returns()
func (_Authority *AuthorityTransactor) CreateDB(opts *bind.TransactOpts, _taxID string, _addr common.Address) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "createDB", _taxID, _addr)
}

// CreateDB is a paid mutator transaction binding the contract method 0x75badbe3.
//
// Solidity: function createDB(string _taxID, address _addr) returns()
func (_Authority *AuthoritySession) CreateDB(_taxID string, _addr common.Address) (*types.Transaction, error) {
	return _Authority.Contract.CreateDB(&_Authority.TransactOpts, _taxID, _addr)
}

// CreateDB is a paid mutator transaction binding the contract method 0x75badbe3.
//
// Solidity: function createDB(string _taxID, address _addr) returns()
func (_Authority *AuthorityTransactorSession) CreateDB(_taxID string, _addr common.Address) (*types.Transaction, error) {
	return _Authority.Contract.CreateDB(&_Authority.TransactOpts, _taxID, _addr)
}

// UnauthorizeRecordWriter is a paid mutator transaction binding the contract method 0xa1df1ccb.
//
// Solidity: function unauthorizeRecordWriter(address _addr) returns()
func (_Authority *AuthorityTransactor) UnauthorizeRecordWriter(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "unauthorizeRecordWriter", _addr)
}

// UnauthorizeRecordWriter is a paid mutator transaction binding the contract method 0xa1df1ccb.
//
// Solidity: function unauthorizeRecordWriter(address _addr) returns()
func (_Authority *AuthoritySession) UnauthorizeRecordWriter(_addr common.Address) (*types.Transaction, error) {
	return _Authority.Contract.UnauthorizeRecordWriter(&_Authority.TransactOpts, _addr)
}

// UnauthorizeRecordWriter is a paid mutator transaction binding the contract method 0xa1df1ccb.
//
// Solidity: function unauthorizeRecordWriter(address _addr) returns()
func (_Authority *AuthorityTransactorSession) UnauthorizeRecordWriter(_addr common.Address) (*types.Transaction, error) {
	return _Authority.Contract.UnauthorizeRecordWriter(&_Authority.TransactOpts, _addr)
}
