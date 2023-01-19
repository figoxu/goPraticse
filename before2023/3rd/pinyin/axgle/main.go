package main

import (
	"github.com/axgle/pinyin"
	"log"
)

//Support Cross Compile
func main(){
	log.Println(pinyin.Convert("世界你好"))
}
