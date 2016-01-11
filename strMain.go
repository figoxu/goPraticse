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
	var v int
	log.Println(v)
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


	log.Println(strings.ToUpper("hello world"))
	a:=[]string{"hello","world"}
	for i,v := range a{
		log.Println("@i:",i," @v:",v)
	}
	va :="hi"
	if va=="hi" {
		log.Println("equals is ==")
	}


	s := "PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9InllcyI/Pgo8YmlzUmVzcG9uc2VEVE8+CiAgICA8cmVzcENvZGU+MTAzMDwvcmVzcENvZGU+CiAgICA8cmVzcEluZm8+5pWw5a2X562+5ZCN6ZSZ6K+vPC9yZXNwSW5mbz4KICAgIDxyZXNwVGltZT4xNDQ3MzgyMDU5NjEyPC9yZXNwVGltZT4KPC9iaXNSZXNwb25zZURUTz4K"
	b,e := base64.StdEncoding.DecodeString(s)
	log.Println("@e:",e,"@v:",string(b))
	 vs := strconv.Itoa(1024)
	log.Println(vs)
	vi,_:=strconv.Atoi(vs)
	log.Println("@vi+1:",(vi+1) )


	mt,err :=strconv.ParseInt("1429345203123",10,64)
	log.Println("@mt:",mt,"@err:",err," @len:",len("testValueIs"))
	log.Println(len("this_is_a_very_long_string__we_just_test_skip_func_is_usefull_does_it_usefull_yes_it_is"))



	ar := strings.Split("hello,world,how,cool,me",",")
	log.Println("@len:",len(ar))
}