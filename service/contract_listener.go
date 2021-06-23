package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"gitee.com/chatcode/simstub/contracts"
	"gitee.com/chatcode/simstub/models"
)

func (this *ContractManager) ListenCrossChainRequest() {
	if this.client == nil {
		this.logger.Error("this.client is nil")
		return
	}
	this.logger.Info("listen on:" + this.bridgeAddress.String())
	bridgeContract, err := contracts.NewBridge(this.bridgeAddress, this.client)
	if err != nil {
		this.logger.Error(err.Error())
		return
	}
	sink := make(chan *contracts.BridgeOnHandleRequest, 1)
	var start uint64 = 1
	opts := &bind.WatchOpts{
		Start: &start,
	}
	filter, err := bridgeContract.WatchOnHandleRequest(opts, sink)
	if err != nil {
		this.logger.Error(err.Error())
		return
	}
	defer filter.Unsubscribe()
	for {
		select {
		case e := <-sink:
			this.logger.Debug("receive CrossChainRequest")
			block, err := this.client.BlockByHash(context.Background(), e.Raw.BlockHash)
			if err != nil {
				this.logger.Error(err.Error())
				continue
			}
			data := strings.ReplaceAll(e.Request, "\x00", "")

			request, err := DecodeInterChainRequest(data)
			if err != nil {
				this.logger.Error(err.Error())
				continue
			}
			obj := &models.CrossChainRequest{
				TxHash:         e.Raw.TxHash.String(),
				BlockNumber:    e.Raw.BlockNumber,
				BlockHash:      e.Raw.BlockHash.String(),
				BlockTime:      block.Time(),
				RequestId:      request.Uid,
				CallType:       request.CallType,
				Path:           request.Path,
				Method:         request.Method,
				Params:         request.Args,
				CallbackPath:   request.CallbackPath,
				CallbackMethod: request.CallbackMethod,
				Status:         1,                 //跨链发起状态
				Caller:         e.Sender.String(), //业务合约地址
				Sender:         request.Identity,  //跨链交易发起者账户地址
				CrossChainType: request.CrossChainType,
			}
			id, err := this.dataBase.AddCrossChainRequest(obj)
			if err != nil {
				this.logger.Error(err.Error())
				continue
			}
			this.logger.Info(fmt.Sprintf("CrossChainRequest database id=%d;sender=%s", id, e.Sender.String()))
		case err := <-filter.Err():
			if err != nil {
				this.logger.Error(err.Error())
			}
			return
		case <-this.stop:
			this.logger.Info("ListenCrossChainRequest receive stop message")
			return
		}
	}
}

type InterChainRequest struct {
	Uid            string
	CallType       string
	Path           string
	Method         string
	Args           string
	CallbackPath   string
	CallbackMethod string
	Identity       string
	CrossChainType string
}

func DecodeInterChainRequest(jsonStr string) (*InterChainRequest, error) {
	var interChainRequest InterChainRequest
	var jsonResult []string
	err := json.Unmarshal([]byte(jsonStr), &jsonResult)
	if err != nil {
		return nil, err
	}
	if len(jsonResult) <= 7 {
		return nil, errors.New("request error")
	}
	interChainRequest.Uid = jsonResult[0]
	interChainRequest.CallType = jsonResult[1]
	interChainRequest.Path = jsonResult[2]
	interChainRequest.Method = jsonResult[3]
	interChainRequest.Args = jsonResult[4]
	interChainRequest.CallbackPath = jsonResult[5]
	interChainRequest.CallbackMethod = jsonResult[6]
	interChainRequest.Identity = jsonResult[7]
	interChainRequest.CrossChainType = jsonResult[8]
	return &interChainRequest, nil

}

func (this *ContractManager) ListenRegisterCallbackResult() {
	if this.client == nil {
		this.logger.Error("this.client is nil")
		return
	}
	this.logger.Info("listen on:" + this.bridgeAddress.String())
	bridgeContract, err := contracts.NewBridge(this.bridgeAddress, this.client)
	if err != nil {
		this.logger.Error(err.Error())
		return
	}
	sink := make(chan *contracts.BridgeOnRegisterCallbackResult, 1)
	var start uint64 = 1
	opts := &bind.WatchOpts{
		Start: &start,
	}
	filter, err := bridgeContract.WatchOnRegisterCallbackResult(opts, sink)
	if err != nil {
		this.logger.Error(err.Error())
		return
	}
	defer filter.Unsubscribe()
	for {
		select {
		case e := <-sink:
			this.logger.Debug("receive RegisterCallbackResult")
			obj := &models.RegisterCallback{
				RequestId:     e.Uid,
				TransactionId: e.Result[0],
				Sequence:      e.Result[1],
				ErrorCode:     e.Result[2],
				ErrorMsg:      e.Result[3],
			}
			if obj.ErrorCode == "0" {
				//成功调用
				this.dataBase.UpdateCrossChainRequestStatus(obj.RequestId, 0)
			} else if obj.ErrorCode == "-1" {
				this.dataBase.UpdateCrossChainRequestStatus(obj.RequestId, -1)
			}
			if e.Result[4] != "" {
				var results []string
				err := json.Unmarshal([]byte(e.Result[4]), &results)
				if err != nil {
					this.logger.Error("json.Unmarshal", "err", err.Error())
				}
				if len(results[0]) >= 2 {
					obj.TxHash = results[0]
					obj.Sender = results[1]
				}
			}
			id, err := this.dataBase.AddRegisterCallback(obj)
			if err != nil {
				this.logger.Error("AddRegisterCallback", "err", err.Error(), "register", fmt.Sprintf("%+v", *obj))
				continue
			}
			this.logger.Info(fmt.Sprintf("RegisterCallback database id=%d", id))
		case err := <-filter.Err():
			this.logger.Error(err.Error())
		case <-this.stop:
			this.logger.Info("ListenRegisterCallbackResult receive stop message")
			return
		}
	}
}
