package models

import (
	"testing"
)

func TestDataBase_AddChainInformation(t *testing.T) {
	conn, err := GetDBConnection(config)
	if err != nil {
		t.Error(err)
		return
	}
	dataBase := NewDataBase(conn)
	defer dataBase.Close()
	info := &ChainInformation{
		ChainId:   1,
		Path:      "judicial.geth.evidence",
		ChainName: "以太坊",
	}
	id, err := dataBase.AddChainInformation(info)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("id=", id)
	info2 := &ChainInformation{
		ChainId:   2,
		Path:      "judicial.sim.evidence",
		ChainName: "简链",
	}
	id, err = dataBase.AddChainInformation(info2)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("id=", id)
}
