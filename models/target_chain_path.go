package models

import "github.com/jinzhu/gorm"

type TargetChainPath struct {
	gorm.Model
	Path            string `gorm:"path"`
	Method          string `gorm:"method"`
	CallbackPath    string `gorm:"callback_path"`
	CallbackMethod  string `gorm:"callback_method"`
	ContractAddress string `gorm:"contract_address"`
}

func (this *TargetChainPath) TableName() string {
	return "target_chain_path"
}

func (this *DataBase) AddTargetChainPath(obj *TargetChainPath) (uint, error) {
	err := this.db.Create(obj).Error
	if err != nil {
		return 0, err
	}
	return obj.ID, nil
}

func (this *DataBase) TargetChainPathExists(contractAddress string, path string, method string, callbackPath string, callbackMethod string) bool {
	var count int
	err := this.db.Model(&TargetChainPath{}).
		Where("path=?", path).
		Where("method=?", method).
		Where("callback_path=?", callbackPath).
		Where("contract_address=?", contractAddress).
		Where("callback_method=?", callbackMethod).Count(&count).Error
	if err == nil && count > 0 {
		return true
	}
	return false
}
