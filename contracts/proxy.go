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
const ProxyABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_argsWithMethodId\",\"type\":\"bytes\"}],\"name\":\"onSendTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"oncallContract\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_path\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_abi\",\"type\":\"string\"}],\"name\":\"addPath\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"bytesToUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_argsWithMethodId\",\"type\":\"bytes\"}],\"name\":\"constantCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAbi\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAbiByName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAddressByName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPaths\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_argsWithMethodId\",\"type\":\"bytes\"}],\"name\":\"sendTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"stringToUint256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ProxyFuncSigs maps the 4-byte function signature to its string representation.
var ProxyFuncSigs = map[string]string{
	"16b1ec2a": "addPath(string,address,string)",
	"02d06d05": "bytesToUint(bytes)",
	"0e36518a": "constantCall(string,bytes)",
	"0fc97bd2": "getAbi(address)",
	"d666f223": "getAbiByName(string)",
	"9a65ddec": "getAddressByName(string)",
	"1767ae33": "getPaths()",
	"aa7fb5a0": "sendTransaction(string,string,bytes)",
	"ac5d3723": "stringToUint256(string)",
}

// ProxyBin is the compiled bytecode used for deploying new contracts.
var ProxyBin = "0x608060405234801561001057600080fd5b50336000818152602081905260409020805460ff1916600190811790915580546001600160a01b0319169091179055610ee08061004e6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80631767ae33116100665780631767ae33146101095780639a65ddec1461011e578063aa7fb5a01461013e578063ac5d372314610151578063d666f2231461016457610093565b806302d06d05146100985780630e36518a146100c15780630fc97bd2146100e157806316b1ec2a146100f4575b600080fd5b6100ab6100a6366004610a7c565b610177565b6040516100b89190610dba565b60405180910390f35b6100d46100cf366004610b2d565b6101bb565b6040516100b89190610d5e565b6100d46100ef366004610a5e565b6101de565b610107610102366004610ab1565b610289565b005b610111610387565b6040516100b89190610d4d565b61013161012c366004610a7c565b610460565b6040516100b89190610d3f565b6100d461014c366004610b96565b610492565b6100ab61015f366004610a7c565b61065d565b6100d4610172366004610a7c565b6106e7565b600080805b83518110156101b4578060010184510360080260020a84828151811061019e57fe5b016020015160f81c02919091019060010161017c565b5092915050565b606060006101c884610460565b90506101d481846107a0565b9150505b92915050565b6001600160a01b03811660009081526003602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084526060939283018282801561027d5780601f106102525761010080835404028352916020019161027d565b820191906000526020600020905b81548152906001019060200180831161026057829003601f168201915b50505050509050919050565b3360009081526020819052604090205460ff166102c15760405162461bcd60e51b81526004016102b890610daa565b60405180910390fd5b60606102cc84610820565b9050826002826040516102df9190610d33565b9081526040516020918190038201902080546001600160a01b0319166001600160a01b039390931692909217909155600480546001810180835560009290925286519192610355927f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9092019190880190610967565b50506001600160a01b0383166000908152600360209081526040909120835161038092850190610967565b5050505050565b60606004805480602002602001604051908101604052809291908181526020016000905b828210156104565760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156104425780601f1061041757610100808354040283529160200191610442565b820191906000526020600020905b81548152906001019060200180831161042557829003601f168201915b5050505050815260200190600101906103ab565b5050505090505b90565b60006002826040516104729190610d33565b908152604051908190036020019020546001600160a01b03169050919050565b3360009081526020819052604090205460609060ff166104c45760405162461bcd60e51b81526004016102b890610daa565b6005846040516104d49190610d33565b9081526040519081900360200190205460ff161561059c576005846040516104fc9190610d33565b90815260408051602092819003830181206001908101805460029281161561010002600019011691909104601f810185900485028301850190935282825290929091908301828280156105905780601f1061056557610100808354040283529160200191610590565b820191906000526020600020905b81548152906001019060200180831161057357829003601f168201915b50505050509050610656565b60006105a784610460565b905060606105b5828561090e565b90507f984ef9d41b788e4f9d6804c2dde4849323bd85aa32a6f05ef2669ae7efdb56f58583866040516105ea93929190610d6f565b60405180910390a160408051808201825260018152602081018390529051600590610616908990610d33565b908152604051602091819003820190208251815460ff1916901515178155828201518051919261064e92600185019290910190610967565b509193505050505b9392505050565b805160009082908290815b818110156106dd57603084828151811061067e57fe5b016020015160f81c108015906106a85750603984828151811061069d57fe5b016020015160f81c11155b156106d55760308482815181106106bb57fe5b602001015160f81c60f81b60f81c0360ff1683600a020192505b600101610668565b5090949350505050565b606060006106f483610460565b6001600160a01b0381166000908152600360209081526040918290208054835160026001831615610100026000190190921691909104601f810184900484028201840190945283815293945091908301828280156107935780601f1061076857610100808354040283529160200191610793565b820191906000526020600020905b81548152906001019060200180831161077657829003601f168201915b5050505050915050919050565b60606000836001600160a01b0316836040516107bc9190610d33565b600060405180830381855afa9150503d80600081146107f7576040519150601f19603f3d011682016040523d82523d6000602084013e6107fc565b606091505b5092509050806101b4578160405162461bcd60e51b81526004016102b89190610d5e565b8051606090829060008060001983015b801561088057601760f91b6001600160f81b03191685828151811061085157fe5b01602001516001600160f81b031916141561087157806001019150610880565b60019092019160001901610830565b506060826040519080825280601f01601f1916602001820160405280156108ae576020820181803883390190505b50905060005b838110156109035785516001840193879181106108cd57fe5b602001015160f81c60f81b8282815181106108e457fe5b60200101906001600160f81b031916908160001a9053506001016108b4565b509695505050505050565b60606000836001600160a01b03168360405161092a9190610d33565b6000604051808303816000865af19150503d80600081146107f7576040519150601f19603f3d011682016040523d82523d6000602084013e6107fc565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109a857805160ff19168380011785556109d5565b828001600101855582156109d5579182015b828111156109d55782518255916020019190600101906109ba565b506109e19291506109e5565b5090565b61045d91905b808211156109e157600081556001016109eb565b80356101d881610e86565b600082601f830112610a1b57600080fd5b8135610a2e610a2982610def565b610dc8565b91508082526020830160208301858383011115610a4a57600080fd5b610a55838284610e40565b50505092915050565b600060208284031215610a7057600080fd5b60006101d484846109ff565b600060208284031215610a8e57600080fd5b813567ffffffffffffffff811115610aa557600080fd5b6101d484828501610a0a565b600080600060608486031215610ac657600080fd5b833567ffffffffffffffff811115610add57600080fd5b610ae986828701610a0a565b9350506020610afa868287016109ff565b925050604084013567ffffffffffffffff811115610b1757600080fd5b610b2386828701610a0a565b9150509250925092565b60008060408385031215610b4057600080fd5b823567ffffffffffffffff811115610b5757600080fd5b610b6385828601610a0a565b925050602083013567ffffffffffffffff811115610b8057600080fd5b610b8c85828601610a0a565b9150509250929050565b600080600060608486031215610bab57600080fd5b833567ffffffffffffffff811115610bc257600080fd5b610bce86828701610a0a565b935050602084013567ffffffffffffffff811115610beb57600080fd5b610afa86828701610a0a565b60006106568383610c80565b610c0c81610e2f565b82525050565b6000610c1d82610e1d565b610c278185610e21565b935083602082028501610c3985610e17565b8060005b85811015610c735784840389528151610c568582610bf7565b9450610c6183610e17565b60209a909a0199925050600101610c3d565b5091979650505050505050565b6000610c8b82610e1d565b610c958185610e21565b9350610ca5818560208601610e4c565b610cae81610e7c565b9093019392505050565b6000610cc382610e1d565b610ccd8185610e2a565b9350610cdd818560208601610e4c565b9290920192915050565b6000610cf4602183610e21565b7f4f6e6c7920757365722063616e2063616c6c20746869732066756e6374696f6e8152601760f91b602082015260400192915050565b610c0c8161045d565b60006106568284610cb8565b602081016101d88284610c03565b602080825281016106568184610c12565b602080825281016106568184610c80565b60608082528101610d808186610c80565b9050610d8f6020830185610c03565b8181036040830152610da18184610c80565b95945050505050565b602080825281016101d881610ce7565b602081016101d88284610d2a565b60405181810167ffffffffffffffff81118282101715610de757600080fd5b604052919050565b600067ffffffffffffffff821115610e0657600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b60006001600160a01b0382166101d8565b82818337506000910152565b60005b83811015610e67578181015183820152602001610e4f565b83811115610e76576000848401525b50505050565b601f01601f191690565b610e8f81610e2f565b8114610e9a57600080fd5b5056fea365627a7a72315820176c1a079923d3369c69c319270066dc5e8d6fc12bf64904ecb8670aa10d43976c6578706572696d656e74616cf564736f6c63430005110040"

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

// BytesToUint is a free data retrieval call binding the contract method 0x02d06d05.
//
// Solidity: function bytesToUint(bytes b) constant returns(uint256)
func (_Proxy *ProxyCaller) BytesToUint(opts *bind.CallOpts, b []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Proxy.contract.Call(opts, out, "bytesToUint", b)
	return *ret0, err
}

// BytesToUint is a free data retrieval call binding the contract method 0x02d06d05.
//
// Solidity: function bytesToUint(bytes b) constant returns(uint256)
func (_Proxy *ProxySession) BytesToUint(b []byte) (*big.Int, error) {
	return _Proxy.Contract.BytesToUint(&_Proxy.CallOpts, b)
}

// BytesToUint is a free data retrieval call binding the contract method 0x02d06d05.
//
// Solidity: function bytesToUint(bytes b) constant returns(uint256)
func (_Proxy *ProxyCallerSession) BytesToUint(b []byte) (*big.Int, error) {
	return _Proxy.Contract.BytesToUint(&_Proxy.CallOpts, b)
}

// ConstantCall is a free data retrieval call binding the contract method 0x0e36518a.
//
// Solidity: function constantCall(string _name, bytes _argsWithMethodId) constant returns(bytes)
func (_Proxy *ProxyCaller) ConstantCall(opts *bind.CallOpts, _name string, _argsWithMethodId []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Proxy.contract.Call(opts, out, "constantCall", _name, _argsWithMethodId)
	return *ret0, err
}

// ConstantCall is a free data retrieval call binding the contract method 0x0e36518a.
//
// Solidity: function constantCall(string _name, bytes _argsWithMethodId) constant returns(bytes)
func (_Proxy *ProxySession) ConstantCall(_name string, _argsWithMethodId []byte) ([]byte, error) {
	return _Proxy.Contract.ConstantCall(&_Proxy.CallOpts, _name, _argsWithMethodId)
}

// ConstantCall is a free data retrieval call binding the contract method 0x0e36518a.
//
// Solidity: function constantCall(string _name, bytes _argsWithMethodId) constant returns(bytes)
func (_Proxy *ProxyCallerSession) ConstantCall(_name string, _argsWithMethodId []byte) ([]byte, error) {
	return _Proxy.Contract.ConstantCall(&_Proxy.CallOpts, _name, _argsWithMethodId)
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

// GetAbiByName is a free data retrieval call binding the contract method 0xd666f223.
//
// Solidity: function getAbiByName(string _name) constant returns(string)
func (_Proxy *ProxyCaller) GetAbiByName(opts *bind.CallOpts, _name string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Proxy.contract.Call(opts, out, "getAbiByName", _name)
	return *ret0, err
}

// GetAbiByName is a free data retrieval call binding the contract method 0xd666f223.
//
// Solidity: function getAbiByName(string _name) constant returns(string)
func (_Proxy *ProxySession) GetAbiByName(_name string) (string, error) {
	return _Proxy.Contract.GetAbiByName(&_Proxy.CallOpts, _name)
}

// GetAbiByName is a free data retrieval call binding the contract method 0xd666f223.
//
// Solidity: function getAbiByName(string _name) constant returns(string)
func (_Proxy *ProxyCallerSession) GetAbiByName(_name string) (string, error) {
	return _Proxy.Contract.GetAbiByName(&_Proxy.CallOpts, _name)
}

// GetAddressByName is a free data retrieval call binding the contract method 0x9a65ddec.
//
// Solidity: function getAddressByName(string _name) constant returns(address)
func (_Proxy *ProxyCaller) GetAddressByName(opts *bind.CallOpts, _name string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Proxy.contract.Call(opts, out, "getAddressByName", _name)
	return *ret0, err
}

// GetAddressByName is a free data retrieval call binding the contract method 0x9a65ddec.
//
// Solidity: function getAddressByName(string _name) constant returns(address)
func (_Proxy *ProxySession) GetAddressByName(_name string) (common.Address, error) {
	return _Proxy.Contract.GetAddressByName(&_Proxy.CallOpts, _name)
}

// GetAddressByName is a free data retrieval call binding the contract method 0x9a65ddec.
//
// Solidity: function getAddressByName(string _name) constant returns(address)
func (_Proxy *ProxyCallerSession) GetAddressByName(_name string) (common.Address, error) {
	return _Proxy.Contract.GetAddressByName(&_Proxy.CallOpts, _name)
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

// ProxyOnSendTransactionIterator is returned from FilterOnSendTransaction and is used to iterate over the raw logs and unpacked data for OnSendTransaction events raised by the Proxy contract.
type ProxyOnSendTransactionIterator struct {
	Event *ProxyOnSendTransaction // Event containing the contract specifics and raw log

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
func (it *ProxyOnSendTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyOnSendTransaction)
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
		it.Event = new(ProxyOnSendTransaction)
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
func (it *ProxyOnSendTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyOnSendTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyOnSendTransaction represents a OnSendTransaction event raised by the Proxy contract.
type ProxyOnSendTransaction struct {
	Name             string
	ContractAddress  common.Address
	ArgsWithMethodId []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnSendTransaction is a free log retrieval operation binding the contract event 0x984ef9d41b788e4f9d6804c2dde4849323bd85aa32a6f05ef2669ae7efdb56f5.
//
// Solidity: event onSendTransaction(string name, address contractAddress, bytes _argsWithMethodId)
func (_Proxy *ProxyFilterer) FilterOnSendTransaction(opts *bind.FilterOpts) (*ProxyOnSendTransactionIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "onSendTransaction")
	if err != nil {
		return nil, err
	}
	return &ProxyOnSendTransactionIterator{contract: _Proxy.contract, event: "onSendTransaction", logs: logs, sub: sub}, nil
}

// WatchOnSendTransaction is a free log subscription operation binding the contract event 0x984ef9d41b788e4f9d6804c2dde4849323bd85aa32a6f05ef2669ae7efdb56f5.
//
// Solidity: event onSendTransaction(string name, address contractAddress, bytes _argsWithMethodId)
func (_Proxy *ProxyFilterer) WatchOnSendTransaction(opts *bind.WatchOpts, sink chan<- *ProxyOnSendTransaction) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "onSendTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyOnSendTransaction)
				if err := _Proxy.contract.UnpackLog(event, "onSendTransaction", log); err != nil {
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

// ParseOnSendTransaction is a log parse operation binding the contract event 0x984ef9d41b788e4f9d6804c2dde4849323bd85aa32a6f05ef2669ae7efdb56f5.
//
// Solidity: event onSendTransaction(string name, address contractAddress, bytes _argsWithMethodId)
func (_Proxy *ProxyFilterer) ParseOnSendTransaction(log types.Log) (*ProxyOnSendTransaction, error) {
	event := new(ProxyOnSendTransaction)
	if err := _Proxy.contract.UnpackLog(event, "onSendTransaction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ProxyOncallContractIterator is returned from FilterOncallContract and is used to iterate over the raw logs and unpacked data for OncallContract events raised by the Proxy contract.
type ProxyOncallContractIterator struct {
	Event *ProxyOncallContract // Event containing the contract specifics and raw log

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
func (it *ProxyOncallContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyOncallContract)
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
		it.Event = new(ProxyOncallContract)
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
func (it *ProxyOncallContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyOncallContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyOncallContract represents a OncallContract event raised by the Proxy contract.
type ProxyOncallContract struct {
	Result []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOncallContract is a free log retrieval operation binding the contract event 0xbe5955717a3ed92eeff6dad42e34417009e08d41c2147c7b1b9c73543578df72.
//
// Solidity: event oncallContract(bytes result)
func (_Proxy *ProxyFilterer) FilterOncallContract(opts *bind.FilterOpts) (*ProxyOncallContractIterator, error) {

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "oncallContract")
	if err != nil {
		return nil, err
	}
	return &ProxyOncallContractIterator{contract: _Proxy.contract, event: "oncallContract", logs: logs, sub: sub}, nil
}

// WatchOncallContract is a free log subscription operation binding the contract event 0xbe5955717a3ed92eeff6dad42e34417009e08d41c2147c7b1b9c73543578df72.
//
// Solidity: event oncallContract(bytes result)
func (_Proxy *ProxyFilterer) WatchOncallContract(opts *bind.WatchOpts, sink chan<- *ProxyOncallContract) (event.Subscription, error) {

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "oncallContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyOncallContract)
				if err := _Proxy.contract.UnpackLog(event, "oncallContract", log); err != nil {
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

// ParseOncallContract is a log parse operation binding the contract event 0xbe5955717a3ed92eeff6dad42e34417009e08d41c2147c7b1b9c73543578df72.
//
// Solidity: event oncallContract(bytes result)
func (_Proxy *ProxyFilterer) ParseOncallContract(log types.Log) (*ProxyOncallContract, error) {
	event := new(ProxyOncallContract)
	if err := _Proxy.contract.UnpackLog(event, "oncallContract", log); err != nil {
		return nil, err
	}
	return event, nil
}
