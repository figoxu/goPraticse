package main

import (
	"database/sql"
	"fmt"
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"log"
	"time"
)

type HotColdInfo struct {
	Id          int
	IdentityNo  string
	Name        string
	Age         int
	Salary      int
	Description string
}

var rsWriter RedisqlWriter

func main() {
	log.Println("main")
	rp := Figo.RedisPool("127.0.0.1:6379", "")
	db, err := sql.Open("mysql", "user:password@/dbname")
	utee.Chk(err)
	rsWriter = NewRedisqlWriter(rp, db)

	for i := 0; i < 1000; i++ {
		info := HotColdInfo{
			Id:          i,
			IdentityNo:  fmt.Sprint("dv_", i),
			Name:        fmt.Sprint("name_", i),
			Age:         i,
			Salary:      i,
			Description: fmt.Sprint("description_", i),
		}
		rsWriter.write(info)
	}

	time.Sleep(time.Hour)
}
