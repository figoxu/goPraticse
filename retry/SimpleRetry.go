package main
import (
	"fmt"
	"log"
	"errors"
)

func main(){
	fmt.Println("hello")


	err := someBusiness()
	retry :=0
	for err != nil && retry<3 {
		retry ++
		err = someBusiness()
		log.Println("no dv stat @retry:",retry)
	}

}

func someBusiness() error{
	return  errors.New("some err")
}
