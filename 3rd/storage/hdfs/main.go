package main

import (
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"log"
)

func main() {
	defer Figo.Catch()
	hdfsClient := Figo.NewHDFSClient("172.17.0.3:9000", "root")
	fullPath := "/user/root/foo.txt"
	err:=hdfsClient.Write(fullPath, []byte("Here We Go,Frightting To Eat"))
	utee.Chk(err)
	v, e := hdfsClient.Read(fullPath)
	utee.Chk(e)
	log.Println("@value I Read is :", string(v))
}
