package main

import "fmt"

func main() {
	bufSize := 1 // Try changing this to 0, and see what happens!
	ch1, ch2, ch3 := make(chan int, bufSize), make(chan int, bufSize), make(chan int, bufSize)
	go func() {
		defer close(ch1)
		defer close(ch2)
		defer close(ch3)
		ch1 <- 1
		ch2 <- 2
		ch3 <- 3
		ch3 <- 6
		ch2 <- 5
		ch1 <- 4 // Notice the reversed order of sends(!)
	}()
	for rs, oks := readOnce(ch1, ch2, ch3); allTrue(oks); rs, oks = readOnce(ch1, ch2, ch3) {
		fmt.Println(rs)
	}
}

// -- Helper funcs --------------------------------------------------------------------------------

func readOnce(chs ...chan int) ([]int, []bool) {
	rs := []int{}
	oks := []bool{}
	for _, ch := range chs {
		r, ok := <-ch
		rs = append(rs, r)
		oks = append(oks, ok)
	}
	return rs, oks
}

func allTrue(vs []bool) bool {
	for _, v := range vs {
		if !v {
			return false
		}
	}
	return true
}
