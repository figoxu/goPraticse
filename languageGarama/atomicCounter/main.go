package main

import "fmt"
import "time"
import "sync/atomic"
import (
	"runtime"
	"log"
)

func main() {

	type S struct{

		v uint8
	}
	s := &S{}
	log.Println("@version:"+fmt.Sprint(s.v))

	// We'll use an unsigned integer to represent our
	// (always-positive) counter.
	var ops uint64 = 0

	// To simulate concurrent updates, we'll start 50
	// goroutines that each increment the counter about
	// once a millisecond.
	for i := 0; i < 50; i++ {
		go func() {
			for j:=0;j<100;j++ {
				// To atomically increment the counter we
				// use `AddUint64`, giving it the memory
				// address of our `ops` counter with the
				// `&` syntax.
				atomic.AddUint64(&ops, 1)

				// Allow other goroutines to proceed.
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second*time.Duration(10))
	fmt.Println("@i:",ops)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}