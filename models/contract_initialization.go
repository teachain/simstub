package models

import "github.com/jinzhu/gorm"

type ContractInitialization struct {
	gorm.Model
	BridgeAddress string `gorm:"bridge_address"`
	AnchorAddress string `gorm:"anchor_address"`
	TxHash        string `gorm:"tx_hash"`
}

func (this *ContractInitialization) TableName() string {
	return "contract_initialization"
}

func (this *DataBase) AddContractInit(obj *ContractInitialization) (uint, error) {
	err := this.db.Create(obj).Error
	if err != nil {
		return 0, err
	}
	return obj.ID, nil
}
func (this *DataBase) ContractInitReady() bool {
	var count int
	err := this.db.Model(&ContractInitialization{}).Count(&count).Error
	if err == nil && count > 0 {
		return true
	}
	return false
}
