package main

import (
	"encoding/json"
	"github.com/figoxu/utee"
	"github.com/jmoiron/jsonq"
	"log"
	"strings"
)

const (
	jsonstring = `{
		    "foo": 1,
		    "bar": 2,
		    "test": "Hello, world!",
		    "baz": 123.1,
		    "array": [
			{"foo": 1},
			{"bar": 2},
			{"baz": 3}
		    ],
		    "subobj": {
			"foo": 1,
			"subarray": [1,2,3],
			"subsubobj": {
			    "bar": 2,
			    "baz": 3,
			    "array": ["hello", "world"]
			}
		    },
		    "bool": true
		}`
)

func main() {
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonstring))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	// data["foo"] -> 1
	i, err := jq.Int("foo")
	utee.Chk(err)
	log.Println("data['foo'] -> 1 : ", i)
	i, err = jq.Int("subobj", "subarray", "1")
	utee.Chk(err)
	log.Println("data['subobj']['subarray'][1] -> 2: ", i)
	s, err := jq.String("subobj", "subsubobj", "array", "0")
	utee.Chk(err)
	log.Println("data['subobj']['subarray']['array'][0] -> 'hello' ", s)

	// data["subobj"] -> map[string]interface{}{"subobj": ...}
	obj, err := jq.Object("subobj")
	utee.Chk(err)
	log.Println("jq.Object('subobj') : ", obj)
}
