package main

import (
	"log"
	"github.com/quexer/utee"
	"github.com/figoxu/Figo"
)

func main(){
	log.Println("Hello")
	pwd:=[]byte{208,130,159,99,231,235,183,11,143,170,60,9,183,15,29,171}
	bs,err:=encrypt([]byte("问世间，是否比此删更高"),pwd)
	utee.Chk(err)
	log.Println(Figo.Bh.BStr(bs))
	bs,err = decrypt(bs,pwd)
	log.Println(string(bs))

	bs=[]byte{35,150,167,147,116,104,115,230,99,200,74,183,227,68,137,100,20,250,140,155,53,62,197,37,176,201,173,105,156,20,204,165,238,30,192,240,189,25,134,67,163,214,31,118,93,186,191,76}
	bs,err=decrypt(bs,pwd)
	utee.Chk(err)
	log.Println("解密后")
	log.Println(Figo.Bh.BStr(bs))
	log.Println(string(bs))
}