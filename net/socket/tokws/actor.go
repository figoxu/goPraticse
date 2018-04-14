package main

import (
	"github.com/figoxu/Figo"
	"github.com/quexer/tok"
	"log"
	"fmt"
)

type WsActor struct{}

func (p *WsActor) BeforeReceive(dv *tok.Device, data []byte) ([]byte, error) {
	return data, nil
}
func (p *WsActor) OnReceive(dv *tok.Device, data []byte) {
	data = Figo.Bh.Append([]byte(fmt.Sprint("回音：",string(data))))
	g_hub.Send(dv.UID(), data, 1024)
	log.Println("收到:", string(data))
}
func (p *WsActor) BeforeSend(dv *tok.Device, data []byte) ([]byte, error) {
	return data, nil
}
func (p *WsActor) OnSent(dv *tok.Device, data []byte) {
	log.Println("发送:", string(data))
}

func (p *WsActor) OnClose(dv *tok.Device) {
	log.Println("关闭:", Figo.JsonString(dv))
}
func (p *WsActor) Ping() []byte {
	return []byte("心跳")
}
func (p *WsActor) Bye(kicker *tok.Device, reason string, dv *tok.Device) []byte {
	return []byte("再见")
}
