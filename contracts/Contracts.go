// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lock_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_required_bond\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getMessage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batch\",\"type\":\"bytes\"}],\"name\":\"newBatch\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_hash\",\"type\":\"bytes\"}],\"name\":\"readHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_hash\",\"type\":\"bytes\"}],\"name\":\"readHashRLP\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"required_bond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(string)
func (_Contracts *ContractsCaller) GetMessage(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getMessage")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(string)
func (_Contracts *ContractsSession) GetMessage() (string, error) {
	return _Contracts.Contract.GetMessage(&_Contracts.CallOpts)
}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(string)
func (_Contracts *ContractsCallerSession) GetMessage() (string, error) {
	return _Contracts.Contract.GetMessage(&_Contracts.CallOpts)
}

// GetStateRoot is a free data retrieval call binding the contract method 0xcd605a1a.
//
// Solidity: function getStateRoot() view returns(bytes32)
func (_Contracts *ContractsCaller) GetStateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getStateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStateRoot is a free data retrieval call binding the contract method 0xcd605a1a.
//
// Solidity: function getStateRoot() view returns(bytes32)
func (_Contracts *ContractsSession) GetStateRoot() ([32]byte, error) {
	return _Contracts.Contract.GetStateRoot(&_Contracts.CallOpts)
}

// GetStateRoot is a free data retrieval call binding the contract method 0xcd605a1a.
//
// Solidity: function getStateRoot() view returns(bytes32)
func (_Contracts *ContractsCallerSession) GetStateRoot() ([32]byte, error) {
	return _Contracts.Contract.GetStateRoot(&_Contracts.CallOpts)
}

// GetToAddress is a free data retrieval call binding the contract method 0xe411842f.
//
// Solidity: function getToAddress() view returns(address)
func (_Contracts *ContractsCaller) GetToAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getToAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToAddress is a free data retrieval call binding the contract method 0xe411842f.
//
// Solidity: function getToAddress() view returns(address)
func (_Contracts *ContractsSession) GetToAddress() (common.Address, error) {
	return _Contracts.Contract.GetToAddress(&_Contracts.CallOpts)
}

// GetToAddress is a free data retrieval call binding the contract method 0xe411842f.
//
// Solidity: function getToAddress() view returns(address)
func (_Contracts *ContractsCallerSession) GetToAddress() (common.Address, error) {
	return _Contracts.Contract.GetToAddress(&_Contracts.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x480bb7c4.
//
// Solidity: function lock_time() view returns(uint256)
func (_Contracts *ContractsCaller) LockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "lock_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0x480bb7c4.
//
// Solidity: function lock_time() view returns(uint256)
func (_Contracts *ContractsSession) LockTime() (*big.Int, error) {
	return _Contracts.Contract.LockTime(&_Contracts.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x480bb7c4.
//
// Solidity: function lock_time() view returns(uint256)
func (_Contracts *ContractsCallerSession) LockTime() (*big.Int, error) {
	return _Contracts.Contract.LockTime(&_Contracts.CallOpts)
}

// RequiredBond is a free data retrieval call binding the contract method 0xb2055400.
//
// Solidity: function required_bond() view returns(uint256)
func (_Contracts *ContractsCaller) RequiredBond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "required_bond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredBond is a free data retrieval call binding the contract method 0xb2055400.
//
// Solidity: function required_bond() view returns(uint256)
func (_Contracts *ContractsSession) RequiredBond() (*big.Int, error) {
	return _Contracts.Contract.RequiredBond(&_Contracts.CallOpts)
}

// RequiredBond is a free data retrieval call binding the contract method 0xb2055400.
//
// Solidity: function required_bond() view returns(uint256)
func (_Contracts *ContractsCallerSession) RequiredBond() (*big.Int, error) {
	return _Contracts.Contract.RequiredBond(&_Contracts.CallOpts)
}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(bytes32)
func (_Contracts *ContractsCaller) StateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "stateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(bytes32)
func (_Contracts *ContractsSession) StateRoot() ([32]byte, error) {
	return _Contracts.Contract.StateRoot(&_Contracts.CallOpts)
}

// StateRoot is a free data retrieval call binding the contract method 0x9588eca2.
//
// Solidity: function stateRoot() view returns(bytes32)
func (_Contracts *ContractsCallerSession) StateRoot() ([32]byte, error) {
	return _Contracts.Contract.StateRoot(&_Contracts.CallOpts)
}

// NewBatch is a paid mutator transaction binding the contract method 0xdbcf9bd2.
//
// Solidity: function newBatch(bytes _batch) returns(string)
func (_Contracts *ContractsTransactor) NewBatch(opts *bind.TransactOpts, _batch []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "newBatch", _batch)
}

// NewBatch is a paid mutator transaction binding the contract method 0xdbcf9bd2.
//
// Solidity: function newBatch(bytes _batch) returns(string)
func (_Contracts *ContractsSession) NewBatch(_batch []byte) (*types.Transaction, error) {
	return _Contracts.Contract.NewBatch(&_Contracts.TransactOpts, _batch)
}

// NewBatch is a paid mutator transaction binding the contract method 0xdbcf9bd2.
//
// Solidity: function newBatch(bytes _batch) returns(string)
func (_Contracts *ContractsTransactorSession) NewBatch(_batch []byte) (*types.Transaction, error) {
	return _Contracts.Contract.NewBatch(&_Contracts.TransactOpts, _batch)
}

// ReadHash is a paid mutator transaction binding the contract method 0xb05146b4.
//
// Solidity: function readHash(bytes _hash) returns(string)
func (_Contracts *ContractsTransactor) ReadHash(opts *bind.TransactOpts, _hash []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "readHash", _hash)
}

// ReadHash is a paid mutator transaction binding the contract method 0xb05146b4.
//
// Solidity: function readHash(bytes _hash) returns(string)
func (_Contracts *ContractsSession) ReadHash(_hash []byte) (*types.Transaction, error) {
	return _Contracts.Contract.ReadHash(&_Contracts.TransactOpts, _hash)
}

// ReadHash is a paid mutator transaction binding the contract method 0xb05146b4.
//
// Solidity: function readHash(bytes _hash) returns(string)
func (_Contracts *ContractsTransactorSession) ReadHash(_hash []byte) (*types.Transaction, error) {
	return _Contracts.Contract.ReadHash(&_Contracts.TransactOpts, _hash)
}

// ReadHashRLP is a paid mutator transaction binding the contract method 0x2b585a54.
//
// Solidity: function readHashRLP(bytes _hash) returns(string)
func (_Contracts *ContractsTransactor) ReadHashRLP(opts *bind.TransactOpts, _hash []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "readHashRLP", _hash)
}

// ReadHashRLP is a paid mutator transaction binding the contract method 0x2b585a54.
//
// Solidity: function readHashRLP(bytes _hash) returns(string)
func (_Contracts *ContractsSession) ReadHashRLP(_hash []byte) (*types.Transaction, error) {
	return _Contracts.Contract.ReadHashRLP(&_Contracts.TransactOpts, _hash)
}

// ReadHashRLP is a paid mutator transaction binding the contract method 0x2b585a54.
//
// Solidity: function readHashRLP(bytes _hash) returns(string)
func (_Contracts *ContractsTransactorSession) ReadHashRLP(_hash []byte) (*types.Transaction, error) {
	return _Contracts.Contract.ReadHashRLP(&_Contracts.TransactOpts, _hash)
}
