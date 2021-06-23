package service

import (
	"crypto/ecdsa"
	"errors"
	"gitee.com/chatcode/simstub/commons"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
)

func LoadAccount(config *commons.AccountConfig) (*bind.TransactOpts, error) {
	keyJson, err := ioutil.ReadFile(config.Keystore)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(keyJson, config.Password)
	if err != nil {
		return nil, err
	}

	publicKey := key.PrivateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	auth := bind.NewKeyedTransactor(key.PrivateKey)

	auth.From = fromAddress

	return auth, nil
}
