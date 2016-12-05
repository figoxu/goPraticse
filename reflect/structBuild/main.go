package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
}
type Bar struct {
}

//用于保存实例化的结构体对象
var regStruct map[string]interface{}

func main() {
	str := "Bar"
	if regStruct[str] != nil {
		t := reflect.ValueOf(regStruct[str]).Type()
		v := reflect.New(t).Elem()
		fmt.Println(v)
		fmt.Println(v.Kind())
		fmt.Println(t.String())
	}

}

func init() {
	regStruct["Foo"] = Foo{}
	regStruct = make(map[string]interface{})
	regStruct["Bar"] = Bar{}
}
