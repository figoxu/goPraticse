package main
import "log"

func main(){

	c := make(chan int)
	go func() {
		log.Println(" method in gorutine")
		c <- 1
	}()
	log.Println(" wait for go rutine done")
	<-c
	log.Println("gorutine is done")
}
