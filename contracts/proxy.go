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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ProxyABI is the input ABI used to generate the binding from.
const ProxyABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_path\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_abi\",\"type\":\"string\"}],\"name\":\"addPath\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_argsWithMethodId\",\"type\":\"bytes\"}],\"name\":\"constantCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAbi\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPaths\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_argsWithMethodId\",\"type\":\"bytes\"}],\"name\":\"sendTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"stringToUint256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ProxyFuncSigs maps the 4-byte function signature to its string representation.
var ProxyFuncSigs = map[string]string{
	"16b1ec2a": "addPath(string,address,string)",
	"0e36518a": "constantCall(string,bytes)",
	"0fc97bd2": "getAbi(address)",
	"1767ae33": "getPaths()",
	"aa7fb5a0": "sendTransaction(string,string,bytes)",
	"ac5d3723": "stringToUint256(string)",
}

// ProxyBin is the compiled bytecode used for deploying new contracts.
var ProxyBin = "0x608060405234801561001057600080fd5b50336000818152602081905260409020805460ff1916600190811790915580546001600160a01b0319169091179055610c938061004e6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630e36518a146100675780630fc97bd21461009057806316b1ec2a146100a35780631767ae33146100b8578063aa7fb5a0146100cd578063ac5d3723146100e0575b600080fd5b61007a610075366004610932565b610100565b6040516100879190610b4c565b60405180910390f35b61007a61009e366004610863565b610123565b6100b66100b13660046108b6565b6101ce565b005b6100c06102cc565b6040516100879190610b3b565b61007a6100db36600461099b565b6103a5565b6100f36100ee366004610881565b610539565b6040516100879190610b6d565b6060600061010d846105c3565b905061011981846105f5565b9150505b92915050565b6001600160a01b03811660009081526003602090815260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156101c25780601f10610197576101008083540402835291602001916101c2565b820191906000526020600020905b8154815290600101906020018083116101a557829003601f168201915b50505050509050919050565b3360009081526020819052604090205460ff166102065760405162461bcd60e51b81526004016101fd90610b5d565b60405180910390fd5b60606102118461067e565b9050826002826040516102249190610b2f565b9081526040516020918190038201902080546001600160a01b0319166001600160a01b03939093169290921790915560048054600181018083556000929092528651919261029a927f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201919088019061076c565b50506001600160a01b038316600090815260036020908152604090912083516102c59285019061076c565b5050505050565b60606004805480602002602001604051908101604052809291908181526020016000905b8282101561039b5760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156103875780601f1061035c57610100808354040283529160200191610387565b820191906000526020600020905b81548152906001019060200180831161036a57829003601f168201915b5050505050815260200190600101906102f0565b5050505090505b90565b3360009081526020819052604090205460609060ff166103d75760405162461bcd60e51b81526004016101fd90610b5d565b6005846040516103e79190610b2f565b9081526040519081900360200190205460ff16156104af5760058460405161040f9190610b2f565b90815260408051602092819003830181206001908101805460029281161561010002600019011691909104601f810185900485028301850190935282825290929091908301828280156104a35780601f10610478576101008083540402835291602001916104a3565b820191906000526020600020905b81548152906001019060200180831161048657829003601f168201915b50505050509050610532565b60006104ba846105c3565b905060606104c882856105f5565b90506040518060400160405280600115158152602001828152506005876040516104f29190610b2f565b908152604051602091819003820190208251815460ff1916901515178155828201518051919261052a9260018501929091019061076c565b509193505050505b9392505050565b805160009082908290815b818110156105b957603084828151811061055a57fe5b016020015160f81c108015906105845750603984828151811061057957fe5b016020015160f81c11155b156105b157603084828151811061059757fe5b602001015160f81c60f81b60f81c0360ff1683600a020192505b600101610544565b5090949350505050565b60006002826040516105d59190610b2f565b908152604051908190036020019020546001600160a01b03169050919050565b60606000836001600160a01b0316836040516106119190610b2f565b6000604051808303816000865af19150503d806000811461064e576040519150601f19603f3d011682016040523d82523d6000602084013e610653565b606091505b509250905080610677578160405162461bcd60e51b81526004016101fd9190610b4c565b5092915050565b8051606090829060008060001983015b80156106de57601760f91b6001600160f81b0319168582815181106106af57fe5b01602001516001600160f81b03191614156106cf578060010191506106de565b6001909201916000190161068e565b506060826040519080825280601f01601f19166020018201604052801561070c576020820181803883390190505b50905060005b8381101561076157855160018401938791811061072b57fe5b602001015160f81c60f81b82828151811061074257fe5b60200101906001600160f81b031916908160001a905350600101610712565b509695505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106107ad57805160ff19168380011785556107da565b828001600101855582156107da579182015b828111156107da5782518255916020019190600101906107bf565b506107e69291506107ea565b5090565b6103a291905b808211156107e657600081556001016107f0565b803561011d81610c39565b600082601f83011261082057600080fd5b813561083361082e82610ba2565b610b7b565b9150808252602083016020830185838301111561084f57600080fd5b61085a838284610bf3565b50505092915050565b60006020828403121561087557600080fd5b60006101198484610804565b60006020828403121561089357600080fd5b813567ffffffffffffffff8111156108aa57600080fd5b6101198482850161080f565b6000806000606084860312156108cb57600080fd5b833567ffffffffffffffff8111156108e257600080fd5b6108ee8682870161080f565b93505060206108ff86828701610804565b925050604084013567ffffffffffffffff81111561091c57600080fd5b6109288682870161080f565b9150509250925092565b6000806040838503121561094557600080fd5b823567ffffffffffffffff81111561095c57600080fd5b6109688582860161080f565b925050602083013567ffffffffffffffff81111561098557600080fd5b6109918582860161080f565b9150509250929050565b6000806000606084860312156109b057600080fd5b833567ffffffffffffffff8111156109c757600080fd5b6109d38682870161080f565b935050602084013567ffffffffffffffff8111156109f057600080fd5b6108ff8682870161080f565b60006105328383610a76565b6000610a1382610bd0565b610a1d8185610bd4565b935083602082028501610a2f85610bca565b8060005b85811015610a695784840389528151610a4c85826109fc565b9450610a5783610bca565b60209a909a0199925050600101610a33565b5091979650505050505050565b6000610a8182610bd0565b610a8b8185610bd4565b9350610a9b818560208601610bff565b610aa481610c2f565b9093019392505050565b6000610ab982610bd0565b610ac38185610bdd565b9350610ad3818560208601610bff565b9290920192915050565b6000610aea602183610bd4565b7f4f6e6c7920757365722063616e2063616c6c20746869732066756e6374696f6e8152601760f91b602082015260400192915050565b610b29816103a2565b82525050565b60006105328284610aae565b602080825281016105328184610a08565b602080825281016105328184610a76565b6020808252810161011d81610add565b6020810161011d8284610b20565b60405181810167ffffffffffffffff81118282101715610b9a57600080fd5b604052919050565b600067ffffffffffffffff821115610bb957600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b60006001600160a01b03821661011d565b82818337506000910152565b60005b83811015610c1a578181015183820152602001610c02565b83811115610c29576000848401525b50505050565b601f01601f191690565b610c4281610be2565b8114610c4d57600080fd5b5056fea365627a7a723158203be4ac2cc55fc009136803b96f538a7afb41a97c4044877182ad42be629e32ed6c6578706572696d656e74616cf564736f6c63430005110040"

// DeployProxy deploys a new Ethereum contract, binding an instance of Proxy to it.
func DeployProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Proxy, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// GetAbi is a free data retrieval call binding the contract method 0x0fc97bd2.
//
// Solidity: function getAbi(address addr) constant returns(string)
func (_Proxy *ProxyCaller) GetAbi(opts *bind.CallOpts, addr common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Proxy.contract.Call(opts, out, "getAbi", addr)
	return *ret0, err
}

// GetAbi is a free data retrieval call binding the contract method 0x0fc97bd2.
//
// Solidity: function getAbi(address addr) constant returns(string)
func (_Proxy *ProxySession) GetAbi(addr common.Address) (string, error) {
	return _Proxy.Contract.GetAbi(&_Proxy.CallOpts, addr)
}

// GetAbi is a free data retrieval call binding the contract method 0x0fc97bd2.
//
// Solidity: function getAbi(address addr) constant returns(string)
func (_Proxy *ProxyCallerSession) GetAbi(addr common.Address) (string, error) {
	return _Proxy.Contract.GetAbi(&_Proxy.CallOpts, addr)
}

// GetPaths is a free data retrieval call binding the contract method 0x1767ae33.
//
// Solidity: function getPaths() constant returns(string[])
func (_Proxy *ProxyCaller) GetPaths(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Proxy.contract.Call(opts, out, "getPaths")
	return *ret0, err
}

// GetPaths is a free data retrieval call binding the contract method 0x1767ae33.
//
// Solidity: function getPaths() constant returns(string[])
func (_Proxy *ProxySession) GetPaths() ([]string, error) {
	return _Proxy.Contract.GetPaths(&_Proxy.CallOpts)
}

// GetPaths is a free data retrieval call binding the contract method 0x1767ae33.
//
// Solidity: function getPaths() constant returns(string[])
func (_Proxy *ProxyCallerSession) GetPaths() ([]string, error) {
	return _Proxy.Contract.GetPaths(&_Proxy.CallOpts)
}

// StringToUint256 is a free data retrieval call binding the contract method 0xac5d3723.
//
// Solidity: function stringToUint256(string _str) constant returns(uint256)
func (_Proxy *ProxyCaller) StringToUint256(opts *bind.CallOpts, _str string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Proxy.contract.Call(opts, out, "stringToUint256", _str)
	return *ret0, err
}

// StringToUint256 is a free data retrieval call binding the contract method 0xac5d3723.
//
// Solidity: function stringToUint256(string _str) constant returns(uint256)
func (_Proxy *ProxySession) StringToUint256(_str string) (*big.Int, error) {
	return _Proxy.Contract.StringToUint256(&_Proxy.CallOpts, _str)
}

// StringToUint256 is a free data retrieval call binding the contract method 0xac5d3723.
//
// Solidity: function stringToUint256(string _str) constant returns(uint256)
func (_Proxy *ProxyCallerSession) StringToUint256(_str string) (*big.Int, error) {
	return _Proxy.Contract.StringToUint256(&_Proxy.CallOpts, _str)
}

// AddPath is a paid mutator transaction binding the contract method 0x16b1ec2a.
//
// Solidity: function addPath(string _path, address _addr, string _abi) returns()
func (_Proxy *ProxyTransactor) AddPath(opts *bind.TransactOpts, _path string, _addr common.Address, _abi string) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addPath", _path, _addr, _abi)
}

// AddPath is a paid mutator transaction binding the contract method 0x16b1ec2a.
//
// Solidity: function addPath(string _path, address _addr, string _abi) returns()
func (_Proxy *ProxySession) AddPath(_path string, _addr common.Address, _abi string) (*types.Transaction, error) {
	return _Proxy.Contract.AddPath(&_Proxy.TransactOpts, _path, _addr, _abi)
}

// AddPath is a paid mutator transaction binding the contract method 0x16b1ec2a.
//
// Solidity: function addPath(string _path, address _addr, string _abi) returns()
func (_Proxy *ProxyTransactorSession) AddPath(_path string, _addr common.Address, _abi string) (*types.Transaction, error) {
	return _Proxy.Contract.AddPath(&_Proxy.TransactOpts, _path, _addr, _abi)
}

// ConstantCall is a paid mutator transaction binding the contract method 0x0e36518a.
//
// Solidity: function constantCall(string _name, bytes _argsWithMethodId) returns(bytes)
func (_Proxy *ProxyTransactor) ConstantCall(opts *bind.TransactOpts, _name string, _argsWithMethodId []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "constantCall", _name, _argsWithMethodId)
}

// ConstantCall is a paid mutator transaction binding the contract method 0x0e36518a.
//
// Solidity: function constantCall(string _name, bytes _argsWithMethodId) returns(bytes)
func (_Proxy *ProxySession) ConstantCall(_name string, _argsWithMethodId []byte) (*types.Transaction, error) {
	return _Proxy.Contract.ConstantCall(&_Proxy.TransactOpts, _name, _argsWithMethodId)
}

// ConstantCall is a paid mutator transaction binding the contract method 0x0e36518a.
//
// Solidity: function constantCall(string _name, bytes _argsWithMethodId) returns(bytes)
func (_Proxy *ProxyTransactorSession) ConstantCall(_name string, _argsWithMethodId []byte) (*types.Transaction, error) {
	return _Proxy.Contract.ConstantCall(&_Proxy.TransactOpts, _name, _argsWithMethodId)
}

// SendTransaction is a paid mutator transaction binding the contract method 0xaa7fb5a0.
//
// Solidity: function sendTransaction(string _uid, string _name, bytes _argsWithMethodId) returns(bytes)
func (_Proxy *ProxyTransactor) SendTransaction(opts *bind.TransactOpts, _uid string, _name string, _argsWithMethodId []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "sendTransaction", _uid, _name, _argsWithMethodId)
}

// SendTransaction is a paid mutator transaction binding the contract method 0xaa7fb5a0.
//
// Solidity: function sendTransaction(string _uid, string _name, bytes _argsWithMethodId) returns(bytes)
func (_Proxy *ProxySession) SendTransaction(_uid string, _name string, _argsWithMethodId []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SendTransaction(&_Proxy.TransactOpts, _uid, _name, _argsWithMethodId)
}

// SendTransaction is a paid mutator transaction binding the contract method 0xaa7fb5a0.
//
// Solidity: function sendTransaction(string _uid, string _name, bytes _argsWithMethodId) returns(bytes)
func (_Proxy *ProxyTransactorSession) SendTransaction(_uid string, _name string, _argsWithMethodId []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SendTransaction(&_Proxy.TransactOpts, _uid, _name, _argsWithMethodId)
}
