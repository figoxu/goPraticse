package main
import (
	"log"
	"fmt"
)

func main(){
	tmpl(ma,10)
	log.Println("--------------------")
	tmpl(mb,10)
}


func ma(val string)(int,error){
	log.Println("ma .... ",val)
	return 0,nil
}

func mb(val string)(int,error){
	log.Println("mb ### ",val)
	return 0,nil
}

func tmpl( a func(string)(int,error),t int){
	for i:=0;i<t;i++ {
		a(fmt.Sprint("test:",i))
	}
}