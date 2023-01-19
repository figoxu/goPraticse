package main

import (
	"log"
	"time"
)

func main() {
	defer log.Println("over")

	wasteTimeChan := make(chan string ,10)
	go func(){
		for {
			time.Sleep(10*time.Second)
			wasteTimeChan<-"hello"
		}
	}()
	timeout := time.After(5 * time.Second)
	for {
		select {
		case v := <-wasteTimeChan:
			log.Println("method waste of time @v:",v)
		case <-timeout:
			return
		}
	}

}
