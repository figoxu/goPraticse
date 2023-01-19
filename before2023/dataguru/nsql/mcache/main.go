package main

import (
	"database/sql"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/figoxu/utee"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"time"
)

func main() {
	log.Println("@val:", getVal())
	time.Sleep(time.Second * time.Duration(10))
	log.Println("@val:", getVal())
	log.Println("@val:", getVal())
}

func getVal() int {
	mc := memcache.New("192.168.56.101:11211")
	k := "Test"
	it, err := mc.Get(k)
	log.Println("@it:", it, "    @err:", err)
	if err == memcache.ErrCacheMiss {
		v := getCountValFromDb()
		mc.Set(&memcache.Item{Key: k, Value: []byte(strconv.Itoa(v))})
		log.Println("get from db")
		return v
	} else {
		utee.Chk(err)
	}
	log.Println("get from cache")
	v, err := strconv.Atoi(string(it.Value))
	utee.Chk(err)
	return v
}

func getCountValFromDb() int {
	pwd := "pwd"
	host := "127.0.0.1"
	url := fmt.Sprint("root:", pwd, "@tcp(", host, ":3306)/test?charset=utf8")

	db, err := sql.Open("mysql", url)
	utee.Chk(err)
	defer db.Close()
	rows, err := db.Query("select count(*) from mtest")
	utee.Chk(err)
	if rows.Next() {
		var id int
		err = rows.Scan(&id)
		utee.Chk(err)
		return id
	}
	return 0
}
