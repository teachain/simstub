package models

import "github.com/jinzhu/gorm"

type AccessControlList struct {
	gorm.Model
	Resource string `gorm:"resource"`
	User     string `gorm:"user"`
	TxHash   string `gorm:"tx_hash"`
}

func (this *DataBase) AddAccessControlList(obj *AccessControlList) (uint, error) {
	err := this.db.Create(obj).Error
	if err != nil {
		return 0, err
	}
	return obj.ID, nil
}
func (this *AccessControlList) TableName() string {
	return "access_control_list"
}

func (this *DataBase) AccessControlListExists(resource string, user string) bool {
	var count int
	err := this.db.Model(&AccessControlList{}).Where("resource=?", resource).Where("user=?", user).Count(&count).Error
	if err == nil && count > 0 {
		return true
	}
	return false
}
