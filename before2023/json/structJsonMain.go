package main
import (
	"log"
	"encoding/json"
	"fmt"
)


func main(){

	//特别注意,结构体的属性要大写开头，否则跨包，取不到
	type Result struct {
		Item  string `json:"item"`
		Value int   `json:"value"`
	}
	//	a := make([]*B , 0, len(m))
	var a []Result
	a = append(a, Result {
		Item:  "this is a test",
		Value: 100,
	})
	a = append(a, Result {
		Item:  "this is a test",
		Value: 100,
	})
	a = append(a, Result {
		Item:  "this is a test",
		Value: 100,
	})
	log.Println("result array:",a)
	val, err := json.Marshal(&a)
	if err!=nil {
		fmt.Println("@err:",err)
	}
	log.Println("json:",string(val))


	b,_:=json.Marshal(Result {
		Item:  "this is a test",
		Value: 100,
	})
	log.Println(string(b))

	log.Println("----step 2--------")
	type R2 struct {
		Name string `json:"name,omitempty"`
		Age  int `json:"age,omitempty"`
		Young bool `json:"young,omitempty"`
	}
	v0 := R2{}
	b,_= json.Marshal(&v0)
	log.Println(string(b))
	v0.Age = 0
	v0.Young = false
	b,_= json.Marshal(&v0)
	log.Println(string(b))
	 m   :=  map[string]interface{}{}
	m["pl"]=v0
	m["hello"]="world"
	b,_= json.Marshal(&m)
	log.Println(string(b))


}