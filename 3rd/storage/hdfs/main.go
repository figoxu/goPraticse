package main

import (
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"log"
)

func main() {
	defer Figo.Catch()
	hdfsClient := Figo.NewHDFSClient("192.168.108.131:9000", "root")
	fullPath := "/figo/github/foo.txt"
	hdfsClient.Write(fullPath, []byte("Here We Go,Frightting To Eat"))
	v, e := hdfsClient.Read(fullPath)
	utee.Chk(e)
	log.Println("@value I Read is :", string(v))
}
