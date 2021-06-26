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
const ContractsABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lock_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_required_bond\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"Fraud_Proved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"Invalid_Proof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"New_Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"New_withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"aggregators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bond\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_batch_submitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_batch_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batch\",\"type\":\"bytes\"}],\"name\":\"newBatch\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prev_stateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_lastBatch\",\"type\":\"bytes\"}],\"name\":\"prove_fraud\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remaining_proof_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"required_bond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"valid_stateRoots\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// Aggregators is a free data retrieval call binding the contract method 0x112cdab9.
//
// Solidity: function aggregators(address ) view returns(address)
func (_Contracts *ContractsCaller) Aggregators(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "aggregators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregators is a free data retrieval call binding the contract method 0x112cdab9.
//
// Solidity: function aggregators(address ) view returns(address)
func (_Contracts *ContractsSession) Aggregators(arg0 common.Address) (common.Address, error) {
	return _Contracts.Contract.Aggregators(&_Contracts.CallOpts, arg0)
}

// Aggregators is a free data retrieval call binding the contract method 0x112cdab9.
//
// Solidity: function aggregators(address ) view returns(address)
func (_Contracts *ContractsCallerSession) Aggregators(arg0 common.Address) (common.Address, error) {
	return _Contracts.Contract.Aggregators(&_Contracts.CallOpts, arg0)
}

// LastBatchSubmitter is a free data retrieval call binding the contract method 0xc357b60b.
//
// Solidity: function last_batch_submitter() view returns(address)
func (_Contracts *ContractsCaller) LastBatchSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "last_batch_submitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastBatchSubmitter is a free data retrieval call binding the contract method 0xc357b60b.
//
// Solidity: function last_batch_submitter() view returns(address)
func (_Contracts *ContractsSession) LastBatchSubmitter() (common.Address, error) {
	return _Contracts.Contract.LastBatchSubmitter(&_Contracts.CallOpts)
}

// LastBatchSubmitter is a free data retrieval call binding the contract method 0xc357b60b.
//
// Solidity: function last_batch_submitter() view returns(address)
func (_Contracts *ContractsCallerSession) LastBatchSubmitter() (common.Address, error) {
	return _Contracts.Contract.LastBatchSubmitter(&_Contracts.CallOpts)
}

// LastBatchTime is a free data retrieval call binding the contract method 0xb194d0ea.
//
// Solidity: function last_batch_time() view returns(uint256)
func (_Contracts *ContractsCaller) LastBatchTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "last_batch_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBatchTime is a free data retrieval call binding the contract method 0xb194d0ea.
//
// Solidity: function last_batch_time() view returns(uint256)
func (_Contracts *ContractsSession) LastBatchTime() (*big.Int, error) {
	return _Contracts.Contract.LastBatchTime(&_Contracts.CallOpts)
}

// LastBatchTime is a free data retrieval call binding the contract method 0xb194d0ea.
//
// Solidity: function last_batch_time() view returns(uint256)
func (_Contracts *ContractsCallerSession) LastBatchTime() (*big.Int, error) {
	return _Contracts.Contract.LastBatchTime(&_Contracts.CallOpts)
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

// PrevStateRoot is a free data retrieval call binding the contract method 0xf53b28aa.
//
// Solidity: function prev_stateRoot() view returns(bytes32)
func (_Contracts *ContractsCaller) PrevStateRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "prev_stateRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PrevStateRoot is a free data retrieval call binding the contract method 0xf53b28aa.
//
// Solidity: function prev_stateRoot() view returns(bytes32)
func (_Contracts *ContractsSession) PrevStateRoot() ([32]byte, error) {
	return _Contracts.Contract.PrevStateRoot(&_Contracts.CallOpts)
}

// PrevStateRoot is a free data retrieval call binding the contract method 0xf53b28aa.
//
// Solidity: function prev_stateRoot() view returns(bytes32)
func (_Contracts *ContractsCallerSession) PrevStateRoot() ([32]byte, error) {
	return _Contracts.Contract.PrevStateRoot(&_Contracts.CallOpts)
}

// RemainingProofTime is a free data retrieval call binding the contract method 0x57e76d72.
//
// Solidity: function remaining_proof_time() view returns(uint256)
func (_Contracts *ContractsCaller) RemainingProofTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "remaining_proof_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingProofTime is a free data retrieval call binding the contract method 0x57e76d72.
//
// Solidity: function remaining_proof_time() view returns(uint256)
func (_Contracts *ContractsSession) RemainingProofTime() (*big.Int, error) {
	return _Contracts.Contract.RemainingProofTime(&_Contracts.CallOpts)
}

// RemainingProofTime is a free data retrieval call binding the contract method 0x57e76d72.
//
// Solidity: function remaining_proof_time() view returns(uint256)
func (_Contracts *ContractsCallerSession) RemainingProofTime() (*big.Int, error) {
	return _Contracts.Contract.RemainingProofTime(&_Contracts.CallOpts)
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

// ValidStateRoots is a free data retrieval call binding the contract method 0xe4481e9c.
//
// Solidity: function valid_stateRoots(bytes32 ) view returns(bool)
func (_Contracts *ContractsCaller) ValidStateRoots(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "valid_stateRoots", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidStateRoots is a free data retrieval call binding the contract method 0xe4481e9c.
//
// Solidity: function valid_stateRoots(bytes32 ) view returns(bool)
func (_Contracts *ContractsSession) ValidStateRoots(arg0 [32]byte) (bool, error) {
	return _Contracts.Contract.ValidStateRoots(&_Contracts.CallOpts, arg0)
}

// ValidStateRoots is a free data retrieval call binding the contract method 0xe4481e9c.
//
// Solidity: function valid_stateRoots(bytes32 ) view returns(bool)
func (_Contracts *ContractsCallerSession) ValidStateRoots(arg0 [32]byte) (bool, error) {
	return _Contracts.Contract.ValidStateRoots(&_Contracts.CallOpts, arg0)
}

// Bond is a paid mutator transaction binding the contract method 0x64c9ec6f.
//
// Solidity: function bond() payable returns()
func (_Contracts *ContractsTransactor) Bond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "bond")
}

// Bond is a paid mutator transaction binding the contract method 0x64c9ec6f.
//
// Solidity: function bond() payable returns()
func (_Contracts *ContractsSession) Bond() (*types.Transaction, error) {
	return _Contracts.Contract.Bond(&_Contracts.TransactOpts)
}

// Bond is a paid mutator transaction binding the contract method 0x64c9ec6f.
//
// Solidity: function bond() payable returns()
func (_Contracts *ContractsTransactorSession) Bond() (*types.Transaction, error) {
	return _Contracts.Contract.Bond(&_Contracts.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Contracts *ContractsTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Contracts *ContractsSession) Deposit() (*types.Transaction, error) {
	return _Contracts.Contract.Deposit(&_Contracts.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Contracts *ContractsTransactorSession) Deposit() (*types.Transaction, error) {
	return _Contracts.Contract.Deposit(&_Contracts.TransactOpts)
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

// ProveFraud is a paid mutator transaction binding the contract method 0x829e6914.
//
// Solidity: function prove_fraud(bytes _key, bytes _value, bytes _proof, bytes32 _root, bytes _lastBatch) returns()
func (_Contracts *ContractsTransactor) ProveFraud(opts *bind.TransactOpts, _key []byte, _value []byte, _proof []byte, _root [32]byte, _lastBatch []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "prove_fraud", _key, _value, _proof, _root, _lastBatch)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x829e6914.
//
// Solidity: function prove_fraud(bytes _key, bytes _value, bytes _proof, bytes32 _root, bytes _lastBatch) returns()
func (_Contracts *ContractsSession) ProveFraud(_key []byte, _value []byte, _proof []byte, _root [32]byte, _lastBatch []byte) (*types.Transaction, error) {
	return _Contracts.Contract.ProveFraud(&_Contracts.TransactOpts, _key, _value, _proof, _root, _lastBatch)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x829e6914.
//
// Solidity: function prove_fraud(bytes _key, bytes _value, bytes _proof, bytes32 _root, bytes _lastBatch) returns()
func (_Contracts *ContractsTransactorSession) ProveFraud(_key []byte, _value []byte, _proof []byte, _root [32]byte, _lastBatch []byte) (*types.Transaction, error) {
	return _Contracts.Contract.ProveFraud(&_Contracts.TransactOpts, _key, _value, _proof, _root, _lastBatch)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7a3e8408.
//
// Solidity: function withdraw(bytes _key, bytes _value, bytes _proof, bytes32 _root) returns()
func (_Contracts *ContractsTransactor) Withdraw(opts *bind.TransactOpts, _key []byte, _value []byte, _proof []byte, _root [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdraw", _key, _value, _proof, _root)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7a3e8408.
//
// Solidity: function withdraw(bytes _key, bytes _value, bytes _proof, bytes32 _root) returns()
func (_Contracts *ContractsSession) Withdraw(_key []byte, _value []byte, _proof []byte, _root [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, _key, _value, _proof, _root)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7a3e8408.
//
// Solidity: function withdraw(bytes _key, bytes _value, bytes _proof, bytes32 _root) returns()
func (_Contracts *ContractsTransactorSession) Withdraw(_key []byte, _value []byte, _proof []byte, _root [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, _key, _value, _proof, _root)
}

// ContractsFraudProvedIterator is returned from FilterFraudProved and is used to iterate over the raw logs and unpacked data for FraudProved events raised by the Contracts contract.
type ContractsFraudProvedIterator struct {
	Event *ContractsFraudProved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsFraudProvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsFraudProved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsFraudProved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsFraudProvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsFraudProvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsFraudProved represents a FraudProved event raised by the Contracts contract.
type ContractsFraudProved struct {
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFraudProved is a free log retrieval operation binding the contract event 0x1f35b91450baf2072ff37f7dbefd90eace09b98ad6d40c13ae7af6f64aacb849.
//
// Solidity: event Fraud_Proved(address challenger)
func (_Contracts *ContractsFilterer) FilterFraudProved(opts *bind.FilterOpts) (*ContractsFraudProvedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Fraud_Proved")
	if err != nil {
		return nil, err
	}
	return &ContractsFraudProvedIterator{contract: _Contracts.contract, event: "Fraud_Proved", logs: logs, sub: sub}, nil
}

// WatchFraudProved is a free log subscription operation binding the contract event 0x1f35b91450baf2072ff37f7dbefd90eace09b98ad6d40c13ae7af6f64aacb849.
//
// Solidity: event Fraud_Proved(address challenger)
func (_Contracts *ContractsFilterer) WatchFraudProved(opts *bind.WatchOpts, sink chan<- *ContractsFraudProved) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Fraud_Proved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsFraudProved)
				if err := _Contracts.contract.UnpackLog(event, "Fraud_Proved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFraudProved is a log parse operation binding the contract event 0x1f35b91450baf2072ff37f7dbefd90eace09b98ad6d40c13ae7af6f64aacb849.
//
// Solidity: event Fraud_Proved(address challenger)
func (_Contracts *ContractsFilterer) ParseFraudProved(log types.Log) (*ContractsFraudProved, error) {
	event := new(ContractsFraudProved)
	if err := _Contracts.contract.UnpackLog(event, "Fraud_Proved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsInvalidProofIterator is returned from FilterInvalidProof and is used to iterate over the raw logs and unpacked data for InvalidProof events raised by the Contracts contract.
type ContractsInvalidProofIterator struct {
	Event *ContractsInvalidProof // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsInvalidProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsInvalidProof)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsInvalidProof)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsInvalidProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsInvalidProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsInvalidProof represents a InvalidProof event raised by the Contracts contract.
type ContractsInvalidProof struct {
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvalidProof is a free log retrieval operation binding the contract event 0x20fc9549bb0aaddcc67903f8d9af4fe5f102d559cd1fcdbaa74dd4b198b44afc.
//
// Solidity: event Invalid_Proof(address challenger)
func (_Contracts *ContractsFilterer) FilterInvalidProof(opts *bind.FilterOpts) (*ContractsInvalidProofIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Invalid_Proof")
	if err != nil {
		return nil, err
	}
	return &ContractsInvalidProofIterator{contract: _Contracts.contract, event: "Invalid_Proof", logs: logs, sub: sub}, nil
}

// WatchInvalidProof is a free log subscription operation binding the contract event 0x20fc9549bb0aaddcc67903f8d9af4fe5f102d559cd1fcdbaa74dd4b198b44afc.
//
// Solidity: event Invalid_Proof(address challenger)
func (_Contracts *ContractsFilterer) WatchInvalidProof(opts *bind.WatchOpts, sink chan<- *ContractsInvalidProof) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Invalid_Proof")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsInvalidProof)
				if err := _Contracts.contract.UnpackLog(event, "Invalid_Proof", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvalidProof is a log parse operation binding the contract event 0x20fc9549bb0aaddcc67903f8d9af4fe5f102d559cd1fcdbaa74dd4b198b44afc.
//
// Solidity: event Invalid_Proof(address challenger)
func (_Contracts *ContractsFilterer) ParseInvalidProof(log types.Log) (*ContractsInvalidProof, error) {
	event := new(ContractsInvalidProof)
	if err := _Contracts.contract.UnpackLog(event, "Invalid_Proof", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsNewDepositIterator is returned from FilterNewDeposit and is used to iterate over the raw logs and unpacked data for NewDeposit events raised by the Contracts contract.
type ContractsNewDepositIterator struct {
	Event *ContractsNewDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsNewDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsNewDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsNewDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsNewDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsNewDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsNewDeposit represents a NewDeposit event raised by the Contracts contract.
type ContractsNewDeposit struct {
	User      common.Address
	StateRoot [32]byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDeposit is a free log retrieval operation binding the contract event 0xcd383a129c9295e144cae64b0726b69050054843ac23c8ca12b84fd69464ed8c.
//
// Solidity: event New_Deposit(address user, bytes32 stateRoot, uint256 value)
func (_Contracts *ContractsFilterer) FilterNewDeposit(opts *bind.FilterOpts) (*ContractsNewDepositIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "New_Deposit")
	if err != nil {
		return nil, err
	}
	return &ContractsNewDepositIterator{contract: _Contracts.contract, event: "New_Deposit", logs: logs, sub: sub}, nil
}

// WatchNewDeposit is a free log subscription operation binding the contract event 0xcd383a129c9295e144cae64b0726b69050054843ac23c8ca12b84fd69464ed8c.
//
// Solidity: event New_Deposit(address user, bytes32 stateRoot, uint256 value)
func (_Contracts *ContractsFilterer) WatchNewDeposit(opts *bind.WatchOpts, sink chan<- *ContractsNewDeposit) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "New_Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsNewDeposit)
				if err := _Contracts.contract.UnpackLog(event, "New_Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewDeposit is a log parse operation binding the contract event 0xcd383a129c9295e144cae64b0726b69050054843ac23c8ca12b84fd69464ed8c.
//
// Solidity: event New_Deposit(address user, bytes32 stateRoot, uint256 value)
func (_Contracts *ContractsFilterer) ParseNewDeposit(log types.Log) (*ContractsNewDeposit, error) {
	event := new(ContractsNewDeposit)
	if err := _Contracts.contract.UnpackLog(event, "New_Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsNewWithdrawIterator is returned from FilterNewWithdraw and is used to iterate over the raw logs and unpacked data for NewWithdraw events raised by the Contracts contract.
type ContractsNewWithdrawIterator struct {
	Event *ContractsNewWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsNewWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsNewWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsNewWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsNewWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsNewWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsNewWithdraw represents a NewWithdraw event raised by the Contracts contract.
type ContractsNewWithdraw struct {
	User      common.Address
	StateRoot [32]byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewWithdraw is a free log retrieval operation binding the contract event 0x33f28a9218883981815a14a0fd9f3d4e88e16b4d67e02bd72815c4dacfb5494f.
//
// Solidity: event New_withdraw(address user, bytes32 stateRoot, uint256 value)
func (_Contracts *ContractsFilterer) FilterNewWithdraw(opts *bind.FilterOpts) (*ContractsNewWithdrawIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "New_withdraw")
	if err != nil {
		return nil, err
	}
	return &ContractsNewWithdrawIterator{contract: _Contracts.contract, event: "New_withdraw", logs: logs, sub: sub}, nil
}

// WatchNewWithdraw is a free log subscription operation binding the contract event 0x33f28a9218883981815a14a0fd9f3d4e88e16b4d67e02bd72815c4dacfb5494f.
//
// Solidity: event New_withdraw(address user, bytes32 stateRoot, uint256 value)
func (_Contracts *ContractsFilterer) WatchNewWithdraw(opts *bind.WatchOpts, sink chan<- *ContractsNewWithdraw) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "New_withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsNewWithdraw)
				if err := _Contracts.contract.UnpackLog(event, "New_withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewWithdraw is a log parse operation binding the contract event 0x33f28a9218883981815a14a0fd9f3d4e88e16b4d67e02bd72815c4dacfb5494f.
//
// Solidity: event New_withdraw(address user, bytes32 stateRoot, uint256 value)
func (_Contracts *ContractsFilterer) ParseNewWithdraw(log types.Log) (*ContractsNewWithdraw, error) {
	event := new(ContractsNewWithdraw)
	if err := _Contracts.contract.UnpackLog(event, "New_withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
