package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"shop/pkg/logging"
	"shop/pkg/setting"
)

var db *gorm.DB

func Setup() {
	var err error

	if db, err = gorm.Open(setting.Dbcfg.DbType,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.Dbcfg.DbUser,
			setting.Dbcfg.DbPass,
			setting.Dbcfg.DbHost,
			setting.Dbcfg.DbPort,
			setting.Dbcfg.DbName)); err != nil {
		logging.Fatal(err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	db.Close()
}
