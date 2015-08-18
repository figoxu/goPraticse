package main
import "fmt"
import (
	atomic "sync/atomic"
	"log"
	"time"
)

var(
	GO_RUTINE_COUNT int64 = 100
)

func init(){
	fmt.Println("sys init")
}

func main(){
	a :=func(){
		for i:=0;i<10000;i++ {
			go increase();
		}
	}

	b := func(){
		for i:=0;i<10000;i++ {
			go decrease();
		}
	}
	//
	for i:=0;i<100;i++{
		go a();
		go b();
	}
	time.Sleep(time.Second*time.Duration(60))
	for i:=0;i<100;i++{
		go a();
		go b();
	}
	time.Sleep(time.Second*time.Duration(60))
	log.Println("@current val is :",GO_RUTINE_COUNT)
}

func increase(){
	atomic.AddInt64(&GO_RUTINE_COUNT,1)
	log.Print("gorutine#open#....@current val is :",GO_RUTINE_COUNT);
}

func decrease(){
	atomic.AddInt64(&GO_RUTINE_COUNT,-1)
	log.Print("gorutine#close#....@current val is :",GO_RUTINE_COUNT);
}
