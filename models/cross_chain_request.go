package models

import "github.com/jinzhu/gorm"

type CrossChainRequest struct {
	gorm.Model
	TxHash         string `gorm:"tx_hash"`
	BlockHash      string `gorm:"block_hash"`
	BlockNumber    uint64 `gorm:"block_number"`
	BlockTime      uint64 `gorm:"block_time"`
	RequestId      string `gorm:"request_id;unique_index:request_id_index"`
	CallType       string `gorm:"call_Type"`        //调用类型（调用或查询）
	Path           string `gorm:"path"`             //路径
	Method         string `gorm:"method"`           //方法
	Params         string `gorm:"params"`           //参数
	CallbackPath   string `gorm:"callback_path"`    //回调路径
	CallbackMethod string `gorm:"callback_method"`  //回调方法
	Status         int    `gorm:"status"`           //跨链请求状态(0初始化，1完成，2失败)
	Caller         string `gorm:"caller"`           //调用桥接合约的合约地址
	Sender         string `gorm:"sender"`           //跨链交易的发起者
	CrossChainType string `gorm:"cross_chain_type"` //跨链类型
}

func (this *CrossChainRequest) TableName() string {
	return "cross_chain_requests"
}

func (this *DataBase) AddCrossChainRequest(obj *CrossChainRequest) (uint, error) {
	err := this.db.Create(obj).Error
	if err != nil {
		return 0, err
	}
	return obj.ID, nil
}
func (this *DataBase) UpdateCrossChainRequestStatus(requestId string, status int) error {
	//需要注意，如果记录不存在，也不会报错
	return this.db.Model(&CrossChainRequest{}).Where("request_id=?", requestId).
		Update("status", status).Error
}
