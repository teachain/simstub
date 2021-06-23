package models

import (
	"fmt"
	"testing"
)

func TestDataBase_LoadTransactionRecordPage(t *testing.T) {
	conn, err := GetDBConnection(config)
	if err != nil {
		t.Error(err)
		return
	}
	dataBase := NewDataBase(conn)
	defer dataBase.Close()

	result, err := dataBase.GetCrossTransactionInformationByPage(0, 10)
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range result {
		fmt.Printf("TransactionRecordPage%+v\n", v)
	}
}
func TestDataBase_GetTotalTransactionRecord(t *testing.T) {
	conn, err := GetDBConnection(config)
	if err != nil {
		t.Error(err)
		return
	}
	dataBase := NewDataBase(conn)
	defer dataBase.Close()

	result, err := dataBase.GetCrossChainTransaction("0xe80e6dca6a8d72959055ae6dc59902536aa4ce58d6a1e6249af3f7991f981ccc")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("CrossTransaction=%+v\n", result)
}
