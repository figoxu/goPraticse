package main

import "fmt"

//只写
func ping(pings chan<- string,msg string){
	pings <- msg
}

//从只写的channel里面，读取到只读的channel里
func pong(pings <-chan string,pongs chan <- string){
	msg:=<-pings
	pongs<-msg
}

func main(){
	pings := make(chan string,1)
	pongs := make(chan string,1)
	ping(pings,"passed message")
	pong(pings,pongs)
	fmt.Println(<-pongs)
}

