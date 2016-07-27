package main

import (
	"github.com/zhenjl/bloom/standard"
	"github.com/zhenjl/cityhash"
	"log"
	"fmt"
)

func main() {
	var capSize uint = 1000000000
	filter := standard.New(capSize)
	filter.SetHasher(cityhash.New64())
	v := []byte("Love")
	b := filter.Add(v).Check(v)
	log.Println("check @v:", b)
	for i:=0;i<100000000;i++ {
		d := []byte( fmt.Sprint("data",i) )
		if filter.Check(d) ==true{
			panic(fmt.Sprint("should not exist @d:",string(d)))
		}
		if flag := filter.Add(d).Check(d) ;flag==false {
			panic(d)
		}
	}
}
