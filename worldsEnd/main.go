package main
import "fmt"

func main(){
	fmt.Println("wait for the words end")
	<-make(chan int)
}
