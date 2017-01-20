package main

import "log"

func main() {
	i := 0
loopLabel:
	for {
		i++
		if i > 10 {
			i = 0
			break loopLabel
			//use goto  ,loop will try again And agian
			//user break ,loop will quit
		}
		log.Println("@i:", i)
	}
}
