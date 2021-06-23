package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"testing"
	"time"
)

func TestAnchorChainTransactor_ConnectChain_Get(t *testing.T) {
	//anchor
	address := common.HexToAddress("0x718Cea95F036B365Fd7fF07e800E51e29F79199f")
	ethclient, err := ethclient.Dial("http://192.168.3.92:8545")
	if err != nil {
		t.Error(err)
	}
	keystorePath := "/Users/daminyang/service/geth/data1/keystore/UTC--2021-05-25T07-28-12.168815000Z--46e6fc87ef7b0c12b96b6098a70701bfc8edb28d"
	if err != nil {
		t.Error(err)
	}
	keyJson, err := ioutil.ReadFile(keystorePath)
	key, err := keystore.DecryptKey(keyJson, "123456")
	if err != nil {
		t.Error(err)
	}

	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Error("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		t.Error(err)
	}

	gasPrice, err := ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err)
	}

	auth := bind.NewKeyedTransactor(key.PrivateKey)
	if err != nil {
		t.Error(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	anchorContract, err := NewAnchorChain(address, ethclient)
	if err != nil {
		t.Error(err)
	}

	//opts *bind.TransactOpts, _path string, _method string, _args []string, _callbackPath string, _callbackMethod string
	//to
	path := "judicial.geth.evidence"
	method := "connect"
	args := []string{"hello1", "world1", "message1"}
	callbackPath := "judicial.sim.evidence"
	callbackMethod := "callback"
	crossChainType:="inner"

	transaction, err := anchorContract.ConnectChain(auth, path, method, args, callbackPath, callbackMethod,crossChainType)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("ConnectChain：" + transaction.Hash().String())
}

func TestAnchorChainFilterer_WatchOncallContract(t *testing.T) {

	address := common.HexToAddress("0xDD361e416b7C372285DAe7c125EE5A92a992Ce9F")

	ethclient, err := ethclient.Dial("http://192.168.3.92:8545")
	if err != nil {
		t.Error(err)
	}
	keystorePath := "/Users/daminyang/service/geth/data1/keystore/UTC--2021-05-25T07-28-12.168815000Z--46e6fc87ef7b0c12b96b6098a70701bfc8edb28d"
	keyJson, err := ioutil.ReadFile(keystorePath)
	key, err := keystore.DecryptKey(keyJson, "123456")
	if err != nil {
		t.Error(err)
	}

	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Error("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		t.Error(err)
	}

	gasPrice, err := ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err)
	}

	auth := bind.NewKeyedTransactor(key.PrivateKey)
	if err != nil {
		t.Error(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	anchorContract, err := NewAnchorChain(address, ethclient)
	tx, err := anchorContract.CallContract(auth, "hello,mygod,123456,789")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("tx:" + tx.Hash().String())
}

func SetNonceAndGasPrice(ethclient *ethclient.Client, auth *bind.TransactOpts, fromAddress common.Address) {
	nonce, err := ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	gasPrice, err := ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
}
func TestDeployAnchorChain_Geth_Init(t *testing.T) {
	ethclient, err := ethclient.Dial("http://192.168.3.92:8545")
	if err != nil {
		t.Error(err)
	}
	keystorePath := "/Users/daminyang/service/geth/data1/keystore/UTC--2021-05-25T07-28-12.168815000Z--46e6fc87ef7b0c12b96b6098a70701bfc8edb28d"
	keyJson, err := ioutil.ReadFile(keystorePath)
	if err != nil {
		t.Error(err)
	}

	key, err := keystore.DecryptKey(keyJson, "123456")
	if err != nil {
		t.Error(err)
	}

	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Error("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		t.Error(err)
	}

	gasPrice, err := ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err)
	}

	auth := bind.NewKeyedTransactor(key.PrivateKey)
	if err != nil {
		t.Error(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address, tx, _, err := DeployAnchorChain(auth, ethclient)
	if err != nil {
		t.Error(err)
	}

	t.Log("geth anchorchain addr:", address.Hex())

	t.Log("geth anchorchain transaction" + tx.Hash().Hex())

	anchorAddr := address

	time.Sleep(time.Second * 1)

	SetNonceAndGasPrice(ethclient, auth, fromAddress)

	address, tx, _, err = DeployBridge(auth, ethclient)
	if err != nil {
		t.Error(err)
	}
	bridgeAddr := address

	t.Log("geth bridge addr:", address.Hex())

	t.Log("geth bridge transaction" + tx.Hash().Hex())

	time.Sleep(time.Second * 1)

	SetNonceAndGasPrice(ethclient, auth, fromAddress)

	address, tx, _, err = DeployProxy(auth, ethclient)
	if err != nil {
		t.Error(err)
	}
	proxyAddr := address

	t.Log("geth proxy addr:", address.Hex())

	t.Log("geth proxy transaction" + tx.Hash().Hex())

	time.Sleep(time.Second * 1)

	//合约部署完毕

	SetNonceAndGasPrice(ethclient, auth, fromAddress)
	//初始化合约
	anchorContract, err := NewAnchorChain(anchorAddr, ethclient)
	if err != nil {
		t.Error(err)
	}
	transaction, err := anchorContract.Init(auth, bridgeAddr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("anchor init method:", transaction.Hash().String())

	time.Sleep(time.Second * 1)
	SetNonceAndGasPrice(ethclient, auth, fromAddress)
	//添加资源
	proxy, err := NewProxy(proxyAddr, ethclient)
	if err != nil {
		t.Error(err)
	}
	resourceAnchorPath := "judicial.geth.evidence"
	transaction, err = proxy.AddPath(auth, resourceAnchorPath, anchorAddr, AnchorChainABI)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("proxy AddPath Anchor:", transaction.Hash().String())

	time.Sleep(time.Second * 1)

	SetNonceAndGasPrice(ethclient, auth, fromAddress)

	resourceBridgePath := "judicial.geth.bridge"

	transaction, err = proxy.AddPath(auth, resourceBridgePath, bridgeAddr, BridgeABI)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("proxy AddPath Anchor:", transaction.Hash().String())

	fmt.Println("Anchor address:", anchorAddr.String())

	fmt.Println("bridge address:", bridgeAddr.String())

	fmt.Println("proxy address:", proxyAddr.String())
}

func TestDeployAnchorChain_Sim_Init(t *testing.T) {

	ethclient, err := ethclient.Dial("http://192.168.4.248:8545")
	if err != nil {
		t.Error(err)
	}
	keystorePath := "/Users/daminyang/service/geth/data/keystore/UTC--2021-05-20T03-37-29.930541000Z--4e136c5649b1ea2b23455dc888dff6927e11ab35"
	keyJson, err := ioutil.ReadFile(keystorePath)
	if err != nil {
		t.Error(err)
	}

	key, err := keystore.DecryptKey(keyJson, "123456")
	if err != nil {
		t.Error(err)
	}

	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Error("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		t.Error(err)
	}

	gasPrice, err := ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err)
	}

	auth := bind.NewKeyedTransactor(key.PrivateKey)
	if err != nil {
		t.Error(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address, tx, _, err := DeployAnchorChain(auth, ethclient)
	if err != nil {
		t.Error(err)
	}

	t.Log("sim anchorchain addr:" + address.Hex())

	t.Log("sim anchorchain transaction:" + tx.Hash().Hex())

	anchorAddr := address

	time.Sleep(time.Second * 1)

	SetNonceAndGasPrice(ethclient, auth, fromAddress)

	address, tx, _, err = DeployBridge(auth, ethclient)
	if err != nil {
		t.Error(err)
	}
	bridgeAddr := address

	t.Log("sim bridge addr:", address.Hex())

	t.Log("sim bridge transaction" + tx.Hash().Hex())

	time.Sleep(time.Second * 1)

	SetNonceAndGasPrice(ethclient, auth, fromAddress)

	address, tx, _, err = DeployProxy(auth, ethclient)
	if err != nil {
		t.Error(err)
	}
	proxyAddr := address

	t.Log("sim proxy addr:", address.Hex())

	t.Log("sim proxy transaction" + tx.Hash().Hex())

	time.Sleep(time.Second * 1)

	//合约部署完毕

	SetNonceAndGasPrice(ethclient, auth, fromAddress)
	//初始化合约
	anchorContract, err := NewAnchorChain(anchorAddr, ethclient)
	if err != nil {
		t.Error(err)
	}
	transaction, err := anchorContract.Init(auth, bridgeAddr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("sim anchor init method:", transaction.Hash().String())

	time.Sleep(time.Second * 1)
	SetNonceAndGasPrice(ethclient, auth, fromAddress)
	//添加资源
	proxy, err := NewProxy(proxyAddr, ethclient)
	if err != nil {
		t.Error(err)
	}
	resourceAnchorPath := "judicial.sim.evidence"
	transaction, err = proxy.AddPath(auth, resourceAnchorPath, anchorAddr, AnchorChainABI)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("sim proxy AddPath Anchor: ", transaction.Hash().String())

	time.Sleep(time.Second * 1)

	SetNonceAndGasPrice(ethclient, auth, fromAddress)

	resourceBridgePath := "judicial.sim.bridge"

	transaction, err = proxy.AddPath(auth, resourceBridgePath, bridgeAddr, BridgeABI)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("sim proxy AddPath Anchor:", transaction.Hash().String())

	fmt.Println("sim Anchor address:", anchorAddr.String())

	fmt.Println("sim bridge address:", bridgeAddr.String())

	fmt.Println("sim proxy address:", proxyAddr.String())
}

func TestAnchorChainTransactor_ConnectChain_sim(t *testing.T) {
	//anchor
	address := common.HexToAddress("")
	ethclient, err := ethclient.Dial("http://192.168.4.248:8545")
	if err != nil {
		t.Error(err)
	}

	keyJson, err := ioutil.ReadFile("/Users/daminyang/service/geth/data/keystore/UTC--2021-05-20T03-37-29.930541000Z--4e136c5649b1ea2b23455dc888dff6927e11ab35")
	if err != nil {
		t.Error(err)
	}

	key, err := keystore.DecryptKey(keyJson, "123456")
	if err != nil {
		t.Error(err)
	}

	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Error("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		t.Error(err)
	}

	gasPrice, err := ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err)
	}

	auth := bind.NewKeyedTransactor(key.PrivateKey)
	if err != nil {
		t.Error(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	anchorContract, err := NewAnchorChain(address, ethclient)
	if err != nil {
		t.Error(err)
	}

	//opts *bind.TransactOpts, _path string, _method string, _args []string, _callbackPath string, _callbackMethod string
	//to
	path := "judicial.geth.evidence"
	method := "connect"
	args := []string{"hello1", "world1", "message1"}
	callbackPath := "judicial.sim.evidence"
	callbackMethod := "callback"
	crossChainType:="inner"

	transaction, err := anchorContract.ConnectChain(auth, path, method, args, callbackPath, callbackMethod,crossChainType)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("ConnectChain：" + transaction.Hash().String())
}
func TestAnchorChainCaller_ViewBlock_sim(t *testing.T) {
	//anchor
	//address := common.HexToAddress("0x3E87B5bB94c0Cde423C98Fb5Aa1105B52Fe1f758")
	//
	//ethclient, err := ethclient.Dial("http://192.168.4.248:8545")
	//if err != nil {
	//	t.Error(err)
	//}

}
func TestAnchorChainCaller_ViewBlock_Geth(t *testing.T) {
	//anchor
	address := common.HexToAddress("0x0dab8fB6B9502Cc8752b571d9E4BA14a16f2686E")

	ethclient, err := ethclient.Dial("http://192.168.3.92:8545")
	if err != nil {
		t.Error(err)
	}
	anchorContract, err := NewAnchorChain(address, ethclient)
	if err != nil {
		t.Error(err)
	}
	result, err := anchorContract.ViewBlock(&bind.CallOpts{}, big.NewInt(0))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("result=" + result)
}
func TestAnchorChainCaller_ScanBlocks_Geth(t *testing.T) {
	//anchor
	address := common.HexToAddress("0x82B48C7891212525AED88785277039A6d1753012")

	ethclient, err := ethclient.Dial("http://192.168.3.92:8545")
	if err != nil {
		t.Error(err)
	}
	anchorContract, err := NewAnchorChain(address, ethclient)
	if err != nil {
		t.Error(err)
	}
	result, err := anchorContract.ScanBlocks(&bind.CallOpts{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
func TestAnchorChainTransactor_Connect_Geth(t *testing.T) {
	address := common.HexToAddress("0x82B48C7891212525AED88785277039A6d1753012")
	ethclient, err := ethclient.Dial("http://192.168.3.92:8545")
	if err != nil {
		t.Error(err)
	}
	keystorePath := "/Users/daminyang/service/geth/data1/keystore/UTC--2021-05-25T07-28-12.168815000Z--46e6fc87ef7b0c12b96b6098a70701bfc8edb28d"
	keyJson, err := ioutil.ReadFile(keystorePath)
	key, err := keystore.DecryptKey(keyJson, "123456")
	if err != nil {
		t.Error(err)
	}

	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Error("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		t.Error(err)
	}

	gasPrice, err := ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err)
	}

	auth := bind.NewKeyedTransactor(key.PrivateKey)
	if err != nil {
		t.Error(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	anchorContract, err := NewAnchorChain(address, ethclient)
	tx, err := anchorContract.Connect(auth, []string{"get", "it", "best"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("tx:" + tx.Hash().String())

}

func TestAnchorChainTransactor_AddUser_Geth(t *testing.T) {
	proxyAddr := common.HexToAddress("0xd8002C0005EeB9C36e33321193bA41200675C96d")
	address := common.HexToAddress("0xF11f7C4f4db93567093CF7BA2996c6c8c0834660")
	ethclient, err := ethclient.Dial("http://192.168.3.92:8545")
	if err != nil {
		t.Error(err)
	}
	keystorePath := "/Users/daminyang/service/geth/data1/keystore/UTC--2021-05-25T07-28-12.168815000Z--46e6fc87ef7b0c12b96b6098a70701bfc8edb28d"
	keyJson, err := ioutil.ReadFile(keystorePath)
	key, err := keystore.DecryptKey(keyJson, "123456")
	if err != nil {
		t.Error(err)
	}

	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Error("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		t.Error(err)
	}

	gasPrice, err := ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err)
	}

	auth := bind.NewKeyedTransactor(key.PrivateKey)
	if err != nil {
		t.Error(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	anchorContract, err := NewAnchorChain(address, ethclient)
	tx, err := anchorContract.AddUser(auth, proxyAddr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("tx:" + tx.Hash().String())
}

func TestAnchorChainTransactor_crossChain_Geth(t *testing.T) {
	address := common.HexToAddress("0x5Cb886953183027DF790458dd183BA688A3effcC")
	ethclient, err := ethclient.Dial("http://192.168.3.92:8545")
	if err != nil {
		t.Error(err)
	}
	keystorePath := "/Users/daminyang/service/geth/data1/keystore/UTC--2021-05-25T07-28-12.168815000Z--46e6fc87ef7b0c12b96b6098a70701bfc8edb28d"
	keyJson, err := ioutil.ReadFile(keystorePath)
	key, err := keystore.DecryptKey(keyJson, "123456")
	if err != nil {
		t.Error(err)
	}

	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Error("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	auth := bind.NewKeyedTransactor(key.PrivateKey)
	if err != nil {
		t.Error(err)
	}
	auth.From = fromAddress

	anchorContract, err := NewAnchorChain(address, ethclient)

	current := time.Now().Unix()

	params := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		params = append(params, fmt.Sprintf("params_%d", current))
		current = current + int64(i)
	}
	tx, err := anchorContract.CrossChain(auth, params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("tx:" + tx.Hash().String())

}

func TestAnchorChainCaller_ScanBlocks_Sim(t *testing.T) {
	//anchor
	address := common.HexToAddress("0xf45d7444CAE80eaEde29975035f53d7Ca62b9dfe")

	ethclient, err := ethclient.Dial("http://192.168.4.248:8545")
	if err != nil {
		t.Error(err)
	}
	anchorContract, err := NewAnchorChain(address, ethclient)
	if err != nil {
		t.Error(err)
	}
	result, err := anchorContract.ScanBlocks(&bind.CallOpts{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
