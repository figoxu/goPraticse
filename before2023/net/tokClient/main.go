package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	log.Println("hello")
	n := 12
	t := time.Now().UnixNano() / 1e6
	for i := 0; i < n; i++ {
		uid := fmt.Sprintf("%d-%05d", t, i+1)
		tokAddr := string("localhost:7000")
		newTcp(uid, tokAddr, 4, 2)
		log.Println("new tcp ", uid)
		if i%1000 == 999 {
			time.Sleep(time.Second)
		}
	}
	time.Sleep(time.Hour)
}

func newTcp(token string, tokAddr string, ping, warm int) {
	conn, err := net.Dial("tcp", tokAddr)
	if err != nil {
		log.Printf("[%s] connect fail, %v", token, err)
		time.Sleep(3 * time.Second)
		return
	}

	adapter := &tcpAdapter{conn}
	adapter.Write([]byte("aes,ack|||" + token))
	go func() {
		//warm up
		time.Sleep(time.Second * time.Duration(rand.Intn(warm)))
		for {
			//ping interval
			time.Sleep(time.Second * time.Duration(ping))
			err = adapter.Write([]byte("ping"))
			if err != nil {
				log.Printf("[%s] write ping err : %v", token, err)
				adapter.Close()
				return
			}

			//log.Println("ping a ", seq)
		}
	}()

	go func() {
		for {
			data, err := adapter.Read()
			if err != nil {
				log.Printf("[%s] read err: %v \n", token, err)
				adapter.Close()
				return
			}
			log.Println("@data:", string(data))

		}
	}()

}

//COPY FROM TOK
const (
	tcp_header_len = 4
)

var (
	TCP_MAX_PACK_LEN uint32 = 4 * 1024 * 1024 //Upper limit for single message
)

type tcpAdapter struct {
	conn net.Conn
}

func (p *tcpAdapter) Read() ([]byte, error) {
	//read header
	b := make([]byte, tcp_header_len)
	if _, err := io.ReadFull(p.conn, b); err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(b)
	var n uint32
	if err := binary.Read(buf, binary.BigEndian, &n); err != nil {
		return nil, err
	}

	if n > TCP_MAX_PACK_LEN {
		return nil, fmt.Errorf("pack length %dM can't greater than %dM", n/1024/1024, TCP_MAX_PACK_LEN/1024/1024)
	}

	b = make([]byte, n)
	_, err := io.ReadFull(p.conn, b)
	return b, err

}

func (p *tcpAdapter) Write(b []byte) error {
	//set write deadline
	if err := p.conn.SetWriteDeadline(time.Now().Add(time.Minute)); err != nil {
		log.Println("[warning] setting write deadline fail: ", err)
		return err
	}

	n := uint32(len(b))

	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, &n); err != nil {
		return err
	}
	_, err := p.conn.Write(append(buf.Bytes(), b...))

	return err
}

func (p *tcpAdapter) Close() {
	p.conn.Close()
}
