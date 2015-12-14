package main
import (
	"fmt"
	"log"
)


func main(){
	fmt.Println("hello")




	fvv := []string{}
	fvv = append(fvv,"hello")
	fvv = append(fvv,"how")
	fvv = append(fvv,"are")
	fvv= append(fvv,"you")
	fvv = append(fvv,"i am")
	fvv = append(fvv,"fine")
	fvv = append(fvv,"thank")
	fvv = append(fvv,"you")
	log.Println(fvv)


	fv2 := fvv
	fvv =[]string{}

	fmt.Println("@fv2:",fv2)
	fmt.Println("@fvv:",fvv)
	a :=[]string{"1","2","3"}
	log.Println("@a:",a)
	log.Println(a[0])
}
