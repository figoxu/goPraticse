package main

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/quexer/utee"
	"github.com/icrowley/fake"
	"github.com/figoxu/Figo"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("postgres", "user=figo password=123456 dbname=figo host=127.0.0.1 port=5432 sslmode=disable")
	utee.Chk(err)
	DB.LogMode(true)

	if DB.CreateTable(&DataInfo{}) == nil {
		log.Fatal("Can't create table")
	}
}

type DataInfo struct {
	Id   int
	Name string
	Language  string
}

func main() {
	sb:=Figo.NewSqlBuffer()
	sb.Append("INSERT INTO data_infos(name,language) VALUES ")
	for i:=0;i<1000;i++ {
		if i>0 {
			sb.Append(" , ")
		}
		sb.Append(" (?,?) ",fake.FullName(),fake.Language())
	}
	DB.Exec(sb.SQL(),sb.Params()...)
}
