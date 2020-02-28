package main

import (
	"log"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type DataItem struct {
	V int
}

func main() {
	var vs []*DataItem
	for i := 0; i < 100; i++ {
		vs = append(vs, &DataItem{V: i})
	}

	var limitChan = make(chan struct{}, 10)
	var waitGroup sync.WaitGroup
	for _, v := range vs {
		limitChan <- struct{}{}
		waitGroup.Add(1)
		item := v // 多一行赋值，可以去除线程安全问题
		go func() {
			defer waitGroup.Done()
			defer func() {
				<-limitChan
			}()
			log.Println(item.V)
			item.V = item.V + 100
			time.Sleep(time.Second * time.Duration(2))
		}()
	}
	waitGroup.Wait()
	logrus.Println("-------")
	for _, v := range vs {
		log.Println(v.V)
	}
}
