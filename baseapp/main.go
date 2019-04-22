package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/quexer/utee"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	sqlite_db *gorm.DB
)

func init() {
	driver := "sqlite3"
	dbLoc := "./test.db"
	sqlitedb, err := gorm.Open(driver, dbLoc)
	utee.Chk(err)
	sqlitedb.DB().SetConnMaxLifetime(time.Minute * 5)
	sqlitedb.DB().SetMaxIdleConns(0)
	sqlitedb.DB().SetMaxOpenConns(5)
	sqlitedb.SingularTable(true)
	sqlitedb.Debug().AutoMigrate(&User{})
	sqlite_db = sqlitedb
}

func main() {
	logrus.Println("Hello")
}
