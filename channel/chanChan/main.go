package main

import (
	"log"
	"fmt"
)

func req(msg string) chan string {
	resp := make(chan string,1)
	resp <-msg
	work(resp)
	return resp
}

func work(param chan string){
	v :=<- param
	param <- fmt.Sprint("Hello",v)
}

func main(){
	log.Println("result:",<-req(" figo"))
}
