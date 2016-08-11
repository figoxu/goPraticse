package main
import (
	"log"
	"gopkg.in/mgo.v2"
	"github.com/figoxu/utee"
)

const mongoURI = "dev:rgP333g7f3@10.9.41.9,10.9.37.208,10.9.39.45/mpush?maxPoolSize=50"

func main(){
	log.Println("Hello")
	dailInfo,err := mgo.ParseURL(mongoURI)
	utee.Chk(err)
	log.Print("@DBName:",dailInfo.Database)
}
