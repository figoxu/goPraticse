package main
import (
//	"fmt"

"github.com/alicebob/qr"
//	"time"
	"log"
	"time"
)

func main(){
//	q, _ := qr.New(
//		d,
//		uid,
//		//			qr.OptionBuffer(1),
//		qr.OptionTimeout(time.Duration(1)*time.Second),
//	)
	q, _ := qr.New(
		"/home/figo/data/sampleTtlQ/q/",
		"example",
		qr.OptionBuffer(1000),
	)
	q.Enqueue("test1")
	q.Close()
	q, _ = qr.New(
		"/home/figo/data/sampleTtlQ/q/",
		"example",
		qr.OptionBuffer(1000),
	)
	q.Enqueue("test2")
	q.Close()
	q, _ = qr.New(
		"/home/figo/data/sampleTtlQ/q/",
		"example",
		qr.OptionBuffer(1000),
	)
	q.Enqueue("test3")
	q.Close()
	q, _ = qr.New(
		"/home/figo/data/sampleTtlQ/q/",
		"example",
		qr.OptionBuffer(1000),
	)
	rt := 0
	for {
		select {
		case v:= <- q.Dequeue():
			log.Println("@v:",v)
		default:
			rt = rt+1
			if rt > 10 {
				return
			}
			log.Println("not found val")
			time.Sleep(time.Second)
		}
	}
}
