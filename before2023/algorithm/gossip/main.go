package main

import (
	"github.com/hashicorp/memberlist"
	"fmt"
	"log"
	"time"
)

const(
	seedPort = "192.168.2.99"
)


func main(){
	cfg := memberlist.DefaultLocalConfig()
	log.Println("@bindPort:",cfg.BindPort,"  @advPort:",cfg.AdvertisePort)
	list, err := memberlist.Create(cfg)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	// Join an existing cluster by specifying at least one known member.tu
	n, err := list.Join([]string{seedPort})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}
	log.Println("@n:",n)

	// Ask for members of the cluster
	for _, member := range list.Members() {
		fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	}
	time.Sleep(time.Hour*time.Duration(1))
}
