package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var proxyAddr = common.HexToAddress("0x1396b2E5253fe9E9A60A973b781422306F77C78C")
var bridgeAddr = common.HexToAddress("0xCE61Cfbec4784a31788877D1F757Cc1ee45b6274")
var anchorAddr = common.HexToAddress("0x49b330e2c7977691A328761201E8741fC7F5c2f1")

func TestDeployProxy(t *testing.T) {
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

	address, tx, _, err := DeployProxy(auth, ethclient)
	if err != nil {
		t.Error(err)
	}

	t.Log("addr:", address.Hex())

	t.Log(tx.Hash().Hex())
}

//注册bridge
func TestNewProxy(t *testing.T) {
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

	proxy, err := NewProxy(proxyAddr, ethclient)
	if err != nil {
		t.Error(err)
	}
	transaction, err := proxy.AddPath(auth, "judicial.sim.bridge", bridgeAddr, BridgeABI)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(transaction.Hash().String())

}

//注册 anchor
func TestNewProxy2(t *testing.T) {
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
	proxy, err := NewProxy(proxyAddr, ethclient)
	if err != nil {
		t.Error(err)
	}

	transaction2, err := proxy.AddPath(auth, "judicial.sim.evidence", anchorAddr, AnchorChainABI)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(transaction2.Hash().String())
}
func TestNewProxy3(t *testing.T) {
	ethclient, err := ethclient.Dial("http://192.168.4.248:8545")
	if err != nil {
		t.Error(err)
	}
	proxy, err := NewProxy(common.HexToAddress("0x51e4B5859A3A88D955504eCEC8181Ac410B6C63A"), ethclient)
	if err != nil {
		t.Error(err)
	}

	call, err := proxy.BytesToUint(&bind.CallOpts{}, big.NewInt(10000).Bytes())
	if err != nil {
		t.Error(err)
	}
	fmt.Println("call", call.String())
}

//0x1862C15A62A491635D327a472890f1CEC68074f3
func TestDeployProxy_Geth(t *testing.T) {
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

	address, tx, _, err := DeployProxy(auth, ethclient)
	if err != nil {
		t.Error(err)
	}

	t.Log("geth proxy addr:", address.Hex())

	t.Log(tx.Hash().Hex())
}
func TestProxySession_AddPath_Anchorchain_Geth(t *testing.T) {
	anchorAddr := common.HexToAddress("0x7c7b0c7D5543d60884b63b6fb61e6C22b94a6fa2")
	resourcePath := "judicial.geth.evidence"
	proxyAddr := common.HexToAddress("0x789E1f7e4A776Cf0479C9703aa05420fB2169Af9")
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
	proxy, err := NewProxy(proxyAddr, ethclient)
	if err != nil {
		t.Error(err)
	}

	transaction, err := proxy.AddPath(auth, resourcePath, anchorAddr, AnchorChainABI)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("proxy add anchor path:" + transaction.Hash().String())
}

//注册bridge
func TestProxySession_AddPath_Bridge_Geth(t *testing.T) {
	bridgeAddr := common.HexToAddress("0x9FE1a176070CD410d60dae01C92C9061BA9e541C")
	resourcePath := "judicial.geth.bridge"

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

	proxy, err := NewProxy(proxyAddr, ethclient)
	if err != nil {
		t.Error(err)
	}
	transaction, err := proxy.AddPath(auth, resourcePath, bridgeAddr, BridgeABI)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("proxy add bridge path" + transaction.Hash().String())

}
