// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permission

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

// NodeManagerABI is the input ABI used to generate the binding from.
const NodeManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_action\",\"type\":\"uint256\"}],\"name\":\"updateNodeStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"enodeId\",\"type\":\"string\"}],\"name\":\"getNodeDetails\",\"outputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_nodeStatus\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addOrgNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"approveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeIndex\",\"type\":\"uint256\"}],\"name\":\"getNodeDetailsFromIndex\",\"outputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_nodeStatus\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addAdminNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NodeBlacklisted\",\"type\":\"event\"}]"

// NodeManager is an auto generated Go binding around an Ethereum contract.
type NodeManager struct {
	NodeManagerCaller     // Read-only binding to the contract
	NodeManagerTransactor // Write-only binding to the contract
	NodeManagerFilterer   // Log filterer for contract events
}

// NodeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeManagerSession struct {
	Contract     *NodeManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeManagerCallerSession struct {
	Contract *NodeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeManagerTransactorSession struct {
	Contract     *NodeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeManagerRaw struct {
	Contract *NodeManager // Generic contract binding to access the raw methods on
}

// NodeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeManagerCallerRaw struct {
	Contract *NodeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NodeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeManagerTransactorRaw struct {
	Contract *NodeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeManager creates a new instance of NodeManager, bound to a specific deployed contract.
func NewNodeManager(address common.Address, backend bind.ContractBackend) (*NodeManager, error) {
	contract, err := bindNodeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeManager{NodeManagerCaller: NodeManagerCaller{contract: contract}, NodeManagerTransactor: NodeManagerTransactor{contract: contract}, NodeManagerFilterer: NodeManagerFilterer{contract: contract}}, nil
}

// NewNodeManagerCaller creates a new read-only instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerCaller(address common.Address, caller bind.ContractCaller) (*NodeManagerCaller, error) {
	contract, err := bindNodeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerCaller{contract: contract}, nil
}

// NewNodeManagerTransactor creates a new write-only instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeManagerTransactor, error) {
	contract, err := bindNodeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerTransactor{contract: contract}, nil
}

// NewNodeManagerFilterer creates a new log filterer instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeManagerFilterer, error) {
	contract, err := bindNodeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeManagerFilterer{contract: contract}, nil
}

// bindNodeManager binds a generic wrapper to an already deployed contract.
func bindNodeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManager *NodeManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeManager.Contract.NodeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManager *NodeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManager.Contract.NodeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManager *NodeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManager.Contract.NodeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManager *NodeManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManager *NodeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManager *NodeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManager.Contract.contract.Transact(opts, method, params...)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(enodeId string) constant returns(_orgId string, _enodeId string, _nodeStatus uint256)
func (_NodeManager *NodeManagerCaller) GetNodeDetails(opts *bind.CallOpts, enodeId string) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	ret := new(struct {
		OrgId      string
		EnodeId    string
		NodeStatus *big.Int
	})
	out := ret
	err := _NodeManager.contract.Call(opts, out, "getNodeDetails", enodeId)
	return *ret, err
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(enodeId string) constant returns(_orgId string, _enodeId string, _nodeStatus uint256)
func (_NodeManager *NodeManagerSession) GetNodeDetails(enodeId string) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	return _NodeManager.Contract.GetNodeDetails(&_NodeManager.CallOpts, enodeId)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(enodeId string) constant returns(_orgId string, _enodeId string, _nodeStatus uint256)
func (_NodeManager *NodeManagerCallerSession) GetNodeDetails(enodeId string) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	return _NodeManager.Contract.GetNodeDetails(&_NodeManager.CallOpts, enodeId)
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(_nodeIndex uint256) constant returns(_orgId string, _enodeId string, _nodeStatus uint256)
func (_NodeManager *NodeManagerCaller) GetNodeDetailsFromIndex(opts *bind.CallOpts, _nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	ret := new(struct {
		OrgId      string
		EnodeId    string
		NodeStatus *big.Int
	})
	out := ret
	err := _NodeManager.contract.Call(opts, out, "getNodeDetailsFromIndex", _nodeIndex)
	return *ret, err
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(_nodeIndex uint256) constant returns(_orgId string, _enodeId string, _nodeStatus uint256)
func (_NodeManager *NodeManagerSession) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	return _NodeManager.Contract.GetNodeDetailsFromIndex(&_NodeManager.CallOpts, _nodeIndex)
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(_nodeIndex uint256) constant returns(_orgId string, _enodeId string, _nodeStatus uint256)
func (_NodeManager *NodeManagerCallerSession) GetNodeDetailsFromIndex(_nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	return _NodeManager.Contract.GetNodeDetailsFromIndex(&_NodeManager.CallOpts, _nodeIndex)
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_NodeManager *NodeManagerCaller) GetNumberOfNodes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeManager.contract.Call(opts, out, "getNumberOfNodes")
	return *ret0, err
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_NodeManager *NodeManagerSession) GetNumberOfNodes() (*big.Int, error) {
	return _NodeManager.Contract.GetNumberOfNodes(&_NodeManager.CallOpts)
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_NodeManager *NodeManagerCallerSession) GetNumberOfNodes() (*big.Int, error) {
	return _NodeManager.Contract.GetNumberOfNodes(&_NodeManager.CallOpts)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0xe3b09d84.
//
// Solidity: function addAdminNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerTransactor) AddAdminNode(opts *bind.TransactOpts, _enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "addAdminNode", _enodeId, _orgId)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0xe3b09d84.
//
// Solidity: function addAdminNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerSession) AddAdminNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddAdminNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddAdminNode is a paid mutator transaction binding the contract method 0xe3b09d84.
//
// Solidity: function addAdminNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerTransactorSession) AddAdminNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddAdminNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerTransactor) AddNode(opts *bind.TransactOpts, _enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "addNode", _enodeId, _orgId)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerSession) AddNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerTransactorSession) AddNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddOrgNode is a paid mutator transaction binding the contract method 0x3f5e1a45.
//
// Solidity: function addOrgNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerTransactor) AddOrgNode(opts *bind.TransactOpts, _enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "addOrgNode", _enodeId, _orgId)
}

// AddOrgNode is a paid mutator transaction binding the contract method 0x3f5e1a45.
//
// Solidity: function addOrgNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerSession) AddOrgNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddOrgNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddOrgNode is a paid mutator transaction binding the contract method 0x3f5e1a45.
//
// Solidity: function addOrgNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerTransactorSession) AddOrgNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddOrgNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0x86bc3652.
//
// Solidity: function approveNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerTransactor) ApproveNode(opts *bind.TransactOpts, _enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "approveNode", _enodeId, _orgId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0x86bc3652.
//
// Solidity: function approveNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerSession) ApproveNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.ApproveNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0x86bc3652.
//
// Solidity: function approveNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerTransactorSession) ApproveNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.ApproveNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x0cc50146.
//
// Solidity: function updateNodeStatus(_enodeId string, _orgId string, _action uint256) returns()
func (_NodeManager *NodeManagerTransactor) UpdateNodeStatus(opts *bind.TransactOpts, _enodeId string, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "updateNodeStatus", _enodeId, _orgId, _action)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x0cc50146.
//
// Solidity: function updateNodeStatus(_enodeId string, _orgId string, _action uint256) returns()
func (_NodeManager *NodeManagerSession) UpdateNodeStatus(_enodeId string, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.UpdateNodeStatus(&_NodeManager.TransactOpts, _enodeId, _orgId, _action)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x0cc50146.
//
// Solidity: function updateNodeStatus(_enodeId string, _orgId string, _action uint256) returns()
func (_NodeManager *NodeManagerTransactorSession) UpdateNodeStatus(_enodeId string, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.UpdateNodeStatus(&_NodeManager.TransactOpts, _enodeId, _orgId, _action)
}

// NodeManagerNodeActivatedIterator is returned from FilterNodeActivated and is used to iterate over the raw logs and unpacked data for NodeActivated events raised by the NodeManager contract.
type NodeManagerNodeActivatedIterator struct {
	Event *NodeManagerNodeActivated // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeActivated)
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
		it.Event = new(NodeManagerNodeActivated)
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
func (it *NodeManagerNodeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeActivated represents a NodeActivated event raised by the NodeManager contract.
type NodeManagerNodeActivated struct {
	EnodeId string
	OrgId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeActivated is a free log retrieval operation binding the contract event 0x49796be3ca168a59c8ae46c75a36a9bb3a84753d3e12a812f93ae010e783b14f.
//
// Solidity: e NodeActivated(_enodeId string, _orgId string)
func (_NodeManager *NodeManagerFilterer) FilterNodeActivated(opts *bind.FilterOpts) (*NodeManagerNodeActivatedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeActivated")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeActivatedIterator{contract: _NodeManager.contract, event: "NodeActivated", logs: logs, sub: sub}, nil
}

// WatchNodeActivated is a free log subscription operation binding the contract event 0x49796be3ca168a59c8ae46c75a36a9bb3a84753d3e12a812f93ae010e783b14f.
//
// Solidity: e NodeActivated(_enodeId string, _orgId string)
func (_NodeManager *NodeManagerFilterer) WatchNodeActivated(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeActivated) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeActivated)
				if err := _NodeManager.contract.UnpackLog(event, "NodeActivated", log); err != nil {
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

// NodeManagerNodeApprovedIterator is returned from FilterNodeApproved and is used to iterate over the raw logs and unpacked data for NodeApproved events raised by the NodeManager contract.
type NodeManagerNodeApprovedIterator struct {
	Event *NodeManagerNodeApproved // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeApproved)
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
		it.Event = new(NodeManagerNodeApproved)
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
func (it *NodeManagerNodeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeApproved represents a NodeApproved event raised by the NodeManager contract.
type NodeManagerNodeApproved struct {
	EnodeId string
	OrgId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeApproved is a free log retrieval operation binding the contract event 0x0413ce00d5de406d9939003416263a7530eaeb13f9d281c8baeba1601def960d.
//
// Solidity: e NodeApproved(_enodeId string, _orgId string)
func (_NodeManager *NodeManagerFilterer) FilterNodeApproved(opts *bind.FilterOpts) (*NodeManagerNodeApprovedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeApprovedIterator{contract: _NodeManager.contract, event: "NodeApproved", logs: logs, sub: sub}, nil
}

// WatchNodeApproved is a free log subscription operation binding the contract event 0x0413ce00d5de406d9939003416263a7530eaeb13f9d281c8baeba1601def960d.
//
// Solidity: e NodeApproved(_enodeId string, _orgId string)
func (_NodeManager *NodeManagerFilterer) WatchNodeApproved(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeApproved) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeApproved)
				if err := _NodeManager.contract.UnpackLog(event, "NodeApproved", log); err != nil {
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

// NodeManagerNodeBlacklistedIterator is returned from FilterNodeBlacklisted and is used to iterate over the raw logs and unpacked data for NodeBlacklisted events raised by the NodeManager contract.
type NodeManagerNodeBlacklistedIterator struct {
	Event *NodeManagerNodeBlacklisted // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeBlacklisted)
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
		it.Event = new(NodeManagerNodeBlacklisted)
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
func (it *NodeManagerNodeBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeBlacklisted represents a NodeBlacklisted event raised by the NodeManager contract.
type NodeManagerNodeBlacklisted struct {
	EnodeId string
	OrgId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeBlacklisted is a free log retrieval operation binding the contract event 0x4714623279994517c446c8fb72c3fdaca26434da1e2490d3976fe0cd880cfa7a.
//
// Solidity: e NodeBlacklisted(_enodeId string, _orgId string)
func (_NodeManager *NodeManagerFilterer) FilterNodeBlacklisted(opts *bind.FilterOpts) (*NodeManagerNodeBlacklistedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeBlacklisted")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeBlacklistedIterator{contract: _NodeManager.contract, event: "NodeBlacklisted", logs: logs, sub: sub}, nil
}

// WatchNodeBlacklisted is a free log subscription operation binding the contract event 0x4714623279994517c446c8fb72c3fdaca26434da1e2490d3976fe0cd880cfa7a.
//
// Solidity: e NodeBlacklisted(_enodeId string, _orgId string)
func (_NodeManager *NodeManagerFilterer) WatchNodeBlacklisted(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeBlacklisted) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeBlacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeBlacklisted)
				if err := _NodeManager.contract.UnpackLog(event, "NodeBlacklisted", log); err != nil {
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

// NodeManagerNodeDeactivatedIterator is returned from FilterNodeDeactivated and is used to iterate over the raw logs and unpacked data for NodeDeactivated events raised by the NodeManager contract.
type NodeManagerNodeDeactivatedIterator struct {
	Event *NodeManagerNodeDeactivated // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeDeactivated)
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
		it.Event = new(NodeManagerNodeDeactivated)
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
func (it *NodeManagerNodeDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeDeactivated represents a NodeDeactivated event raised by the NodeManager contract.
type NodeManagerNodeDeactivated struct {
	EnodeId string
	OrgId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeDeactivated is a free log retrieval operation binding the contract event 0xc6c3720fe673e87bb26e06be713d514278aa94c3939cfe7c64b9bea4d486824a.
//
// Solidity: e NodeDeactivated(_enodeId string, _orgId string)
func (_NodeManager *NodeManagerFilterer) FilterNodeDeactivated(opts *bind.FilterOpts) (*NodeManagerNodeDeactivatedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeDeactivated")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeDeactivatedIterator{contract: _NodeManager.contract, event: "NodeDeactivated", logs: logs, sub: sub}, nil
}

// WatchNodeDeactivated is a free log subscription operation binding the contract event 0xc6c3720fe673e87bb26e06be713d514278aa94c3939cfe7c64b9bea4d486824a.
//
// Solidity: e NodeDeactivated(_enodeId string, _orgId string)
func (_NodeManager *NodeManagerFilterer) WatchNodeDeactivated(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeDeactivated) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeDeactivated)
				if err := _NodeManager.contract.UnpackLog(event, "NodeDeactivated", log); err != nil {
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

// NodeManagerNodeProposedIterator is returned from FilterNodeProposed and is used to iterate over the raw logs and unpacked data for NodeProposed events raised by the NodeManager contract.
type NodeManagerNodeProposedIterator struct {
	Event *NodeManagerNodeProposed // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeProposed)
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
		it.Event = new(NodeManagerNodeProposed)
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
func (it *NodeManagerNodeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeProposed represents a NodeProposed event raised by the NodeManager contract.
type NodeManagerNodeProposed struct {
	EnodeId string
	OrgId   string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeProposed is a free log retrieval operation binding the contract event 0xb1a7eec7dd1a516c3132d6d1f770758b19aa34c3a07c138caf662688b7e3556f.
//
// Solidity: e NodeProposed(_enodeId string, _orgId string)
func (_NodeManager *NodeManagerFilterer) FilterNodeProposed(opts *bind.FilterOpts) (*NodeManagerNodeProposedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeProposed")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeProposedIterator{contract: _NodeManager.contract, event: "NodeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeProposed is a free log subscription operation binding the contract event 0xb1a7eec7dd1a516c3132d6d1f770758b19aa34c3a07c138caf662688b7e3556f.
//
// Solidity: e NodeProposed(_enodeId string, _orgId string)
func (_NodeManager *NodeManagerFilterer) WatchNodeProposed(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeProposed) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeProposed)
				if err := _NodeManager.contract.UnpackLog(event, "NodeProposed", log); err != nil {
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