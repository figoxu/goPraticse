package main
import (
	"fmt"
	"log"
	"time"
)

const (
	MAX_OPQ_AS_BASIC = 200000
)

var (
	g_as_basic_q = make(chan string, MAX_OPQ_AS_BASIC)
)


func processCore() {
	for op := range g_as_basic_q {
		time.Sleep(time.Millisecond*time.Duration(10))
		log.Println(op)
	}
}

func execute(op string) {
	select {
	case g_as_basic_q <- op:
	default:
		log.Println("[warn] op as basic overstock")
	}
}

func main(){
	go processCore()

	for i:=0;i<100;i++ {
		execute(fmt.Sprint("hello",i))
	}
	log.Println("Sleep 10 Second")
	time.Sleep(time.Second*time.Duration(10))
	for i:=0;i<100;i++ {
		time.Sleep(time.Second*time.Duration(1))
		execute(fmt.Sprint("hello",i))
	}
	log.Println("Sleep 10 Second")
	time.Sleep(time.Second*time.Duration(10))
}