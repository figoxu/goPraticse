package main

import "log"

func main() {
	log.Println("hello ", test())
}

func test() (result string) {
	result = "world"
	return
}
