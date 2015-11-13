package main
import (
	"log"
	"bytes"
	"github.com/quexer/utee"
	"strconv"
	"fmt"
	"strings"
	"encoding/base64"
)

func main(){
	for i:=0;i<5;i++{
		fmt.Println(i)
	}

	t:=utee.Tick()
	log.Println("@t:",t)
	buf := bytes.NewBufferString("hello")
	buf.WriteString(" world @time:"+strconv.FormatInt(t,10))
	log.Println(buf.String())
	if v,e :=strconv.Atoi("");e!=nil{
		log.Println("@err:",e," v:",v)
	}else{
		log.Println(v)
	}


	fmt.Println(strings.ToUpper("hello world"))
	a:=[]string{"hello","world"}
	for i,v := range a{
		fmt.Println("@i:",i," @v:",v)
	}
	va :="hi"
	if va=="hi" {
		fmt.Println("equals is ==")
	}


	s := "PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9InllcyI/Pgo8YmlzUmVzcG9uc2VEVE8+CiAgICA8cmVzcENvZGU+MTAzMDwvcmVzcENvZGU+CiAgICA8cmVzcEluZm8+5pWw5a2X562+5ZCN6ZSZ6K+vPC9yZXNwSW5mbz4KICAgIDxyZXNwVGltZT4xNDQ3MzgyMDU5NjEyPC9yZXNwVGltZT4KPC9iaXNSZXNwb25zZURUTz4K"
	b,e := base64.StdEncoding.DecodeString(s)
	fmt.Println("@e:",e,"@v:",string(b))
}