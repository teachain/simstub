package contracts

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"strings"
	"testing"
)

//0x9FE1a176070CD410d60dae01C92C9061BA9e541C
func TestDeployBridge_Geth(t *testing.T) {
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

	address, tx, _, err := DeployBridge(auth, ethclient)
	if err != nil {
		t.Error(err)
	}

	t.Log("geth bridge addr:", address.Hex())

	t.Log(tx.Hash().Hex())
}

func TestDeployBridge(t *testing.T) {
	ethclient, err := ethclient.Dial("http://192.168.4.248:8545")
	if err != nil {
		t.Error(err)
	}
	proxy, err := NewBridge(bridgeAddr, ethclient)
	if err != nil {
		t.Error(err)
	}

	requests, err := proxy.GetInterchainRequests(&bind.CallOpts{}, big.NewInt(2).Bytes())
	if err != nil {
		t.Error(err)
	}
	fmt.Println("requests=", requests)

	var s []string

	requests = strings.ReplaceAll(requests, "\x00", "")

	err = json.Unmarshal([]byte(requests), &s)

	if err != nil {
		t.Error(err)
		return
	}
	for _, str := range s {
		fmt.Println("str=", str)
	}
}
