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

// Main4ABI is the input ABI used to generate the binding from.
const Main4ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buydata\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"isbaid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Payment\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"KeyAddress\",\"type\":\"string\"},{\"name\":\"DataSchemaAddress\",\"type\":\"string\"},{\"name\":\"computionAddress\",\"type\":\"string\"},{\"name\":\"TransactionDetailAddress\",\"type\":\"string\"}],\"name\":\"askfordata\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"DataSchemaAddress\",\"type\":\"string\"},{\"name\":\"MataDataAddress\",\"type\":\"string\"},{\"name\":\"_Payment\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FallbackCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"DataSchemaAddress\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"MataDataAddress\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"Payment\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_this\",\"type\":\"address\"}],\"name\":\"BidData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Aborted\",\"type\":\"event\"}]"

// Main4Bin is the compiled bytecode used for deploying new contracts.
const Main4Bin = `608060405234801561001057600080fd5b506040516108a33803806108a3833981018060405281019080805182019291906020018051820192919060200180519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600460006101000a81548160ff021916908360018111156100a957fe5b0217905550806001819055507fb1ce085c6a9a9a2c2a4af0e92549b3e7cafa61e5bd6ac98f57c384720aaa37cf8383600154306040518080602001806020018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838103835287818151815260200191508051906020019080838360005b8381101561015857808201518184015260208101905061013d565b50505050905090810190601f1680156101855780820380516001836020036101000a031916815260200191505b50838103825286818151815260200191508051906020019080838360005b838110156101be5780820151818401526020810190506101a3565b50505050905090810190601f1680156101eb5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a15050506106958061020e6000396000f30060806040526004361061008e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806308551a53146100fb578063200d2ed21461015257806335a063b4146101815780634257630e146101985780637c193d11146101ba578063a33b126014610215578063c19d93fb14610240578063d4a381b214610279575b7f4fdfaf7cd2ae67f3639149bddef2addc991d3f26e0121df03517ec62523649ca3334604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1005b34801561010757600080fd5b506101106103c8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561015e57600080fd5b506101676103ed565b604051808215151515815260200191505060405180910390f35b34801561018d57600080fd5b50610196610400565b005b6101a0610562565b604051808215151515815260200191505060405180910390f35b3480156101c657600080fd5b506101fb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105ec565b604051808215151515815260200191505060405180910390f35b34801561022157600080fd5b5061022a610642565b6040518082815260200191505060405180910390f35b34801561024c57600080fd5b50610255610648565b6040518082600181111561026557fe5b60ff16815260200191505060405180910390f35b34801561028557600080fd5b506103b2600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061065b565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561045b57600080fd5b600080600181111561046957fe5b600460009054906101000a900460ff16600181111561048457fe5b14151561049057600080fd5b7f72c874aeff0b183a56e2b79c71b46e1aed4dee5e09862134b8821ba2fddbf8bf60405160405180910390a16001600460006101000a81548160ff021916908360018111156104db57fe5b02179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f1935050505015801561055e573d6000803e3d6000fd5b5050565b60006001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600260006101000a81548160ff021916908315150217905550600260009054906101000a900460ff16905090565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60015481565b600460009054906101000a900460ff1681565b600060c890509493505050505600a165627a7a723058206b15db539e55bbd1c22dbdadeec1f990559834fa958bbee9fb78d755bf334ccf0029`

// DeployMain4 deploys a new Ethereum contract, binding an instance of Main4 to it.
func DeployMain4(auth *bind.TransactOpts, backend bind.ContractBackend, DataSchemaAddress string, MataDataAddress string, _Payment *big.Int) (common.Address, *types.Transaction, *Main4, error) {
	parsed, err := abi.JSON(strings.NewReader(Main4ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Main4Bin), backend, DataSchemaAddress, MataDataAddress, _Payment)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main4{Main4Caller: Main4Caller{contract: contract}, Main4Transactor: Main4Transactor{contract: contract}, Main4Filterer: Main4Filterer{contract: contract}}, nil
}

// Main4 is an auto generated Go binding around an Ethereum contract.
type Main4 struct {
	Main4Caller     // Read-only binding to the contract
	Main4Transactor // Write-only binding to the contract
	Main4Filterer   // Log filterer for contract events
}

// Main4Caller is an auto generated read-only Go binding around an Ethereum contract.
type Main4Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Main4Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Main4Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Main4Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Main4Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Main4Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Main4Session struct {
	Contract     *Main4            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Main4CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Main4CallerSession struct {
	Contract *Main4Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Main4TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Main4TransactorSession struct {
	Contract     *Main4Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Main4Raw is an auto generated low-level Go binding around an Ethereum contract.
type Main4Raw struct {
	Contract *Main4 // Generic contract binding to access the raw methods on
}

// Main4CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Main4CallerRaw struct {
	Contract *Main4Caller // Generic read-only contract binding to access the raw methods on
}

// Main4TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Main4TransactorRaw struct {
	Contract *Main4Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMain4 creates a new instance of Main4, bound to a specific deployed contract.
func NewMain4(address common.Address, backend bind.ContractBackend) (*Main4, error) {
	contract, err := bindMain4(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main4{Main4Caller: Main4Caller{contract: contract}, Main4Transactor: Main4Transactor{contract: contract}, Main4Filterer: Main4Filterer{contract: contract}}, nil
}

// NewMain4Caller creates a new read-only instance of Main4, bound to a specific deployed contract.
func NewMain4Caller(address common.Address, caller bind.ContractCaller) (*Main4Caller, error) {
	contract, err := bindMain4(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Main4Caller{contract: contract}, nil
}

// NewMain4Transactor creates a new write-only instance of Main4, bound to a specific deployed contract.
func NewMain4Transactor(address common.Address, transactor bind.ContractTransactor) (*Main4Transactor, error) {
	contract, err := bindMain4(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Main4Transactor{contract: contract}, nil
}

// NewMain4Filterer creates a new log filterer instance of Main4, bound to a specific deployed contract.
func NewMain4Filterer(address common.Address, filterer bind.ContractFilterer) (*Main4Filterer, error) {
	contract, err := bindMain4(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Main4Filterer{contract: contract}, nil
}

// bindMain4 binds a generic wrapper to an already deployed contract.
func bindMain4(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Main4ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main4 *Main4Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main4.Contract.Main4Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main4 *Main4Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main4.Contract.Main4Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main4 *Main4Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main4.Contract.Main4Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main4 *Main4CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main4.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main4 *Main4TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main4.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main4 *Main4TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main4.Contract.contract.Transact(opts, method, params...)
}

// Payment is a free data retrieval call binding the contract method 0xa33b1260.
//
// Solidity: function Payment() constant returns(uint256)
func (_Main4 *Main4Caller) Payment(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Main4.contract.Call(opts, out, "Payment")
	return *ret0, err
}

// Payment is a free data retrieval call binding the contract method 0xa33b1260.
//
// Solidity: function Payment() constant returns(uint256)
func (_Main4 *Main4Session) Payment() (*big.Int, error) {
	return _Main4.Contract.Payment(&_Main4.CallOpts)
}

// Payment is a free data retrieval call binding the contract method 0xa33b1260.
//
// Solidity: function Payment() constant returns(uint256)
func (_Main4 *Main4CallerSession) Payment() (*big.Int, error) {
	return _Main4.Contract.Payment(&_Main4.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() constant returns(address)
func (_Main4 *Main4Caller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Main4.contract.Call(opts, out, "seller")
	return *ret0, err
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() constant returns(address)
func (_Main4 *Main4Session) Seller() (common.Address, error) {
	return _Main4.Contract.Seller(&_Main4.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() constant returns(address)
func (_Main4 *Main4CallerSession) Seller() (common.Address, error) {
	return _Main4.Contract.Seller(&_Main4.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Main4 *Main4Caller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Main4.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Main4 *Main4Session) State() (uint8, error) {
	return _Main4.Contract.State(&_Main4.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Main4 *Main4CallerSession) State() (uint8, error) {
	return _Main4.Contract.State(&_Main4.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(bool)
func (_Main4 *Main4Caller) Status(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Main4.contract.Call(opts, out, "status")
	return *ret0, err
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(bool)
func (_Main4 *Main4Session) Status() (bool, error) {
	return _Main4.Contract.Status(&_Main4.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(bool)
func (_Main4 *Main4CallerSession) Status() (bool, error) {
	return _Main4.Contract.Status(&_Main4.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Main4 *Main4Transactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main4.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Main4 *Main4Session) Abort() (*types.Transaction, error) {
	return _Main4.Contract.Abort(&_Main4.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_Main4 *Main4TransactorSession) Abort() (*types.Transaction, error) {
	return _Main4.Contract.Abort(&_Main4.TransactOpts)
}

// Askfordata is a paid mutator transaction binding the contract method 0xd4a381b2.
//
// Solidity: function askfordata(KeyAddress string, DataSchemaAddress string, computionAddress string, TransactionDetailAddress string) returns(uint256)
func (_Main4 *Main4Transactor) Askfordata(opts *bind.TransactOpts, KeyAddress string, DataSchemaAddress string, computionAddress string, TransactionDetailAddress string) (*types.Transaction, error) {
	return _Main4.contract.Transact(opts, "askfordata", KeyAddress, DataSchemaAddress, computionAddress, TransactionDetailAddress)
}

// Askfordata is a paid mutator transaction binding the contract method 0xd4a381b2.
//
// Solidity: function askfordata(KeyAddress string, DataSchemaAddress string, computionAddress string, TransactionDetailAddress string) returns(uint256)
func (_Main4 *Main4Session) Askfordata(KeyAddress string, DataSchemaAddress string, computionAddress string, TransactionDetailAddress string) (*types.Transaction, error) {
	return _Main4.Contract.Askfordata(&_Main4.TransactOpts, KeyAddress, DataSchemaAddress, computionAddress, TransactionDetailAddress)
}

// Askfordata is a paid mutator transaction binding the contract method 0xd4a381b2.
//
// Solidity: function askfordata(KeyAddress string, DataSchemaAddress string, computionAddress string, TransactionDetailAddress string) returns(uint256)
func (_Main4 *Main4TransactorSession) Askfordata(KeyAddress string, DataSchemaAddress string, computionAddress string, TransactionDetailAddress string) (*types.Transaction, error) {
	return _Main4.Contract.Askfordata(&_Main4.TransactOpts, KeyAddress, DataSchemaAddress, computionAddress, TransactionDetailAddress)
}

// Buydata is a paid mutator transaction binding the contract method 0x4257630e.
//
// Solidity: function buydata() returns(bool)
func (_Main4 *Main4Transactor) Buydata(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main4.contract.Transact(opts, "buydata")
}

// Buydata is a paid mutator transaction binding the contract method 0x4257630e.
//
// Solidity: function buydata() returns(bool)
func (_Main4 *Main4Session) Buydata() (*types.Transaction, error) {
	return _Main4.Contract.Buydata(&_Main4.TransactOpts)
}

// Buydata is a paid mutator transaction binding the contract method 0x4257630e.
//
// Solidity: function buydata() returns(bool)
func (_Main4 *Main4TransactorSession) Buydata() (*types.Transaction, error) {
	return _Main4.Contract.Buydata(&_Main4.TransactOpts)
}

// Isbaid is a paid mutator transaction binding the contract method 0x7c193d11.
//
// Solidity: function isbaid(buyer address) returns(bool)
func (_Main4 *Main4Transactor) Isbaid(opts *bind.TransactOpts, buyer common.Address) (*types.Transaction, error) {
	return _Main4.contract.Transact(opts, "isbaid", buyer)
}

// Isbaid is a paid mutator transaction binding the contract method 0x7c193d11.
//
// Solidity: function isbaid(buyer address) returns(bool)
func (_Main4 *Main4Session) Isbaid(buyer common.Address) (*types.Transaction, error) {
	return _Main4.Contract.Isbaid(&_Main4.TransactOpts, buyer)
}

// Isbaid is a paid mutator transaction binding the contract method 0x7c193d11.
//
// Solidity: function isbaid(buyer address) returns(bool)
func (_Main4 *Main4TransactorSession) Isbaid(buyer common.Address) (*types.Transaction, error) {
	return _Main4.Contract.Isbaid(&_Main4.TransactOpts, buyer)
}

// Main4AbortedIterator is returned from FilterAborted and is used to iterate over the raw logs and unpacked data for Aborted events raised by the Main4 contract.
type Main4AbortedIterator struct {
	Event *Main4Aborted // Event containing the contract specifics and raw log

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
func (it *Main4AbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Main4Aborted)
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
		it.Event = new(Main4Aborted)
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
func (it *Main4AbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Main4AbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Main4Aborted represents a Aborted event raised by the Main4 contract.
type Main4Aborted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAborted is a free log retrieval operation binding the contract event 0x72c874aeff0b183a56e2b79c71b46e1aed4dee5e09862134b8821ba2fddbf8bf.
//
// Solidity: e Aborted()
func (_Main4 *Main4Filterer) FilterAborted(opts *bind.FilterOpts) (*Main4AbortedIterator, error) {

	logs, sub, err := _Main4.contract.FilterLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return &Main4AbortedIterator{contract: _Main4.contract, event: "Aborted", logs: logs, sub: sub}, nil
}

// WatchAborted is a free log subscription operation binding the contract event 0x72c874aeff0b183a56e2b79c71b46e1aed4dee5e09862134b8821ba2fddbf8bf.
//
// Solidity: e Aborted()
func (_Main4 *Main4Filterer) WatchAborted(opts *bind.WatchOpts, sink chan<- *Main4Aborted) (event.Subscription, error) {

	logs, sub, err := _Main4.contract.WatchLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Main4Aborted)
				if err := _Main4.contract.UnpackLog(event, "Aborted", log); err != nil {
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

// Main4BidDataIterator is returned from FilterBidData and is used to iterate over the raw logs and unpacked data for BidData events raised by the Main4 contract.
type Main4BidDataIterator struct {
	Event *Main4BidData // Event containing the contract specifics and raw log

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
func (it *Main4BidDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Main4BidData)
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
		it.Event = new(Main4BidData)
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
func (it *Main4BidDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Main4BidDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Main4BidData represents a BidData event raised by the Main4 contract.
type Main4BidData struct {
	DataSchemaAddress string
	MataDataAddress   string
	Payment           *big.Int
	This              common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBidData is a free log retrieval operation binding the contract event 0xb1ce085c6a9a9a2c2a4af0e92549b3e7cafa61e5bd6ac98f57c384720aaa37cf.
//
// Solidity: e BidData(DataSchemaAddress string, MataDataAddress string, Payment uint256, _this address)
func (_Main4 *Main4Filterer) FilterBidData(opts *bind.FilterOpts) (*Main4BidDataIterator, error) {

	logs, sub, err := _Main4.contract.FilterLogs(opts, "BidData")
	if err != nil {
		return nil, err
	}
	return &Main4BidDataIterator{contract: _Main4.contract, event: "BidData", logs: logs, sub: sub}, nil
}

// WatchBidData is a free log subscription operation binding the contract event 0xb1ce085c6a9a9a2c2a4af0e92549b3e7cafa61e5bd6ac98f57c384720aaa37cf.
//
// Solidity: e BidData(DataSchemaAddress string, MataDataAddress string, Payment uint256, _this address)
func (_Main4 *Main4Filterer) WatchBidData(opts *bind.WatchOpts, sink chan<- *Main4BidData) (event.Subscription, error) {

	logs, sub, err := _Main4.contract.WatchLogs(opts, "BidData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Main4BidData)
				if err := _Main4.contract.UnpackLog(event, "BidData", log); err != nil {
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

// Main4FallbackCalledIterator is returned from FilterFallbackCalled and is used to iterate over the raw logs and unpacked data for FallbackCalled events raised by the Main4 contract.
type Main4FallbackCalledIterator struct {
	Event *Main4FallbackCalled // Event containing the contract specifics and raw log

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
func (it *Main4FallbackCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Main4FallbackCalled)
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
		it.Event = new(Main4FallbackCalled)
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
func (it *Main4FallbackCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Main4FallbackCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Main4FallbackCalled represents a FallbackCalled event raised by the Main4 contract.
type Main4FallbackCalled struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFallbackCalled is a free log retrieval operation binding the contract event 0x4fdfaf7cd2ae67f3639149bddef2addc991d3f26e0121df03517ec62523649ca.
//
// Solidity: e FallbackCalled(from address, amount uint256)
func (_Main4 *Main4Filterer) FilterFallbackCalled(opts *bind.FilterOpts) (*Main4FallbackCalledIterator, error) {

	logs, sub, err := _Main4.contract.FilterLogs(opts, "FallbackCalled")
	if err != nil {
		return nil, err
	}
	return &Main4FallbackCalledIterator{contract: _Main4.contract, event: "FallbackCalled", logs: logs, sub: sub}, nil
}

// WatchFallbackCalled is a free log subscription operation binding the contract event 0x4fdfaf7cd2ae67f3639149bddef2addc991d3f26e0121df03517ec62523649ca.
//
// Solidity: e FallbackCalled(from address, amount uint256)
func (_Main4 *Main4Filterer) WatchFallbackCalled(opts *bind.WatchOpts, sink chan<- *Main4FallbackCalled) (event.Subscription, error) {

	logs, sub, err := _Main4.contract.WatchLogs(opts, "FallbackCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Main4FallbackCalled)
				if err := _Main4.contract.UnpackLog(event, "FallbackCalled", log); err != nil {
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
