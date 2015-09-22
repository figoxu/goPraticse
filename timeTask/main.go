package main


import (
	"fmt"
	"time"
)

func testTimer1() {
	go func() {
		fmt.Println("test timer1")
	}()

}

func testTimer2() {
	go func() {
		fmt.Println("test timer2")
	}()
}

func timer1() {
	timer1 := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer1.C:
			testTimer1()
		}
	}
}

func timer2() {
	timer2 := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-timer2.C:
			testTimer2()
		}
	}
}

func main() {
	go timer1()
	timer2()
}
