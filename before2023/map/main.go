package main

import (
	"log"
	"encoding/json"
	"github.com/figoxu/utee"
)

var bm map[string]interface{} //key : made

type Info struct {
	Made     string `json:"made"`					//首字母需要大写
	Daily    int    `json:"daily"`
	Monthly  int    `json:"monthly"`
	Total    int    `json:"total"`
	UnqTotal int    `json:"unqTotal"`
}

func main() {
	m := make(map[string]string)
	m["hello"] = "echo hello"
	m["world"] = "echo world"
	m["go"] = "echo go"
	m["is"] = "echo is"
	m["cool"] = "echo cool"

	for k, v := range m {
		log.Printf("k=%v, v=%v\n", k, v)
	}
	log.Println("---------------")
	for k := range m {
		log.Println(k)
	}

	bm = make(map[string]interface{})
	bm["a"] = Info{
		Made:"test",
		Daily:100,
	}
	b,e:=json.Marshal(bm)
	utee.Chk(e)
	log.Println(string(b))
}
