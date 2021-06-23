package models

import "github.com/jinzhu/gorm"

type ChainInformation struct {
	gorm.Model
	ChainId   int    `gorm:"column:chain_id"`
	Path      string `gorm:"column:path"`
	ChainName string `gorm:"column:chain_name"`
}

func (this *ChainInformation) TableName() string {
	return "chain_information"
}
func (this *DataBase) AddChainInformation(obj *ChainInformation) (uint, error) {
	err := this.db.Create(obj).Error
	if err != nil {
		return 0, err
	}
	return obj.ID, nil
}
