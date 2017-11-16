package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/figoxu/Figo"
	"log"
)

// 可通过openssl产生
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`
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
`)

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCrGh1sc5AKD1EQ8WdA1iWF4m7w
XtO6WoS7Dtfd0Jm2ud+LKBQ+e7R6YIXnwfEKB/4Jm+jNtCi7/Zrx5gtEpUuVAyrE
o5+qr5al5KibeJq3xyI/626IBsDMFX5o3WOoXceTF7+lgi6r+OuokqFJgpeh7YAN
XQ8Y8mn8ucw+Ly+LbQIDAQAB
-----END PUBLIC KEY-----
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func main() {
	data, _ := RsaEncrypt([]byte("test dataΩ......"))
	fmt.Println(base64.StdEncoding.EncodeToString(data))

	log.Println(Figo.Bh.BStr(data))
	origData, _ := RsaDecrypt(data)
	fmt.Println(string(origData))
}