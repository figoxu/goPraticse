package main
import (
	"fmt"
	"io/ioutil"
	"golang.org/x/crypto/pkcs12"
	"encoding/pem"
	"crypto/tls"
	"bytes"
	"log"
	"crypto/x509"
	"crypto/x509/pkix"
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
	log.Println("------------------")
//	fmt.Println("-- cer buffer is -")
//	fmt.Println("------------------")
//	fmt.Println(buf.String())
//	fmt.Println("------------------")
	block, _ := pem.Decode([]byte(buf.String()))
	if block == nil {
		log.Println("failed to parse certificate PEM")
	}
	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Println("failed to parse certificate: " + err.Error())
	}
	log.Println("------------------")
	log.Println("@commonName:",c.Subject.CommonName)
	log.Println("------------------")
	print(c.Subject)
	log.Println("------------------")
	print(c.Issuer)
}

func print(info pkix.Name){
	log.Println("@country:",info.Country)
	log.Println("@Organization:",info.Organization)
	log.Println("@OrganizationalUnit:",info.OrganizationalUnit)
	log.Println("@Locality:",info.Locality)
	log.Println("@Province:",info.Province)
	log.Println("@StreetAddress:",info.StreetAddress)
	log.Println("@PostalCode:",info.PostalCode)
	log.Println("@SerialNumber:",info.SerialNumber)
	log.Println("@CommonName:",info.CommonName)


	log.Println("@Names:",info.Names)
	log.Println("@ExtraNames:",info.ExtraNames)
}