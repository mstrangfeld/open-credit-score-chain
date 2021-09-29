// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditscoredb

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

// CreditScoreDBMetaData contains all meta data concerning the CreditScoreDB contract.
var CreditScoreDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_score\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_scoreHash\",\"type\":\"string\"}],\"name\":\"createRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"invalidateRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"score\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nonce\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"scoreHash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf7e214f": "authority()",
		"b797dcb5": "createRecord(string,string,string,string)",
		"5bbe446a": "invalidateRecord(uint256)",
		"8da5cb5b": "owner()",
		"34461067": "records(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161092238038061092283398101604081905261002f9161005d565b600080546001600160a01b039092166001600160a01b0319928316179055600180549091163317905561008b565b60006020828403121561006e578081fd5b81516001600160a01b0381168114610084578182fd5b9392505050565b6108888061009a6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063344610671461005c5780635bbe446a1461008b5780638da5cb5b146100a0578063b797dcb5146100cb578063bf7e214f146100de575b600080fd5b61006f61006a36600461073d565b6100f1565b60405161008297969594939291906107a0565b60405180910390f35b61009e61009936600461073d565b61036c565b005b6000546100b3906001600160a01b031681565b6040516001600160a01b039091168152602001610082565b61009e6100d936600461067e565b6103b1565b6001546100b3906001600160a01b031681565b6002818154811061010157600080fd5b600091825260209091206006909102018054600182015460028301805460ff841695506101009093046001600160a01b031693919261013f90610817565b80601f016020809104026020016040519081016040528092919081815260200182805461016b90610817565b80156101b85780601f1061018d576101008083540402835291602001916101b8565b820191906000526020600020905b81548152906001019060200180831161019b57829003601f168201915b5050505050908060030180546101cd90610817565b80601f01602080910402602001604051908101604052809291908181526020018280546101f990610817565b80156102465780601f1061021b57610100808354040283529160200191610246565b820191906000526020600020905b81548152906001019060200180831161022957829003601f168201915b50505050509080600401805461025b90610817565b80601f016020809104026020016040519081016040528092919081815260200182805461028790610817565b80156102d45780601f106102a9576101008083540402835291602001916102d4565b820191906000526020600020905b8154815290600101906020018083116102b757829003601f168201915b5050505050908060050180546102e990610817565b80601f016020809104026020016040519081016040528092919081815260200182805461031590610817565b80156103625780601f1061033757610100808354040283529160200191610362565b820191906000526020600020905b81548152906001019060200180831161034557829003601f168201915b5050505050905087565b60006002828154811061038f57634e487b7160e01b600052603260045260246000fd5b60009182526020909120600690910201805460ff191691151591909117905550565b60026040518060e00160405280600115158152602001336001600160a01b031681526020014281526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f8b0181900481028201810190925289815291810191908a908a9081908401838280828437600092019190915250505090825250604080516020601f890181900481028201810190925287815291810191908890889081908401838280828437600092019190915250505090825250604080516020601f870181900481028201810190925285815291810191908690869081908401838280828437600092018290525093909452505083546001818101865594825260209182902084516006909202018054838601516001600160a01b031661010002610100600160a81b031993151560ff1990921691909117929092169190911781556040840151948101949094556060830151805193949361053d93506002850192919091019061059e565b506080820151805161055991600384019160209091019061059e565b5060a0820151805161057591600484019160209091019061059e565b5060c0820151805161059191600584019160209091019061059e565b5050505050505050505050565b8280546105aa90610817565b90600052602060002090601f0160209004810192826105cc5760008555610612565b82601f106105e557805160ff1916838001178555610612565b82800160010185558215610612579182015b828111156106125782518255916020019190600101906105f7565b5061061e929150610622565b5090565b5b8082111561061e5760008155600101610623565b60008083601f840112610648578182fd5b50813567ffffffffffffffff81111561065f578182fd5b60208301915083602082850101111561067757600080fd5b9250929050565b6000806000806000806000806080898b031215610699578384fd5b883567ffffffffffffffff808211156106b0578586fd5b6106bc8c838d01610637565b909a50985060208b01359150808211156106d4578586fd5b6106e08c838d01610637565b909850965060408b01359150808211156106f8578586fd5b6107048c838d01610637565b909650945060608b013591508082111561071c578384fd5b506107298b828c01610637565b999c989b5096995094979396929594505050565b60006020828403121561074e578081fd5b5035919050565b60008151808452815b8181101561077a5760208185018101518683018201520161075e565b8181111561078b5782602083870101525b50601f01601f19169290920160200192915050565b6000881515825260018060a01b038816602083015286604083015260e060608301526107cf60e0830187610755565b82810360808401526107e18187610755565b905082810360a08401526107f58186610755565b905082810360c08401526108098185610755565b9a9950505050505050505050565b60028104600182168061082b57607f821691505b6020821081141561084c57634e487b7160e01b600052602260045260246000fd5b5091905056fea2646970667358221220d8525549cafdd3469d6f103583ae3df6deae6a085c490c2b2dd8299ed039c9ac64736f6c63430008020033",
}

// CreditScoreDBABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditScoreDBMetaData.ABI instead.
var CreditScoreDBABI = CreditScoreDBMetaData.ABI

// Deprecated: Use CreditScoreDBMetaData.Sigs instead.
// CreditScoreDBFuncSigs maps the 4-byte function signature to its string representation.
var CreditScoreDBFuncSigs = CreditScoreDBMetaData.Sigs

// CreditScoreDBBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CreditScoreDBMetaData.Bin instead.
var CreditScoreDBBin = CreditScoreDBMetaData.Bin

// DeployCreditScoreDB deploys a new Ethereum contract, binding an instance of CreditScoreDB to it.
func DeployCreditScoreDB(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *CreditScoreDB, error) {
	parsed, err := CreditScoreDBMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CreditScoreDBBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CreditScoreDB{CreditScoreDBCaller: CreditScoreDBCaller{contract: contract}, CreditScoreDBTransactor: CreditScoreDBTransactor{contract: contract}, CreditScoreDBFilterer: CreditScoreDBFilterer{contract: contract}}, nil
}

// CreditScoreDB is an auto generated Go binding around an Ethereum contract.
type CreditScoreDB struct {
	CreditScoreDBCaller     // Read-only binding to the contract
	CreditScoreDBTransactor // Write-only binding to the contract
	CreditScoreDBFilterer   // Log filterer for contract events
}

// CreditScoreDBCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditScoreDBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditScoreDBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditScoreDBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditScoreDBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditScoreDBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditScoreDBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditScoreDBSession struct {
	Contract     *CreditScoreDB    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditScoreDBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditScoreDBCallerSession struct {
	Contract *CreditScoreDBCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CreditScoreDBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditScoreDBTransactorSession struct {
	Contract     *CreditScoreDBTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CreditScoreDBRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditScoreDBRaw struct {
	Contract *CreditScoreDB // Generic contract binding to access the raw methods on
}

// CreditScoreDBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditScoreDBCallerRaw struct {
	Contract *CreditScoreDBCaller // Generic read-only contract binding to access the raw methods on
}

// CreditScoreDBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditScoreDBTransactorRaw struct {
	Contract *CreditScoreDBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditScoreDB creates a new instance of CreditScoreDB, bound to a specific deployed contract.
func NewCreditScoreDB(address common.Address, backend bind.ContractBackend) (*CreditScoreDB, error) {
	contract, err := bindCreditScoreDB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditScoreDB{CreditScoreDBCaller: CreditScoreDBCaller{contract: contract}, CreditScoreDBTransactor: CreditScoreDBTransactor{contract: contract}, CreditScoreDBFilterer: CreditScoreDBFilterer{contract: contract}}, nil
}

// NewCreditScoreDBCaller creates a new read-only instance of CreditScoreDB, bound to a specific deployed contract.
func NewCreditScoreDBCaller(address common.Address, caller bind.ContractCaller) (*CreditScoreDBCaller, error) {
	contract, err := bindCreditScoreDB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditScoreDBCaller{contract: contract}, nil
}

// NewCreditScoreDBTransactor creates a new write-only instance of CreditScoreDB, bound to a specific deployed contract.
func NewCreditScoreDBTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditScoreDBTransactor, error) {
	contract, err := bindCreditScoreDB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditScoreDBTransactor{contract: contract}, nil
}

// NewCreditScoreDBFilterer creates a new log filterer instance of CreditScoreDB, bound to a specific deployed contract.
func NewCreditScoreDBFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditScoreDBFilterer, error) {
	contract, err := bindCreditScoreDB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditScoreDBFilterer{contract: contract}, nil
}

// bindCreditScoreDB binds a generic wrapper to an already deployed contract.
func bindCreditScoreDB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditScoreDBABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditScoreDB *CreditScoreDBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditScoreDB.Contract.CreditScoreDBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditScoreDB *CreditScoreDBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditScoreDB.Contract.CreditScoreDBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditScoreDB *CreditScoreDBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditScoreDB.Contract.CreditScoreDBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditScoreDB *CreditScoreDBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditScoreDB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditScoreDB *CreditScoreDBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditScoreDB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditScoreDB *CreditScoreDBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditScoreDB.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_CreditScoreDB *CreditScoreDBCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditScoreDB.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_CreditScoreDB *CreditScoreDBSession) Authority() (common.Address, error) {
	return _CreditScoreDB.Contract.Authority(&_CreditScoreDB.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_CreditScoreDB *CreditScoreDBCallerSession) Authority() (common.Address, error) {
	return _CreditScoreDB.Contract.Authority(&_CreditScoreDB.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreditScoreDB *CreditScoreDBCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditScoreDB.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreditScoreDB *CreditScoreDBSession) Owner() (common.Address, error) {
	return _CreditScoreDB.Contract.Owner(&_CreditScoreDB.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreditScoreDB *CreditScoreDBCallerSession) Owner() (common.Address, error) {
	return _CreditScoreDB.Contract.Owner(&_CreditScoreDB.CallOpts)
}

// Records is a free data retrieval call binding the contract method 0x34461067.
//
// Solidity: function records(uint256 ) view returns(bool valid, address issuer, uint256 timestamp, string reason, string score, string nonce, string scoreHash)
func (_CreditScoreDB *CreditScoreDBCaller) Records(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Valid     bool
	Issuer    common.Address
	Timestamp *big.Int
	Reason    string
	Score     string
	Nonce     string
	ScoreHash string
}, error) {
	var out []interface{}
	err := _CreditScoreDB.contract.Call(opts, &out, "records", arg0)

	outstruct := new(struct {
		Valid     bool
		Issuer    common.Address
		Timestamp *big.Int
		Reason    string
		Score     string
		Nonce     string
		ScoreHash string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Valid = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Issuer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Reason = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Score = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.ScoreHash = *abi.ConvertType(out[6], new(string)).(*string)

	return *outstruct, err

}

// Records is a free data retrieval call binding the contract method 0x34461067.
//
// Solidity: function records(uint256 ) view returns(bool valid, address issuer, uint256 timestamp, string reason, string score, string nonce, string scoreHash)
func (_CreditScoreDB *CreditScoreDBSession) Records(arg0 *big.Int) (struct {
	Valid     bool
	Issuer    common.Address
	Timestamp *big.Int
	Reason    string
	Score     string
	Nonce     string
	ScoreHash string
}, error) {
	return _CreditScoreDB.Contract.Records(&_CreditScoreDB.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x34461067.
//
// Solidity: function records(uint256 ) view returns(bool valid, address issuer, uint256 timestamp, string reason, string score, string nonce, string scoreHash)
func (_CreditScoreDB *CreditScoreDBCallerSession) Records(arg0 *big.Int) (struct {
	Valid     bool
	Issuer    common.Address
	Timestamp *big.Int
	Reason    string
	Score     string
	Nonce     string
	ScoreHash string
}, error) {
	return _CreditScoreDB.Contract.Records(&_CreditScoreDB.CallOpts, arg0)
}

// CreateRecord is a paid mutator transaction binding the contract method 0xb797dcb5.
//
// Solidity: function createRecord(string _reason, string _score, string _nonce, string _scoreHash) returns()
func (_CreditScoreDB *CreditScoreDBTransactor) CreateRecord(opts *bind.TransactOpts, _reason string, _score string, _nonce string, _scoreHash string) (*types.Transaction, error) {
	return _CreditScoreDB.contract.Transact(opts, "createRecord", _reason, _score, _nonce, _scoreHash)
}

// CreateRecord is a paid mutator transaction binding the contract method 0xb797dcb5.
//
// Solidity: function createRecord(string _reason, string _score, string _nonce, string _scoreHash) returns()
func (_CreditScoreDB *CreditScoreDBSession) CreateRecord(_reason string, _score string, _nonce string, _scoreHash string) (*types.Transaction, error) {
	return _CreditScoreDB.Contract.CreateRecord(&_CreditScoreDB.TransactOpts, _reason, _score, _nonce, _scoreHash)
}

// CreateRecord is a paid mutator transaction binding the contract method 0xb797dcb5.
//
// Solidity: function createRecord(string _reason, string _score, string _nonce, string _scoreHash) returns()
func (_CreditScoreDB *CreditScoreDBTransactorSession) CreateRecord(_reason string, _score string, _nonce string, _scoreHash string) (*types.Transaction, error) {
	return _CreditScoreDB.Contract.CreateRecord(&_CreditScoreDB.TransactOpts, _reason, _score, _nonce, _scoreHash)
}

// InvalidateRecord is a paid mutator transaction binding the contract method 0x5bbe446a.
//
// Solidity: function invalidateRecord(uint256 _index) returns()
func (_CreditScoreDB *CreditScoreDBTransactor) InvalidateRecord(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _CreditScoreDB.contract.Transact(opts, "invalidateRecord", _index)
}

// InvalidateRecord is a paid mutator transaction binding the contract method 0x5bbe446a.
//
// Solidity: function invalidateRecord(uint256 _index) returns()
func (_CreditScoreDB *CreditScoreDBSession) InvalidateRecord(_index *big.Int) (*types.Transaction, error) {
	return _CreditScoreDB.Contract.InvalidateRecord(&_CreditScoreDB.TransactOpts, _index)
}

// InvalidateRecord is a paid mutator transaction binding the contract method 0x5bbe446a.
//
// Solidity: function invalidateRecord(uint256 _index) returns()
func (_CreditScoreDB *CreditScoreDBTransactorSession) InvalidateRecord(_index *big.Int) (*types.Transaction, error) {
	return _CreditScoreDB.Contract.InvalidateRecord(&_CreditScoreDB.TransactOpts, _index)
}
