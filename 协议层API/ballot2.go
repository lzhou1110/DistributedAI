// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Main2ABI is the input ABI used to generate the binding from.
const Main2ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getIpfsHashByPool\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"ipfsHash\",\"type\":\"string\"}],\"name\":\"add2Ipfspool\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ipfsHashs\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sayHello\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Main2Bin is the compiled bytecode used for deploying new contracts.
const Main2Bin = `608060405234801561001057600080fd5b50610ac1806100206000396000f300608060405260043610610083576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630864f21614610088578063316bc9ac1461016a57806373d4a13a1461028557806382ab890a14610315578063b69ef8a81461037c578063dcd8f9cf146103a7578063ef5fb05b1461044d575b600080fd5b34801561009457600080fd5b506100ef600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506104d0565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561012f578082015181840152602081019050610114565b50505050905090810190601f16801561015c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61020a600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061066a565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561024a57808201518184015260208101905061022f565b50505050905090810190601f1680156102775780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561029157600080fd5b5061029a6107ae565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102da5780820151818401526020810190506102bf565b50505050905090810190601f1680156103075780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103336004803603810190808035906020019092919050505061084c565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561038857600080fd5b5061039161086b565b6040518082815260200191505060405180910390f35b3480156103b357600080fd5b506103d260048036038101908080359060200190929190505050610871565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104125780820151818401526020810190506103f7565b50505050905090810190601f16801561043f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61045561092c565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561049557808201518184015260208101905061047a565b50505050905090810190601f1680156104c25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60606003826040518082805190602001908083835b60208310151561050a57805182526020820191506020810190506020830392506104e5565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206001908054600181600116156101000203166002900461055e929190610969565b506003826040518082805190602001908083835b6020831015156105975780518252602082019150602081019050602083039250610572565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561065e5780601f106106335761010080835404028352916020019161065e565b820191906000526020600020905b81548152906001019060200180831161064157829003601f168201915b50505050509050919050565b6060816003846040518082805190602001908083835b6020831015156106a55780518252602082019150602081019050602083039250610680565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902090805190602001906106eb9291906109f0565b506000600283908060018154018082558091505090600182039060005260206000200160009091929091909150908051906020019061072b9291906109f0565b50111561076f576040805190810160405280600281526020017f6f6b00000000000000000000000000000000000000000000000000000000000081525090506107a8565b6040805190810160405280600481526020017f6661696c0000000000000000000000000000000000000000000000000000000081525090505b92915050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108445780601f1061081957610100808354040283529160200191610844565b820191906000526020600020905b81548152906001019060200180831161082757829003601f168201915b505050505081565b6000808260008082825401925050819055503360005491509150915091565b60005481565b60028181548110151561088057fe5b906000526020600020016000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109245780601f106108f957610100808354040283529160200191610924565b820191906000526020600020905b81548152906001019060200180831161090757829003601f168201915b505050505081565b60606040805190810160405280600b81526020017f68656c6c6f20776f726c64000000000000000000000000000000000000000000815250905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109a257805485556109df565b828001600101855582156109df57600052602060002091601f016020900482015b828111156109de5782548255916001019190600101906109c3565b5b5090506109ec9190610a70565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a3157805160ff1916838001178555610a5f565b82800160010185558215610a5f579182015b82811115610a5e578251825591602001919060010190610a43565b5b509050610a6c9190610a70565b5090565b610a9291905b80821115610a8e576000816000905550600101610a76565b5090565b905600a165627a7a723058200ca12051c2e477554034338def32f444bdaf29edaee22b338d72f4bcbb34cdfe0029`

// DeployMain2 deploys a new Ethereum contract, binding an instance of Main2 to it.
func DeployMain2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Main2, error) {
	parsed, err := abi.JSON(strings.NewReader(Main2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Main2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main2{Main2Caller: Main2Caller{contract: contract}, Main2Transactor: Main2Transactor{contract: contract}, Main2Filterer: Main2Filterer{contract: contract}}, nil
}

// Main2 is an auto generated Go binding around an Ethereum contract.
type Main2 struct {
	Main2Caller     // Read-only binding to the contract
	Main2Transactor // Write-only binding to the contract
	Main2Filterer   // Log filterer for contract events
}

// Main2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Main2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Main2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Main2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Main2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Main2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Main2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Main2Session struct {
	Contract     *Main2            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Main2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Main2CallerSession struct {
	Contract *Main2Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Main2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Main2TransactorSession struct {
	Contract     *Main2Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Main2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Main2Raw struct {
	Contract *Main2 // Generic contract binding to access the raw methods on
}

// Main2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Main2CallerRaw struct {
	Contract *Main2Caller // Generic read-only contract binding to access the raw methods on
}

// Main2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Main2TransactorRaw struct {
	Contract *Main2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMain2 creates a new instance of Main2, bound to a specific deployed contract.
func NewMain2(address common.Address, backend bind.ContractBackend) (*Main2, error) {
	contract, err := bindMain2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main2{Main2Caller: Main2Caller{contract: contract}, Main2Transactor: Main2Transactor{contract: contract}, Main2Filterer: Main2Filterer{contract: contract}}, nil
}

// NewMain2Caller creates a new read-only instance of Main2, bound to a specific deployed contract.
func NewMain2Caller(address common.Address, caller bind.ContractCaller) (*Main2Caller, error) {
	contract, err := bindMain2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Main2Caller{contract: contract}, nil
}

// NewMain2Transactor creates a new write-only instance of Main2, bound to a specific deployed contract.
func NewMain2Transactor(address common.Address, transactor bind.ContractTransactor) (*Main2Transactor, error) {
	contract, err := bindMain2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Main2Transactor{contract: contract}, nil
}

// NewMain2Filterer creates a new log filterer instance of Main2, bound to a specific deployed contract.
func NewMain2Filterer(address common.Address, filterer bind.ContractFilterer) (*Main2Filterer, error) {
	contract, err := bindMain2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Main2Filterer{contract: contract}, nil
}

// bindMain2 binds a generic wrapper to an already deployed contract.
func bindMain2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Main2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main2 *Main2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main2.Contract.Main2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main2 *Main2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main2.Contract.Main2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main2 *Main2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main2.Contract.Main2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main2 *Main2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main2 *Main2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main2 *Main2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main2.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_Main2 *Main2Caller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main2.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_Main2 *Main2Session) Balance() (*big.Int, error) {
	return _Main2.Contract.Balance(&_Main2.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_Main2 *Main2CallerSession) Balance() (*big.Int, error) {
	return _Main2.Contract.Balance(&_Main2.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Main2 *Main2Caller) Data(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Main2.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Main2 *Main2Session) Data() (string, error) {
	return _Main2.Contract.Data(&_Main2.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Main2 *Main2CallerSession) Data() (string, error) {
	return _Main2.Contract.Data(&_Main2.CallOpts)
}

// GetIpfsHashByPool is a free data retrieval call binding the contract method 0x0864f216.
//
// Solidity: function getIpfsHashByPool(name string) constant returns(string)
func (_Main2 *Main2Caller) GetIpfsHashByPool(opts *bind.CallOpts, name string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Main2.contract.Call(opts, out, "getIpfsHashByPool", name)
	return *ret0, err
}

// GetIpfsHashByPool is a free data retrieval call binding the contract method 0x0864f216.
//
// Solidity: function getIpfsHashByPool(name string) constant returns(string)
func (_Main2 *Main2Session) GetIpfsHashByPool(name string) (string, error) {
	return _Main2.Contract.GetIpfsHashByPool(&_Main2.CallOpts, name)
}

// GetIpfsHashByPool is a free data retrieval call binding the contract method 0x0864f216.
//
// Solidity: function getIpfsHashByPool(name string) constant returns(string)
func (_Main2 *Main2CallerSession) GetIpfsHashByPool(name string) (string, error) {
	return _Main2.Contract.GetIpfsHashByPool(&_Main2.CallOpts, name)
}

// IpfsHashs is a free data retrieval call binding the contract method 0xdcd8f9cf.
//
// Solidity: function ipfsHashs( uint256) constant returns(string)
func (_Main2 *Main2Caller) IpfsHashs(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Main2.contract.Call(opts, out, "ipfsHashs", arg0)
	return *ret0, err
}

// IpfsHashs is a free data retrieval call binding the contract method 0xdcd8f9cf.
//
// Solidity: function ipfsHashs( uint256) constant returns(string)
func (_Main2 *Main2Session) IpfsHashs(arg0 *big.Int) (string, error) {
	return _Main2.Contract.IpfsHashs(&_Main2.CallOpts, arg0)
}

// IpfsHashs is a free data retrieval call binding the contract method 0xdcd8f9cf.
//
// Solidity: function ipfsHashs( uint256) constant returns(string)
func (_Main2 *Main2CallerSession) IpfsHashs(arg0 *big.Int) (string, error) {
	return _Main2.Contract.IpfsHashs(&_Main2.CallOpts, arg0)
}

// Add2Ipfspool is a paid mutator transaction binding the contract method 0x316bc9ac.
//
// Solidity: function add2Ipfspool(name string, ipfsHash string) returns(string)
func (_Main2 *Main2Transactor) Add2Ipfspool(opts *bind.TransactOpts, name string, ipfsHash string) (*types.Transaction, error) {
	return _Main2.contract.Transact(opts, "add2Ipfspool", name, ipfsHash)
}

// Add2Ipfspool is a paid mutator transaction binding the contract method 0x316bc9ac.
//
// Solidity: function add2Ipfspool(name string, ipfsHash string) returns(string)
func (_Main2 *Main2Session) Add2Ipfspool(name string, ipfsHash string) (*types.Transaction, error) {
	return _Main2.Contract.Add2Ipfspool(&_Main2.TransactOpts, name, ipfsHash)
}

// Add2Ipfspool is a paid mutator transaction binding the contract method 0x316bc9ac.
//
// Solidity: function add2Ipfspool(name string, ipfsHash string) returns(string)
func (_Main2 *Main2TransactorSession) Add2Ipfspool(name string, ipfsHash string) (*types.Transaction, error) {
	return _Main2.Contract.Add2Ipfspool(&_Main2.TransactOpts, name, ipfsHash)
}

// SayHello is a paid mutator transaction binding the contract method 0xef5fb05b.
//
// Solidity: function sayHello() returns(string)
func (_Main2 *Main2Transactor) SayHello(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main2.contract.Transact(opts, "sayHello")
}

// SayHello is a paid mutator transaction binding the contract method 0xef5fb05b.
//
// Solidity: function sayHello() returns(string)
func (_Main2 *Main2Session) SayHello() (*types.Transaction, error) {
	return _Main2.Contract.SayHello(&_Main2.TransactOpts)
}

// SayHello is a paid mutator transaction binding the contract method 0xef5fb05b.
//
// Solidity: function sayHello() returns(string)
func (_Main2 *Main2TransactorSession) SayHello() (*types.Transaction, error) {
	return _Main2.Contract.SayHello(&_Main2.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(amount uint256) returns(address, uint256)
func (_Main2 *Main2Transactor) Update(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Main2.contract.Transact(opts, "update", amount)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(amount uint256) returns(address, uint256)
func (_Main2 *Main2Session) Update(amount *big.Int) (*types.Transaction, error) {
	return _Main2.Contract.Update(&_Main2.TransactOpts, amount)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(amount uint256) returns(address, uint256)
func (_Main2 *Main2TransactorSession) Update(amount *big.Int) (*types.Transaction, error) {
	return _Main2.Contract.Update(&_Main2.TransactOpts, amount)
}
