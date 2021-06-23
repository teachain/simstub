package service

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

func buildArgsWithMethodId(method string, args []interface{}, abiContent string) ([]byte, error) {
	if len(abiContent) == 0 {
		return nil, errors.New("abi is null")
	}
	if len(method) == 0 {
		return nil, errors.New("request Method must be set")
	}
	abiObject, err := abi.JSON(strings.NewReader(abiContent))
	if err != nil {
		return nil, errors.New("abi read:" + err.Error())
	}
	if len(args) == 0 {
		data, err := abiObject.Pack(method)
		if err != nil {
			return nil, errors.New("abi pack no params:" + err.Error())
		}
		return data, nil
	}
	data, err := abiObject.Pack(method, args...)
	if err != nil {
		return nil, errors.New("abi pack params:" + err.Error())
	}
	return data, nil
}
