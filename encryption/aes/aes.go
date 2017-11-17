package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"bytes"
	"io"
	"crypto/rand"
)

//cbc encrypt with PKCS5Padding
func encrypt(data,pwd []byte) ([]byte, error) {
	blk, err := aes.NewCipher(pwd)
	if err != nil {
		return nil, err
	}
	data = PKCS5Padding(data, aes.BlockSize)

	dst := make([]byte, aes.BlockSize+len(data))
	iv := dst[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(blk, iv)
	mode.CryptBlocks(dst[aes.BlockSize:], data)

	return dst, nil
}

//cbc decrypt with PKCS5Padding
func decrypt(data,pwd []byte) ([]byte, error) {
	blk, err := aes.NewCipher(pwd)
	if err != nil {
		return nil, err
	}

	if len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("bad ciphertext fmt")
	}

	if len(data) <= aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(blk, iv)
	mode.CryptBlocks(data, data)
	data = PKCS5UnPadding(data)
	if len(data) == 0 {
		return nil, fmt.Errorf("bad unpadding value")
	}
	return data, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err, " (recover)")
		}
	}()
	d := []byte{}
	length := len(src)
	if length < 1 {
		return d
	}
	unpadding := int(src[length-1])
	if unpadding > length {
		return d
	}
	d = src[:(length - unpadding)]
	return d
}
