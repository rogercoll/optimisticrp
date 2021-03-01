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
const ContractsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lock_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_required_bond\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accuser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"suspect\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transition_index\",\"type\":\"uint256\"}],\"name\":\"ORI_Fraud_Proven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"suspect\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accuser\",\"type\":\"address\"}],\"name\":\"ORI_Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"block_time\",\"type\":\"uint256\"}],\"name\":\"ORI_New_Optimistic_State\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"block_time\",\"type\":\"uint256\"}],\"name\":\"ORI_New_Optimistic_States\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"new_state\",\"type\":\"bytes32\"}],\"name\":\"ORI_New_State\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tree_size\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"block_time\",\"type\":\"uint256\"}],\"name\":\"ORI_Rolled_Back\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"suspect\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accuser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"block_time\",\"type\":\"uint256\"}],\"name\":\"ORI_Unlocked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"account_states\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"call_data_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"state\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"last_time\",\"type\":\"uint256\"}],\"name\":\"archive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"bond\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"suspect\",\"type\":\"address\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"locked_timestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logic_contract\",\"outputs\":[{\"internalType\":\"contractOptimistic_Roll_In_Compatible\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"call_data\",\"type\":\"bytes\"}],\"name\":\"perform\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"call_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"call_data_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"last_time\",\"type\":\"uint256\"}],\"name\":\"perform_and_exit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"call_data\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"new_state\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"call_data_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"last_time\",\"type\":\"uint256\"}],\"name\":\"perform_many_optimistically\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"call_data\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"new_state\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"perform_many_optimistically_and_enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"call_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"new_state\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"call_data_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"last_time\",\"type\":\"uint256\"}],\"name\":\"perform_optimistically\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"call_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"new_state\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"perform_optimistically_and_enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"suspect\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"call_data\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"state\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"call_data_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"last_time\",\"type\":\"uint256\"}],\"name\":\"prove_fraud\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"required_bond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rolled_back_call_data_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"rolled_back_call_data\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"roll_back_proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"current_size\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"current_size_proof\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"call_data_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"state\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"last_time\",\"type\":\"uint256\"}],\"name\":\"rollback\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rollback_sizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unarchive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"unbond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"suspect\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"state\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"call_data_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"last_time\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

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

// AccountStates is a free data retrieval call binding the contract method 0xaf7b5a41.
//
// Solidity: function account_states(address ) view returns(bytes32)
func (_Contracts *ContractsCaller) AccountStates(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "account_states", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AccountStates is a free data retrieval call binding the contract method 0xaf7b5a41.
//
// Solidity: function account_states(address ) view returns(bytes32)
func (_Contracts *ContractsSession) AccountStates(arg0 common.Address) ([32]byte, error) {
	return _Contracts.Contract.AccountStates(&_Contracts.CallOpts, arg0)
}

// AccountStates is a free data retrieval call binding the contract method 0xaf7b5a41.
//
// Solidity: function account_states(address ) view returns(bytes32)
func (_Contracts *ContractsCallerSession) AccountStates(arg0 common.Address) ([32]byte, error) {
	return _Contracts.Contract.AccountStates(&_Contracts.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Contracts *ContractsCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Contracts *ContractsSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.Balances(&_Contracts.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.Balances(&_Contracts.CallOpts, arg0)
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

// LockedTimestamps is a free data retrieval call binding the contract method 0x7a9f05b5.
//
// Solidity: function locked_timestamps(address ) view returns(uint256)
func (_Contracts *ContractsCaller) LockedTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "locked_timestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedTimestamps is a free data retrieval call binding the contract method 0x7a9f05b5.
//
// Solidity: function locked_timestamps(address ) view returns(uint256)
func (_Contracts *ContractsSession) LockedTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.LockedTimestamps(&_Contracts.CallOpts, arg0)
}

// LockedTimestamps is a free data retrieval call binding the contract method 0x7a9f05b5.
//
// Solidity: function locked_timestamps(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) LockedTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.LockedTimestamps(&_Contracts.CallOpts, arg0)
}

// Lockers is a free data retrieval call binding the contract method 0x2c8e7a21.
//
// Solidity: function lockers(address ) view returns(address)
func (_Contracts *ContractsCaller) Lockers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "lockers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lockers is a free data retrieval call binding the contract method 0x2c8e7a21.
//
// Solidity: function lockers(address ) view returns(address)
func (_Contracts *ContractsSession) Lockers(arg0 common.Address) (common.Address, error) {
	return _Contracts.Contract.Lockers(&_Contracts.CallOpts, arg0)
}

// Lockers is a free data retrieval call binding the contract method 0x2c8e7a21.
//
// Solidity: function lockers(address ) view returns(address)
func (_Contracts *ContractsCallerSession) Lockers(arg0 common.Address) (common.Address, error) {
	return _Contracts.Contract.Lockers(&_Contracts.CallOpts, arg0)
}

// LogicContract is a free data retrieval call binding the contract method 0x4de94320.
//
// Solidity: function logic_contract() view returns(address)
func (_Contracts *ContractsCaller) LogicContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "logic_contract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LogicContract is a free data retrieval call binding the contract method 0x4de94320.
//
// Solidity: function logic_contract() view returns(address)
func (_Contracts *ContractsSession) LogicContract() (common.Address, error) {
	return _Contracts.Contract.LogicContract(&_Contracts.CallOpts)
}

// LogicContract is a free data retrieval call binding the contract method 0x4de94320.
//
// Solidity: function logic_contract() view returns(address)
func (_Contracts *ContractsCallerSession) LogicContract() (common.Address, error) {
	return _Contracts.Contract.LogicContract(&_Contracts.CallOpts)
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

// RollbackSizes is a free data retrieval call binding the contract method 0xc45481bb.
//
// Solidity: function rollback_sizes(address ) view returns(uint256)
func (_Contracts *ContractsCaller) RollbackSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "rollback_sizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RollbackSizes is a free data retrieval call binding the contract method 0xc45481bb.
//
// Solidity: function rollback_sizes(address ) view returns(uint256)
func (_Contracts *ContractsSession) RollbackSizes(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.RollbackSizes(&_Contracts.CallOpts, arg0)
}

// RollbackSizes is a free data retrieval call binding the contract method 0xc45481bb.
//
// Solidity: function rollback_sizes(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) RollbackSizes(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.RollbackSizes(&_Contracts.CallOpts, arg0)
}

// Archive is a paid mutator transaction binding the contract method 0xe6a64799.
//
// Solidity: function archive(address destination, bytes32 call_data_root, bytes32 state, uint256 last_time) returns()
func (_Contracts *ContractsTransactor) Archive(opts *bind.TransactOpts, destination common.Address, call_data_root [32]byte, state [32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "archive", destination, call_data_root, state, last_time)
}

// Archive is a paid mutator transaction binding the contract method 0xe6a64799.
//
// Solidity: function archive(address destination, bytes32 call_data_root, bytes32 state, uint256 last_time) returns()
func (_Contracts *ContractsSession) Archive(destination common.Address, call_data_root [32]byte, state [32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Archive(&_Contracts.TransactOpts, destination, call_data_root, state, last_time)
}

// Archive is a paid mutator transaction binding the contract method 0xe6a64799.
//
// Solidity: function archive(address destination, bytes32 call_data_root, bytes32 state, uint256 last_time) returns()
func (_Contracts *ContractsTransactorSession) Archive(destination common.Address, call_data_root [32]byte, state [32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Archive(&_Contracts.TransactOpts, destination, call_data_root, state, last_time)
}

// Bond is a paid mutator transaction binding the contract method 0x247ce85b.
//
// Solidity: function bond(address user) payable returns()
func (_Contracts *ContractsTransactor) Bond(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "bond", user)
}

// Bond is a paid mutator transaction binding the contract method 0x247ce85b.
//
// Solidity: function bond(address user) payable returns()
func (_Contracts *ContractsSession) Bond(user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Bond(&_Contracts.TransactOpts, user)
}

// Bond is a paid mutator transaction binding the contract method 0x247ce85b.
//
// Solidity: function bond(address user) payable returns()
func (_Contracts *ContractsTransactorSession) Bond(user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Bond(&_Contracts.TransactOpts, user)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_Contracts *ContractsSession) Initialize() (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() payable returns()
func (_Contracts *ContractsTransactorSession) Initialize() (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf435f5a7.
//
// Solidity: function lock(address suspect) payable returns()
func (_Contracts *ContractsTransactor) Lock(opts *bind.TransactOpts, suspect common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "lock", suspect)
}

// Lock is a paid mutator transaction binding the contract method 0xf435f5a7.
//
// Solidity: function lock(address suspect) payable returns()
func (_Contracts *ContractsSession) Lock(suspect common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Lock(&_Contracts.TransactOpts, suspect)
}

// Lock is a paid mutator transaction binding the contract method 0xf435f5a7.
//
// Solidity: function lock(address suspect) payable returns()
func (_Contracts *ContractsTransactorSession) Lock(suspect common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Lock(&_Contracts.TransactOpts, suspect)
}

// Perform is a paid mutator transaction binding the contract method 0xbb6ae2cb.
//
// Solidity: function perform(bytes call_data) payable returns()
func (_Contracts *ContractsTransactor) Perform(opts *bind.TransactOpts, call_data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "perform", call_data)
}

// Perform is a paid mutator transaction binding the contract method 0xbb6ae2cb.
//
// Solidity: function perform(bytes call_data) payable returns()
func (_Contracts *ContractsSession) Perform(call_data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Perform(&_Contracts.TransactOpts, call_data)
}

// Perform is a paid mutator transaction binding the contract method 0xbb6ae2cb.
//
// Solidity: function perform(bytes call_data) payable returns()
func (_Contracts *ContractsTransactorSession) Perform(call_data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Perform(&_Contracts.TransactOpts, call_data)
}

// PerformAndExit is a paid mutator transaction binding the contract method 0x2a415454.
//
// Solidity: function perform_and_exit(bytes call_data, bytes32 call_data_root, uint256 last_time) payable returns()
func (_Contracts *ContractsTransactor) PerformAndExit(opts *bind.TransactOpts, call_data []byte, call_data_root [32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "perform_and_exit", call_data, call_data_root, last_time)
}

// PerformAndExit is a paid mutator transaction binding the contract method 0x2a415454.
//
// Solidity: function perform_and_exit(bytes call_data, bytes32 call_data_root, uint256 last_time) payable returns()
func (_Contracts *ContractsSession) PerformAndExit(call_data []byte, call_data_root [32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.PerformAndExit(&_Contracts.TransactOpts, call_data, call_data_root, last_time)
}

// PerformAndExit is a paid mutator transaction binding the contract method 0x2a415454.
//
// Solidity: function perform_and_exit(bytes call_data, bytes32 call_data_root, uint256 last_time) payable returns()
func (_Contracts *ContractsTransactorSession) PerformAndExit(call_data []byte, call_data_root [32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.PerformAndExit(&_Contracts.TransactOpts, call_data, call_data_root, last_time)
}

// PerformManyOptimistically is a paid mutator transaction binding the contract method 0x6a8dddef.
//
// Solidity: function perform_many_optimistically(bytes[] call_data, bytes32 new_state, bytes32 call_data_root, bytes32[] proof, uint256 last_time) returns()
func (_Contracts *ContractsTransactor) PerformManyOptimistically(opts *bind.TransactOpts, call_data [][]byte, new_state [32]byte, call_data_root [32]byte, proof [][32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "perform_many_optimistically", call_data, new_state, call_data_root, proof, last_time)
}

// PerformManyOptimistically is a paid mutator transaction binding the contract method 0x6a8dddef.
//
// Solidity: function perform_many_optimistically(bytes[] call_data, bytes32 new_state, bytes32 call_data_root, bytes32[] proof, uint256 last_time) returns()
func (_Contracts *ContractsSession) PerformManyOptimistically(call_data [][]byte, new_state [32]byte, call_data_root [32]byte, proof [][32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.PerformManyOptimistically(&_Contracts.TransactOpts, call_data, new_state, call_data_root, proof, last_time)
}

// PerformManyOptimistically is a paid mutator transaction binding the contract method 0x6a8dddef.
//
// Solidity: function perform_many_optimistically(bytes[] call_data, bytes32 new_state, bytes32 call_data_root, bytes32[] proof, uint256 last_time) returns()
func (_Contracts *ContractsTransactorSession) PerformManyOptimistically(call_data [][]byte, new_state [32]byte, call_data_root [32]byte, proof [][32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.PerformManyOptimistically(&_Contracts.TransactOpts, call_data, new_state, call_data_root, proof, last_time)
}

// PerformManyOptimisticallyAndEnter is a paid mutator transaction binding the contract method 0x08542bb1.
//
// Solidity: function perform_many_optimistically_and_enter(bytes[] call_data, bytes32 new_state, bytes32[] proof) returns()
func (_Contracts *ContractsTransactor) PerformManyOptimisticallyAndEnter(opts *bind.TransactOpts, call_data [][]byte, new_state [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "perform_many_optimistically_and_enter", call_data, new_state, proof)
}

// PerformManyOptimisticallyAndEnter is a paid mutator transaction binding the contract method 0x08542bb1.
//
// Solidity: function perform_many_optimistically_and_enter(bytes[] call_data, bytes32 new_state, bytes32[] proof) returns()
func (_Contracts *ContractsSession) PerformManyOptimisticallyAndEnter(call_data [][]byte, new_state [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.PerformManyOptimisticallyAndEnter(&_Contracts.TransactOpts, call_data, new_state, proof)
}

// PerformManyOptimisticallyAndEnter is a paid mutator transaction binding the contract method 0x08542bb1.
//
// Solidity: function perform_many_optimistically_and_enter(bytes[] call_data, bytes32 new_state, bytes32[] proof) returns()
func (_Contracts *ContractsTransactorSession) PerformManyOptimisticallyAndEnter(call_data [][]byte, new_state [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.PerformManyOptimisticallyAndEnter(&_Contracts.TransactOpts, call_data, new_state, proof)
}

// PerformOptimistically is a paid mutator transaction binding the contract method 0x1646d051.
//
// Solidity: function perform_optimistically(bytes call_data, bytes32 new_state, bytes32 call_data_root, bytes32[] proof, uint256 last_time) returns()
func (_Contracts *ContractsTransactor) PerformOptimistically(opts *bind.TransactOpts, call_data []byte, new_state [32]byte, call_data_root [32]byte, proof [][32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "perform_optimistically", call_data, new_state, call_data_root, proof, last_time)
}

// PerformOptimistically is a paid mutator transaction binding the contract method 0x1646d051.
//
// Solidity: function perform_optimistically(bytes call_data, bytes32 new_state, bytes32 call_data_root, bytes32[] proof, uint256 last_time) returns()
func (_Contracts *ContractsSession) PerformOptimistically(call_data []byte, new_state [32]byte, call_data_root [32]byte, proof [][32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.PerformOptimistically(&_Contracts.TransactOpts, call_data, new_state, call_data_root, proof, last_time)
}

// PerformOptimistically is a paid mutator transaction binding the contract method 0x1646d051.
//
// Solidity: function perform_optimistically(bytes call_data, bytes32 new_state, bytes32 call_data_root, bytes32[] proof, uint256 last_time) returns()
func (_Contracts *ContractsTransactorSession) PerformOptimistically(call_data []byte, new_state [32]byte, call_data_root [32]byte, proof [][32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.PerformOptimistically(&_Contracts.TransactOpts, call_data, new_state, call_data_root, proof, last_time)
}

// PerformOptimisticallyAndEnter is a paid mutator transaction binding the contract method 0x177f15c5.
//
// Solidity: function perform_optimistically_and_enter(bytes call_data, bytes32 new_state, bytes32[] proof) returns()
func (_Contracts *ContractsTransactor) PerformOptimisticallyAndEnter(opts *bind.TransactOpts, call_data []byte, new_state [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "perform_optimistically_and_enter", call_data, new_state, proof)
}

// PerformOptimisticallyAndEnter is a paid mutator transaction binding the contract method 0x177f15c5.
//
// Solidity: function perform_optimistically_and_enter(bytes call_data, bytes32 new_state, bytes32[] proof) returns()
func (_Contracts *ContractsSession) PerformOptimisticallyAndEnter(call_data []byte, new_state [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.PerformOptimisticallyAndEnter(&_Contracts.TransactOpts, call_data, new_state, proof)
}

// PerformOptimisticallyAndEnter is a paid mutator transaction binding the contract method 0x177f15c5.
//
// Solidity: function perform_optimistically_and_enter(bytes call_data, bytes32 new_state, bytes32[] proof) returns()
func (_Contracts *ContractsTransactorSession) PerformOptimisticallyAndEnter(call_data []byte, new_state [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.PerformOptimisticallyAndEnter(&_Contracts.TransactOpts, call_data, new_state, proof)
}

// ProveFraud is a paid mutator transaction binding the contract method 0xa68f5913.
//
// Solidity: function prove_fraud(address suspect, bytes[] call_data, bytes32 state, bytes32 call_data_root, bytes32[] proof, uint256 last_time) returns()
func (_Contracts *ContractsTransactor) ProveFraud(opts *bind.TransactOpts, suspect common.Address, call_data [][]byte, state [32]byte, call_data_root [32]byte, proof [][32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "prove_fraud", suspect, call_data, state, call_data_root, proof, last_time)
}

// ProveFraud is a paid mutator transaction binding the contract method 0xa68f5913.
//
// Solidity: function prove_fraud(address suspect, bytes[] call_data, bytes32 state, bytes32 call_data_root, bytes32[] proof, uint256 last_time) returns()
func (_Contracts *ContractsSession) ProveFraud(suspect common.Address, call_data [][]byte, state [32]byte, call_data_root [32]byte, proof [][32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ProveFraud(&_Contracts.TransactOpts, suspect, call_data, state, call_data_root, proof, last_time)
}

// ProveFraud is a paid mutator transaction binding the contract method 0xa68f5913.
//
// Solidity: function prove_fraud(address suspect, bytes[] call_data, bytes32 state, bytes32 call_data_root, bytes32[] proof, uint256 last_time) returns()
func (_Contracts *ContractsTransactorSession) ProveFraud(suspect common.Address, call_data [][]byte, state [32]byte, call_data_root [32]byte, proof [][32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ProveFraud(&_Contracts.TransactOpts, suspect, call_data, state, call_data_root, proof, last_time)
}

// Rollback is a paid mutator transaction binding the contract method 0xe1124200.
//
// Solidity: function rollback(bytes32 rolled_back_call_data_root, bytes[] rolled_back_call_data, bytes32[] roll_back_proof, uint256 current_size, bytes32 current_size_proof, bytes32 call_data_root, bytes32 state, uint256 last_time) payable returns()
func (_Contracts *ContractsTransactor) Rollback(opts *bind.TransactOpts, rolled_back_call_data_root [32]byte, rolled_back_call_data [][]byte, roll_back_proof [][32]byte, current_size *big.Int, current_size_proof [32]byte, call_data_root [32]byte, state [32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "rollback", rolled_back_call_data_root, rolled_back_call_data, roll_back_proof, current_size, current_size_proof, call_data_root, state, last_time)
}

// Rollback is a paid mutator transaction binding the contract method 0xe1124200.
//
// Solidity: function rollback(bytes32 rolled_back_call_data_root, bytes[] rolled_back_call_data, bytes32[] roll_back_proof, uint256 current_size, bytes32 current_size_proof, bytes32 call_data_root, bytes32 state, uint256 last_time) payable returns()
func (_Contracts *ContractsSession) Rollback(rolled_back_call_data_root [32]byte, rolled_back_call_data [][]byte, roll_back_proof [][32]byte, current_size *big.Int, current_size_proof [32]byte, call_data_root [32]byte, state [32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Rollback(&_Contracts.TransactOpts, rolled_back_call_data_root, rolled_back_call_data, roll_back_proof, current_size, current_size_proof, call_data_root, state, last_time)
}

// Rollback is a paid mutator transaction binding the contract method 0xe1124200.
//
// Solidity: function rollback(bytes32 rolled_back_call_data_root, bytes[] rolled_back_call_data, bytes32[] roll_back_proof, uint256 current_size, bytes32 current_size_proof, bytes32 call_data_root, bytes32 state, uint256 last_time) payable returns()
func (_Contracts *ContractsTransactorSession) Rollback(rolled_back_call_data_root [32]byte, rolled_back_call_data [][]byte, roll_back_proof [][32]byte, current_size *big.Int, current_size_proof [32]byte, call_data_root [32]byte, state [32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Rollback(&_Contracts.TransactOpts, rolled_back_call_data_root, rolled_back_call_data, roll_back_proof, current_size, current_size_proof, call_data_root, state, last_time)
}

// Unarchive is a paid mutator transaction binding the contract method 0x7ad246f8.
//
// Solidity: function unarchive() payable returns()
func (_Contracts *ContractsTransactor) Unarchive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unarchive")
}

// Unarchive is a paid mutator transaction binding the contract method 0x7ad246f8.
//
// Solidity: function unarchive() payable returns()
func (_Contracts *ContractsSession) Unarchive() (*types.Transaction, error) {
	return _Contracts.Contract.Unarchive(&_Contracts.TransactOpts)
}

// Unarchive is a paid mutator transaction binding the contract method 0x7ad246f8.
//
// Solidity: function unarchive() payable returns()
func (_Contracts *ContractsTransactorSession) Unarchive() (*types.Transaction, error) {
	return _Contracts.Contract.Unarchive(&_Contracts.TransactOpts)
}

// Unbond is a paid mutator transaction binding the contract method 0xe4da61ab.
//
// Solidity: function unbond(address destination) returns()
func (_Contracts *ContractsTransactor) Unbond(opts *bind.TransactOpts, destination common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unbond", destination)
}

// Unbond is a paid mutator transaction binding the contract method 0xe4da61ab.
//
// Solidity: function unbond(address destination) returns()
func (_Contracts *ContractsSession) Unbond(destination common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Unbond(&_Contracts.TransactOpts, destination)
}

// Unbond is a paid mutator transaction binding the contract method 0xe4da61ab.
//
// Solidity: function unbond(address destination) returns()
func (_Contracts *ContractsTransactorSession) Unbond(destination common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Unbond(&_Contracts.TransactOpts, destination)
}

// Unlock is a paid mutator transaction binding the contract method 0x8ee16d13.
//
// Solidity: function unlock(address suspect, bytes32 state, bytes32 call_data_root, uint256 last_time) returns()
func (_Contracts *ContractsTransactor) Unlock(opts *bind.TransactOpts, suspect common.Address, state [32]byte, call_data_root [32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unlock", suspect, state, call_data_root, last_time)
}

// Unlock is a paid mutator transaction binding the contract method 0x8ee16d13.
//
// Solidity: function unlock(address suspect, bytes32 state, bytes32 call_data_root, uint256 last_time) returns()
func (_Contracts *ContractsSession) Unlock(suspect common.Address, state [32]byte, call_data_root [32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Unlock(&_Contracts.TransactOpts, suspect, state, call_data_root, last_time)
}

// Unlock is a paid mutator transaction binding the contract method 0x8ee16d13.
//
// Solidity: function unlock(address suspect, bytes32 state, bytes32 call_data_root, uint256 last_time) returns()
func (_Contracts *ContractsTransactorSession) Unlock(suspect common.Address, state [32]byte, call_data_root [32]byte, last_time *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Unlock(&_Contracts.TransactOpts, suspect, state, call_data_root, last_time)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address destination) returns()
func (_Contracts *ContractsTransactor) Withdraw(opts *bind.TransactOpts, destination common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdraw", destination)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address destination) returns()
func (_Contracts *ContractsSession) Withdraw(destination common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, destination)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address destination) returns()
func (_Contracts *ContractsTransactorSession) Withdraw(destination common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, destination)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactorSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// ContractsORIFraudProvenIterator is returned from FilterORIFraudProven and is used to iterate over the raw logs and unpacked data for ORIFraudProven events raised by the Contracts contract.
type ContractsORIFraudProvenIterator struct {
	Event *ContractsORIFraudProven // Event containing the contract specifics and raw log

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
func (it *ContractsORIFraudProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsORIFraudProven)
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
		it.Event = new(ContractsORIFraudProven)
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
func (it *ContractsORIFraudProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsORIFraudProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsORIFraudProven represents a ORIFraudProven event raised by the Contracts contract.
type ContractsORIFraudProven struct {
	Accuser         common.Address
	Suspect         common.Address
	TransitionIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterORIFraudProven is a free log retrieval operation binding the contract event 0x55469cdf3fd81a478fed2a1491832e3d145bf2ab3c6a6fbb187022b3885b90fd.
//
// Solidity: event ORI_Fraud_Proven(address indexed accuser, address indexed suspect, uint256 indexed transition_index)
func (_Contracts *ContractsFilterer) FilterORIFraudProven(opts *bind.FilterOpts, accuser []common.Address, suspect []common.Address, transition_index []*big.Int) (*ContractsORIFraudProvenIterator, error) {

	var accuserRule []interface{}
	for _, accuserItem := range accuser {
		accuserRule = append(accuserRule, accuserItem)
	}
	var suspectRule []interface{}
	for _, suspectItem := range suspect {
		suspectRule = append(suspectRule, suspectItem)
	}
	var transition_indexRule []interface{}
	for _, transition_indexItem := range transition_index {
		transition_indexRule = append(transition_indexRule, transition_indexItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ORI_Fraud_Proven", accuserRule, suspectRule, transition_indexRule)
	if err != nil {
		return nil, err
	}
	return &ContractsORIFraudProvenIterator{contract: _Contracts.contract, event: "ORI_Fraud_Proven", logs: logs, sub: sub}, nil
}

// WatchORIFraudProven is a free log subscription operation binding the contract event 0x55469cdf3fd81a478fed2a1491832e3d145bf2ab3c6a6fbb187022b3885b90fd.
//
// Solidity: event ORI_Fraud_Proven(address indexed accuser, address indexed suspect, uint256 indexed transition_index)
func (_Contracts *ContractsFilterer) WatchORIFraudProven(opts *bind.WatchOpts, sink chan<- *ContractsORIFraudProven, accuser []common.Address, suspect []common.Address, transition_index []*big.Int) (event.Subscription, error) {

	var accuserRule []interface{}
	for _, accuserItem := range accuser {
		accuserRule = append(accuserRule, accuserItem)
	}
	var suspectRule []interface{}
	for _, suspectItem := range suspect {
		suspectRule = append(suspectRule, suspectItem)
	}
	var transition_indexRule []interface{}
	for _, transition_indexItem := range transition_index {
		transition_indexRule = append(transition_indexRule, transition_indexItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ORI_Fraud_Proven", accuserRule, suspectRule, transition_indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsORIFraudProven)
				if err := _Contracts.contract.UnpackLog(event, "ORI_Fraud_Proven", log); err != nil {
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

// ParseORIFraudProven is a log parse operation binding the contract event 0x55469cdf3fd81a478fed2a1491832e3d145bf2ab3c6a6fbb187022b3885b90fd.
//
// Solidity: event ORI_Fraud_Proven(address indexed accuser, address indexed suspect, uint256 indexed transition_index)
func (_Contracts *ContractsFilterer) ParseORIFraudProven(log types.Log) (*ContractsORIFraudProven, error) {
	event := new(ContractsORIFraudProven)
	if err := _Contracts.contract.UnpackLog(event, "ORI_Fraud_Proven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsORILockedIterator is returned from FilterORILocked and is used to iterate over the raw logs and unpacked data for ORILocked events raised by the Contracts contract.
type ContractsORILockedIterator struct {
	Event *ContractsORILocked // Event containing the contract specifics and raw log

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
func (it *ContractsORILockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsORILocked)
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
		it.Event = new(ContractsORILocked)
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
func (it *ContractsORILockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsORILockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsORILocked represents a ORILocked event raised by the Contracts contract.
type ContractsORILocked struct {
	Suspect common.Address
	Accuser common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterORILocked is a free log retrieval operation binding the contract event 0x8773bde6581ad6ddd421210de867340039fb65ce3df41edba7b5de6d24ae7a51.
//
// Solidity: event ORI_Locked(address indexed suspect, address indexed accuser)
func (_Contracts *ContractsFilterer) FilterORILocked(opts *bind.FilterOpts, suspect []common.Address, accuser []common.Address) (*ContractsORILockedIterator, error) {

	var suspectRule []interface{}
	for _, suspectItem := range suspect {
		suspectRule = append(suspectRule, suspectItem)
	}
	var accuserRule []interface{}
	for _, accuserItem := range accuser {
		accuserRule = append(accuserRule, accuserItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ORI_Locked", suspectRule, accuserRule)
	if err != nil {
		return nil, err
	}
	return &ContractsORILockedIterator{contract: _Contracts.contract, event: "ORI_Locked", logs: logs, sub: sub}, nil
}

// WatchORILocked is a free log subscription operation binding the contract event 0x8773bde6581ad6ddd421210de867340039fb65ce3df41edba7b5de6d24ae7a51.
//
// Solidity: event ORI_Locked(address indexed suspect, address indexed accuser)
func (_Contracts *ContractsFilterer) WatchORILocked(opts *bind.WatchOpts, sink chan<- *ContractsORILocked, suspect []common.Address, accuser []common.Address) (event.Subscription, error) {

	var suspectRule []interface{}
	for _, suspectItem := range suspect {
		suspectRule = append(suspectRule, suspectItem)
	}
	var accuserRule []interface{}
	for _, accuserItem := range accuser {
		accuserRule = append(accuserRule, accuserItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ORI_Locked", suspectRule, accuserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsORILocked)
				if err := _Contracts.contract.UnpackLog(event, "ORI_Locked", log); err != nil {
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

// ParseORILocked is a log parse operation binding the contract event 0x8773bde6581ad6ddd421210de867340039fb65ce3df41edba7b5de6d24ae7a51.
//
// Solidity: event ORI_Locked(address indexed suspect, address indexed accuser)
func (_Contracts *ContractsFilterer) ParseORILocked(log types.Log) (*ContractsORILocked, error) {
	event := new(ContractsORILocked)
	if err := _Contracts.contract.UnpackLog(event, "ORI_Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsORINewOptimisticStateIterator is returned from FilterORINewOptimisticState and is used to iterate over the raw logs and unpacked data for ORINewOptimisticState events raised by the Contracts contract.
type ContractsORINewOptimisticStateIterator struct {
	Event *ContractsORINewOptimisticState // Event containing the contract specifics and raw log

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
func (it *ContractsORINewOptimisticStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsORINewOptimisticState)
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
		it.Event = new(ContractsORINewOptimisticState)
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
func (it *ContractsORINewOptimisticStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsORINewOptimisticStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsORINewOptimisticState represents a ORINewOptimisticState event raised by the Contracts contract.
type ContractsORINewOptimisticState struct {
	User      common.Address
	BlockTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterORINewOptimisticState is a free log retrieval operation binding the contract event 0x4779c4b07abff82b16061ec9a47d081e7f4981c29088395cdb7ff87e322cbbc6.
//
// Solidity: event ORI_New_Optimistic_State(address indexed user, uint256 indexed block_time)
func (_Contracts *ContractsFilterer) FilterORINewOptimisticState(opts *bind.FilterOpts, user []common.Address, block_time []*big.Int) (*ContractsORINewOptimisticStateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var block_timeRule []interface{}
	for _, block_timeItem := range block_time {
		block_timeRule = append(block_timeRule, block_timeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ORI_New_Optimistic_State", userRule, block_timeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsORINewOptimisticStateIterator{contract: _Contracts.contract, event: "ORI_New_Optimistic_State", logs: logs, sub: sub}, nil
}

// WatchORINewOptimisticState is a free log subscription operation binding the contract event 0x4779c4b07abff82b16061ec9a47d081e7f4981c29088395cdb7ff87e322cbbc6.
//
// Solidity: event ORI_New_Optimistic_State(address indexed user, uint256 indexed block_time)
func (_Contracts *ContractsFilterer) WatchORINewOptimisticState(opts *bind.WatchOpts, sink chan<- *ContractsORINewOptimisticState, user []common.Address, block_time []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var block_timeRule []interface{}
	for _, block_timeItem := range block_time {
		block_timeRule = append(block_timeRule, block_timeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ORI_New_Optimistic_State", userRule, block_timeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsORINewOptimisticState)
				if err := _Contracts.contract.UnpackLog(event, "ORI_New_Optimistic_State", log); err != nil {
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

// ParseORINewOptimisticState is a log parse operation binding the contract event 0x4779c4b07abff82b16061ec9a47d081e7f4981c29088395cdb7ff87e322cbbc6.
//
// Solidity: event ORI_New_Optimistic_State(address indexed user, uint256 indexed block_time)
func (_Contracts *ContractsFilterer) ParseORINewOptimisticState(log types.Log) (*ContractsORINewOptimisticState, error) {
	event := new(ContractsORINewOptimisticState)
	if err := _Contracts.contract.UnpackLog(event, "ORI_New_Optimistic_State", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsORINewOptimisticStatesIterator is returned from FilterORINewOptimisticStates and is used to iterate over the raw logs and unpacked data for ORINewOptimisticStates events raised by the Contracts contract.
type ContractsORINewOptimisticStatesIterator struct {
	Event *ContractsORINewOptimisticStates // Event containing the contract specifics and raw log

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
func (it *ContractsORINewOptimisticStatesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsORINewOptimisticStates)
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
		it.Event = new(ContractsORINewOptimisticStates)
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
func (it *ContractsORINewOptimisticStatesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsORINewOptimisticStatesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsORINewOptimisticStates represents a ORINewOptimisticStates event raised by the Contracts contract.
type ContractsORINewOptimisticStates struct {
	User      common.Address
	BlockTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterORINewOptimisticStates is a free log retrieval operation binding the contract event 0x0b87b136840d19f5f25329273082c00833265a189b70137e06df6315ddc7839e.
//
// Solidity: event ORI_New_Optimistic_States(address indexed user, uint256 indexed block_time)
func (_Contracts *ContractsFilterer) FilterORINewOptimisticStates(opts *bind.FilterOpts, user []common.Address, block_time []*big.Int) (*ContractsORINewOptimisticStatesIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var block_timeRule []interface{}
	for _, block_timeItem := range block_time {
		block_timeRule = append(block_timeRule, block_timeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ORI_New_Optimistic_States", userRule, block_timeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsORINewOptimisticStatesIterator{contract: _Contracts.contract, event: "ORI_New_Optimistic_States", logs: logs, sub: sub}, nil
}

// WatchORINewOptimisticStates is a free log subscription operation binding the contract event 0x0b87b136840d19f5f25329273082c00833265a189b70137e06df6315ddc7839e.
//
// Solidity: event ORI_New_Optimistic_States(address indexed user, uint256 indexed block_time)
func (_Contracts *ContractsFilterer) WatchORINewOptimisticStates(opts *bind.WatchOpts, sink chan<- *ContractsORINewOptimisticStates, user []common.Address, block_time []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var block_timeRule []interface{}
	for _, block_timeItem := range block_time {
		block_timeRule = append(block_timeRule, block_timeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ORI_New_Optimistic_States", userRule, block_timeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsORINewOptimisticStates)
				if err := _Contracts.contract.UnpackLog(event, "ORI_New_Optimistic_States", log); err != nil {
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

// ParseORINewOptimisticStates is a log parse operation binding the contract event 0x0b87b136840d19f5f25329273082c00833265a189b70137e06df6315ddc7839e.
//
// Solidity: event ORI_New_Optimistic_States(address indexed user, uint256 indexed block_time)
func (_Contracts *ContractsFilterer) ParseORINewOptimisticStates(log types.Log) (*ContractsORINewOptimisticStates, error) {
	event := new(ContractsORINewOptimisticStates)
	if err := _Contracts.contract.UnpackLog(event, "ORI_New_Optimistic_States", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsORINewStateIterator is returned from FilterORINewState and is used to iterate over the raw logs and unpacked data for ORINewState events raised by the Contracts contract.
type ContractsORINewStateIterator struct {
	Event *ContractsORINewState // Event containing the contract specifics and raw log

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
func (it *ContractsORINewStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsORINewState)
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
		it.Event = new(ContractsORINewState)
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
func (it *ContractsORINewStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsORINewStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsORINewState represents a ORINewState event raised by the Contracts contract.
type ContractsORINewState struct {
	User     common.Address
	NewState [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterORINewState is a free log retrieval operation binding the contract event 0x0f5025cc4f20aa47a346d1b7d9da6ba8c68cc8e83b75e813da4b4490d55365ae.
//
// Solidity: event ORI_New_State(address indexed user, bytes32 indexed new_state)
func (_Contracts *ContractsFilterer) FilterORINewState(opts *bind.FilterOpts, user []common.Address, new_state [][32]byte) (*ContractsORINewStateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var new_stateRule []interface{}
	for _, new_stateItem := range new_state {
		new_stateRule = append(new_stateRule, new_stateItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ORI_New_State", userRule, new_stateRule)
	if err != nil {
		return nil, err
	}
	return &ContractsORINewStateIterator{contract: _Contracts.contract, event: "ORI_New_State", logs: logs, sub: sub}, nil
}

// WatchORINewState is a free log subscription operation binding the contract event 0x0f5025cc4f20aa47a346d1b7d9da6ba8c68cc8e83b75e813da4b4490d55365ae.
//
// Solidity: event ORI_New_State(address indexed user, bytes32 indexed new_state)
func (_Contracts *ContractsFilterer) WatchORINewState(opts *bind.WatchOpts, sink chan<- *ContractsORINewState, user []common.Address, new_state [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var new_stateRule []interface{}
	for _, new_stateItem := range new_state {
		new_stateRule = append(new_stateRule, new_stateItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ORI_New_State", userRule, new_stateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsORINewState)
				if err := _Contracts.contract.UnpackLog(event, "ORI_New_State", log); err != nil {
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

// ParseORINewState is a log parse operation binding the contract event 0x0f5025cc4f20aa47a346d1b7d9da6ba8c68cc8e83b75e813da4b4490d55365ae.
//
// Solidity: event ORI_New_State(address indexed user, bytes32 indexed new_state)
func (_Contracts *ContractsFilterer) ParseORINewState(log types.Log) (*ContractsORINewState, error) {
	event := new(ContractsORINewState)
	if err := _Contracts.contract.UnpackLog(event, "ORI_New_State", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsORIRolledBackIterator is returned from FilterORIRolledBack and is used to iterate over the raw logs and unpacked data for ORIRolledBack events raised by the Contracts contract.
type ContractsORIRolledBackIterator struct {
	Event *ContractsORIRolledBack // Event containing the contract specifics and raw log

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
func (it *ContractsORIRolledBackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsORIRolledBack)
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
		it.Event = new(ContractsORIRolledBack)
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
func (it *ContractsORIRolledBackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsORIRolledBackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsORIRolledBack represents a ORIRolledBack event raised by the Contracts contract.
type ContractsORIRolledBack struct {
	User      common.Address
	TreeSize  *big.Int
	BlockTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterORIRolledBack is a free log retrieval operation binding the contract event 0x4d7ed8c49e6b03daee23a18f4bd14bd7e4628e5ed54c57bf84407a693867eca9.
//
// Solidity: event ORI_Rolled_Back(address indexed user, uint256 indexed tree_size, uint256 indexed block_time)
func (_Contracts *ContractsFilterer) FilterORIRolledBack(opts *bind.FilterOpts, user []common.Address, tree_size []*big.Int, block_time []*big.Int) (*ContractsORIRolledBackIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tree_sizeRule []interface{}
	for _, tree_sizeItem := range tree_size {
		tree_sizeRule = append(tree_sizeRule, tree_sizeItem)
	}
	var block_timeRule []interface{}
	for _, block_timeItem := range block_time {
		block_timeRule = append(block_timeRule, block_timeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ORI_Rolled_Back", userRule, tree_sizeRule, block_timeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsORIRolledBackIterator{contract: _Contracts.contract, event: "ORI_Rolled_Back", logs: logs, sub: sub}, nil
}

// WatchORIRolledBack is a free log subscription operation binding the contract event 0x4d7ed8c49e6b03daee23a18f4bd14bd7e4628e5ed54c57bf84407a693867eca9.
//
// Solidity: event ORI_Rolled_Back(address indexed user, uint256 indexed tree_size, uint256 indexed block_time)
func (_Contracts *ContractsFilterer) WatchORIRolledBack(opts *bind.WatchOpts, sink chan<- *ContractsORIRolledBack, user []common.Address, tree_size []*big.Int, block_time []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tree_sizeRule []interface{}
	for _, tree_sizeItem := range tree_size {
		tree_sizeRule = append(tree_sizeRule, tree_sizeItem)
	}
	var block_timeRule []interface{}
	for _, block_timeItem := range block_time {
		block_timeRule = append(block_timeRule, block_timeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ORI_Rolled_Back", userRule, tree_sizeRule, block_timeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsORIRolledBack)
				if err := _Contracts.contract.UnpackLog(event, "ORI_Rolled_Back", log); err != nil {
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

// ParseORIRolledBack is a log parse operation binding the contract event 0x4d7ed8c49e6b03daee23a18f4bd14bd7e4628e5ed54c57bf84407a693867eca9.
//
// Solidity: event ORI_Rolled_Back(address indexed user, uint256 indexed tree_size, uint256 indexed block_time)
func (_Contracts *ContractsFilterer) ParseORIRolledBack(log types.Log) (*ContractsORIRolledBack, error) {
	event := new(ContractsORIRolledBack)
	if err := _Contracts.contract.UnpackLog(event, "ORI_Rolled_Back", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsORIUnlockedIterator is returned from FilterORIUnlocked and is used to iterate over the raw logs and unpacked data for ORIUnlocked events raised by the Contracts contract.
type ContractsORIUnlockedIterator struct {
	Event *ContractsORIUnlocked // Event containing the contract specifics and raw log

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
func (it *ContractsORIUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsORIUnlocked)
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
		it.Event = new(ContractsORIUnlocked)
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
func (it *ContractsORIUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsORIUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsORIUnlocked represents a ORIUnlocked event raised by the Contracts contract.
type ContractsORIUnlocked struct {
	Suspect   common.Address
	Accuser   common.Address
	BlockTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterORIUnlocked is a free log retrieval operation binding the contract event 0xd896484469c47833eb445d62b95b9dd6b28d0c050a9b7f2ce8902d137316500e.
//
// Solidity: event ORI_Unlocked(address indexed suspect, address indexed accuser, uint256 indexed block_time)
func (_Contracts *ContractsFilterer) FilterORIUnlocked(opts *bind.FilterOpts, suspect []common.Address, accuser []common.Address, block_time []*big.Int) (*ContractsORIUnlockedIterator, error) {

	var suspectRule []interface{}
	for _, suspectItem := range suspect {
		suspectRule = append(suspectRule, suspectItem)
	}
	var accuserRule []interface{}
	for _, accuserItem := range accuser {
		accuserRule = append(accuserRule, accuserItem)
	}
	var block_timeRule []interface{}
	for _, block_timeItem := range block_time {
		block_timeRule = append(block_timeRule, block_timeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ORI_Unlocked", suspectRule, accuserRule, block_timeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsORIUnlockedIterator{contract: _Contracts.contract, event: "ORI_Unlocked", logs: logs, sub: sub}, nil
}

// WatchORIUnlocked is a free log subscription operation binding the contract event 0xd896484469c47833eb445d62b95b9dd6b28d0c050a9b7f2ce8902d137316500e.
//
// Solidity: event ORI_Unlocked(address indexed suspect, address indexed accuser, uint256 indexed block_time)
func (_Contracts *ContractsFilterer) WatchORIUnlocked(opts *bind.WatchOpts, sink chan<- *ContractsORIUnlocked, suspect []common.Address, accuser []common.Address, block_time []*big.Int) (event.Subscription, error) {

	var suspectRule []interface{}
	for _, suspectItem := range suspect {
		suspectRule = append(suspectRule, suspectItem)
	}
	var accuserRule []interface{}
	for _, accuserItem := range accuser {
		accuserRule = append(accuserRule, accuserItem)
	}
	var block_timeRule []interface{}
	for _, block_timeItem := range block_time {
		block_timeRule = append(block_timeRule, block_timeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ORI_Unlocked", suspectRule, accuserRule, block_timeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsORIUnlocked)
				if err := _Contracts.contract.UnpackLog(event, "ORI_Unlocked", log); err != nil {
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

// ParseORIUnlocked is a log parse operation binding the contract event 0xd896484469c47833eb445d62b95b9dd6b28d0c050a9b7f2ce8902d137316500e.
//
// Solidity: event ORI_Unlocked(address indexed suspect, address indexed accuser, uint256 indexed block_time)
func (_Contracts *ContractsFilterer) ParseORIUnlocked(log types.Log) (*ContractsORIUnlocked, error) {
	event := new(ContractsORIUnlocked)
	if err := _Contracts.contract.UnpackLog(event, "ORI_Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
