package main

import "fmt"

func main() {
	c := make (chan  int ,4)
	c<-15
	c<-14
	c<-65
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	//all of follow data are zero when channel closed
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
