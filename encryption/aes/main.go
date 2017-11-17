package main

import (
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"log"
)

func main() {
	pwd := []byte{208, 130, 159, 99, 231, 235, 183, 11, 143, 170, 60, 9, 183, 15, 29, 171}
	aesHelp := Figo.NewAesHelp(pwd)
	bs, err := aesHelp.Encrypt([]byte("问世间，是否比此删更高"))
	utee.Chk(err)
	log.Println(Figo.Bh.BStr(bs))
	bs, err = aesHelp.Decrypt(bs)
	log.Println(string(bs))

	bs = []byte{35, 150, 167, 147, 116, 104, 115, 230, 99, 200, 74, 183, 227, 68, 137, 100, 20, 250, 140, 155, 53, 62, 197, 37, 176, 201, 173, 105, 156, 20, 204, 165, 238, 30, 192, 240, 189, 25, 134, 67, 163, 214, 31, 118, 93, 186, 191, 76}
	bs, err = aesHelp.Decrypt(bs)
	utee.Chk(err)
	log.Println("解密后")
	log.Println(Figo.Bh.BStr(bs))
	log.Println(string(bs))
}
