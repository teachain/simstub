package models

import "github.com/jinzhu/gorm"

type Contract struct {
	gorm.Model
	Name    string `gorm:"name"`    //合约自定义名称
	Address string `gorm:"address"` //合约地址
}

func (this *Contract) TableName() string {
	return "contract"
}

func (this *DataBase) AddContract(obj *Contract) (uint, error) {
	err := this.db.Create(obj).Error
	if err != nil {
		return 0, err
	}
	return obj.ID, nil
}
func (this *DataBase) ContractExists(name string) bool {
	var count int
	err := this.db.Model(&Contract{}).Where("name=?", name).Count(&count).Error
	if err == nil && count > 0 {
		return true
	}
	return false
}
func (this *DataBase) GetContractAddressByName(name string) string {
	var contract Contract
	err := this.db.Model(&Contract{}).First(&contract, "name=?", name).Error
	if err == nil {
		return contract.Address
	}
	return ""
}
