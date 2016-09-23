package main
import (
	"log"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func from(connection chan int){
	connection <- rand.Intn(100)
}

func to(connection chan int){
	i := <- connection
	fmt.Printf("Some one sent me %d \n",i)
}


func main(){
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	log.Println("CPU Num:",cpus)
	connection := make(chan int)
	go from(connection)
	go to(connection)
	time.Sleep(time.Duration(2)*time.Second)
}
