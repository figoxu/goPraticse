package main
import (
	"fmt"
	"encoding/hex"
)

func main(){
	msg := "----"
	if v,e:=hex.DecodeString("----");e!=nil {
		fmt.Println("@e:",e)
	}else{
		fmt.Println("@v:",v)
	}
	v := fmt.Sprintf("%x",msg)
	fmt.Println("@val:",v)
}
