package main

import (
	"github.com/quexer/utee"
	"log"
	"net/url"
	"sync"
)

var wg sync.WaitGroup

func main() {
	log.Println("hello")
	t := 10000

	st := utee.Tick()
	for i := 0; i < t; i++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
	cost := utee.Tick() - st
	log.Println("@times:", t, " invoke cost:", cost, " second")

}

func test() {
	defer wg.Done()
	if _, err := utee.Http2Post("https://127.0.0.1:10443/helloPost", url.Values{}); err != nil {
		log.Println("utee h2 post @err:", err)
	}
	//else {
	//		log.Println("utee h2 post @rsp:", string(v))
	//	}
}
