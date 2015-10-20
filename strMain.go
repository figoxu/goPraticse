package main
import (
	"log"
	"bytes"
)

func main(){
	buf := bytes.NewBufferString("hello")
	buf.WriteString(" world")
	log.Println(buf.String())
}