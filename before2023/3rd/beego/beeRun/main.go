package main
import (
	"log"
	"time"
)

func main(){
	for {
		log.Println("hello world ")
		time.Sleep(time.Duration(3)*time.Second)
	}
}
