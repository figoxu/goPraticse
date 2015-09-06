package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	REDIS_SERVER := "127.0.0.1:6379"
	REDIS_PASSWD := ""
	fmt.Print("hello")
	pool := createPool(REDIS_SERVER, REDIS_PASSWD)
	c := pool.Get()
	defer c.Close()

	data := []interface{}{}
	data = append(data, "meSet")
	for i := 0; i < 1000; i++ {
		v := fmt.Sprint("hello", i)
		d := []byte(v)
		data = append(data, i)
		data = append(data, d)
	}

	v2, err := redis.Int(c.Do("ZADD",  data...))
	if err != nil {
		fmt.Println("@err:", err)
	}
	fmt.Println("@v:", v2)

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
