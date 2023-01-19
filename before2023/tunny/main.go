package main
import (
	"fmt"
	"log"
	"github.com/jeffail/tunny"
	"github.com/figoxu/utee"
	"time"
)


func main(){
	fmt.Println("test")


	pool, err := tunny.CreatePool(10, func(input interface{}) interface{} {
		v := input.(string)
		time.Sleep(time.Second*time.Duration(5))
		log.Println("execute @v:",v)
		return nil
	}).Open()
	utee.Chk(err)


	messages := make(chan string,1000000)
	for i:=0;i<1000000;i++{
		messages <- fmt.Sprint("test",i)
	}

	for line := range messages {
		go pool.SendWorkTimed(5.1*1000, line)
	}

}