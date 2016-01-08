package main
import (
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
		log.Println(i)
	}
	log.Println("---------------------------")
	for i:=0;i<=100;i++{
		if i%2==1 {
			continue
		}
		log.Println(i)
	}
	log.Println("#####################")
	for i:=0;i<=100;i++{
		if i>=20 {
			break
		}
		log.Println("@index:",i)
	}
}
