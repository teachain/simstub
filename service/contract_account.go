package service

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/chatcode/simstub/contracts"
	"gitee.com/chatcode/simstub/models"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

func (this *ContractManager) CheckAnchorUser(resource string) error {
	start := time.Now()
	defer func() {
		this.logger.Info("CheckAnchorUser exited")
		this.logger.Info(fmt.Sprintf("CheckAnchorUser time elapsed:%+v", time.Now().Sub(start)))
	}()
	anchorExists := this.dataBase.ContractExists(resource)
	if anchorExists {
		anchorAddress := this.dataBase.GetContractAddressByName(resource)
		exists := this.dataBase.AccessControlListExists(anchorAddress, this.proxyAddress.String())
		if exists {
			return nil
		}
		anchorContract, err := contracts.NewAnchorChain(common.HexToAddress(anchorAddress), this.client)
		if err != nil {
			return err
		}
		this.logger.Info("CheckAnchorUser anchorAddress=" + anchorAddress)
		this.logger.Info("CheckAnchorUser proxyAddress=" + this.proxyAddress.String())
		tx, err := anchorContract.AddUser(this.auth, this.proxyAddress)
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
							obj := &models.AccessControlList{
								Resource: anchorAddress,
								User:     this.proxyAddress.String(),
								TxHash:   tx.Hash().String(),
							}
							this.logger.Info(fmt.Sprintf("anchor add user(proxy) success;tx=%s", tx.Hash().String()))
							_, err = this.dataBase.AddAccessControlList(obj)
							if err != nil {
								return err
							}
							break label
						} else if receipt.Status == 0 {
							this.logger.Error(fmt.Sprintf("anchor add user(proxy) fail;tx=%s", tx.Hash().String()))
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
	} else {
		return errors.New("anchor contract not exists")
	}
	return nil
}
