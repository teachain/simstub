package models

import (
	"testing"
	"time"
)

func TestDataBase_AddCrossChainRequest(t *testing.T) {
	conn, err := GetDBConnection(config)
	if err != nil {
		t.Error(err)
		return
	}
	dataBase := NewDataBase(conn)
	defer dataBase.Close()
	obj := &CrossChainRequest{
		TxHash:                  "0xddddddddafdfadfdsf",
		BlockNumber:             100,
		BlockHash:               "0xdddddddddddfadsfdf",
		BlockTime:               uint64(time.Now().Unix()),
		RequestId:               "0xddddd",
		CallType:                "invoke",
		Path:                    "system.out.println",
		Method:                  "connect",
		Params:                  "hello",
		CallbackPath:            "system.in.scan",
		CallbackMethod:          "callback",
		Status:                  0,
	}
	id, err := dataBase.AddCrossChainRequest(obj)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("id=", id)

}
func TestDataBase_UpdateCrossChainRequestStatus(t *testing.T) {
	conn, err := GetDBConnection(config)
	if err != nil {
		t.Error(err)
		return
	}
	dataBase := NewDataBase(conn)
	defer dataBase.Close()
	err = dataBase.UpdateCrossChainRequestStatus("0xdddddx", 2)
	if err != nil {
		t.Error(err)
	}
}
