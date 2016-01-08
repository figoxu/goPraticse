package main

import (
	"fmt"
	"github.com/hashicorp/memberlist"
	"log"
	"time"
)

func main() {
	fmt.Println("main A")
	config := memberlist.DefaultLANConfig()
	config.BindAddr = "192.168.50.25"
	config.BindPort = 9898
	list, err := memberlist.Create(config)
	n, err := list.Join([]string{"192.168.50.25:9898"})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}
	log.Println("@n:", n)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}
	for {
		checkCluster(list)
		time.Sleep(time.Second)
	}
}

type memberlistBroadcast struct {
	node   string
	msg    []byte
	//	notify chan struct{}
}

func (p *memberlistBroadcast) Invalidates(other memberlist.Broadcast) bool {
	// Check if that broadcast is a memberlist type
	mb, ok := other.(*memberlistBroadcast)
	if !ok {
		return false
	}

	// Invalidates any message about the same node
	return p.node == mb.node
}

func (p *memberlistBroadcast) Message() []byte {
	return p.msg
}

func (p *memberlistBroadcast) Finished() {
	//	select {
	//	case p.notify <- struct{}{}:
	//	default:
	//	}
}

func checkCluster(list *memberlist.Memberlist) {

	// Join an existing cluster by specifying at least one known member.
	//	memberlist.Broadcast(&memberlist.memberlistBroadcast{"test", []byte("1. this is a test."))
	// Ask for members of the cluster

	//	m := &memberlistBroadcast{"test", []byte("1. this is a test.")}
	//	memberlist.Broadcast(m)
	broadcast();
	list.UpdateNode(time.Millisecond)
	fmt.Println("==============================")
	for _, member := range list.Members() {
		fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
	}
	fmt.Println("==============================")
}


func broadcast(){

	q := &memberlist.TransmitLimitedQueue{RetransmitMult: 3, NumNodes: func() int { return 10 }}

	// 18 bytes per message
	q.QueueBroadcast(&memberlistBroadcast{"test", []byte("1. this is a test.")})
	q.QueueBroadcast(&memberlistBroadcast{"foo", []byte("2. this is a test.")})
	q.QueueBroadcast(&memberlistBroadcast{"bar", []byte("3. this is a test.")})
	q.QueueBroadcast(&memberlistBroadcast{"baz", []byte("4. this is a test.")})

}