package main

import (
	"github.com/figoxu/utee"
	"github.com/yuin/gopher-lua"
	"log"
)

func main() {
	log.Println("Hello")

	L := lua.NewState()
	defer L.Close()
	err := L.DoString(`print("hello by lua")`)
	utee.Chk(err)
}
