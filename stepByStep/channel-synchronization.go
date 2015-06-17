package main

import "fmt"
import "time"

func worker(done chan bool){
	fmt.Print("working....")
	time.Sleep(time.Second)
	fmt.Println("done")

	done<-true
}

//通过通道来同步只写。使用阻塞接收等待一个goroutine里完成的一个例子
func main(){
	done := make(chan bool,1)
	//运行一个goruntime
	go worker(done)

	//等待goruntime的回执
	val := <- done

	fmt.Println(val)
}