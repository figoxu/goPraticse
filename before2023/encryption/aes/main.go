package main

import (
	"encoding/base64"
	"encoding/hex"
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

	key, err := hex.DecodeString("535b02b78c856fe97c1d05fd9c177742")
	utee.Chk(err)
	iv, err := hex.DecodeString("c6997d77fdbff4dcb6bc3afc559bd6b5")
	aesHelp = Figo.NewAesHelp(key, iv...)
	content1, err := hex.DecodeString("2061F59B269FCA6535886706FCB82CF25FCC891F4AC116D7A4B33110E14F0DC857D2C3D44554B0AFBAA84C034F3AC7FD731B7B301AA1727F53CD76688DE9A1C99424081A3D5CADD5471871EC4E6284215AFDEF4A66548A44C226D9ED196CF029")
	utee.Chk(err)
	bs, err = aesHelp.Decrypt(content1)
	utee.Chk(err)
	log.Println(Figo.Bh.BStr(bs))
	log.Println(string(bs))

	content2, err := base64.StdEncoding.DecodeString(`IGH1myafymU1iGcG/Lgs8l/MiR9KwRbXpLMxEOFPDchX0sPURVSwr7qoTANPOsf9cxt7MBqhcn9TzXZojemhyZQkCBo9XK3VRxhx7E5ihCFa/e9KZlSKRMIm2e0ZbPAp`)
	utee.Chk(err)
	bs, err = aesHelp.Decrypt(content2)
	log.Println(Figo.Bh.BStr(bs))
	log.Println(string(bs))
}
