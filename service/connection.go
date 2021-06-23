package service

import (
	"errors"
	"fmt"
	"gitee.com/chatcode/chainrouter/stub"
	"gitee.com/chatcode/simstub/commons"
	"gitee.com/chatcode/simstub/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/go-hclog"
	"os"
	"strings"
)

type SimpleConnection struct {
	client        *ethclient.Client
	proxyAddress  common.Address //代理合约地址
	bridgeAddress common.Address //桥接合约地址
	config        *commons.Config
	logger        hclog.Logger
}

func NewSimpleConnection(client *ethclient.Client, config *commons.Config) *SimpleConnection {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: false,
		Name:       config.Common.Name,
		TimeFormat: commons.TimeFormat,
	})
	return &SimpleConnection{
		client:       client,
		config:       config,
		logger:       logger,
	}
}
func (this *SimpleConnection) SetProxyAddress(proxyAddress common.Address) {
	this.proxyAddress = proxyAddress
}

func (this *SimpleConnection) GetResources() []stub.ResourceInfo {
	paths, err := this.ListPath()
	if err != nil {
		return nil
	}
	this.logger.Debug(fmt.Sprintf("path=%+v", paths))
	result := make([]stub.ResourceInfo, 0)
	names := make(map[string]struct{})
	for _, v := range paths {
		name := strings.Split(v, ".")[2]
		if _, ok := names[name]; ok {
			continue
		}
		names[name] = struct{}{}
		result = append(result, stub.ResourceInfo{
			Name:     name,
			StubType: this.config.Common.Type,
		})
	}
	if _, ok := names[stub.PROXY_NAME]; !ok {
		result = append(result, stub.ResourceInfo{
			Name:     stub.PROXY_NAME,
			StubType: this.config.Common.Type,
		})
	}
	return result
}

//获取所有的资源路径
func (this *SimpleConnection) ListPath() ([]string, error) {
	if this.client == nil {
		return nil, errors.New("client is nil")
	}
	proxyContract, err := contracts.NewProxy(this.proxyAddress, this.client)
	if err != nil {
		return nil, err
	}
	return proxyContract.GetPaths(&bind.CallOpts{})
}
func (this *SimpleConnection) GetAbiByName(name string) (string, error) {
	if this.client == nil {
		return "", errors.New("client is nil")
	}
	proxyContract, err := contracts.NewProxy(this.proxyAddress, this.client)
	if err != nil {
		return "", err
	}
	return proxyContract.GetAbiByName(&bind.CallOpts{}, name)
}
func (this *SimpleConnection) CallByProxy(name string, argsWithMethodId []byte) ([]byte, error) {
	if this.client == nil {
		return nil, errors.New("client is nil")
	}
	proxyContract, err := contracts.NewProxy(this.proxyAddress, this.client)
	if err != nil {
		return nil, err
	}
	return proxyContract.ConstantCall(&bind.CallOpts{}, name, argsWithMethodId)
}

func (this *SimpleConnection) SendTransactionByProxy(auth *bind.TransactOpts, uid string, name string, argsWithMethodId []byte) (*types.Transaction, error) {
	if this.client == nil {
		return nil, errors.New("client is nil")
	}
	proxyContract, err := contracts.NewProxy(this.proxyAddress, this.client)
	if err != nil {
		return nil, err
	}
	return proxyContract.SendTransaction(auth, uid, name, argsWithMethodId)
}
