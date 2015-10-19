package main
import (
	"fmt"
	"time"
	"log"
)


func main(){
	fmt.Println("hello")
	now := time.Now();
	now.Format("2006010215");

	for i:=0;i<1000;i++{
//		now = now.Add(time.Hour*time.Duration(1))
		now = now.Add(time.Hour*time.Duration(24))
		v := now.Format("2006010215");
		log.Println("@val:",v)
	}
}
