package main

import (
	"github.com/figoxu/Figo"
	"log"
)

func main() {
	name := Figo.ReadInput("What's your name ?",Figo.THEME_Blue,Figo.THEME_Green)
	log.Println("My Name Is ", name)
}
