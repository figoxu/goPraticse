package main

import (
	"github.com/figoxu/Figo"
	"log"
	"github.com/quexer/utee"
)

// 可通过openssl产生
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey4Test = `
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCrGh1sc5AKD1EQ8WdA1iWF4m7wXtO6WoS7Dtfd0Jm2ud+LKBQ+
e7R6YIXnwfEKB/4Jm+jNtCi7/Zrx5gtEpUuVAyrEo5+qr5al5KibeJq3xyI/626I
BsDMFX5o3WOoXceTF7+lgi6r+OuokqFJgpeh7YANXQ8Y8mn8ucw+Ly+LbQIDAQAB
AoGAGgoxbC3yP/WwyrlSk4WD1Gpvo9lqs7PO+4D4zWNP4YVMRitlWVUOVImYF3tm
qbYprWCy/4tpn6KrECGImXvmkplXPxd4x3W+haZftx3VjTwh5fvT9yHp4swXxN+h
LMItDdIOWS4U6wVJa77Dy7VfK303LZrPLqnxkf4oEywp5YECQQDZOz1WD7nOqOiy
AlwDhfeLTmArN0f+gV6RLrxMp2XRqC2DN5nMq5O5BVVMK9LBgArNqYfxWYuMa3K2
qliRDPPxAkEAyaNWq/fDvjpK9TgztqsHIiG+cUQpWI759zt5qHNA+QF4L43dtAVZ
zBR/uam1jnRuM6K0ZCSZo2ITiqapmk8bPQJAEd9d3IbOssIS4xJun5uWElAQeX3C
3p2mOiuuMmBTcDx2AiXA8aXsMXzO18WDQYhXWzRniuPjJ1pvxbeeMdDvAQJBAMDh
uZAJEzrOAlQurfFICyvQQZ+Rx0dKhbzFLOxBS96mVDSRLYn+MFbzKPcOa3lY0O4d
7xd4l2td7zmLkePlVjUCQQCY8VuIfKc0+AWvPnktKXbx9bBdJZSDginZM5cu7pdx
W0uB9KZoLqgbGLIvWrLyA6SBqo87Q1j1//wFgLP+A2Gn
-----END RSA PRIVATE KEY-----
`

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey4Test = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrGh1sc5AKD1EQ8WdA1iWF4m7w
XtO6WoS7Dtfd0Jm2ud+LKBQ+e7R6YIXnwfEKB/4Jm+jNtCi7/Zrx5gtEpUuVAyrE
o5+qr5al5KibeJq3xyI/626IBsDMFX5o3WOoXceTF7+lgi6r+OuokqFJgpeh7YAN
XQ8Y8mn8ucw+Ly+LbQIDAQAB
-----END PUBLIC KEY-----
`


func main() {
	rsaHelp := Figo.NewRsaHelp(publicKey4Test, privateKey4Test)

	bs, err := rsaHelp.PriEnc([]byte("hello figo"))
	utee.Chk(err)
	log.Println(Figo.Bh.BStr(bs))
	result, err := rsaHelp.PubDec(bs)
	utee.Chk(err)
	log.Println(string(result))

	bs, err = rsaHelp.PubEnc([]byte("nice world"))
	utee.Chk(err)
	log.Println(Figo.Bh.BStr(bs))
	result, err = rsaHelp.PriDec(bs)
	utee.Chk(err)
	log.Println(string(result))
}