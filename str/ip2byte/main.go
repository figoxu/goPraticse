package main

import (
	"fmt"
	"github.com/figoxu/Figo"
	"log"
	"strings"
)

func main() {
	log.Println(parseIp("59.110.166.246:5052"))
	log.Println(parseIp("8.8.8.8:65535"))
	log.Println(parseIp("255.255.255.0:255"))
	bs := parseIp2bytes("59.110.166.246:5052")
	log.Println(bs, "   ", parsebytes2IpAddr(bs))
	bs = parseIp2bytes("8.8.8.8:65535")
	log.Println(bs, "   ", parsebytes2IpAddr(bs))
	bs = parseIp2bytes("255.255.255.0:255")
	log.Println(bs, "   ", parsebytes2IpAddr(bs))
	bs = parseIp2bytes("127.0.0.1:256")
	log.Println(bs, "   ", parsebytes2IpAddr(bs))
}

func parseIp(ipAddr string) (ips []byte, port uint16) {
	vs := strings.Split(ipAddr, ":")
	host := vs[0]
	port = 80
	if len(vs) > 1 {
		portV, _ := Figo.TpInt(vs[1])
		port = uint16(portV)
	}
	ipSegments := strings.Split(host, ".")
	ips = make([]byte, 0)
	for _, ipStr := range ipSegments {
		v, _ := Figo.TpInt(ipStr)
		ips = append(ips, byte(v))
	}
	return ips, port
}

func parseIp2bytes(ipAddr string) (bs []byte) {
	ips, port := parseIp(ipAddr)
	h, l := byte(port>>8), byte(port&0xff)
	bs = make([]byte, 0)
	bs = append(bs, ips...)
	bs = append(bs, h, l)
	return bs
}

func parsebytes2IpAddr(bs []byte) (addr string) {
	addr = fmt.Sprint(bs[0], ".", bs[1], ".", bs[2], ".", bs[3])
	addr = fmt.Sprint(addr, ":", int(bs[4])*256+int(bs[5]))
	return addr
}
