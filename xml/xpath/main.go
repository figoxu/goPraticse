package main
import (
	"fmt"
	"bytes"
	"gopkg.in/xmlpath.v1"
	"log"
)


func main(){


	xs := `
	<bisResponseDTO>
    <respCode>00001</respCode>
    <respInfo>OK</respInfo>
    <respTime>1447386329994</respTime>
</bisResponseDTO>`
	node, err := xmlpath.Parse(bytes.NewBuffer( []byte(xs)))
	if err!=nil {
		log.Println("@err:",err)
	}

	path := xmlpath.MustCompile("/bisResponseDTO/respCode")

	v,f := path.String(node)
	if !f {
		log.Println("@f:",f)
	}
	fmt.Println(v)
	if v!="0000"{
		fmt.Println("failure")
	}
}
