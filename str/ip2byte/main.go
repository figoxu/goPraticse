package main

import (
	"github.com/figoxu/Figo"
	"log"
	"strings"
)

func main() {
	log.Println(parseIp("59.110.166.246:5052"))
	log.Println(parseIp("8.8.8.8:65535"))
	log.Println(parseIp("255.255.255.0:65535"))
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