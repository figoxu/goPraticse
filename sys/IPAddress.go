package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("udp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	ips :=strings.Split(conn.LocalAddr().String(), ":")

	for _,ip := range ips{
		fmt.Println(ip)
	}

}