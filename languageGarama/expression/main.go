package main
import (
	"log"
)

func main(){
	m :=map[string]int{}
	m["a"],m["b"],m["c"] = m["a"]+1,m["b"]+2,m["c"]+3
	m["a"],m["b"],m["c"] = m["a"]+4,m["b"]+5,m["c"]+6
	log.Println("@a:",m["a"],"@b:",m["b"],"@c",m["c"])
}
