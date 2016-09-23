package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, die chan bool) {
	defer wg.Done()
	i := 0
	for {

		select {
		case v := <-die:
			log.Println("@v i receive:", v)
			return
		default:
			i++
			fmt.Print(".")
			if i == 10 {
				fmt.Println("")
			}
		}
	}
}

func main() {
	die := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(&wg, die)
	}
	time.Sleep(time.Duration(2) * time.Second)
	log.Println("Let's ShutDown")
	close(die)
	log.Println("Wait For Over")
	wg.Wait()
	log.Println("Over")
}
