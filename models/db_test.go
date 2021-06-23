package models

import (
	"gitee.com/chatcode/simstub/commons"
	"testing"
)

var config = &commons.DBConfig{
	Username: "root",
	Password: "root",
	Host:     "192.168.4.248",
	Port:     3306,
	DbName:   "gethstub",
	Charset:  "utf8mb4",
	MaxIdle:  1000,
	MaxOpen:  2000,
	LogMode:  true,
	Loc:      "Asia/Shanghai",
}

func TestGetDBConnection(t *testing.T) {
	conn, err := GetDBConnection(config)
	if err != nil {
		t.Error(err)
		return
	}
	dataBase := NewDataBase(conn)

	defer dataBase.Close()

}
