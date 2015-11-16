package main
import (
	"fmt"
	"io/ioutil"
	"golang.org/x/crypto/pkcs12"
	"encoding/pem"
	"crypto/tls"
	"bytes"
)

func main(){

	b,e:= ioutil.ReadFile("/home/figo/data/testCer/Certificates.p12")
	if e!=nil {
		fmt.Println("@error:",e)
		return
	}
	blocks, err := pkcs12.ToPEM(b,"111");
	if err!=nil {
		fmt.Println("@error:",err)
		return
	}
	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		panic(err)
	}



	fmt.Println("success @cert:",cert)
	buf := bytes.NewBufferString("")
	for _, b := range blocks {
		pem.Encode(buf,b)
	}

	fmt.Println("------------------")
	fmt.Println("-- cer buffer is -")
	fmt.Println("------------------")
	fmt.Println(buf.String())
	fmt.Println("------------------")
	fmt.Println("hello")
}
