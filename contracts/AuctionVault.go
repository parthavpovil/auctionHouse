// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// AuctionVaultMetaData contains all meta data concerning the AuctionVault contract.
var AuctionVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AuctionVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionVaultMetaData.ABI instead.
var AuctionVaultABI = AuctionVaultMetaData.ABI

// AuctionVault is an auto generated Go binding around an Ethereum contract.
type AuctionVault struct {
	AuctionVaultCaller     // Read-only binding to the contract
	AuctionVaultTransactor // Write-only binding to the contract
	AuctionVaultFilterer   // Log filterer for contract events
}

// AuctionVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionVaultSession struct {
	Contract     *AuctionVault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionVaultCallerSession struct {
	Contract *AuctionVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AuctionVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionVaultTransactorSession struct {
	Contract     *AuctionVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AuctionVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionVaultRaw struct {
	Contract *AuctionVault // Generic contract binding to access the raw methods on
}

// AuctionVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionVaultCallerRaw struct {
	Contract *AuctionVaultCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionVaultTransactorRaw struct {
	Contract *AuctionVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionVault creates a new instance of AuctionVault, bound to a specific deployed contract.
func NewAuctionVault(address common.Address, backend bind.ContractBackend) (*AuctionVault, error) {
	contract, err := bindAuctionVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionVault{AuctionVaultCaller: AuctionVaultCaller{contract: contract}, AuctionVaultTransactor: AuctionVaultTransactor{contract: contract}, AuctionVaultFilterer: AuctionVaultFilterer{contract: contract}}, nil
}

// NewAuctionVaultCaller creates a new read-only instance of AuctionVault, bound to a specific deployed contract.
func NewAuctionVaultCaller(address common.Address, caller bind.ContractCaller) (*AuctionVaultCaller, error) {
	contract, err := bindAuctionVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionVaultCaller{contract: contract}, nil
}

// NewAuctionVaultTransactor creates a new write-only instance of AuctionVault, bound to a specific deployed contract.
func NewAuctionVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionVaultTransactor, error) {
	contract, err := bindAuctionVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionVaultTransactor{contract: contract}, nil
}

// NewAuctionVaultFilterer creates a new log filterer instance of AuctionVault, bound to a specific deployed contract.
func NewAuctionVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionVaultFilterer, error) {
	contract, err := bindAuctionVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionVaultFilterer{contract: contract}, nil
}

// bindAuctionVault binds a generic wrapper to an already deployed contract.
func bindAuctionVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionVault *AuctionVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionVault.Contract.AuctionVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionVault *AuctionVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionVault.Contract.AuctionVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionVault *AuctionVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionVault.Contract.AuctionVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionVault *AuctionVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionVault *AuctionVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionVault *AuctionVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionVault.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_AuctionVault *AuctionVaultCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuctionVault.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_AuctionVault *AuctionVaultSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _AuctionVault.Contract.Balances(&_AuctionVault.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_AuctionVault *AuctionVaultCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _AuctionVault.Contract.Balances(&_AuctionVault.CallOpts, arg0)
}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_AuctionVault *AuctionVaultCaller) ContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionVault.contract.Call(opts, &out, "contractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_AuctionVault *AuctionVaultSession) ContractBalance() (*big.Int, error) {
	return _AuctionVault.Contract.ContractBalance(&_AuctionVault.CallOpts)
}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_AuctionVault *AuctionVaultCallerSession) ContractBalance() (*big.Int, error) {
	return _AuctionVault.Contract.ContractBalance(&_AuctionVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionVault *AuctionVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionVault *AuctionVaultSession) Owner() (common.Address, error) {
	return _AuctionVault.Contract.Owner(&_AuctionVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionVault *AuctionVaultCallerSession) Owner() (common.Address, error) {
	return _AuctionVault.Contract.Owner(&_AuctionVault.CallOpts)
}

// ApproveWithdrawal is a paid mutator transaction binding the contract method 0xf4970e71.
//
// Solidity: function approveWithdrawal(address user, uint256 amount) returns()
func (_AuctionVault *AuctionVaultTransactor) ApproveWithdrawal(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AuctionVault.contract.Transact(opts, "approveWithdrawal", user, amount)
}

// ApproveWithdrawal is a paid mutator transaction binding the contract method 0xf4970e71.
//
// Solidity: function approveWithdrawal(address user, uint256 amount) returns()
func (_AuctionVault *AuctionVaultSession) ApproveWithdrawal(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AuctionVault.Contract.ApproveWithdrawal(&_AuctionVault.TransactOpts, user, amount)
}

// ApproveWithdrawal is a paid mutator transaction binding the contract method 0xf4970e71.
//
// Solidity: function approveWithdrawal(address user, uint256 amount) returns()
func (_AuctionVault *AuctionVaultTransactorSession) ApproveWithdrawal(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AuctionVault.Contract.ApproveWithdrawal(&_AuctionVault.TransactOpts, user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_AuctionVault *AuctionVaultTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionVault.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_AuctionVault *AuctionVaultSession) Deposit() (*types.Transaction, error) {
	return _AuctionVault.Contract.Deposit(&_AuctionVault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_AuctionVault *AuctionVaultTransactorSession) Deposit() (*types.Transaction, error) {
	return _AuctionVault.Contract.Deposit(&_AuctionVault.TransactOpts)
}

// AuctionVaultDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the AuctionVault contract.
type AuctionVaultDepositedIterator struct {
	Event *AuctionVaultDeposited // Event containing the contract specifics and raw log

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
func (it *AuctionVaultDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionVaultDeposited)
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
		it.Event = new(AuctionVaultDeposited)
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
func (it *AuctionVaultDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionVaultDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionVaultDeposited represents a Deposited event raised by the AuctionVault contract.
type AuctionVaultDeposited struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed user, uint256 amount)
func (_AuctionVault *AuctionVaultFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address) (*AuctionVaultDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AuctionVault.contract.FilterLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return &AuctionVaultDepositedIterator{contract: _AuctionVault.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed user, uint256 amount)
func (_AuctionVault *AuctionVaultFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *AuctionVaultDeposited, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AuctionVault.contract.WatchLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionVaultDeposited)
				if err := _AuctionVault.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed user, uint256 amount)
func (_AuctionVault *AuctionVaultFilterer) ParseDeposited(log types.Log) (*AuctionVaultDeposited, error) {
	event := new(AuctionVaultDeposited)
	if err := _AuctionVault.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionVaultWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the AuctionVault contract.
type AuctionVaultWithdrawnIterator struct {
	Event *AuctionVaultWithdrawn // Event containing the contract specifics and raw log

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
func (it *AuctionVaultWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionVaultWithdrawn)
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
		it.Event = new(AuctionVaultWithdrawn)
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
func (it *AuctionVaultWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionVaultWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionVaultWithdrawn represents a Withdrawn event raised by the AuctionVault contract.
type AuctionVaultWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_AuctionVault *AuctionVaultFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*AuctionVaultWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AuctionVault.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &AuctionVaultWithdrawnIterator{contract: _AuctionVault.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_AuctionVault *AuctionVaultFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *AuctionVaultWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AuctionVault.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionVaultWithdrawn)
				if err := _AuctionVault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_AuctionVault *AuctionVaultFilterer) ParseWithdrawn(log types.Log) (*AuctionVaultWithdrawn, error) {
	event := new(AuctionVaultWithdrawn)
	if err := _AuctionVault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
