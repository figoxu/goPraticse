package main

import (
	"flag"
	"fmt"
	"github.com/alicebob/qr"
	"github.com/figoxu/utee"
	"log"
	"os"
	"time"
)

type FileQueue struct {
	qr *qr.Qr
}

func NewFileQueue(bufferCap, concurrent int, diskLoc, queueName string, worker func(interface{})) (*FileQueue, error) {
	err := os.MkdirAll(diskLoc, 0777)
	if err != nil {
		return nil, err
	}
	q, err := qr.New(
		diskLoc,
		queueName,
		qr.OptionBuffer(bufferCap),
	)
	if err != nil {
		return nil, err
	}
	fq := &FileQueue{
		qr: q,
	}
	f := func() {
		for {
			worker(fq.Deq())
		}
	}
	for i := 0; i < concurrent; i++ {
		go f()
	}
	return fq, nil
}

func (p FileQueue) Enq(data interface{}) {
	p.qr.Enqueue(data)
}

func (p FileQueue) Deq() interface{} {
	return <-p.qr.Dequeue()
}

var (
	diskLoc   string
	queueName string
)

func init() {
	diskLoc = *flag.String("loc", "./fqueue", " storage the data of queue")
	queueName = *flag.String("name", "sample", " name of queue ,must be uniqueue in the diskLoc")
	flag.Parse()
}

func main() {
	log.Println("Hello @loc:", diskLoc)
	queue, err := NewFileQueue(1000, 10, diskLoc, queueName, consume)
	utee.Chk(err)
	for i := 0; i < 10000000; i++ {
		queue.Enq(fmt.Sprint("data", i))
	}
	time.Sleep(time.Second * time.Duration(60))
	log.Println("Quit now")
}

func consume(data interface{}) {
	time.Sleep(time.Second)
	log.Println("Consume @data:", data)
}
