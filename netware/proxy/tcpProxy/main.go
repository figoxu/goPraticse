package main

import (
	"github.com/figoxu/Figo"
)

func main() {
	localAddr := "127.0.0.1:8080"
	remoteAddr := "127.0.0.1:6379"
	redisProxy := Figo.NewTcpProxy(localAddr, remoteAddr, 100)
	redisProxy.Listen()
}
