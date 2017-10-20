package main

import (
	"log"
	"github.com/figoxu/Figo"
	"github.com/Pallinder/go-randomdata"
	"fmt"
)


func main() {
	bq := Figo.NewBlockExecuteQ(1000, 3, 3, func(v interface{})bool {
		b := randomdata.Boolean()
		if b {
			log.Println("execute @v:", v, " SUCCESS")
		} else {
			log.Println("execute @v:", v, " FAILURE")
		}
		return b
	})

	mockInput := func() {
		for i := 0; i < 30; i++ {
			bq.Enq(fmt.Sprint("testData ", i))
		}
	}
	mockInput()
	<-make(chan bool)
}
