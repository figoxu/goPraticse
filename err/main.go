package main
import (
	"log"
	"github.com/gogap/errors"
)

func main(){
	log.Println("hello")
	e:=errors.New("test")
	log.Println(e.Error())
	log.Println(e.Error()=="test")
}
