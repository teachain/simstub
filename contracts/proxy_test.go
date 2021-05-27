package contracts

import (
	"context"
	"crypto/ecdsa"
	"io/ioutil"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"


)

func TestDeployProxy(t *testing.T) {
	ethclient, err := ethclient.Dial("http://localhost:8545")
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

	auth:= bind.NewKeyedTransactor(key.PrivateKey)
	if err != nil {
		t.Error(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, _, err := DeployProxy(auth, ethclient)
	if err != nil {
		t.Error(err)
	}

	t.Log(address.Hex())

	t.Log(tx.Hash().Hex())


}
