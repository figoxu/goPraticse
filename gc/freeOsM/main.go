package main

import (
	"fmt"
	"log"
	"sync"
	_ "expvar"
	_ "net/http/pprof"
	"net/http"
)

//go tool pprof -alloc_space http://localhost:6666/debug/pprof/heap
//go tool pprof -inuse_space http://localhost:6666/debug/pprof/heap
func main() {
	go makeMem()
	log.Fatal(http.ListenAndServe(":6666", nil))
}

func makeMem() {
	log.Println("begin genrate garbage")
	c := 0
	for {
		var wg sync.WaitGroup
		for i := 0; i < 33; i++ {
			wg.Add(1)
			go generateGarbage(1, &wg)
		}
		wg.Wait()
		c++
		log.Println("execute @c:", c)
	}
}

type Garbage struct {
	content string
}

func generateGarbage(seed int, wg *sync.WaitGroup) []Garbage {
	defer wg.Done()
	gs := []Garbage{}
	for i := 0; i < 1000000; i++ {
		gs = append(gs, Garbage{content: fmt.Sprint("i am gabadge No:", i, " @seed:", seed)})
	}
	return gs
}
