// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package eth

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EthSSHKeyABI is the input ABI used to generate the binding from.
const EthSSHKeyABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sshPublicKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sshkey\",\"type\":\"string\"}],\"name\":\"addSshKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"mortal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// EthSSHKey is an auto generated Go binding around an Ethereum contract.
type EthSSHKey struct {
	EthSSHKeyCaller     // Read-only binding to the contract
	EthSSHKeyTransactor // Write-only binding to the contract
}

// EthSSHKeyCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthSSHKeyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSSHKeyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthSSHKeyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSSHKeySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthSSHKeySession struct {
	Contract     *EthSSHKey        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthSSHKeyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthSSHKeyCallerSession struct {
	Contract *EthSSHKeyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EthSSHKeyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthSSHKeyTransactorSession struct {
	Contract     *EthSSHKeyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthSSHKeyRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthSSHKeyRaw struct {
	Contract *EthSSHKey // Generic contract binding to access the raw methods on
}

// EthSSHKeyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthSSHKeyCallerRaw struct {
	Contract *EthSSHKeyCaller // Generic read-only contract binding to access the raw methods on
}

// EthSSHKeyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthSSHKeyTransactorRaw struct {
	Contract *EthSSHKeyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthSSHKey creates a new instance of EthSSHKey, bound to a specific deployed contract.
func NewEthSSHKey(address common.Address, backend bind.ContractBackend) (*EthSSHKey, error) {
	contract, err := bindEthSSHKey(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthSSHKey{EthSSHKeyCaller: EthSSHKeyCaller{contract: contract}, EthSSHKeyTransactor: EthSSHKeyTransactor{contract: contract}}, nil
}

// NewEthSSHKeyCaller creates a new read-only instance of EthSSHKey, bound to a specific deployed contract.
func NewEthSSHKeyCaller(address common.Address, caller bind.ContractCaller) (*EthSSHKeyCaller, error) {
	contract, err := bindEthSSHKey(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EthSSHKeyCaller{contract: contract}, nil
}

// NewEthSSHKeyTransactor creates a new write-only instance of EthSSHKey, bound to a specific deployed contract.
func NewEthSSHKeyTransactor(address common.Address, transactor bind.ContractTransactor) (*EthSSHKeyTransactor, error) {
	contract, err := bindEthSSHKey(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EthSSHKeyTransactor{contract: contract}, nil
}

// bindEthSSHKey binds a generic wrapper to an already deployed contract.
func bindEthSSHKey(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthSSHKeyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthSSHKey *EthSSHKeyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthSSHKey.Contract.EthSSHKeyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthSSHKey *EthSSHKeyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSSHKey.Contract.EthSSHKeyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthSSHKey *EthSSHKeyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthSSHKey.Contract.EthSSHKeyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthSSHKey *EthSSHKeyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthSSHKey.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthSSHKey *EthSSHKeyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSSHKey.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthSSHKey *EthSSHKeyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthSSHKey.Contract.contract.Transact(opts, method, params...)
}

// SshPublicKeys is a free data retrieval call binding the contract method 0x53918ad1.
//
// Solidity: function sshPublicKeys( uint256) constant returns(string)
func (_EthSSHKey *EthSSHKeyCaller) SshPublicKeys(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EthSSHKey.contract.Call(opts, out, "sshPublicKeys", arg0)
	return *ret0, err
}

// SshPublicKeys is a free data retrieval call binding the contract method 0x53918ad1.
//
// Solidity: function sshPublicKeys( uint256) constant returns(string)
func (_EthSSHKey *EthSSHKeySession) SshPublicKeys(arg0 *big.Int) (string, error) {
	return _EthSSHKey.Contract.SshPublicKeys(&_EthSSHKey.CallOpts, arg0)
}

// SshPublicKeys is a free data retrieval call binding the contract method 0x53918ad1.
//
// Solidity: function sshPublicKeys( uint256) constant returns(string)
func (_EthSSHKey *EthSSHKeyCallerSession) SshPublicKeys(arg0 *big.Int) (string, error) {
	return _EthSSHKey.Contract.SshPublicKeys(&_EthSSHKey.CallOpts, arg0)
}

// AddSshKey is a paid mutator transaction binding the contract method 0x690252c8.
//
// Solidity: function addSshKey(sshkey string) returns(bool)
func (_EthSSHKey *EthSSHKeyTransactor) AddSshKey(opts *bind.TransactOpts, sshkey string) (*types.Transaction, error) {
	return _EthSSHKey.contract.Transact(opts, "addSshKey", sshkey)
}

// AddSshKey is a paid mutator transaction binding the contract method 0x690252c8.
//
// Solidity: function addSshKey(sshkey string) returns(bool)
func (_EthSSHKey *EthSSHKeySession) AddSshKey(sshkey string) (*types.Transaction, error) {
	return _EthSSHKey.Contract.AddSshKey(&_EthSSHKey.TransactOpts, sshkey)
}

// AddSshKey is a paid mutator transaction binding the contract method 0x690252c8.
//
// Solidity: function addSshKey(sshkey string) returns(bool)
func (_EthSSHKey *EthSSHKeyTransactorSession) AddSshKey(sshkey string) (*types.Transaction, error) {
	return _EthSSHKey.Contract.AddSshKey(&_EthSSHKey.TransactOpts, sshkey)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_EthSSHKey *EthSSHKeyTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSSHKey.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_EthSSHKey *EthSSHKeySession) Kill() (*types.Transaction, error) {
	return _EthSSHKey.Contract.Kill(&_EthSSHKey.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_EthSSHKey *EthSSHKeyTransactorSession) Kill() (*types.Transaction, error) {
	return _EthSSHKey.Contract.Kill(&_EthSSHKey.TransactOpts)
}

// Mortal is a paid mutator transaction binding the contract method 0xf1eae25c.
//
// Solidity: function mortal() returns()
func (_EthSSHKey *EthSSHKeyTransactor) Mortal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSSHKey.contract.Transact(opts, "mortal")
}

// Mortal is a paid mutator transaction binding the contract method 0xf1eae25c.
//
// Solidity: function mortal() returns()
func (_EthSSHKey *EthSSHKeySession) Mortal() (*types.Transaction, error) {
	return _EthSSHKey.Contract.Mortal(&_EthSSHKey.TransactOpts)
}

// Mortal is a paid mutator transaction binding the contract method 0xf1eae25c.
//
// Solidity: function mortal() returns()
func (_EthSSHKey *EthSSHKeyTransactorSession) Mortal() (*types.Transaction, error) {
	return _EthSSHKey.Contract.Mortal(&_EthSSHKey.TransactOpts)
}
