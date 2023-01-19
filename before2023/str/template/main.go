package main

import (
	"github.com/cloudfly/template"
	"fmt"
)

type Sample struct {
	State string
	Id    int
}
func (p Sample) Value(key string) interface{} {
	if key=="State" {
		return p.State
	}
	if key=="Id"{
		return p.Id
	}
	return ""
}

func main() {
	v1 := Sample{
		State:"Hello",
		Id:10000001,
	}
	result, err := template.Parse("http://www.baidu.com/{{.State}}/open/{{.Id}}", v1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.(string))

}
