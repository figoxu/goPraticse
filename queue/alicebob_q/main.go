package main

import(
	libq "github.com/alicebob/q"
	"os"
	"fmt"
	"log"
	"sync"
	"strings"
	"github.com/quexer/utee"
)

func main(){
	// Read and write a lot of messages, as fast as possible.
	i:=utee.Fint("1000000")
	fmt.Println(i)
	i++
	fmt.Println(i)
	fmt.Println("---------------")

	eventCount := 1000000
	clients := 10

	d := setupDataDir()
	q, err := libq.NewQ(d, "events", libq.BlockCount(1000))
	if err != nil {
		log.Fatal(err)
	}
	defer q.Close()
	wg := sync.WaitGroup{}
	payload := strings.Repeat("0xDEADBEEF", 30)

	for i := 0; i < clients; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < eventCount/clients; j++ {
				q.Enqueue(payload)
			}
		}()
	}
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range make([]struct{}, eventCount-5001) {
			// The inserts are non-derministic, so we can't have an interesting
			// payload.
			if got := <-q.Queue(); payload != got {
				log.Fatalf("Want for %d: %#v, got %#v", i, payload, got)
			}
		}
	}()
	wg.Wait()
}


func setupDataDir() string {
	os.RemoveAll("./d")
	if err := os.Mkdir("./d/", 0700); err != nil {
		panic(fmt.Sprintf("Can't make ./d/: %v", err))
	}
	return "./d"
}