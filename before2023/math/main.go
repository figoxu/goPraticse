package main
import (
	"strconv"
	"log"
	"fmt"
)

func main(){
	i := 10001 %100
	log.Println("@i:",i)


	v,err:=strconv.Atoi("111")
	log.Println("@v:",v," @err:",err)
	v,err=strconv.Atoi(fmt.Sprint(nil))
	log.Println("@v:",v," @err:",err)
}
