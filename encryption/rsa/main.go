package main

import (
	"encoding/base64"
	"fmt"
	"github.com/figoxu/Figo"
	"log"
)

func main() {
	data, _ := PubEncrypt([]byte("test dataΩ......"))
	fmt.Println(base64.StdEncoding.EncodeToString(data))

	log.Println(Figo.Bh.BStr(data))
	origData, _ := PriDecrypt(data)
	fmt.Println(string(origData))
}