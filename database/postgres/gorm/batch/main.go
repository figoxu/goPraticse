package main

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/quexer/utee"
	"github.com/figoxu/Figo"
	"time"
	"github.com/icrowley/fake"
)

var (
	DB *gorm.DB
	BMQ Figo.BatchMemQueue
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
	Id       int
	Name     string
	Language string
}

func main() {
	BMQ=Figo.NewBatchMemQueue(10e5, 6, 1000, bacthWorker)
	log.Println("Begin")
	for i:=0;i<10e5;i++ {
		BMQ.Enq(DataInfo{
			Name:fake.FullName(),
			Language:fake.Language(),
		})
	}
	log.Println("END")
	time.Sleep(time.Minute*time.Duration(2))
}

func bacthWorker(vs []interface{}) {
	defer Figo.Catch()
	if len(vs) <= 0 {
		return
	}
	sb := Figo.NewSqlBuffer()
	sb.Append("INSERT INTO data_infos(name,language) VALUES ")
	for i, v := range vs {
		data := v.(DataInfo)
		if i > 0 {
			sb.Append(" , ")
		}
		sb.Append(" (?,?) ", data.Name, data.Language)
	}
	DB.Exec(sb.SQL(), sb.Params()...)
	log.Println("SQL RUN")
}
