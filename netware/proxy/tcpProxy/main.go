package main

import (
	"fmt"
	"github.com/quexer/utee"
	"io"
	"net"
	"os"
)

func main() {
	localAddr := "127.0.0.1:8080"
	remoteAddr := "127.0.0.1:6379"
	local, err := net.Listen("tcp", localAddr)
	if local == nil {
		utee.Chk(err)
	}
	for {
		conn, err := local.Accept()
		if conn == nil {
			utee.Chk(err)
		}
		go forward(conn, remoteAddr)
	}
}

func forward(local net.Conn, remoteAddr string) {
	remote, err := net.Dial("tcp", remoteAddr)
	if remote == nil {
		fmt.Fprintf(os.Stderr, "remote dial failed: %v\n", err)
		return
	}
	go io.Copy(local, remote)
	go io.Copy(remote, local)
}
