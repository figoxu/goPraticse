package main

import (
	"fmt"

	"figoxu.me/redis/pkg/ds"
	"figoxu.me/redis/pkg/ut"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("hello")

	ps := NewPubSub(ds.DefaultRedis())
	topic := "foo"
	ps.Sub(topic, SubscriberHdl)
	for i := 0; i < 100; i++ {
		err := ps.Pub(topic, fmt.Sprint("hello ", i))
		ut.Chk(err)
	}
	var v chan string
	<-v
}

func SubscriberHdl(v string) {
	logrus.Println("消费 ", v)
}
