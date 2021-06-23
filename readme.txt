judicial.sim.evidence 存证
judicial.sim.proxy  代理
judicial.sim.bridge 桥接

http://192.168.4.248:8545
proxy address:0xf8e81D47203A594245E36C48e151709F0C19fBe8
bridge address:0xd9145CCE52D386f254917e481eB44e9943F39138
contract address:0xd8b934580fcE35a11B58C6D73aDeE468a2833fa8

anchor:
[
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "address",
				"name": "user",
				"type": "address"
			}
		],
		"name": "addUser",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "bool",
				"name": "state",
				"type": "bool"
			},
			{
				"internalType": "string[]",
				"name": "_result",
				"type": "string[]"
			}
		],
		"name": "callback",
		"outputs": [
			{
				"internalType": "string[]",
				"name": "",
				"type": "string[]"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "string[]",
				"name": "_data",
				"type": "string[]"
			}
		],
		"name": "connect",
		"outputs": [
			{
				"internalType": "string[]",
				"name": "",
				"type": "string[]"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "string",
				"name": "_path",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_method",
				"type": "string"
			},
			{
				"internalType": "string[]",
				"name": "_args",
				"type": "string[]"
			},
			{
				"internalType": "string",
				"name": "_callbackPath",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_callbackMethod",
				"type": "string"
			}
		],
		"name": "connectChain",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "address",
				"name": "hub",
				"type": "address"
			}
		],
		"name": "init",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "address",
				"name": "user",
				"type": "address"
			}
		],
		"name": "removeUser",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "scanBlocks",
		"outputs": [
			{
				"internalType": "string[]",
				"name": "",
				"type": "string[]"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"internalType": "uint256",
				"name": "i",
				"type": "uint256"
			}
		],
		"name": "viewBlock",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]



bridge :

[
	{
		"constant": true,
		"inputs": [],
		"name": "getIncrement",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_num",
				"type": "uint256"
			}
		],
		"name": "getInterchainRequests",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "getVersion",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "string",
				"name": "_path",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_method",
				"type": "string"
			},
			{
				"internalType": "string[]",
				"name": "_args",
				"type": "string[]"
			},
			{
				"internalType": "string",
				"name": "_callbackPath",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_callbackMethod",
				"type": "string"
			}
		],
		"name": "interchainInvoke",
		"outputs": [
			{
				"internalType": "string",
				"name": "uid",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "string",
				"name": "_path",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_method",
				"type": "string"
			},
			{
				"internalType": "string[]",
				"name": "_args",
				"type": "string[]"
			},
			{
				"internalType": "string",
				"name": "_callbackPath",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_callbackMethod",
				"type": "string"
			}
		],
		"name": "interchainQuery",
		"outputs": [
			{
				"internalType": "string",
				"name": "uid",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "string",
				"name": "_uid",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_tid",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_seq",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_errorCode",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_errorMsg",
				"type": "string"
			},
			{
				"internalType": "string[]",
				"name": "_result",
				"type": "string[]"
			}
		],
		"name": "registerCallbackResult",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"internalType": "string",
				"name": "_uid",
				"type": "string"
			}
		],
		"name": "selectCallbackResult",
		"outputs": [
			{
				"internalType": "string[]",
				"name": "",
				"type": "string[]"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_index",
				"type": "uint256"
			}
		],
		"name": "updateCurrentRequestIndex",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]
