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

// Main3ABI is the input ABI used to generate the binding from.
const Main3ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"confirmresult\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"isbaid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ResultAddress\",\"type\":\"string\"},{\"name\":\"TransactionDetailAddress\",\"type\":\"string\"}],\"name\":\"uploadTrainRes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Payment\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"DataSchemaAddress\",\"type\":\"string\"},{\"name\":\"MetadataAddress\",\"type\":\"string\"},{\"name\":\"ModelAddress\",\"type\":\"string\"},{\"name\":\"StrategyAddress\",\"type\":\"string\"},{\"name\":\"ComputionAddress\",\"type\":\"string\"}],\"name\":\"askTraining\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"OperationSchemasAddress\",\"type\":\"string\"},{\"name\":\"ComputingAddress\",\"type\":\"string\"},{\"name\":\"ComputerAttributesAddress\",\"type\":\"string\"},{\"name\":\"_Payment\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FallbackCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"OperationSchemasAddress\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"ComputingAddress\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"ComputerAttributesAddress\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_Payment\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_this\",\"type\":\"address\"}],\"name\":\"Bidcomputing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"ResultAddress\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"TransactionDetailAddress\",\"type\":\"string\"}],\"name\":\"UploadTrainRes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"DataSchemaAddress\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"MetadataAddress\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"ModelAddress\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"StrategyAddress\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"ComputionAddress\",\"type\":\"string\"}],\"name\":\"AskTraining\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Aborted\",\"type\":\"event\"}]"

// Main3Bin is the compiled bytecode used for deploying new contracts.
const Main3Bin = `608060405234801561001057600080fd5b50604051610bb5380380610bb583398101806040528101908080518201929190602001805182019291906020018051820192919060200180519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600560006101000a81548160ff021916908360018111156100b357fe5b0217905550806001819055507f912cc211503ae8ca8d50e63e4add398570877aece91348b63d9d2ec238c516398484848430604051808060200180602001806020018681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001848103845289818151815260200191508051906020019080838360005b8381101561016557808201518184015260208101905061014a565b50505050905090810190601f1680156101925780820380516001836020036101000a031916815260200191505b50848103835288818151815260200191508051906020019080838360005b838110156101cb5780820151818401526020810190506101b0565b50505050905090810190601f1680156101f85780820380516001836020036101000a031916815260200191505b50848103825287818151815260200191508051906020019080838360005b83811015610231578082015181840152602081019050610216565b50505050905090810190601f16801561025e5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390a150505050610931806102846000396000f300608060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806308551a5314610106578063200d2ed21461015d57806335a063b4146101885780633d49fb031461019f5780637c193d11146101b65780637dc48c0514610211578063a33b1260146102d4578063c19d93fb146102ff578063f75e149e14610338575b7f4fdfaf7cd2ae67f3639149bddef2addc991d3f26e0121df03517ec62523649ca3334604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1005b34801561011257600080fd5b5061011b6104c0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016957600080fd5b506101726104e5565b6040518082815260200191505060405180910390f35b34801561019457600080fd5b5061019d6104eb565b005b3480156101ab57600080fd5b506101b461064d565b005b3480156101c257600080fd5b506101f7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610711565b604051808215151515815260200191505060405180910390f35b34801561021d57600080fd5b506102be600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610767565b6040518082815260200191505060405180910390f35b3480156102e057600080fd5b506102e961087c565b6040518082815260200191505060405180910390f35b34801561030b57600080fd5b50610314610882565b6040518082600181111561032457fe5b60ff16815260200191505060405180910390f35b6104aa600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610895565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561054657600080fd5b600080600181111561055457fe5b600560009054906101000a900460ff16600181111561056f57fe5b14151561057b57600080fd5b7f72c874aeff0b183a56e2b79c71b46e1aed4dee5e09862134b8821ba2fddbf8bf60405160405180910390a16001600560006101000a81548160ff021916908360018111156105c657fe5b02179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f19350505050158015610649573d6000803e3d6000fd5b5050565b6001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6001549081150290604051600060405180830381858888f1935050505015801561070e573d6000803e3d6000fd5b50565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60007f58cc13e7b32e0dc6c686ca61dc32f63b4c1eb79e751fad3e3ac87cceec1065458383604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156107d05780820151818401526020810190506107b5565b50505050905090810190601f1680156107fd5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561083657808201518184015260208101905061081b565b50505050905090810190601f1680156108635780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a160c8905092915050565b60015481565b600560009054906101000a900460ff1681565b60006001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060c86002819055506002549050959450505050505600a165627a7a72305820d906adeeba887366862331edeed2912182357c59f3c98c3dbb2ecfcf6daa5bf20029`

// DeployMain3 deploys a new Ethereum contract, binding an instance of Main3 to it.
func DeployMain3(auth *bind.TransactOpts, backend bind.ContractBackend, OperationSchemasAddress string, ComputingAddress string, ComputerAttributesAddress string, _Payment *big.Int) (common.Address, *types.Transaction, *Main3, error) {
	parsed, err := abi.JSON(strings.NewReader(Main3ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Main3Bin), backend, OperationSchemasAddress, ComputingAddress, ComputerAttributesAddress, _Payment)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main3{Main3Caller: Main3Caller{contract: contract}, Main3Transactor: Main3Transactor{contract: contract}, Main3Filterer: Main3Filterer{contract: contract}}, nil
}

// Main3 is an auto generated Go binding around an Ethereum contract.
type Main3 struct {
	Main3Caller     // Read-only binding to the contract
	Main3Transactor // Write-only binding to the contract
	Main3Filterer   // Log filterer for contract events
}

// Main3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Main3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Main3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Main3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Main3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Main3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Main3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Main3Session struct {
	Contract     *Main3            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Main3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Main3CallerSession struct {
	Contract *Main3Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Main3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Main3TransactorSession struct {
	Contract     *Main3Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Main3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Main3Raw struct {
	Contract *Main3 // Generic contract binding to access the raw methods on
}

// Main3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Main3CallerRaw struct {
	Contract *Main3Caller // Generic read-only contract binding to access the raw methods on
}

// Main3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Main3TransactorRaw struct {
	Contract *Main3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMain3 creates a new instance of Main3, bound to a specific deployed contract.
func NewMain3(address common.Address, backend bind.ContractBackend) (*Main3, error) {
	contract, err := bindMain3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main3{Main3Caller: Main3Caller{contract: contract}, Main3Transactor: Main3Transactor{contract: contract}, Main3Filterer: Main3Filterer{contract: contract}}, nil
}

// NewMain3Caller creates a new read-only instance of Main3, bound to a specific deployed contract.
func NewMain3Caller(address common.Address, caller bind.ContractCaller) (*Main3Caller, error) {
	contract, err := bindMain3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Main3Caller{contract: contract}, nil
}

// NewMain3Transactor creates a new write-only instance of Main3, bound to a specific deployed contract.
func NewMain3Transactor(address common.Address, transactor bind.ContractTransactor) (*Main3Transactor, error) {
	contract, err := bindMain3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Main3Transactor{contract: contract}, nil
}

// NewMain3Filterer creates a new log filterer instance of Main3, bound to a specific deployed contract.
func NewMain3Filterer(address common.Address, filterer bind.ContractFilterer) (*Main3Filterer, error) {
	contract, err := bindMain3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Main3Filterer{contract: contract}, nil
}

// bindMain3 binds a generic wrapper to an already deployed contract.
func bindMain3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Main3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main3 *Main3Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main3.Contract.Main3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main3 *Main3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main3.Contract.Main3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main3 *Main3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main3.Contract.Main3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main3 *Main3CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main3 *Main3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main3 *Main3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main3.Contract.contract.Transact(opts, method, params...)
}

// Payment is a free data retrieval call binding the contract method 0xa33b1260.
//
// Solidity: function Payment() constant returns(uint256)
func (_Main3 *Main3Caller) Payment(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main3.contract.Call(opts, out, "Payment")
	return *ret0, err
}

// Payment is a free data retrieval call binding the contract method 0xa33b1260.
//
// Solidity: function Payment() constant returns(uint256)
func (_Main3 *Main3Session) Payment() (*big.Int, error) {
	return _Main3.Contract.Payment(&_Main3.CallOpts)
}

// Payment is a free data retrieval call binding the contract method 0xa33b1260.
//
// Solidity: function Payment() constant returns(uint256)
func (_Main3 *Main3CallerSession) Payment() (*big.Int, error) {
	return _Main3.Contract.Payment(&_Main3.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() constant returns(address)
func (_Main3 *Main3Caller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Main3.contract.Call(opts, out, "seller")
	return *ret0, err
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() constant returns(address)
func (_Main3 *Main3Session) Seller() (common.Address, error) {
	return _Main3.Contract.Seller(&_Main3.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() constant returns(address)
func (_Main3 *Main3CallerSession) Seller() (common.Address, error) {
	return _Main3.Contract.Seller(&_Main3.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Main3 *Main3Caller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Main3.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Main3 *Main3Session) State() (uint8, error) {
	return _Main3.Contract.State(&_Main3.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Main3 *Main3CallerSession) State() (uint8, error) {
	return _Main3.Contract.State(&_Main3.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(uint256)
func (_Main3 *Main3Caller) Status(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main3.contract.Call(opts, out, "status")
	return *ret0, err
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(uint256)
func (_Main3 *Main3Session) Status() (*big.Int, error) {
	return _Main3.Contract.Status(&_Main3.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(uint256)
func (_Main3 *Main3CallerSession) Status() (*big.Int, error) {
	return _Main3.Contract.Status(&_Main3.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Main3 *Main3Transactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main3.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Main3 *Main3Session) Abort() (*types.Transaction, error) {
	return _Main3.Contract.Abort(&_Main3.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Main3 *Main3TransactorSession) Abort() (*types.Transaction, error) {
	return _Main3.Contract.Abort(&_Main3.TransactOpts)
}

// AskTraining is a paid mutator transaction binding the contract method 0xf75e149e.
//
// Solidity: function askTraining(DataSchemaAddress string, MetadataAddress string, ModelAddress string, StrategyAddress string, ComputionAddress string) returns(uint256)
func (_Main3 *Main3Transactor) AskTraining(opts *bind.TransactOpts, DataSchemaAddress string, MetadataAddress string, ModelAddress string, StrategyAddress string, ComputionAddress string) (*types.Transaction, error) {
	return _Main3.contract.Transact(opts, "askTraining", DataSchemaAddress, MetadataAddress, ModelAddress, StrategyAddress, ComputionAddress)
}

// AskTraining is a paid mutator transaction binding the contract method 0xf75e149e.
//
// Solidity: function askTraining(DataSchemaAddress string, MetadataAddress string, ModelAddress string, StrategyAddress string, ComputionAddress string) returns(uint256)
func (_Main3 *Main3Session) AskTraining(DataSchemaAddress string, MetadataAddress string, ModelAddress string, StrategyAddress string, ComputionAddress string) (*types.Transaction, error) {
	return _Main3.Contract.AskTraining(&_Main3.TransactOpts, DataSchemaAddress, MetadataAddress, ModelAddress, StrategyAddress, ComputionAddress)
}

// AskTraining is a paid mutator transaction binding the contract method 0xf75e149e.
//
// Solidity: function askTraining(DataSchemaAddress string, MetadataAddress string, ModelAddress string, StrategyAddress string, ComputionAddress string) returns(uint256)
func (_Main3 *Main3TransactorSession) AskTraining(DataSchemaAddress string, MetadataAddress string, ModelAddress string, StrategyAddress string, ComputionAddress string) (*types.Transaction, error) {
	return _Main3.Contract.AskTraining(&_Main3.TransactOpts, DataSchemaAddress, MetadataAddress, ModelAddress, StrategyAddress, ComputionAddress)
}

// Confirmresult is a paid mutator transaction binding the contract method 0x3d49fb03.
//
// Solidity: function confirmresult() returns()
func (_Main3 *Main3Transactor) Confirmresult(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main3.contract.Transact(opts, "confirmresult")
}

// Confirmresult is a paid mutator transaction binding the contract method 0x3d49fb03.
//
// Solidity: function confirmresult() returns()
func (_Main3 *Main3Session) Confirmresult() (*types.Transaction, error) {
	return _Main3.Contract.Confirmresult(&_Main3.TransactOpts)
}

// Confirmresult is a paid mutator transaction binding the contract method 0x3d49fb03.
//
// Solidity: function confirmresult() returns()
func (_Main3 *Main3TransactorSession) Confirmresult() (*types.Transaction, error) {
	return _Main3.Contract.Confirmresult(&_Main3.TransactOpts)
}

// Isbaid is a paid mutator transaction binding the contract method 0x7c193d11.
//
// Solidity: function isbaid(buyer address) returns(bool)
func (_Main3 *Main3Transactor) Isbaid(opts *bind.TransactOpts, buyer common.Address) (*types.Transaction, error) {
	return _Main3.contract.Transact(opts, "isbaid", buyer)
}

// Isbaid is a paid mutator transaction binding the contract method 0x7c193d11.
//
// Solidity: function isbaid(buyer address) returns(bool)
func (_Main3 *Main3Session) Isbaid(buyer common.Address) (*types.Transaction, error) {
	return _Main3.Contract.Isbaid(&_Main3.TransactOpts, buyer)
}

// Isbaid is a paid mutator transaction binding the contract method 0x7c193d11.
//
// Solidity: function isbaid(buyer address) returns(bool)
func (_Main3 *Main3TransactorSession) Isbaid(buyer common.Address) (*types.Transaction, error) {
	return _Main3.Contract.Isbaid(&_Main3.TransactOpts, buyer)
}

// UploadTrainRes is a paid mutator transaction binding the contract method 0x7dc48c05.
//
// Solidity: function uploadTrainRes(ResultAddress string, TransactionDetailAddress string) returns(uint256)
func (_Main3 *Main3Transactor) UploadTrainRes(opts *bind.TransactOpts, ResultAddress string, TransactionDetailAddress string) (*types.Transaction, error) {
	return _Main3.contract.Transact(opts, "uploadTrainRes", ResultAddress, TransactionDetailAddress)
}

// UploadTrainRes is a paid mutator transaction binding the contract method 0x7dc48c05.
//
// Solidity: function uploadTrainRes(ResultAddress string, TransactionDetailAddress string) returns(uint256)
func (_Main3 *Main3Session) UploadTrainRes(ResultAddress string, TransactionDetailAddress string) (*types.Transaction, error) {
	return _Main3.Contract.UploadTrainRes(&_Main3.TransactOpts, ResultAddress, TransactionDetailAddress)
}

// UploadTrainRes is a paid mutator transaction binding the contract method 0x7dc48c05.
//
// Solidity: function uploadTrainRes(ResultAddress string, TransactionDetailAddress string) returns(uint256)
func (_Main3 *Main3TransactorSession) UploadTrainRes(ResultAddress string, TransactionDetailAddress string) (*types.Transaction, error) {
	return _Main3.Contract.UploadTrainRes(&_Main3.TransactOpts, ResultAddress, TransactionDetailAddress)
}

// Main3AbortedIterator is returned from FilterAborted and is used to iterate over the raw logs and unpacked data for Aborted events raised by the Main3 contract.
type Main3AbortedIterator struct {
	Event *Main3Aborted // Event containing the contract specifics and raw log

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
func (it *Main3AbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Main3Aborted)
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
		it.Event = new(Main3Aborted)
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
func (it *Main3AbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Main3AbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Main3Aborted represents a Aborted event raised by the Main3 contract.
type Main3Aborted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAborted is a free log retrieval operation binding the contract event 0x72c874aeff0b183a56e2b79c71b46e1aed4dee5e09862134b8821ba2fddbf8bf.
//
// Solidity: e Aborted()
func (_Main3 *Main3Filterer) FilterAborted(opts *bind.FilterOpts) (*Main3AbortedIterator, error) {

	logs, sub, err := _Main3.contract.FilterLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return &Main3AbortedIterator{contract: _Main3.contract, event: "Aborted", logs: logs, sub: sub}, nil
}

// WatchAborted is a free log subscription operation binding the contract event 0x72c874aeff0b183a56e2b79c71b46e1aed4dee5e09862134b8821ba2fddbf8bf.
//
// Solidity: e Aborted()
func (_Main3 *Main3Filterer) WatchAborted(opts *bind.WatchOpts, sink chan<- *Main3Aborted) (event.Subscription, error) {

	logs, sub, err := _Main3.contract.WatchLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Main3Aborted)
				if err := _Main3.contract.UnpackLog(event, "Aborted", log); err != nil {
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

// Main3AskTrainingIterator is returned from FilterAskTraining and is used to iterate over the raw logs and unpacked data for AskTraining events raised by the Main3 contract.
type Main3AskTrainingIterator struct {
	Event *Main3AskTraining // Event containing the contract specifics and raw log

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
func (it *Main3AskTrainingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Main3AskTraining)
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
		it.Event = new(Main3AskTraining)
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
func (it *Main3AskTrainingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Main3AskTrainingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Main3AskTraining represents a AskTraining event raised by the Main3 contract.
type Main3AskTraining struct {
	DataSchemaAddress string
	MetadataAddress   string
	ModelAddress      string
	StrategyAddress   string
	ComputionAddress  string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAskTraining is a free log retrieval operation binding the contract event 0xe11530c50043630fa2e0173e42a425e9d44cc904f976afcd503f4fb3ecb228f3.
//
// Solidity: e AskTraining(DataSchemaAddress string, MetadataAddress string, ModelAddress string, StrategyAddress string, ComputionAddress string)
func (_Main3 *Main3Filterer) FilterAskTraining(opts *bind.FilterOpts) (*Main3AskTrainingIterator, error) {

	logs, sub, err := _Main3.contract.FilterLogs(opts, "AskTraining")
	if err != nil {
		return nil, err
	}
	return &Main3AskTrainingIterator{contract: _Main3.contract, event: "AskTraining", logs: logs, sub: sub}, nil
}

// WatchAskTraining is a free log subscription operation binding the contract event 0xe11530c50043630fa2e0173e42a425e9d44cc904f976afcd503f4fb3ecb228f3.
//
// Solidity: e AskTraining(DataSchemaAddress string, MetadataAddress string, ModelAddress string, StrategyAddress string, ComputionAddress string)
func (_Main3 *Main3Filterer) WatchAskTraining(opts *bind.WatchOpts, sink chan<- *Main3AskTraining) (event.Subscription, error) {

	logs, sub, err := _Main3.contract.WatchLogs(opts, "AskTraining")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Main3AskTraining)
				if err := _Main3.contract.UnpackLog(event, "AskTraining", log); err != nil {
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

// Main3BidcomputingIterator is returned from FilterBidcomputing and is used to iterate over the raw logs and unpacked data for Bidcomputing events raised by the Main3 contract.
type Main3BidcomputingIterator struct {
	Event *Main3Bidcomputing // Event containing the contract specifics and raw log

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
func (it *Main3BidcomputingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Main3Bidcomputing)
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
		it.Event = new(Main3Bidcomputing)
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
func (it *Main3BidcomputingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Main3BidcomputingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Main3Bidcomputing represents a Bidcomputing event raised by the Main3 contract.
type Main3Bidcomputing struct {
	OperationSchemasAddress   string
	ComputingAddress          string
	ComputerAttributesAddress string
	Payment                   *big.Int
	This                      common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterBidcomputing is a free log retrieval operation binding the contract event 0x912cc211503ae8ca8d50e63e4add398570877aece91348b63d9d2ec238c51639.
//
// Solidity: e Bidcomputing(OperationSchemasAddress string, ComputingAddress string, ComputerAttributesAddress string, _Payment uint256, _this address)
func (_Main3 *Main3Filterer) FilterBidcomputing(opts *bind.FilterOpts) (*Main3BidcomputingIterator, error) {

	logs, sub, err := _Main3.contract.FilterLogs(opts, "Bidcomputing")
	if err != nil {
		return nil, err
	}
	return &Main3BidcomputingIterator{contract: _Main3.contract, event: "Bidcomputing", logs: logs, sub: sub}, nil
}

// WatchBidcomputing is a free log subscription operation binding the contract event 0x912cc211503ae8ca8d50e63e4add398570877aece91348b63d9d2ec238c51639.
//
// Solidity: e Bidcomputing(OperationSchemasAddress string, ComputingAddress string, ComputerAttributesAddress string, _Payment uint256, _this address)
func (_Main3 *Main3Filterer) WatchBidcomputing(opts *bind.WatchOpts, sink chan<- *Main3Bidcomputing) (event.Subscription, error) {

	logs, sub, err := _Main3.contract.WatchLogs(opts, "Bidcomputing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Main3Bidcomputing)
				if err := _Main3.contract.UnpackLog(event, "Bidcomputing", log); err != nil {
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

// Main3FallbackCalledIterator is returned from FilterFallbackCalled and is used to iterate over the raw logs and unpacked data for FallbackCalled events raised by the Main3 contract.
type Main3FallbackCalledIterator struct {
	Event *Main3FallbackCalled // Event containing the contract specifics and raw log

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
func (it *Main3FallbackCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Main3FallbackCalled)
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
		it.Event = new(Main3FallbackCalled)
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
func (it *Main3FallbackCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Main3FallbackCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Main3FallbackCalled represents a FallbackCalled event raised by the Main3 contract.
type Main3FallbackCalled struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFallbackCalled is a free log retrieval operation binding the contract event 0x4fdfaf7cd2ae67f3639149bddef2addc991d3f26e0121df03517ec62523649ca.
//
// Solidity: e FallbackCalled(from address, amount uint256)
func (_Main3 *Main3Filterer) FilterFallbackCalled(opts *bind.FilterOpts) (*Main3FallbackCalledIterator, error) {

	logs, sub, err := _Main3.contract.FilterLogs(opts, "FallbackCalled")
	if err != nil {
		return nil, err
	}
	return &Main3FallbackCalledIterator{contract: _Main3.contract, event: "FallbackCalled", logs: logs, sub: sub}, nil
}

// WatchFallbackCalled is a free log subscription operation binding the contract event 0x4fdfaf7cd2ae67f3639149bddef2addc991d3f26e0121df03517ec62523649ca.
//
// Solidity: e FallbackCalled(from address, amount uint256)
func (_Main3 *Main3Filterer) WatchFallbackCalled(opts *bind.WatchOpts, sink chan<- *Main3FallbackCalled) (event.Subscription, error) {

	logs, sub, err := _Main3.contract.WatchLogs(opts, "FallbackCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Main3FallbackCalled)
				if err := _Main3.contract.UnpackLog(event, "FallbackCalled", log); err != nil {
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

// Main3UploadTrainResIterator is returned from FilterUploadTrainRes and is used to iterate over the raw logs and unpacked data for UploadTrainRes events raised by the Main3 contract.
type Main3UploadTrainResIterator struct {
	Event *Main3UploadTrainRes // Event containing the contract specifics and raw log

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
func (it *Main3UploadTrainResIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Main3UploadTrainRes)
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
		it.Event = new(Main3UploadTrainRes)
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
func (it *Main3UploadTrainResIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Main3UploadTrainResIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Main3UploadTrainRes represents a UploadTrainRes event raised by the Main3 contract.
type Main3UploadTrainRes struct {
	ResultAddress            string
	TransactionDetailAddress string
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterUploadTrainRes is a free log retrieval operation binding the contract event 0x58cc13e7b32e0dc6c686ca61dc32f63b4c1eb79e751fad3e3ac87cceec106545.
//
// Solidity: e UploadTrainRes(ResultAddress string, TransactionDetailAddress string)
func (_Main3 *Main3Filterer) FilterUploadTrainRes(opts *bind.FilterOpts) (*Main3UploadTrainResIterator, error) {

	logs, sub, err := _Main3.contract.FilterLogs(opts, "UploadTrainRes")
	if err != nil {
		return nil, err
	}
	return &Main3UploadTrainResIterator{contract: _Main3.contract, event: "UploadTrainRes", logs: logs, sub: sub}, nil
}

// WatchUploadTrainRes is a free log subscription operation binding the contract event 0x58cc13e7b32e0dc6c686ca61dc32f63b4c1eb79e751fad3e3ac87cceec106545.
//
// Solidity: e UploadTrainRes(ResultAddress string, TransactionDetailAddress string)
func (_Main3 *Main3Filterer) WatchUploadTrainRes(opts *bind.WatchOpts, sink chan<- *Main3UploadTrainRes) (event.Subscription, error) {

	logs, sub, err := _Main3.contract.WatchLogs(opts, "UploadTrainRes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Main3UploadTrainRes)
				if err := _Main3.contract.UnpackLog(event, "UploadTrainRes", log); err != nil {
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
