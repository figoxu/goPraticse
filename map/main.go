package main

import "log"

func main(){
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
	for v := range m {
		log.Println(v)
	}
}
