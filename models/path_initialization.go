package models

import "github.com/jinzhu/gorm"

type PathInitialization struct {
	gorm.Model
	ProxyAddress string `gorm:"proxy_address"`
	PathAddress  string `gorm:"path_address"`
	Path         string `gorm:"path"`
	TxHash       string `gorm:"tx_hash"`
}

func (this *DataBase) AddPathInit(obj *PathInitialization) (uint, error) {
	err := this.db.Create(obj).Error
	if err != nil {
		return 0, err
	}
	return obj.ID, nil
}
func (this *DataBase) PathInitReady(proxyAddress string, pathAddress string, path string) bool {
	var count int
	err := this.db.Model(&PathInitialization{}).Where("proxy_address=?", proxyAddress).Where("path_address=?", pathAddress).Where("path=?", path).Count(&count).Error
	if err == nil && count > 0 {
		return true
	}
	return false
}
