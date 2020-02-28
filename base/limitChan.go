package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	var limitChan = make(chan struct{}, 10)
	var waitGroup sync.WaitGroup
	for i := 0; i < 100; i++ {
		limitChan <- struct{}{}
		waitGroup.Add(1)
		j := i
		go func() {
			defer waitGroup.Done()
			defer func() {
				<-limitChan
			}()
			log.Println(j)
			time.Sleep(time.Second * time.Duration(2))
		}()
	}
	waitGroup.Wait()
}
