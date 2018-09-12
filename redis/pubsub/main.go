package main

import (
	"github.com/figoxu/Figo"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/quexer/utee"
	"time"
)

type SysEnv struct {
	rp *redis.Pool
}

var sysEnv = SysEnv{
}

func main(){
	REDIS_SERVER := "127.0.0.1:6379"
	REDIS_PASSWD := ""
	sysEnv.rp = Figo.RedisPool(REDIS_SERVER, REDIS_PASSWD)
	channelName := "ch_test"
	go sub("客户端1", channelName)
	go sub("客户端2", channelName)
	go sub("客户端3", channelName)
	time.Sleep(time.Second*time.Duration(1))
	 pub("世界你好",channelName)
	 pub("感动中国",channelName)
}


func sub(clientName,channelName string){
	c:=sysEnv.rp.Get()
	defer c.Close()
	psc := redis.PubSubConn{c}
	psc.Subscribe(channelName)
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s => %s: message: %s\n",clientName, v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s => %s: %s %d\n",clientName, v.Channel, v.Kind, v.Count)
		case error:
			utee.Chk(v.(error))
		}
	}
}

func pub(content,channelName string){
	c:=sysEnv.rp.Get()
	defer c.Close()
	c.Do("PUBLISH", channelName, content)
}

