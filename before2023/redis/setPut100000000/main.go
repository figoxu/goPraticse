package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"log"
	"github.com/pborman/uuid"
	"fmt"
	"sync/atomic"
)


var(
	COUNT int64 = 100
	pool *redis.Pool
)
const (
	CMD_REDIS_SET_ADD = "ZADD"
)

func main() {
	REDIS_SERVER := "10.10.81.163:6379"
	REDIS_PASSWD := ""
//	REDIS_SERVER := "127.0.0.1:6379"
//	REDIS_PASSWD := ""
	pool = createPool(REDIS_SERVER, REDIS_PASSWD)
	d := []byte("test")
	Add("test_Range_set",d,100)
	log.Printf("SUCCESS ")


	setName := "SetTest"
	piece := 16697
//	piece := 166
//	16697*100*60  = 100182000
	for i:=0;i<60;i++ {
		dayVal := getDay(-1*i,DATE_FORMAT_YmDH)
		for j:=0;j<100;j++{
			go Gen(setName,dayVal,piece)
		}
	}
	<-make(chan int)
}

func Gen(setName string,dayVal,total int){
	for i:=0;i<total;i++{
		v := fmt.Sprintf("mk%v", uuid.NewUUID())
		d := []byte(v)
		Add(setName,d,dayVal)
		atomic.AddInt64(&COUNT,1)
		fmt.Println("success to add @key",v," @total:",COUNT)
	}
}

func Add(uid interface{}, data []byte, score int) error {
	c := pool.Get()
	defer c.Close()
	redis.Int(c.Do(CMD_REDIS_SET_ADD, uid, score, data))
	return nil
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
