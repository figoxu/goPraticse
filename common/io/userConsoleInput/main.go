package main

import (
	"github.com/figoxu/Figo"
	"log"
)

func main() {
	name := Figo.ReadInput("What's your name ?")
	log.Println("My Name Is ", name)
}
