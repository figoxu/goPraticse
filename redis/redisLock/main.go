package main

import (
	"log"
	"github.com/figoxu/Figo"
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
)

var(
	rp = Figo.RedisPool("127.0.0.1:6379", "")
)

func main() {
	for i:=0;i<100;i++{
		go TestMethod(fmt.Sprint("Helllo @i:",i))
	}
	time.Sleep(time.Hour*time.Duration(1))
}

func TestMethod(msg string){
	ll:=RedisMutex{
		rp:rp,
		resource:"testLock",
	}
	defer ll.Unlock()
	lock := func(){
		for{
			if b,_:=ll.Lock(60);b==true{
				return
			}
			time.Sleep(time.Nanosecond*time.Duration(3))
		}
	}
	lock()
	log.Println(msg)
	time.Sleep(time.Second*time.Duration(10))
}

type RedisMutex struct {
	rp       *redis.Pool
	resource string
}

func (p *RedisMutex) Lock(ttlSec int) (bool, error) {
	c := p.rp.Get()
	defer c.Close()
	if _, err := redis.String(c.Do("SET", p.resource, "1", "EX", ttlSec, "NX")); err != nil {
		if err == redis.ErrNil {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (p *RedisMutex) Unlock() {
	c := p.rp.Get()
	defer c.Close()
	c.Do("del", p.resource)
}
