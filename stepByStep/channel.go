package main

import "fmt"

func main(){
	//创建一个chan  语法：make(chan val-type)
	message := make(chan string)

	//往 message 里面写数据
	go func(){ message <- "ping" }()

	//从 message里面读数据
	msg := <-message
	fmt.Println(msg)
}
