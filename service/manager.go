package service

import (
	"errors"
	"fmt"
	"gitee.com/chatcode/chainrouter/stub"
	"gitee.com/chatcode/simstub/commons"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/go-hclog"
	"os"
	"strings"
)

type Manager struct {
	auth            *bind.TransactOpts
	conn            *SimpleConnection
	logger          hclog.Logger
	abiCache        map[string]string
	config          *commons.Config
	contractManager *ContractManager
}

func NewManager(config *commons.Config) (*Manager, error) {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: false,
		Name:       config.Common.Name,
		TimeFormat: commons.TimeFormat,
	})
	url := fmt.Sprintf("%s://%s:%d", config.Node.Protocol, config.Node.Host, config.Node.Port)
	logger.Debug("node url:" + url)
	ethclient, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	account, err := LoadAccount(config.Account)
	if err != nil {
		return nil, err
	}
	conn := NewSimpleConnection(ethclient, config)
	contractManager, err := NewContractManager(ethclient, account, config)
	if err != nil {
		return nil, err
	}
	go contractManager.Run(func() {
		conn.SetProxyAddress(contractManager.GetProxyAddress())
	})
	return &Manager{
		config:          config,
		logger:          logger,
		abiCache:        make(map[string]string),
		conn:            conn,
		auth:            account,
		contractManager: contractManager,
	}, nil
}

func (this *Manager) GetAbiByName(name string) (string, error) {
	if _, ok := this.abiCache[name]; ok {
		return this.abiCache[name], nil
	}
	return this.conn.GetAbiByName(name)
}

func (this *Manager) CallByProxy(request stub.TransactionRequest) stub.TransactionResponse {
	path := request.TransactionContext.Path
	abi, err := this.GetAbiByName(path.Resource)
	if err != nil {
		return stub.TransactionResponse{
			ErrorCode: -1,
			Message:   err.Error(),
		}
	}
	argsWithMethod, err := buildArgsWithMethodId(request.Method, request.Args, abi)
	if err != nil {
		return stub.TransactionResponse{
			ErrorCode: -1,
			Message:   err.Error(),
		}
	}
	result, err := this.conn.CallByProxy(path.Resource, argsWithMethod)
	if err != nil {
		return stub.TransactionResponse{
			ErrorCode: -1,
			Message:   err.Error(),
		}
	}
	res, err := Unpack(request.Method, request.OutputType, result, abi)
	if err != nil {
		return stub.TransactionResponse{
			ErrorCode: -1,
			Message:   err.Error(),
		}
	}
	return stub.TransactionResponse{ErrorCode: 0, Result: res}
}

func Unpack(method string, outputType string, data []byte, abiContent string) (interface{}, error) {
	abiObject, err := abi.JSON(strings.NewReader(abiContent))
	if err != nil {
		return nil, err
	}
	if outputType == "string" {
		var (
			ret0 = new(string)
		)
		out := ret0
		err := abiObject.Unpack(out, method, data)
		if err != nil {
			return nil, err
		}
		return *ret0, nil

	}
	if outputType == "[]string" {
		var (
			ret0 = new([]string)
		)
		out := ret0
		err := abiObject.Unpack(out, method, data)
		if err != nil {
			return nil, err
		}
		return *ret0, nil
	}
	return nil, errors.New("unknow outputType:" + outputType)

}

func (this *Manager) SendTransactionByProxy(request stub.TransactionRequest) stub.TransactionResponse {
	this.logger.Debug("[Manager] SendTransactionByProxy")
	path := request.TransactionContext.Path
	abi, err := this.GetAbiByName(path.Resource)
	if err != nil {
		this.logger.Error("[Manager] SendTransactionByProxy GetAbiByName err=" + err.Error())
		return stub.TransactionResponse{
			ErrorCode: -1,
			Message:   err.Error(),
		}
	}
	this.logger.Debug("[Manager] SendTransactionByProxy method=" + request.Method)
	argsWithMethod, err := buildArgsWithMethodId(request.Method, request.Args, abi)
	if err != nil {
		return stub.TransactionResponse{
			ErrorCode: -1,
			Message:   err.Error(),
		}
	}
	uid := request.Options[stub.TRANSACTION_UNIQUE_ID]
	result, err := this.conn.SendTransactionByProxy(this.auth, uid.(string), path.Resource, argsWithMethod)
	if err != nil {
		return stub.TransactionResponse{
			ErrorCode: -1,
			Message:   err.Error(),
		}
	}
	response := stub.TransactionResponse{
		ErrorCode: 0,
		Hash:      result.Hash().String(),
		Sender:    this.auth.From.String(),
	}
	//返回交易哈希值
	return response
}
func (this *Manager) GetResources() []stub.ResourceInfo {
	this.logger.Debug("[Manager] GetResources")
	return this.conn.GetResources()
}
func (this *Manager) GetCrossTransactionInformationPage(request stub.CrossTransactionInformationPageRequest) stub.CrossTransactionInformationPage {
	result, err := this.contractManager.GetDataBase().GetCrossTransactionInformationByPage(request.Start, request.Limit)
	if err != nil {
		this.logger.Error(err.Error())
		return stub.CrossTransactionInformationPage{}
	}
	total, err := this.contractManager.GetDataBase().GetTotalCrossTransactionInformation()
	if err != nil {
		this.logger.Error(err.Error())
		return stub.CrossTransactionInformationPage{}
	}
	return stub.CrossTransactionInformationPage{
		Total: total,
		Data:  result,
	}
}
func (this *Manager) GetCrossChainTransaction(request stub.CrossChainTransactionRequest) stub.CrossChainTransaction {
	result, err := this.contractManager.GetDataBase().GetCrossChainTransaction(request.TxHash)
	if err != nil {
		this.logger.Error(err.Error())
		return stub.CrossChainTransaction{}
	}
	return result
}
