package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/pborman/uuid"
	"log"
	"time"
)

const (
	OPQ_SEND = "TEST_OPQ"
	BATCH_SIZE = 100
)

var (
	pool        *redis.Pool
	ch_send_opq = make(chan OpSend, 100000)
)
//marshall不能处理小写的属性
type OpSend struct {
	Tp  int
	Msg string
}

func main() {
	fmt.Println("hello")
	REDIS_SERVER := "127.0.0.1:6379"
	REDIS_PASSWD := ""
	pool = createPool(REDIS_SERVER, REDIS_PASSWD)

	go product()
	go consumer()
	go dEnq()
	time.Sleep(time.Second * time.Duration(10))
}

func dEnq(){
	for{

		if v,e:=Deq();e!=nil{
			log.Println("====deq err ==== @err:",e)
		}else{
			o := &OpSend{}
			json.Unmarshal(v,o);
			log.Println("====deq=== @tp:",o.Tp,"  @msg:",o.Msg)
		}

	}

}

func Deq() ([]byte, error) {
	c := pool.Get()
	defer c.Close()

	for {
		k, err := redis.String(c.Do("LPOP", OPQ_SEND))
		if err != nil && err != redis.ErrNil {
			return nil, err
		}

		if len(k) == 0 {
			break
		}
		b, err := redis.Bytes(c.Do("GET", k))
		if err != nil && err != redis.ErrNil {
			return nil, err
		}
		if b != nil {
			c.Send("DEL", k)
			return b, nil
		}
	}
	return nil, nil
}


func consumer() {
	v := []OpSend{}
	for {
		select {
		case o := <-ch_send_opq:
			v = append(v,o)
			if len(v) > BATCH_SIZE {
				BatchEnq(v,1000*60*60*24)
				v = []OpSend{}
			}
		default:
			if len(v) >0 {
				BatchEnq(v,1000*60*60*24)
				v = []OpSend{}
			}
			time.Sleep(time.Millisecond * time.Duration(10))
		}
	}
}

func product() {

	for {
		o := OpSend{
			Tp:  100,
			Msg: "test",
		}
		select {
		case ch_send_opq <- o:
		default:
			log.Println("[warn] opq overstock")
			time.Sleep(time.Second * time.Duration(5))
		}
	}
}

func BatchEnq(d []OpSend, ttl ...uint32) {
	if len(d) <= 0 {
		return
	}
	c := pool.Get()
	defer c.Close()
	for _,v  := range d {
		if data, err := json.Marshal(v); err != nil {
			fmt.Println("err:", err)
		} else {
			k := fmt.Sprintf("mk%v", uuid.NewUUID())
			if len(ttl) > 0 && ttl[0] > 0 {
				if err := c.Send("SETEX", k, ttl[0], data); err != nil {
					log.Println("[Batch Q SetEx] err :", err)
				}
			} else {
				if err := c.Send("SET", k, data); err != nil {
					log.Println("[Batch Q SET] err :", err)
				}
			}
			if err := c.Send("RPUSH", OPQ_SEND, k); err != nil {
				log.Println("[Batch Q RPUSH] err :", err)
			}
		}
	}
	c.Flush()
}

func createPool(server, auth string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     500,
		MaxActive:   500,
		Wait:        true,
		IdleTimeout: 4 * time.Minute,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if auth != "" {
				if _, err := c.Do("AUTH", auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
