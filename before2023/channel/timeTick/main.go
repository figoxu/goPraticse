package main

import (
	"log"
	"time"
)

func main() {
	heartbeat := time.Tick(1 * time.Second)
	for {
		select {
		case <-heartbeat:
			log.Println("tick")
		}
	}
}
