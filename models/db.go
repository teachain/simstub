package models

import (
	"fmt"
	"gitee.com/chatcode/simstub/commons"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/url"
)

func GetDBConnection(conf *commons.DBConfig) (*gorm.DB, error) {
	format := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s"
	dsn := fmt.Sprintf(format, conf.Username, conf.Password, conf.Host, conf.Port, conf.DbName, conf.Charset, url.QueryEscape(conf.Loc))
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.LogMode(conf.LogMode)
	db.DB().SetMaxIdleConns(conf.MaxIdle)
	db.DB().SetMaxOpenConns(conf.MaxOpen)
	return db, nil
}

type DataBase struct {
	db *gorm.DB
}

func NewDataBase(db *gorm.DB) *DataBase {
	autoMigrate(db)
	return &DataBase{db: db}
}
func (this *DataBase) Close() {
	if this.db != nil {
		this.db.Close()
		this.db = nil
	}
}

func autoMigrate(db *gorm.DB) {
	if db == nil {
		return
	}
	db.AutoMigrate(&Contract{})
	db.AutoMigrate(&CrossChainRequest{})
	db.AutoMigrate(&ContractInitialization{})
	db.AutoMigrate(&PathInitialization{})
	db.AutoMigrate(&AccessControlList{})
	db.AutoMigrate(&TargetChainPath{})
	db.AutoMigrate(&RegisterCallback{})
	db.AutoMigrate(&ChainInformation{})
}
