package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello World!")
	barrier := NewBarrier(3)
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("A")
			barrier.BarrierWait()
			fmt.Println("B")
			barrier.BarrierWait()
			fmt.Println("C")
		}()
	}
	wg.Wait()
}

type Barrier struct {
	curCnt int
	maxCnt int
	cond   *sync.Cond
}

func NewBarrier(maxCnt int) *Barrier {
	mutex := new(sync.Mutex)
	cond := sync.NewCond(mutex)
	return &Barrier{curCnt: maxCnt, maxCnt: maxCnt, cond: cond}
}

func (barrier *Barrier) BarrierWait() {
	barrier.cond.L.Lock()
	if barrier.curCnt--; barrier.curCnt > 0 {
		barrier.cond.Wait()
	} else {
		barrier.cond.Broadcast()
		barrier.curCnt = barrier.maxCnt
	}
	barrier.cond.L.Unlock()
}
