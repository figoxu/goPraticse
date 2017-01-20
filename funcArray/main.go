package main

import "fmt"

func main() {
	var fn [10]func(int)
	for i := 0; i < len(fn); i++ {
		fn[i] = make_fn()
	}
	for i, f := range fn {
		f(i)
	}
}
func make_fn() func(i int) {
	return func(i int) {
		fmt.Println(i)
	}
}
