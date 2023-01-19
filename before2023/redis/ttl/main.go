package main
import (
	"log"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main(){
	log.Println("hello")
	REDIS_SERVER := "127.0.0.1:6379"
	REDIS_PASSWD := ""
	pool := createPool(REDIS_SERVER, REDIS_PASSWD)
	c:=pool.Get()
	defer c.Close()
	if i,e:=redis.Int( c.Do("ttl","notExistKey"));e!=nil{
		log.Println("@err:",e)
	}else{
		log.Println("@v:",i)
	}
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
