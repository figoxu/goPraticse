package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var host = "127.0.0.1"

func main() {
	num := 1000      //指定并发个数
	portNum := 65535 //扫描到最大的端口
	var sy sync.WaitGroup
	for i := 0; i < portNum; {
		j := i + num
		if j > portNum {
			j = portNum
		}
		sy.Add(1)
		go func(i, j int) {
			defer sy.Done()
			ch(i, j)
		}(i, j)
		i = j
	}
	sy.Wait()
}

func ch(a, b int) {
	for i := a; i <= b; i++ {
		s := strconv.Itoa(i)
		_, err := net.DialTimeout("tcp", host+":"+s, time.Second*1)
		if err != nil {
			if i == 80 {
				log.Println("@i:", err)
			}
			continue
		}
		fmt.Println(s)
	}
}
