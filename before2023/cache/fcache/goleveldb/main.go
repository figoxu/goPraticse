package main

import (
	"github.com/figoxu/utee"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"fmt"
	"encoding/json"
	"os"
)


type Test struct {
	Name  string	 `json:"name"`
	Tp    string	 `json:"tp"`
	Count int	 `json:"count"`
}

func main() {
	err := os.MkdirAll("./gldb", 0777)
	utee.Chk(err)
	db, err := leveldb.OpenFile("./gldb", nil)
	utee.Chk(err)
	defer db.Close()

	err = db.Put([]byte("key"), []byte("value"), nil)
	utee.Chk(err)
	data, err := db.Get([]byte("key"), nil)
	utee.Chk(err)
	log.Println("@data:", string(data))
	iter := db.NewIterator(nil, nil)
	i :=0
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		log.Println("@key:", string(key), "@value:", string(value))
		i++
	}
	log.Println("read @i:",i)
	iter.Release()
	err = iter.Error()
	err = db.Delete([]byte("key"), nil)
	utee.Chk(err)

	log.Println("热身完毕")

	st := utee.Tick()
	for i:=0;i<10*10000;i++ {
		b, _ := json.Marshal(Test{Name:"figo",Tp:"android",Count:1024})
		db.Put([]byte(fmt.Sprint("test",i)),b,nil)
	}
	writeCost := utee.Tick()-st
	st = utee.Tick()
	iter = db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		log.Println("@key:", string(key), "@value:", string(value))
		err = db.Delete([]byte("key"), nil)
		utee.Chk(err)
	}
	log.Println("100,0000  read cost ",(utee.Tick()-st),"m second")
	log.Println("100,0000  write cost ",writeCost,"m second")

	iter.Release()
	if err := iter.Error(); err != nil {
		log.Println("iter @err:", err)
	}
	log.Println("finish")
}
