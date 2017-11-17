package main

import (
	"log"
	"github.com/quexer/utee"
	"github.com/figoxu/Figo"
)

func main(){
	log.Println("Hello")
	pwd:=[]byte{206,157,169,84,64,190,253,50,154,107,154,136,199,175,89,59}
	bs,err:=encrypt([]byte("问世间，是否比此删更高"),pwd)
	utee.Chk(err)
	log.Println(Figo.Bh.BStr(bs))
	bs,err = decrypt(bs,pwd)
	log.Println(string(bs))
}