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

// AnchorChainABI is the input ABI used to generate the binding from.
const AnchorChainABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"_data\",\"type\":\"string[]\"}],\"name\":\"onConnect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"name\":\"oncallContract\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"name\":\"callContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_result\",\"type\":\"string[]\"}],\"name\":\"callback\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_data\",\"type\":\"string[]\"}],\"name\":\"connect\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_method\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_args\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_callbackPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_callbackMethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_crossChainType\",\"type\":\"string\"}],\"name\":\"connectChain\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_args\",\"type\":\"string[]\"}],\"name\":\"crossChain\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"hub\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"scanBlocks\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_method\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_callbackPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_callbackMethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_crossChainType\",\"type\":\"string\"}],\"name\":\"setTargetChainPath\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"viewBlock\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AnchorChainFuncSigs maps the 4-byte function signature to its string representation.
var AnchorChainFuncSigs = map[string]string{
	"421b2d8b": "addUser(address)",
	"9ad936b0": "callContract(string)",
	"8b9aebed": "callback(bool,string[])",
	"2fe630f6": "connect(string[])",
	"71d25b1a": "connectChain(string,string,string[],string,string,string)",
	"d27c6533": "crossChain(string[])",
	"19ab453c": "init(address)",
	"98575188": "removeUser(address)",
	"803416ce": "scanBlocks()",
	"2334105c": "setTargetChainPath(string,string,string,string,string)",
	"394a7362": "viewBlock(uint256)",
}

// AnchorChainBin is the compiled bytecode used for deploying new contracts.
var AnchorChainBin = "0x608060405234801561001057600080fd5b50336000818152600660205260409020805460ff19166001179055600780546001600160a01b03191690911790556113aa8061004d6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806371d25b1a1161007157806371d25b1a14610132578063803416ce146101455780638b9aebed1461014d57806398575188146101605780639ad936b014610173578063d27c653314610186576100a9565b806319ab453c146100ae5780632334105c146100c35780632fe630f6146100d6578063394a7362146100ff578063421b2d8b1461011f575b600080fd5b6100c16100bc366004610b74565b610199565b005b6100c16100d1366004610d97565b6101ee565b6100e96100e4366004610b9a565b610283565b6040516100f691906110f6565b60405180910390f35b61011261010d366004610e7e565b6103f9565b6040516100f69190611107565b6100c161012d366004610b74565b6104c4565b610112610140366004610c87565b610512565b6100e9610619565b6100e961015b366004610bce565b6106f2565b6100c161016e366004610b74565b61074b565b6100c1610181366004610c1f565b6107ba565b610112610194366004610b9a565b6107f4565b6007546001600160a01b031633146101cc5760405162461bcd60e51b81526004016101c39061120e565b60405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6007546001600160a01b031633146102185760405162461bcd60e51b81526004016101c39061120e565b845161022b9060019060208801906108fa565b50835161023f9060029060208701906108fa565b5082516102539060039060208601906108fa565b5081516102679060049060208501906108fa565b50805161027b9060059060208401906108fa565b505050505050565b3360009081526006602052604090205460609060ff166102b55760405162461bcd60e51b81526004016101c39061122e565b7fc32864ce20a82077c97cf7b6f75aac4c666318b4ebc19aa8a877625cd7f89704826040516102e491906110f6565b60405180910390a160005b825181101561035357600983828151811061030657fe5b602002602001015160405161031b91906110ea565b9081526040519081900360200190205460ff161561034b5760405162461bcd60e51b81526004016101c39061123e565b6001016102ef565b5060005b82518110156103f2576001600984838151811061037057fe5b602002602001015160405161038591906110ea565b908152602001604051809103902060006101000a81548160ff021916908315150217905550600a8382815181106103b857fe5b60209081029190910181015182546001810180855560009485529383902082516103e894919092019201906108fa565b5050600101610357565b5090919050565b600a54606090821061041d5760405162461bcd60e51b81526004016101c39061121e565b600a828154811061042a57fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156104b85780601f1061048d576101008083540402835291602001916104b8565b820191906000526020600020905b81548152906001019060200180831161049b57829003601f168201915b50505050509050919050565b6007546001600160a01b031633146104ee5760405162461bcd60e51b81526004016101c39061120e565b6001600160a01b03166000908152600660205260409020805460ff19166001179055565b3360009081526006602052604090205460609060ff166105445760405162461bcd60e51b81526004016101c39061122e565b600054604051631cde0c3d60e21b81526060916001600160a01b03169063737830f49061057f908b908b908b908b908b908b90600401611118565b600060405180830381600087803b15801561059957600080fd5b505af11580156105ad573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526105d59190810190610c53565b90507f24cca36728e0487747cbf830bd8ad5c42a077fa2c2299a6c6da3ffa5224ae45a816040516106069190611107565b60405180910390a1979650505050505050565b6060600a805480602002602001604051908101604052809291908181526020016000905b828210156106e85760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156106d45780601f106106a9576101008083540402835291602001916106d4565b820191906000526020600020905b8154815290600101906020018083116106b757829003601f168201915b50505050508152602001906001019061063d565b5050505090505b90565b60608215610742578160088360008151811061070a57fe5b602002602001015160405161071f91906110ea565b90815260200160405180910390209080519060200190610740929190610978565b505b50805b92915050565b6007546001600160a01b031633146107755760405162461bcd60e51b81526004016101c39061120e565b6001600160a01b03811660009081526006602052604090205460ff16156107b7576001600160a01b0381166000908152600660205260409020805460ff191690555b50565b7f24cca36728e0487747cbf830bd8ad5c42a077fa2c2299a6c6da3ffa5224ae45a816040516107e99190611107565b60405180910390a150565b3360009081526006602052604090205460609060ff166108265760405162461bcd60e51b81526004016101c39061122e565b600054604051631cde0c3d60e21b81526060916001600160a01b03169063737830f4906108659060019060029088906003906004906005908201611199565b600060405180830381600087803b15801561087f57600080fd5b505af1158015610893573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108bb9190810190610c53565b90507f24cca36728e0487747cbf830bd8ad5c42a077fa2c2299a6c6da3ffa5224ae45a816040516108ec9190611107565b60405180910390a192915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061093b57805160ff1916838001178555610968565b82800160010185558215610968579182015b8281111561096857825182559160200191906001019061094d565b506109749291506109d1565b5090565b8280548282559060005260206000209081019282156109c5579160200282015b828111156109c557825180516109b59184916020909101906108fa565b5091602001919060010190610998565b506109749291506109eb565b6106ef91905b8082111561097457600081556001016109d7565b6106ef91905b80821115610974576000610a058282610a0e565b506001016109f1565b50805460018160011615610100020316600290046000825580601f10610a3457506107b7565b601f0160209004906000526020600020908101906107b791906109d1565b803561074581611341565b600082601f830112610a6e57600080fd5b8135610a81610a7c82611274565b61124e565b81815260209384019390925082018360005b83811015610abf5781358601610aa98882610ad4565b8452506020928301929190910190600101610a93565b5050505092915050565b803561074581611355565b600082601f830112610ae557600080fd5b8135610af3610a7c82611294565b91508082526020830160208301858383011115610b0f57600080fd5b610b1a8382846112fb565b50505092915050565b600082601f830112610b3457600080fd5b8151610b42610a7c82611294565b91508082526020830160208301858383011115610b5e57600080fd5b610b1a838284611307565b80356107458161135e565b600060208284031215610b8657600080fd5b6000610b928484610a52565b949350505050565b600060208284031215610bac57600080fd5b81356001600160401b03811115610bc257600080fd5b610b9284828501610a5d565b60008060408385031215610be157600080fd5b6000610bed8585610ac9565b92505060208301356001600160401b03811115610c0957600080fd5b610c1585828601610a5d565b9150509250929050565b600060208284031215610c3157600080fd5b81356001600160401b03811115610c4757600080fd5b610b9284828501610ad4565b600060208284031215610c6557600080fd5b81516001600160401b03811115610c7b57600080fd5b610b9284828501610b23565b60008060008060008060c08789031215610ca057600080fd5b86356001600160401b03811115610cb657600080fd5b610cc289828a01610ad4565b96505060208701356001600160401b03811115610cde57600080fd5b610cea89828a01610ad4565b95505060408701356001600160401b03811115610d0657600080fd5b610d1289828a01610a5d565b94505060608701356001600160401b03811115610d2e57600080fd5b610d3a89828a01610ad4565b93505060808701356001600160401b03811115610d5657600080fd5b610d6289828a01610ad4565b92505060a08701356001600160401b03811115610d7e57600080fd5b610d8a89828a01610ad4565b9150509295509295509295565b600080600080600060a08688031215610daf57600080fd5b85356001600160401b03811115610dc557600080fd5b610dd188828901610ad4565b95505060208601356001600160401b03811115610ded57600080fd5b610df988828901610ad4565b94505060408601356001600160401b03811115610e1557600080fd5b610e2188828901610ad4565b93505060608601356001600160401b03811115610e3d57600080fd5b610e4988828901610ad4565b92505060808601356001600160401b03811115610e6557600080fd5b610e7188828901610ad4565b9150509295509295909350565b600060208284031215610e9057600080fd5b6000610b928484610b69565b6000610ea88383610f1d565b9392505050565b6000610eba826112cd565b610ec481856112d1565b935083602082028501610ed6856112bb565b8060005b85811015610f105784840389528151610ef38582610e9c565b9450610efe836112bb565b60209a909a0199925050600101610eda565b5091979650505050505050565b6000610f28826112cd565b610f3281856112d1565b9350610f42818560208601611307565b610f4b81611337565b9093019392505050565b6000610f60826112cd565b610f6a81856112da565b9350610f7a818560208601611307565b9290920192915050565b600081546001811660008114610fa15760018114610fc757611006565b607f6002830416610fb281876112d1565b60ff1984168152955050602085019250611006565b60028204610fd581876112d1565b9550610fe0856112c1565b60005b82811015610fff57815488820152600190910190602001610fe3565b8701945050505b505092915050565b600061101b6022836112d1565b7f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f815261371760f11b602082015260400192915050565b600061105f600d836112d1565b6c1a5b99195e081a5b9d985b1a59609a1b815260200192915050565b60006110886021836112d1565b7f4f6e6c7920757365722063616e2063616c6c20746869732066756e6374696f6e8152601760f91b602082015260400192915050565b60006110cb6010836112d1565b6f1a185cda081a185908195e1a5cdd195960821b815260200192915050565b6000610ea88284610f55565b60208082528101610ea88184610eaf565b60208082528101610ea88184610f1d565b60c080825281016111298189610f1d565b9050818103602083015261113d8188610f1d565b905081810360408301526111518187610eaf565b905081810360608301526111658186610f1d565b905081810360808301526111798185610f1d565b905081810360a083015261118d8184610f1d565b98975050505050505050565b60c080825281016111aa8189610f84565b905081810360208301526111be8188610f84565b905081810360408301526111d28187610eaf565b905081810360608301526111e68186610f84565b905081810360808301526111fa8185610f84565b905081810360a083015261118d8184610f84565b602080825281016107458161100e565b6020808252810161074581611052565b602080825281016107458161107b565b60208082528101610745816110be565b6040518181016001600160401b038111828210171561126c57600080fd5b604052919050565b60006001600160401b0382111561128a57600080fd5b5060209081020190565b60006001600160401b038211156112aa57600080fd5b506020601f91909101601f19160190565b60200190565b60009081526020902090565b5190565b90815260200190565b919050565b6000610745826112ef565b151590565b6001600160a01b031690565b82818337506000910152565b60005b8381101561132257818101518382015260200161130a565b83811115611331576000848401525b50505050565b601f01601f191690565b61134a816112df565b81146107b757600080fd5b61134a816112ea565b61134a816106ef56fea365627a7a723158208a13dba6a41f70974c775c4406335d2e9f632c00876e23ac22ee554c48ac9bfd6c6578706572696d656e74616cf564736f6c63430005110040"

// DeployAnchorChain deploys a new Ethereum contract, binding an instance of AnchorChain to it.
func DeployAnchorChain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AnchorChain, error) {
	parsed, err := abi.JSON(strings.NewReader(AnchorChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AnchorChainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AnchorChain{AnchorChainCaller: AnchorChainCaller{contract: contract}, AnchorChainTransactor: AnchorChainTransactor{contract: contract}, AnchorChainFilterer: AnchorChainFilterer{contract: contract}}, nil
}

// AnchorChain is an auto generated Go binding around an Ethereum contract.
type AnchorChain struct {
	AnchorChainCaller     // Read-only binding to the contract
	AnchorChainTransactor // Write-only binding to the contract
	AnchorChainFilterer   // Log filterer for contract events
}

// AnchorChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnchorChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchorChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnchorChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchorChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnchorChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchorChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnchorChainSession struct {
	Contract     *AnchorChain      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnchorChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnchorChainCallerSession struct {
	Contract *AnchorChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AnchorChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnchorChainTransactorSession struct {
	Contract     *AnchorChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AnchorChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnchorChainRaw struct {
	Contract *AnchorChain // Generic contract binding to access the raw methods on
}

// AnchorChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnchorChainCallerRaw struct {
	Contract *AnchorChainCaller // Generic read-only contract binding to access the raw methods on
}

// AnchorChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnchorChainTransactorRaw struct {
	Contract *AnchorChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnchorChain creates a new instance of AnchorChain, bound to a specific deployed contract.
func NewAnchorChain(address common.Address, backend bind.ContractBackend) (*AnchorChain, error) {
	contract, err := bindAnchorChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AnchorChain{AnchorChainCaller: AnchorChainCaller{contract: contract}, AnchorChainTransactor: AnchorChainTransactor{contract: contract}, AnchorChainFilterer: AnchorChainFilterer{contract: contract}}, nil
}

// NewAnchorChainCaller creates a new read-only instance of AnchorChain, bound to a specific deployed contract.
func NewAnchorChainCaller(address common.Address, caller bind.ContractCaller) (*AnchorChainCaller, error) {
	contract, err := bindAnchorChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnchorChainCaller{contract: contract}, nil
}

// NewAnchorChainTransactor creates a new write-only instance of AnchorChain, bound to a specific deployed contract.
func NewAnchorChainTransactor(address common.Address, transactor bind.ContractTransactor) (*AnchorChainTransactor, error) {
	contract, err := bindAnchorChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnchorChainTransactor{contract: contract}, nil
}

// NewAnchorChainFilterer creates a new log filterer instance of AnchorChain, bound to a specific deployed contract.
func NewAnchorChainFilterer(address common.Address, filterer bind.ContractFilterer) (*AnchorChainFilterer, error) {
	contract, err := bindAnchorChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnchorChainFilterer{contract: contract}, nil
}

// bindAnchorChain binds a generic wrapper to an already deployed contract.
func bindAnchorChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AnchorChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnchorChain *AnchorChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AnchorChain.Contract.AnchorChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnchorChain *AnchorChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnchorChain.Contract.AnchorChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnchorChain *AnchorChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnchorChain.Contract.AnchorChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnchorChain *AnchorChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AnchorChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnchorChain *AnchorChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnchorChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnchorChain *AnchorChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnchorChain.Contract.contract.Transact(opts, method, params...)
}

// ScanBlocks is a free data retrieval call binding the contract method 0x803416ce.
//
// Solidity: function scanBlocks() constant returns(string[])
func (_AnchorChain *AnchorChainCaller) ScanBlocks(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _AnchorChain.contract.Call(opts, out, "scanBlocks")
	return *ret0, err
}

// ScanBlocks is a free data retrieval call binding the contract method 0x803416ce.
//
// Solidity: function scanBlocks() constant returns(string[])
func (_AnchorChain *AnchorChainSession) ScanBlocks() ([]string, error) {
	return _AnchorChain.Contract.ScanBlocks(&_AnchorChain.CallOpts)
}

// ScanBlocks is a free data retrieval call binding the contract method 0x803416ce.
//
// Solidity: function scanBlocks() constant returns(string[])
func (_AnchorChain *AnchorChainCallerSession) ScanBlocks() ([]string, error) {
	return _AnchorChain.Contract.ScanBlocks(&_AnchorChain.CallOpts)
}

// ViewBlock is a free data retrieval call binding the contract method 0x394a7362.
//
// Solidity: function viewBlock(uint256 i) constant returns(string)
func (_AnchorChain *AnchorChainCaller) ViewBlock(opts *bind.CallOpts, i *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AnchorChain.contract.Call(opts, out, "viewBlock", i)
	return *ret0, err
}

// ViewBlock is a free data retrieval call binding the contract method 0x394a7362.
//
// Solidity: function viewBlock(uint256 i) constant returns(string)
func (_AnchorChain *AnchorChainSession) ViewBlock(i *big.Int) (string, error) {
	return _AnchorChain.Contract.ViewBlock(&_AnchorChain.CallOpts, i)
}

// ViewBlock is a free data retrieval call binding the contract method 0x394a7362.
//
// Solidity: function viewBlock(uint256 i) constant returns(string)
func (_AnchorChain *AnchorChainCallerSession) ViewBlock(i *big.Int) (string, error) {
	return _AnchorChain.Contract.ViewBlock(&_AnchorChain.CallOpts, i)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address user) returns()
func (_AnchorChain *AnchorChainTransactor) AddUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _AnchorChain.contract.Transact(opts, "addUser", user)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address user) returns()
func (_AnchorChain *AnchorChainSession) AddUser(user common.Address) (*types.Transaction, error) {
	return _AnchorChain.Contract.AddUser(&_AnchorChain.TransactOpts, user)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address user) returns()
func (_AnchorChain *AnchorChainTransactorSession) AddUser(user common.Address) (*types.Transaction, error) {
	return _AnchorChain.Contract.AddUser(&_AnchorChain.TransactOpts, user)
}

// CallContract is a paid mutator transaction binding the contract method 0x9ad936b0.
//
// Solidity: function callContract(string result) returns()
func (_AnchorChain *AnchorChainTransactor) CallContract(opts *bind.TransactOpts, result string) (*types.Transaction, error) {
	return _AnchorChain.contract.Transact(opts, "callContract", result)
}

// CallContract is a paid mutator transaction binding the contract method 0x9ad936b0.
//
// Solidity: function callContract(string result) returns()
func (_AnchorChain *AnchorChainSession) CallContract(result string) (*types.Transaction, error) {
	return _AnchorChain.Contract.CallContract(&_AnchorChain.TransactOpts, result)
}

// CallContract is a paid mutator transaction binding the contract method 0x9ad936b0.
//
// Solidity: function callContract(string result) returns()
func (_AnchorChain *AnchorChainTransactorSession) CallContract(result string) (*types.Transaction, error) {
	return _AnchorChain.Contract.CallContract(&_AnchorChain.TransactOpts, result)
}

// Callback is a paid mutator transaction binding the contract method 0x8b9aebed.
//
// Solidity: function callback(bool state, string[] _result) returns(string[])
func (_AnchorChain *AnchorChainTransactor) Callback(opts *bind.TransactOpts, state bool, _result []string) (*types.Transaction, error) {
	return _AnchorChain.contract.Transact(opts, "callback", state, _result)
}

// Callback is a paid mutator transaction binding the contract method 0x8b9aebed.
//
// Solidity: function callback(bool state, string[] _result) returns(string[])
func (_AnchorChain *AnchorChainSession) Callback(state bool, _result []string) (*types.Transaction, error) {
	return _AnchorChain.Contract.Callback(&_AnchorChain.TransactOpts, state, _result)
}

// Callback is a paid mutator transaction binding the contract method 0x8b9aebed.
//
// Solidity: function callback(bool state, string[] _result) returns(string[])
func (_AnchorChain *AnchorChainTransactorSession) Callback(state bool, _result []string) (*types.Transaction, error) {
	return _AnchorChain.Contract.Callback(&_AnchorChain.TransactOpts, state, _result)
}

// Connect is a paid mutator transaction binding the contract method 0x2fe630f6.
//
// Solidity: function connect(string[] _data) returns(string[])
func (_AnchorChain *AnchorChainTransactor) Connect(opts *bind.TransactOpts, _data []string) (*types.Transaction, error) {
	return _AnchorChain.contract.Transact(opts, "connect", _data)
}

// Connect is a paid mutator transaction binding the contract method 0x2fe630f6.
//
// Solidity: function connect(string[] _data) returns(string[])
func (_AnchorChain *AnchorChainSession) Connect(_data []string) (*types.Transaction, error) {
	return _AnchorChain.Contract.Connect(&_AnchorChain.TransactOpts, _data)
}

// Connect is a paid mutator transaction binding the contract method 0x2fe630f6.
//
// Solidity: function connect(string[] _data) returns(string[])
func (_AnchorChain *AnchorChainTransactorSession) Connect(_data []string) (*types.Transaction, error) {
	return _AnchorChain.Contract.Connect(&_AnchorChain.TransactOpts, _data)
}

// ConnectChain is a paid mutator transaction binding the contract method 0x71d25b1a.
//
// Solidity: function connectChain(string _path, string _method, string[] _args, string _callbackPath, string _callbackMethod, string _crossChainType) returns(string)
func (_AnchorChain *AnchorChainTransactor) ConnectChain(opts *bind.TransactOpts, _path string, _method string, _args []string, _callbackPath string, _callbackMethod string, _crossChainType string) (*types.Transaction, error) {
	return _AnchorChain.contract.Transact(opts, "connectChain", _path, _method, _args, _callbackPath, _callbackMethod, _crossChainType)
}

// ConnectChain is a paid mutator transaction binding the contract method 0x71d25b1a.
//
// Solidity: function connectChain(string _path, string _method, string[] _args, string _callbackPath, string _callbackMethod, string _crossChainType) returns(string)
func (_AnchorChain *AnchorChainSession) ConnectChain(_path string, _method string, _args []string, _callbackPath string, _callbackMethod string, _crossChainType string) (*types.Transaction, error) {
	return _AnchorChain.Contract.ConnectChain(&_AnchorChain.TransactOpts, _path, _method, _args, _callbackPath, _callbackMethod, _crossChainType)
}

// ConnectChain is a paid mutator transaction binding the contract method 0x71d25b1a.
//
// Solidity: function connectChain(string _path, string _method, string[] _args, string _callbackPath, string _callbackMethod, string _crossChainType) returns(string)
func (_AnchorChain *AnchorChainTransactorSession) ConnectChain(_path string, _method string, _args []string, _callbackPath string, _callbackMethod string, _crossChainType string) (*types.Transaction, error) {
	return _AnchorChain.Contract.ConnectChain(&_AnchorChain.TransactOpts, _path, _method, _args, _callbackPath, _callbackMethod, _crossChainType)
}

// CrossChain is a paid mutator transaction binding the contract method 0xd27c6533.
//
// Solidity: function crossChain(string[] _args) returns(string)
func (_AnchorChain *AnchorChainTransactor) CrossChain(opts *bind.TransactOpts, _args []string) (*types.Transaction, error) {
	return _AnchorChain.contract.Transact(opts, "crossChain", _args)
}

// CrossChain is a paid mutator transaction binding the contract method 0xd27c6533.
//
// Solidity: function crossChain(string[] _args) returns(string)
func (_AnchorChain *AnchorChainSession) CrossChain(_args []string) (*types.Transaction, error) {
	return _AnchorChain.Contract.CrossChain(&_AnchorChain.TransactOpts, _args)
}

// CrossChain is a paid mutator transaction binding the contract method 0xd27c6533.
//
// Solidity: function crossChain(string[] _args) returns(string)
func (_AnchorChain *AnchorChainTransactorSession) CrossChain(_args []string) (*types.Transaction, error) {
	return _AnchorChain.Contract.CrossChain(&_AnchorChain.TransactOpts, _args)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address hub) returns()
func (_AnchorChain *AnchorChainTransactor) Init(opts *bind.TransactOpts, hub common.Address) (*types.Transaction, error) {
	return _AnchorChain.contract.Transact(opts, "init", hub)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address hub) returns()
func (_AnchorChain *AnchorChainSession) Init(hub common.Address) (*types.Transaction, error) {
	return _AnchorChain.Contract.Init(&_AnchorChain.TransactOpts, hub)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address hub) returns()
func (_AnchorChain *AnchorChainTransactorSession) Init(hub common.Address) (*types.Transaction, error) {
	return _AnchorChain.Contract.Init(&_AnchorChain.TransactOpts, hub)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address user) returns()
func (_AnchorChain *AnchorChainTransactor) RemoveUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _AnchorChain.contract.Transact(opts, "removeUser", user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address user) returns()
func (_AnchorChain *AnchorChainSession) RemoveUser(user common.Address) (*types.Transaction, error) {
	return _AnchorChain.Contract.RemoveUser(&_AnchorChain.TransactOpts, user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address user) returns()
func (_AnchorChain *AnchorChainTransactorSession) RemoveUser(user common.Address) (*types.Transaction, error) {
	return _AnchorChain.Contract.RemoveUser(&_AnchorChain.TransactOpts, user)
}

// SetTargetChainPath is a paid mutator transaction binding the contract method 0x2334105c.
//
// Solidity: function setTargetChainPath(string _path, string _method, string _callbackPath, string _callbackMethod, string _crossChainType) returns()
func (_AnchorChain *AnchorChainTransactor) SetTargetChainPath(opts *bind.TransactOpts, _path string, _method string, _callbackPath string, _callbackMethod string, _crossChainType string) (*types.Transaction, error) {
	return _AnchorChain.contract.Transact(opts, "setTargetChainPath", _path, _method, _callbackPath, _callbackMethod, _crossChainType)
}

// SetTargetChainPath is a paid mutator transaction binding the contract method 0x2334105c.
//
// Solidity: function setTargetChainPath(string _path, string _method, string _callbackPath, string _callbackMethod, string _crossChainType) returns()
func (_AnchorChain *AnchorChainSession) SetTargetChainPath(_path string, _method string, _callbackPath string, _callbackMethod string, _crossChainType string) (*types.Transaction, error) {
	return _AnchorChain.Contract.SetTargetChainPath(&_AnchorChain.TransactOpts, _path, _method, _callbackPath, _callbackMethod, _crossChainType)
}

// SetTargetChainPath is a paid mutator transaction binding the contract method 0x2334105c.
//
// Solidity: function setTargetChainPath(string _path, string _method, string _callbackPath, string _callbackMethod, string _crossChainType) returns()
func (_AnchorChain *AnchorChainTransactorSession) SetTargetChainPath(_path string, _method string, _callbackPath string, _callbackMethod string, _crossChainType string) (*types.Transaction, error) {
	return _AnchorChain.Contract.SetTargetChainPath(&_AnchorChain.TransactOpts, _path, _method, _callbackPath, _callbackMethod, _crossChainType)
}

// AnchorChainOnConnectIterator is returned from FilterOnConnect and is used to iterate over the raw logs and unpacked data for OnConnect events raised by the AnchorChain contract.
type AnchorChainOnConnectIterator struct {
	Event *AnchorChainOnConnect // Event containing the contract specifics and raw log

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
func (it *AnchorChainOnConnectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnchorChainOnConnect)
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
		it.Event = new(AnchorChainOnConnect)
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
func (it *AnchorChainOnConnectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnchorChainOnConnectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnchorChainOnConnect represents a OnConnect event raised by the AnchorChain contract.
type AnchorChainOnConnect struct {
	Data []string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOnConnect is a free log retrieval operation binding the contract event 0xc32864ce20a82077c97cf7b6f75aac4c666318b4ebc19aa8a877625cd7f89704.
//
// Solidity: event onConnect(string[] _data)
func (_AnchorChain *AnchorChainFilterer) FilterOnConnect(opts *bind.FilterOpts) (*AnchorChainOnConnectIterator, error) {

	logs, sub, err := _AnchorChain.contract.FilterLogs(opts, "onConnect")
	if err != nil {
		return nil, err
	}
	return &AnchorChainOnConnectIterator{contract: _AnchorChain.contract, event: "onConnect", logs: logs, sub: sub}, nil
}

// WatchOnConnect is a free log subscription operation binding the contract event 0xc32864ce20a82077c97cf7b6f75aac4c666318b4ebc19aa8a877625cd7f89704.
//
// Solidity: event onConnect(string[] _data)
func (_AnchorChain *AnchorChainFilterer) WatchOnConnect(opts *bind.WatchOpts, sink chan<- *AnchorChainOnConnect) (event.Subscription, error) {

	logs, sub, err := _AnchorChain.contract.WatchLogs(opts, "onConnect")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnchorChainOnConnect)
				if err := _AnchorChain.contract.UnpackLog(event, "onConnect", log); err != nil {
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

// ParseOnConnect is a log parse operation binding the contract event 0xc32864ce20a82077c97cf7b6f75aac4c666318b4ebc19aa8a877625cd7f89704.
//
// Solidity: event onConnect(string[] _data)
func (_AnchorChain *AnchorChainFilterer) ParseOnConnect(log types.Log) (*AnchorChainOnConnect, error) {
	event := new(AnchorChainOnConnect)
	if err := _AnchorChain.contract.UnpackLog(event, "onConnect", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AnchorChainOncallContractIterator is returned from FilterOncallContract and is used to iterate over the raw logs and unpacked data for OncallContract events raised by the AnchorChain contract.
type AnchorChainOncallContractIterator struct {
	Event *AnchorChainOncallContract // Event containing the contract specifics and raw log

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
func (it *AnchorChainOncallContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnchorChainOncallContract)
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
		it.Event = new(AnchorChainOncallContract)
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
func (it *AnchorChainOncallContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnchorChainOncallContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnchorChainOncallContract represents a OncallContract event raised by the AnchorChain contract.
type AnchorChainOncallContract struct {
	Result string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOncallContract is a free log retrieval operation binding the contract event 0x24cca36728e0487747cbf830bd8ad5c42a077fa2c2299a6c6da3ffa5224ae45a.
//
// Solidity: event oncallContract(string result)
func (_AnchorChain *AnchorChainFilterer) FilterOncallContract(opts *bind.FilterOpts) (*AnchorChainOncallContractIterator, error) {

	logs, sub, err := _AnchorChain.contract.FilterLogs(opts, "oncallContract")
	if err != nil {
		return nil, err
	}
	return &AnchorChainOncallContractIterator{contract: _AnchorChain.contract, event: "oncallContract", logs: logs, sub: sub}, nil
}

// WatchOncallContract is a free log subscription operation binding the contract event 0x24cca36728e0487747cbf830bd8ad5c42a077fa2c2299a6c6da3ffa5224ae45a.
//
// Solidity: event oncallContract(string result)
func (_AnchorChain *AnchorChainFilterer) WatchOncallContract(opts *bind.WatchOpts, sink chan<- *AnchorChainOncallContract) (event.Subscription, error) {

	logs, sub, err := _AnchorChain.contract.WatchLogs(opts, "oncallContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnchorChainOncallContract)
				if err := _AnchorChain.contract.UnpackLog(event, "oncallContract", log); err != nil {
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

// ParseOncallContract is a log parse operation binding the contract event 0x24cca36728e0487747cbf830bd8ad5c42a077fa2c2299a6c6da3ffa5224ae45a.
//
// Solidity: event oncallContract(string result)
func (_AnchorChain *AnchorChainFilterer) ParseOncallContract(log types.Log) (*AnchorChainOncallContract, error) {
	event := new(AnchorChainOncallContract)
	if err := _AnchorChain.contract.UnpackLog(event, "oncallContract", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"request\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"onHandleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[5]\",\"name\":\"result\",\"type\":\"string[5]\"}],\"name\":\"onRegisterCallbackResult\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"bytesToUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"getInterchainRequests\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_method\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_args\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_callbackPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_callbackMethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_crossChainType\",\"type\":\"string\"}],\"name\":\"interchainInvoke\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_method\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_args\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"_callbackPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_callbackMethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_crossChainType\",\"type\":\"string\"}],\"name\":\"interchainQuery\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_seq\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_errorCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_errorMsg\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_result\",\"type\":\"string[]\"}],\"name\":\"registerCallbackResult\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uid\",\"type\":\"string\"}],\"name\":\"selectCallbackResult\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"updateCurrentRequestIndex\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeFuncSigs maps the 4-byte function signature to its string representation.
var BridgeFuncSigs = map[string]string{
	"02d06d05": "bytesToUint(bytes)",
	"6ac4691f": "getIncrement()",
	"d27e8515": "getInterchainRequests(bytes)",
	"0d8e6e2c": "getVersion()",
	"737830f4": "interchainInvoke(string,string,string[],string,string,string)",
	"723b9eb3": "interchainQuery(string,string,string[],string,string,string)",
	"88a67acf": "registerCallbackResult(string,string,string,string,string,string[])",
	"dfd3523d": "selectCallbackResult(string)",
	"4de9f17b": "updateCurrentRequestIndex(bytes)",
}

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x608060405260008055600060015534801561001957600080fd5b50611554806100296000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063723b9eb311610066578063723b9eb3146100f3578063737830f41461010657806388a67acf14610119578063d27e85151461012c578063dfd3523d1461013f57610093565b806302d06d05146100985780630d8e6e2c146100c15780634de9f17b146100d65780636ac4691f146100eb575b600080fd5b6100ab6100a6366004610e72565b61015f565b6040516100b89190611402565b60405180910390f35b6100c96101a5565b6040516100b891906113ac565b6100e96100e4366004610e72565b6101c6565b005b6100ab6101e7565b6100c9610101366004610eae565b6101ed565b6100c9610114366004610eae565b610222565b6100e9610127366004610fbe565b61024c565b6100c961013a366004610e72565b6102f3565b61015261014d366004610e72565b610460565b6040516100b8919061139b565b600080805b835181101561019c578060010184510360080260020a84828151811061018657fe5b016020015160f81c029190910190600101610164565b5090505b919050565b60408051808201909152600681526576302e302e3160d01b60208201525b90565b60006101d18261015f565b90508060015410156101e35760018190555b5050565b60005490565b6060610217604051806040016040528060018152602001600360fc1b815250888888888888610556565b979650505050505050565b6060610217604051806040016040528060018152602001603160f81b815250888888888888610556565b610254610c35565b6040518060a0016040528087815260200186815260200185815260200184815260200161028084610714565b815250905080600388604051610296919061131b565b9081526040519081900360200190206102b0916005610c5c565b507f8d94f3a7b0fc3693d4d57b6e40fa24a0120e823bb942561db780bb52aadf4f0b87826040516102e29291906113bd565b60405180910390a150505050505050565b606060006103008361015f565b905060005460015414156103305750506040805180820190915260048152631b9d5b1b60e21b60208201526101a0565b600060015460005403821061034b576001546000540361034d565b815b905060608160405190808252806020026020018201604052801561038557816020015b60608152602001906001900390816103705790505b50905060005b8281101561044d576001805482018101600090815260026020818152604092839020805484519581161561010002600019011692909204601f81018290048202850182019093528284529091908301828280156104295780601f106103fe57610100808354040283529160200191610429565b820191906000526020600020905b81548152906001019060200180831161040c57829003601f168201915b505050505082828151811061043a57fe5b602090810291909101015260010161038b565b5061045781610714565b95945050505050565b6060600382604051610472919061131b565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101561054b5760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156105375780601f1061050c57610100808354040283529160200191610537565b820191906000526020600020905b81548152906001019060200180831161051a57829003601f168201915b5050505050815260200190600101906104a0565b505050509050919050565b600080546001019081905560609061056d90610894565b604080516009808252610140820190925291925060609190816020015b606081526020019060019003908161058a57905050905081816000815181106105af57fe5b602002602001018190525088816001815181106105c857fe5b602002602001018190525087816002815181106105e157fe5b602002602001018190525086816003815181106105fa57fe5b602002602001018190525061060e86610714565b8160048151811061061b57fe5b6020026020010181905250848160058151811061063457fe5b6020026020010181905250838160068151811061064d57fe5b6020026020010181905250610661326108ff565b8160078151811061066e57fe5b6020026020010181905250828160088151811061068757fe5b602002602001018190525061069b81610714565b60008054815260026020908152604090912082516106bf9391929190910190610cb9565b50600080548152600260205260409081902090517f16b1def4e124306bba7341ae8f88d81981f84c000f745603b0148353cbbe62e6916107009133906113e2565b60405180910390a150979650505050505050565b80516060908061073e5750506040805180820190915260028152615b5d60f01b60208201526101a0565b6040805180820190915260018152605b60f81b6020820152915060005b600182038110156107f257826040516020016107779190611356565b6040516020818303038152906040529250826107a585838151811061079857fe5b60200260200101516109fe565b6040516020016107b6929190611327565b6040516020818303038152906040529250826040516020016107d8919061133f565b60408051601f19818403018152919052925060010161075b565b50816040516020016108049190611356565b60405160208183030381529060405291508161082884600184038151811061079857fe5b604051602001610839929190611327565b60405160208183030381529060405291508160405160200161085b9190611356565b60405160208183030381529060405291508160405160200161087d919061136d565b604051602081830303815290604052915050919050565b60606000826108bc5750506040805180820190915260018152600360fc1b60208201526101a0565b82156108ef5761010081049050600a8306603001600160f81b0260001b81179050600a83816108e757fe5b0492506108bc565b6108f881610b8f565b9392505050565b604080516028808252606082810190935282919060208201818038833901905050905060005b60148110156109ec5760008160130360080260020a856001600160a01b03168161094b57fe5b0460f81b9050600060108260f81c60ff168161096357fe5b0460f81b905060008160f81c6010028360f81c0360f81b905061098582610c04565b85856002028151811061099457fe5b60200101906001600160f81b031916908160001a9053506109b481610c04565b8585600202600101815181106109c657fe5b60200101906001600160f81b031916908160001a90535050600190920191506109259050565b508060405160200161087d9190611384565b80516040805160028302808252601f19601f82011682016020019092526060928492909184918015610a37576020820181803883390190505b5090506000805b83811015610b0d57848181518110610a5257fe5b6020910101516001600160f81b031916601760fa1b1480610a915750848181518110610a7a57fe5b6020910101516001600160f81b031916601160f91b145b15610ac557601760fa1b838380600101945081518110610aad57fe5b60200101906001600160f81b031916908160001a9053505b848181518110610ad157fe5b602001015160f81c60f81b838380600101945081518110610aee57fe5b60200101906001600160f81b031916908160001a905350600101610a3e565b6060826040519080825280601f01601f191660200182016040528015610b3a576020820181803883390190505b509050600091505b8282101561021757838281518110610b5657fe5b602001015160f81c60f81b818381518110610b6d57fe5b60200101906001600160f81b031916908160001a905350600190910190610b42565b604080516020808252818301909252606091829190602082018180388339019050509050602060005b81811015610bfb57848160208110610bcc57fe5b1a60f81b838281518110610bdc57fe5b60200101906001600160f81b031916908160001a905350600101610bb8565b50909392505050565b6000600a60f883901c1015610c24578160f81c60300160f81b90506101a0565b8160f81c60570160f81b90506101a0565b6040518060a001604052806005905b6060815260200190600190039081610c445790505090565b828054828255906000526020600020908101928215610ca9579160200282015b82811115610ca95782518051610c99918491602090910190610cb9565b5091602001919060010190610c7c565b50610cb5929150610d33565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610cfa57805160ff1916838001178555610d27565b82800160010185558215610d27579182015b82811115610d27578251825591602001919060010190610d0c565b50610cb5929150610d56565b6101c391905b80821115610cb5576000610d4d8282610d70565b50600101610d39565b6101c391905b80821115610cb55760008155600101610d5c565b50805460018160011615610100020316600290046000825580601f10610d965750610db4565b601f016020900490600052602060002090810190610db49190610d56565b50565b600082601f830112610dc857600080fd5b8135610ddb610dd68261143c565b611416565b81815260209384019390925082018360005b83811015610e195781358601610e038882610e23565b8452506020928301929190910190600101610ded565b5050505092915050565b600082601f830112610e3457600080fd5b8135610e42610dd68261145c565b91508082526020830160208301858383011115610e5e57600080fd5b610e698382846114cb565b50505092915050565b600060208284031215610e8457600080fd5b81356001600160401b03811115610e9a57600080fd5b610ea684828501610e23565b949350505050565b60008060008060008060c08789031215610ec757600080fd5b86356001600160401b03811115610edd57600080fd5b610ee989828a01610e23565b96505060208701356001600160401b03811115610f0557600080fd5b610f1189828a01610e23565b95505060408701356001600160401b03811115610f2d57600080fd5b610f3989828a01610db7565b94505060608701356001600160401b03811115610f5557600080fd5b610f6189828a01610e23565b93505060808701356001600160401b03811115610f7d57600080fd5b610f8989828a01610e23565b92505060a08701356001600160401b03811115610fa557600080fd5b610fb189828a01610e23565b9150509295509295509295565b60008060008060008060c08789031215610fd757600080fd5b86356001600160401b03811115610fed57600080fd5b610ff989828a01610e23565b96505060208701356001600160401b0381111561101557600080fd5b61102189828a01610e23565b95505060408701356001600160401b0381111561103d57600080fd5b61104989828a01610e23565b94505060608701356001600160401b0381111561106557600080fd5b61107189828a01610e23565b93505060808701356001600160401b0381111561108d57600080fd5b61109989828a01610e23565b92505060a08701356001600160401b038111156110b557600080fd5b610fb189828a01610db7565b60006108f883836111ab565b6110d6816114b4565b82525050565b60006110e782611495565b6110f181856101a0565b935083602082028501611103856101c3565b8060005b8581101561113d578484038952815161112085826110c1565b945061112b83611483565b60209a909a0199925050600101611107565b5091979650505050505050565b60006111558261149b565b61115f818561149f565b93508360208202850161117185611483565b8060005b8581101561113d578484038952815161118e85826110c1565b945061119983611483565b60209a909a0199925050600101611175565b60006111b68261149b565b6111c0818561149f565b93506111d08185602086016114d7565b6111d981611507565b9093019392505050565b60006111ee8261149b565b6111f881856101a0565b93506112088185602086016114d7565b9290920192915050565b60008154600181166000811461122f576001811461125557611294565b607f6002830416611240818761149f565b60ff1984168152955050602085019250611294565b60028204611263818761149f565b955061126e85611489565b60005b8281101561128d57815488820152600190910190602001611271565b8701945050505b505092915050565b60006112a96002836101a0565b61088b60f21b815260020192915050565b60006112c76002836101a0565b61060f60f31b815260020192915050565b60006112e56001836101a0565b601160f91b815260010192915050565b60006113026001836101a0565b605d60f81b815260010192915050565b6110d6816101c3565b60006108f882846111e3565b600061133382856111e3565b9150610ea682846111e3565b600061134b82846111e3565b91506108f88261129c565b600061136282846111e3565b91506108f8826112d8565b600061137982846111e3565b91506108f8826112f5565b600061138f826112ba565b91506108f882846111e3565b602080825281016108f8818461114a565b602080825281016108f881846111ab565b604080825281016113ce81856111ab565b90508181036020830152610ea681846110dc565b604080825281016113f38185611212565b90506108f860208301846110cd565b602081016114108284611312565b92915050565b6040518181016001600160401b038111828210171561143457600080fd5b604052919050565b60006001600160401b0382111561145257600080fd5b5060209081020190565b60006001600160401b0382111561147257600080fd5b506020601f91909101601f19160190565b60200190565b60009081526020902090565b50600590565b5190565b90815260200190565b6001600160a01b031690565b6000611410826000611410826000611410826114a8565b82818337506000910152565b60005b838110156114f25781810151838201526020016114da565b83811115611501576000848401525b50505050565b601f01601f19169056fea365627a7a72315820645c3927d24da871fd5fec36e6a5ba06f6c3beebf35d050420bad31dffe4d1b16c6578706572696d656e74616cf564736f6c63430005110040"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// BytesToUint is a free data retrieval call binding the contract method 0x02d06d05.
//
// Solidity: function bytesToUint(bytes b) constant returns(uint256)
func (_Bridge *BridgeCaller) BytesToUint(opts *bind.CallOpts, b []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "bytesToUint", b)
	return *ret0, err
}

// BytesToUint is a free data retrieval call binding the contract method 0x02d06d05.
//
// Solidity: function bytesToUint(bytes b) constant returns(uint256)
func (_Bridge *BridgeSession) BytesToUint(b []byte) (*big.Int, error) {
	return _Bridge.Contract.BytesToUint(&_Bridge.CallOpts, b)
}

// BytesToUint is a free data retrieval call binding the contract method 0x02d06d05.
//
// Solidity: function bytesToUint(bytes b) constant returns(uint256)
func (_Bridge *BridgeCallerSession) BytesToUint(b []byte) (*big.Int, error) {
	return _Bridge.Contract.BytesToUint(&_Bridge.CallOpts, b)
}

// GetIncrement is a free data retrieval call binding the contract method 0x6ac4691f.
//
// Solidity: function getIncrement() constant returns(uint256)
func (_Bridge *BridgeCaller) GetIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "getIncrement")
	return *ret0, err
}

// GetIncrement is a free data retrieval call binding the contract method 0x6ac4691f.
//
// Solidity: function getIncrement() constant returns(uint256)
func (_Bridge *BridgeSession) GetIncrement() (*big.Int, error) {
	return _Bridge.Contract.GetIncrement(&_Bridge.CallOpts)
}

// GetIncrement is a free data retrieval call binding the contract method 0x6ac4691f.
//
// Solidity: function getIncrement() constant returns(uint256)
func (_Bridge *BridgeCallerSession) GetIncrement() (*big.Int, error) {
	return _Bridge.Contract.GetIncrement(&_Bridge.CallOpts)
}

// GetInterchainRequests is a free data retrieval call binding the contract method 0xd27e8515.
//
// Solidity: function getInterchainRequests(bytes b) constant returns(string)
func (_Bridge *BridgeCaller) GetInterchainRequests(opts *bind.CallOpts, b []byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "getInterchainRequests", b)
	return *ret0, err
}

// GetInterchainRequests is a free data retrieval call binding the contract method 0xd27e8515.
//
// Solidity: function getInterchainRequests(bytes b) constant returns(string)
func (_Bridge *BridgeSession) GetInterchainRequests(b []byte) (string, error) {
	return _Bridge.Contract.GetInterchainRequests(&_Bridge.CallOpts, b)
}

// GetInterchainRequests is a free data retrieval call binding the contract method 0xd27e8515.
//
// Solidity: function getInterchainRequests(bytes b) constant returns(string)
func (_Bridge *BridgeCallerSession) GetInterchainRequests(b []byte) (string, error) {
	return _Bridge.Contract.GetInterchainRequests(&_Bridge.CallOpts, b)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() constant returns(string)
func (_Bridge *BridgeCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "getVersion")
	return *ret0, err
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() constant returns(string)
func (_Bridge *BridgeSession) GetVersion() (string, error) {
	return _Bridge.Contract.GetVersion(&_Bridge.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() constant returns(string)
func (_Bridge *BridgeCallerSession) GetVersion() (string, error) {
	return _Bridge.Contract.GetVersion(&_Bridge.CallOpts)
}

// SelectCallbackResult is a free data retrieval call binding the contract method 0xdfd3523d.
//
// Solidity: function selectCallbackResult(string _uid) constant returns(string[])
func (_Bridge *BridgeCaller) SelectCallbackResult(opts *bind.CallOpts, _uid string) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "selectCallbackResult", _uid)
	return *ret0, err
}

// SelectCallbackResult is a free data retrieval call binding the contract method 0xdfd3523d.
//
// Solidity: function selectCallbackResult(string _uid) constant returns(string[])
func (_Bridge *BridgeSession) SelectCallbackResult(_uid string) ([]string, error) {
	return _Bridge.Contract.SelectCallbackResult(&_Bridge.CallOpts, _uid)
}

// SelectCallbackResult is a free data retrieval call binding the contract method 0xdfd3523d.
//
// Solidity: function selectCallbackResult(string _uid) constant returns(string[])
func (_Bridge *BridgeCallerSession) SelectCallbackResult(_uid string) ([]string, error) {
	return _Bridge.Contract.SelectCallbackResult(&_Bridge.CallOpts, _uid)
}

// InterchainInvoke is a paid mutator transaction binding the contract method 0x737830f4.
//
// Solidity: function interchainInvoke(string _path, string _method, string[] _args, string _callbackPath, string _callbackMethod, string _crossChainType) returns(string uid)
func (_Bridge *BridgeTransactor) InterchainInvoke(opts *bind.TransactOpts, _path string, _method string, _args []string, _callbackPath string, _callbackMethod string, _crossChainType string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "interchainInvoke", _path, _method, _args, _callbackPath, _callbackMethod, _crossChainType)
}

// InterchainInvoke is a paid mutator transaction binding the contract method 0x737830f4.
//
// Solidity: function interchainInvoke(string _path, string _method, string[] _args, string _callbackPath, string _callbackMethod, string _crossChainType) returns(string uid)
func (_Bridge *BridgeSession) InterchainInvoke(_path string, _method string, _args []string, _callbackPath string, _callbackMethod string, _crossChainType string) (*types.Transaction, error) {
	return _Bridge.Contract.InterchainInvoke(&_Bridge.TransactOpts, _path, _method, _args, _callbackPath, _callbackMethod, _crossChainType)
}

// InterchainInvoke is a paid mutator transaction binding the contract method 0x737830f4.
//
// Solidity: function interchainInvoke(string _path, string _method, string[] _args, string _callbackPath, string _callbackMethod, string _crossChainType) returns(string uid)
func (_Bridge *BridgeTransactorSession) InterchainInvoke(_path string, _method string, _args []string, _callbackPath string, _callbackMethod string, _crossChainType string) (*types.Transaction, error) {
	return _Bridge.Contract.InterchainInvoke(&_Bridge.TransactOpts, _path, _method, _args, _callbackPath, _callbackMethod, _crossChainType)
}

// InterchainQuery is a paid mutator transaction binding the contract method 0x723b9eb3.
//
// Solidity: function interchainQuery(string _path, string _method, string[] _args, string _callbackPath, string _callbackMethod, string _crossChainType) returns(string uid)
func (_Bridge *BridgeTransactor) InterchainQuery(opts *bind.TransactOpts, _path string, _method string, _args []string, _callbackPath string, _callbackMethod string, _crossChainType string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "interchainQuery", _path, _method, _args, _callbackPath, _callbackMethod, _crossChainType)
}

// InterchainQuery is a paid mutator transaction binding the contract method 0x723b9eb3.
//
// Solidity: function interchainQuery(string _path, string _method, string[] _args, string _callbackPath, string _callbackMethod, string _crossChainType) returns(string uid)
func (_Bridge *BridgeSession) InterchainQuery(_path string, _method string, _args []string, _callbackPath string, _callbackMethod string, _crossChainType string) (*types.Transaction, error) {
	return _Bridge.Contract.InterchainQuery(&_Bridge.TransactOpts, _path, _method, _args, _callbackPath, _callbackMethod, _crossChainType)
}

// InterchainQuery is a paid mutator transaction binding the contract method 0x723b9eb3.
//
// Solidity: function interchainQuery(string _path, string _method, string[] _args, string _callbackPath, string _callbackMethod, string _crossChainType) returns(string uid)
func (_Bridge *BridgeTransactorSession) InterchainQuery(_path string, _method string, _args []string, _callbackPath string, _callbackMethod string, _crossChainType string) (*types.Transaction, error) {
	return _Bridge.Contract.InterchainQuery(&_Bridge.TransactOpts, _path, _method, _args, _callbackPath, _callbackMethod, _crossChainType)
}

// RegisterCallbackResult is a paid mutator transaction binding the contract method 0x88a67acf.
//
// Solidity: function registerCallbackResult(string _uid, string _tid, string _seq, string _errorCode, string _errorMsg, string[] _result) returns()
func (_Bridge *BridgeTransactor) RegisterCallbackResult(opts *bind.TransactOpts, _uid string, _tid string, _seq string, _errorCode string, _errorMsg string, _result []string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "registerCallbackResult", _uid, _tid, _seq, _errorCode, _errorMsg, _result)
}

// RegisterCallbackResult is a paid mutator transaction binding the contract method 0x88a67acf.
//
// Solidity: function registerCallbackResult(string _uid, string _tid, string _seq, string _errorCode, string _errorMsg, string[] _result) returns()
func (_Bridge *BridgeSession) RegisterCallbackResult(_uid string, _tid string, _seq string, _errorCode string, _errorMsg string, _result []string) (*types.Transaction, error) {
	return _Bridge.Contract.RegisterCallbackResult(&_Bridge.TransactOpts, _uid, _tid, _seq, _errorCode, _errorMsg, _result)
}

// RegisterCallbackResult is a paid mutator transaction binding the contract method 0x88a67acf.
//
// Solidity: function registerCallbackResult(string _uid, string _tid, string _seq, string _errorCode, string _errorMsg, string[] _result) returns()
func (_Bridge *BridgeTransactorSession) RegisterCallbackResult(_uid string, _tid string, _seq string, _errorCode string, _errorMsg string, _result []string) (*types.Transaction, error) {
	return _Bridge.Contract.RegisterCallbackResult(&_Bridge.TransactOpts, _uid, _tid, _seq, _errorCode, _errorMsg, _result)
}

// UpdateCurrentRequestIndex is a paid mutator transaction binding the contract method 0x4de9f17b.
//
// Solidity: function updateCurrentRequestIndex(bytes b) returns()
func (_Bridge *BridgeTransactor) UpdateCurrentRequestIndex(opts *bind.TransactOpts, b []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateCurrentRequestIndex", b)
}

// UpdateCurrentRequestIndex is a paid mutator transaction binding the contract method 0x4de9f17b.
//
// Solidity: function updateCurrentRequestIndex(bytes b) returns()
func (_Bridge *BridgeSession) UpdateCurrentRequestIndex(b []byte) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateCurrentRequestIndex(&_Bridge.TransactOpts, b)
}

// UpdateCurrentRequestIndex is a paid mutator transaction binding the contract method 0x4de9f17b.
//
// Solidity: function updateCurrentRequestIndex(bytes b) returns()
func (_Bridge *BridgeTransactorSession) UpdateCurrentRequestIndex(b []byte) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateCurrentRequestIndex(&_Bridge.TransactOpts, b)
}

// BridgeOnHandleRequestIterator is returned from FilterOnHandleRequest and is used to iterate over the raw logs and unpacked data for OnHandleRequest events raised by the Bridge contract.
type BridgeOnHandleRequestIterator struct {
	Event *BridgeOnHandleRequest // Event containing the contract specifics and raw log

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
func (it *BridgeOnHandleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOnHandleRequest)
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
		it.Event = new(BridgeOnHandleRequest)
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
func (it *BridgeOnHandleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOnHandleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOnHandleRequest represents a OnHandleRequest event raised by the Bridge contract.
type BridgeOnHandleRequest struct {
	Request string
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOnHandleRequest is a free log retrieval operation binding the contract event 0x16b1def4e124306bba7341ae8f88d81981f84c000f745603b0148353cbbe62e6.
//
// Solidity: event onHandleRequest(string request, address sender)
func (_Bridge *BridgeFilterer) FilterOnHandleRequest(opts *bind.FilterOpts) (*BridgeOnHandleRequestIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "onHandleRequest")
	if err != nil {
		return nil, err
	}
	return &BridgeOnHandleRequestIterator{contract: _Bridge.contract, event: "onHandleRequest", logs: logs, sub: sub}, nil
}

// WatchOnHandleRequest is a free log subscription operation binding the contract event 0x16b1def4e124306bba7341ae8f88d81981f84c000f745603b0148353cbbe62e6.
//
// Solidity: event onHandleRequest(string request, address sender)
func (_Bridge *BridgeFilterer) WatchOnHandleRequest(opts *bind.WatchOpts, sink chan<- *BridgeOnHandleRequest) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "onHandleRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOnHandleRequest)
				if err := _Bridge.contract.UnpackLog(event, "onHandleRequest", log); err != nil {
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

// ParseOnHandleRequest is a log parse operation binding the contract event 0x16b1def4e124306bba7341ae8f88d81981f84c000f745603b0148353cbbe62e6.
//
// Solidity: event onHandleRequest(string request, address sender)
func (_Bridge *BridgeFilterer) ParseOnHandleRequest(log types.Log) (*BridgeOnHandleRequest, error) {
	event := new(BridgeOnHandleRequest)
	if err := _Bridge.contract.UnpackLog(event, "onHandleRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeOnRegisterCallbackResultIterator is returned from FilterOnRegisterCallbackResult and is used to iterate over the raw logs and unpacked data for OnRegisterCallbackResult events raised by the Bridge contract.
type BridgeOnRegisterCallbackResultIterator struct {
	Event *BridgeOnRegisterCallbackResult // Event containing the contract specifics and raw log

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
func (it *BridgeOnRegisterCallbackResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOnRegisterCallbackResult)
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
		it.Event = new(BridgeOnRegisterCallbackResult)
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
func (it *BridgeOnRegisterCallbackResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOnRegisterCallbackResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOnRegisterCallbackResult represents a OnRegisterCallbackResult event raised by the Bridge contract.
type BridgeOnRegisterCallbackResult struct {
	Uid    string
	Result [5]string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOnRegisterCallbackResult is a free log retrieval operation binding the contract event 0x8d94f3a7b0fc3693d4d57b6e40fa24a0120e823bb942561db780bb52aadf4f0b.
//
// Solidity: event onRegisterCallbackResult(string uid, string[5] result)
func (_Bridge *BridgeFilterer) FilterOnRegisterCallbackResult(opts *bind.FilterOpts) (*BridgeOnRegisterCallbackResultIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "onRegisterCallbackResult")
	if err != nil {
		return nil, err
	}
	return &BridgeOnRegisterCallbackResultIterator{contract: _Bridge.contract, event: "onRegisterCallbackResult", logs: logs, sub: sub}, nil
}

// WatchOnRegisterCallbackResult is a free log subscription operation binding the contract event 0x8d94f3a7b0fc3693d4d57b6e40fa24a0120e823bb942561db780bb52aadf4f0b.
//
// Solidity: event onRegisterCallbackResult(string uid, string[5] result)
func (_Bridge *BridgeFilterer) WatchOnRegisterCallbackResult(opts *bind.WatchOpts, sink chan<- *BridgeOnRegisterCallbackResult) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "onRegisterCallbackResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOnRegisterCallbackResult)
				if err := _Bridge.contract.UnpackLog(event, "onRegisterCallbackResult", log); err != nil {
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

// ParseOnRegisterCallbackResult is a log parse operation binding the contract event 0x8d94f3a7b0fc3693d4d57b6e40fa24a0120e823bb942561db780bb52aadf4f0b.
//
// Solidity: event onRegisterCallbackResult(string uid, string[5] result)
func (_Bridge *BridgeFilterer) ParseOnRegisterCallbackResult(log types.Log) (*BridgeOnRegisterCallbackResult, error) {
	event := new(BridgeOnRegisterCallbackResult)
	if err := _Bridge.contract.UnpackLog(event, "onRegisterCallbackResult", log); err != nil {
		return nil, err
	}
	return event, nil
}
