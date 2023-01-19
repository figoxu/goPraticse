package main

import (
	"flag"
	"github.com/quexer/tok"
	"github.com/quexer/utee"
	"log"
	"time"
)

func main() {
	var socketPort string
	flag.StringVar(&socketPort, "socket_port", ":7000", "the port for socket")
	flag.Parse()

	g_hub = initActor(socketPort)
	log.Println("Hello")
	time.Sleep(time.Hour)
}

func initActor(port string) *tok.Hub {

	hc := &tok.HubConfig{
		Actor: NewTcpActor(),
		Sso:   true,
	}
	hb, err := tok.Listen(nil, hc, port)
	utee.Chk(err)
	return hb
}
