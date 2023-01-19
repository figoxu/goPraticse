package main

import (
	"github.com/quexer/utee"
	"github.com/robertkrimen/otto"
	"log"
)

func main() {
	log.Println("hello")
	vm := otto.New()
	vm.Run(`
    abc = 2 + 2;
	console.log("The value of abc is " + abc); // 4
`)
	value, err := vm.Get("abc")
	utee.Chk(err)
	v, _ := value.ToInteger()
	log.Println("Value Get From VM  is :", v)
}
