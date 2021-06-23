package service

import (
	"context"
	"fmt"
	"gitee.com/chatcode/simstub/contracts"
	"gitee.com/chatcode/simstub/models"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

func (this *ContractManager) CheckTargetChainPath(resource string, path string, method string, callbackPath string, callbackMethod string,crossChainType string) error {
	start := time.Now()
	defer func() {
		this.logger.Info("CheckTargetChainPath exited")
		this.logger.Info(fmt.Sprintf("CheckTargetChainPath time elapsed:%+v", time.Now().Sub(start)))
	}()
	anchorExists := this.dataBase.ContractExists(resource)
	if anchorExists {
		anchorAddress := this.dataBase.GetContractAddressByName(resource)
		if this.dataBase.TargetChainPathExists(anchorAddress, path, method, callbackPath, callbackMethod) {
			return nil
		}
		anchorContract, err := contracts.NewAnchorChain(common.HexToAddress(anchorAddress), this.client)
		if err != nil {
			return err
		}
		this.logger.Info("setTargetchain anchorAddress=" + anchorAddress)
		tx, err := anchorContract.SetTargetChainPath(this.auth, path, method, callbackPath, callbackMethod,crossChainType)
		if err != nil {
			return err
		}
		timer := time.NewTimer(startInterval)
		defer timer.Stop()
	label:
		for {
			select {
			case <-timer.C:
				receipt, err := this.client.TransactionReceipt(context.Background(), tx.Hash())
				if err == nil {
					if receipt != nil {
						if receipt.Status == 1 {
							obj := &models.TargetChainPath{
								ContractAddress: anchorAddress,
								Path:            path,
								Method:          method,
								CallbackPath:    callbackPath,
								CallbackMethod:  callbackMethod,
							}
							this.logger.Info(fmt.Sprintf("anchor SetTargetChainPath success;tx=%s", tx.Hash().String()))
							_, err = this.dataBase.AddTargetChainPath(obj)
							if err != nil {
								return err
							}
							break label
						} else if receipt.Status == 0 {
							this.logger.Error(fmt.Sprintf("anchor SetTargetChainPath fail;tx=%s", tx.Hash().String()))
							break label
						}
					} else {
						timer.Reset(checkInterval)
					}
				} else {
					timer.Reset(checkInterval)
				}
			}
		}
	}
	return nil
}
