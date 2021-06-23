package models

import (
	"testing"
)

func TestDataBase_AddContract(t *testing.T) {
	conn, err := GetDBConnection(config)
	if err != nil {
		t.Error(err)
		return
	}
	dataBase := NewDataBase(conn)
	defer dataBase.Close()
	obj := &Contract{
		Name:    "proxy",
		Address: "0xkdaoisjfioadsjfsfjsdfiodjsf",
	}
	id, err := dataBase.AddContract(obj)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("id=", id)
}

func TestDataBase_ContractExists(t *testing.T) {
	conn, err := GetDBConnection(config)
	if err != nil {
		t.Error(err)
		return
	}
	dataBase := NewDataBase(conn)
	defer dataBase.Close()
	exist := dataBase.ContractExists("proxy1")

	t.Log("exist=", exist)
}
