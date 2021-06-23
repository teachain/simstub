package models

import "github.com/jinzhu/gorm"

type RegisterCallback struct {
	gorm.Model
	RequestId     string `gorm:"request_id"`
	Sequence      string `gorm:"sequence"`
	TransactionId string `gorm:"transaction_id"`
	ErrorCode     string `gorm:"error_code"`
	ErrorMsg      string `gorm:"error_msg"`
	Sender        string `gorm:"sender"`
	TxHash        string `gorm:"tx_hash"`
}

func (this *DataBase) AddRegisterCallback(obj *RegisterCallback) (uint, error) {
	err := this.db.Create(obj).Error
	if err != nil {
		return 0, err
	}
	return obj.ID, nil
}
func (this *RegisterCallback) TableName() string {
	return "register_callback"
}
