

package models

import (
    "log"
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var db *gorm.DB

type Model struct {
    ID int `gorm:"primary_key" json:"id"`
    CreatedOn int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
}

func init() {
    var (
        err error
		dbType, dbName, user, password, host, tablePrefix string
    )

	dbType = os.Getenv("DBTYPE")
    dbName = os.Getenv("DBNAME")
    user = os.Getenv("DBUSER")
    password = os.Getenv("DBPASSWORD")
    host = os.Getenv("DBHOST")
    tablePrefix = os.Getenv("DBTABLE_PREFIX")

    db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", 
        user, 
        password, 
        host, 
        dbName))

    if err != nil {
        log.Println(err)
    }

    gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
        return tablePrefix + defaultTableName;
    }

    db.SingularTable(true)
    db.LogMode(true)
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
    defer db.Close()
}

