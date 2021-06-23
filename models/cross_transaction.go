package models

import (
	"fmt"
	"gitee.com/chatcode/chainrouter/stub"
)

func (this *DataBase) GetCrossTransactionInformationByPage(start int, limit int) ([]stub.CrossTransactionInformation, error) {
	transactionRecords := make([]stub.CrossTransactionInformation, 0)
	crossChainRequestTableName := (&CrossChainRequest{}).TableName()
	registerCallbackTableName := (&RegisterCallback{}).TableName()
	chainInformationTableName := (&ChainInformation{}).TableName()
	format := `SELECT 
    a.tx_hash as tx_hash,
    (select chain_name from %s  where path=a.callback_path) as from_chain_name,
    (select chain_name from %s  where path=a.path) as to_chain_name,
	a.cross_chain_type as cross_chain_type,
	a.status as status,
	a.block_time   as block_time FROM %s as a ,%s as b where a.request_id=b.request_id limit %d offset %d`
	//在这种方式下，column关键字必须在tag中，否则有可能读不出值来
	err := this.db.Raw(fmt.Sprintf(format, chainInformationTableName, chainInformationTableName, crossChainRequestTableName, registerCallbackTableName, limit, start)).Scan(&transactionRecords).Error
	return transactionRecords, err
}
func (this *DataBase) GetTotalCrossTransactionInformation() (int, error) {
	type TotalParam struct {
		Total int `gorm:"column:total"`
	}
	var totalParam TotalParam
	crossChainRequestTableName := (&CrossChainRequest{}).TableName()
	registerCallbackTableName := (&RegisterCallback{}).TableName()
	format := `SELECT count(*) as total FROM %s as a ,%s as b where a.request_id=b.request_id`
	err := this.db.Raw(fmt.Sprintf(format, crossChainRequestTableName, registerCallbackTableName)).Scan(&totalParam).Error
	if err != nil {
		return 0, err
	}
	return totalParam.Total, nil
}

func (this *DataBase) GetCrossChainTransaction(txHash string) (stub.CrossChainTransaction, error) {
	var result stub.CrossChainTransaction
	crossChainRequestTableName := (&CrossChainRequest{}).TableName()
	registerCallbackTableName := (&RegisterCallback{}).TableName()
	chainInformationTableName := (&ChainInformation{}).TableName()
	format := `SELECT 
    (select chain_name from %s where path=a.callback_path) as from_chain_name,
    (select chain_name from %s where path=a.path) as to_chain_name,
    a.block_time as from_block_time,
    a.tx_hash as from_tx_hash,
    a.status as status,
    a.sender as from_sender,
    a.params as params,
    b.sender as to_sender,
    b.tx_hash as  to_tx_hash,
	a.params as data FROM %s as a ,%s as b where a.request_id=b.request_id and a.tx_hash='%s'`
	//在这种方式下，column关键字必须在tag中，否则有可能读不出值来
	err := this.db.Raw(fmt.Sprintf(format, chainInformationTableName, chainInformationTableName, crossChainRequestTableName, registerCallbackTableName,txHash)).Scan(&result).Error
	return result, err
}
