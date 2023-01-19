package main

import (
	"encoding/json"
	"fmt"
)

type Connector struct {
	Id     string
	Weight int
}

var (
	m = map[string]*Connector{}
)

func main() {
	append := func(id string, weight int) {
		m[id] = &Connector{
			Id:     id,
			Weight: weight,
		}
	}
	for i := 0; i < 100; i++ {
		v := fmt.Sprint("test", i)
		append(v, i)
	}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("@err:", err)
	}
	conf := string(b)

	fmt.Println("@val:" + conf)

	var result map[string]*Connector
	json.Unmarshal([]byte(conf), &result)



	fmt.Println("@result:",result)
	fmt.Println(result["test98"].Weight)
}
