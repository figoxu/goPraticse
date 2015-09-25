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
}