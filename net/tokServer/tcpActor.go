package main

import (
	"github.com/pborman/uuid"
	"github.com/quexer/tok"
	"github.com/quexer/utee"
	"log"
	"net/http"
)

var (
	g_hub *tok.Hub
)

type TcpActor struct {
	m utee.SyncMap //storage dv conf info  [version,ability info]
}

func NewTcpActor() *TcpActor {
	actor := &TcpActor{
		m: utee.SyncMap{},
	}
	return actor
}

type Frame struct {
	Uid  interface{}
	Data []byte
}

//save conn with
func (p *TcpActor) put(uid interface{}, meta string) {
	p.m.Put(uid, meta)
}

func (p *TcpActor) get(uid interface{}) string {
	b, _ := p.m.Get(uid)
	if s, ok := b.(string); ok {
		return s
	} else {
		return ""
	}
}

func (p *TcpActor) Auth(r *http.Request) (interface{}, error) {
	//auth against http request. return uid if auth success

	//request base on  tcp_conn.go -> buildReq
	//	id := r.Header.Get("Cookie")
	//	p.put(id, r.Header.Get(tok.META_HEADER))
	//	return id, nil

	return uuid.New(), nil
}

func (p *TcpActor) BeforeReceive(uid interface{}, data []byte) []byte {
	//is invoked every time the server receive valid payload

	//time to  Decode

	return data
}

func (p *TcpActor) OnReceive(uid interface{}, data []byte) {
	//is invoked every time the server receive valid payload

	//time to process to receive
	if string(data) == "ping" {
		g_hub.Send(uid, []byte("pong"))
	}

	log.Println("receive @uid:", uid, " @data:", string(data))

}

func (p *TcpActor) BeforeSend(uid interface{}, data []byte) []byte {
	//is invoked if message is sent successfully. count mean copy quantity

	//time to Encode
	return data
}

func (p *TcpActor) OnSent(uid interface{}, data []byte, count int) {
	//is invoked if message is sent successfully. count mean copy quantity

	//time to process data to send
	log.Println("send @uid:", uid, "@data:", string(data))
}

func (p *TcpActor) OnCache(uid interface{}) {
	//is invoked after message caching
}

func (p *TcpActor) OnClose(uid interface{}, active int) {
	//is invoked after a connection has been closed
	//active, count of active connections for this user
	log.Println("close @uid:", uid, "@active:", active)
}

func (p *TcpActor) Ping() []byte {
	//Build ping payload.  auto ping feature will be disabled if this method return nil
	return []byte("ping")
}

func (p *TcpActor) Bye(reason string) []byte {
	//Build payload for different reason before connection is closed
	return []byte("bye")
}
