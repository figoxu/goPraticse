package main
import (
	"fmt"
	"log"
)

func main(){
	lb:
	for i:=0;i<=100;i++{
		log.Println(" i/3=",(i/3)," @i:",i)
		log.Println(" i/3+1=",(i/3+1)," @i:",i)
		if i%2==0 {
			continue lb
		}
		fmt.Println(i)
	}
}
