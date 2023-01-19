package main

import (
	"fmt"
	"github.com/quexer/utee"
	"log"
	"net/url"
	"time"
)

func main() {
	//	workerCount := 10;
	//	taskCount := 100000;
	workerCount := 10
	taskCount := 3

	for i:=0; i < workerCount; i++ {
		go worker(taskCount)
	}
	time.Sleep(time.Minute * time.Duration(10))
}

func worker(taskCount int) {

	for j := 0 ; j < taskCount; j++ {
		postMsg()
	}
}

func postMsg() {
	appId := "XXXX"
	appKey := "XXXXX"
	restapi := fmt.Sprint("http://XXXX:XXX", "/XX/XX/XXX/", appId, "/XX/XX")
	q := url.Values{}
	q.Add("id", "3fe84919c1de91e35fbcd2d5bdb45fbc2")
	q.Add("title", "test")
	q.Add("content", "testMsg")
	b, err := utee.HttpPost(restapi, q, appId, appKey)
	if err != nil {
		log.Println("error:", err)
	} else {
		v := string(b)
		log.Println("success to finish task with return value @v", v)
	}
}
