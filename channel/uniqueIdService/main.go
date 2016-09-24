package main

import (
	"fmt"
	"log"
)

func main() {
	id:=make(chan string)
	go func(){
		var counter int64 = 0
		for{
			id <- fmt.Sprintf("%x",counter)
			counter+=1
		}
	}()
	x:=<-id
	log.Println("@x:",x)
	x=<-id
	log.Println("@x:",x)
}
