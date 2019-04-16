package main

import (
	"errors"
	"github.com/figoxu/goPraticse/err/a"
	"log"
)

func main() {
	log.Println("hello")
	e := errors.New("test")
	log.Println(e.Error())
	log.Println(e.Error() == "test")
	a.TMethod()
}
