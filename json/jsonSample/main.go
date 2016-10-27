package main

import (
	"encoding/json"
	"github.com/figoxu/utee"
	"log"
	"reflect"
)

func main() {
	a := make(map[string]interface{})
	val := `{
"a":1,
"b":"bilibilibibi",
"c":1.02,
"d":false
}`
	utee.Chk(json.Unmarshal([]byte(val), &a))
	log.Println("Hello")
	log.Println(a)

	for k, v := range a {
		log.Println("@k:", k, " @v:", v, "  @type:", reflect.TypeOf(v))
	}
}
