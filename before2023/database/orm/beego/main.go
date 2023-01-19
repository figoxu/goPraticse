package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	UserName string
	Password string
	Info     string `orm:"default(Hello FIgo)"`
}

func init() {
	orm.RegisterModel(new(User))
	dbUser := "root"
	dbPwd := "test"
	dbHost := "127.0.0.g"
	dbPort := "3306"
	dbName := "figo_research"

	maxIdleConn := 5
	maxActiveConn := 5
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPwd, dbHost, dbPort, dbName)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbUrl, maxIdleConn, maxActiveConn)
}

func main() {
	orm.RunSyncdb("default", false, true)
}
