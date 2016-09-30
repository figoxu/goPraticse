package main

import (
	"fmt"
	"github.com/figoxu/utee"
	"github.com/garyburd/redigo/redis"
	bloom "github.com/kristinn/go-bloom"
	"log"
	"time"
)

func newRedisPool(maxIdle int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:6379", "127.0.0.1"))
		},
	}
}

func main() {
	redisPool := newRedisPool(10)
	filter, err := bloom.NewRedis(redisPool, "redis-save-test", 15000, 7)
	utee.Chk(err)

	filter.Append([]byte("afi"))
	filter.Save()

	exists, err := filter.Exists([]byte("afi"))
	if !exists {
		log.Fatal("afi should exist in the Redis backend")
	}
	if err != nil {
		log.Fatal(err)
	}

	exists, err = filter.Exists([]byte("amma"))
	if exists {
		log.Fatal("amma shouldn't exist in the Redis backend")
	}
	if err != nil {
		log.Fatal(err)
	}

//	conn := redisPool.Get()
//	defer conn.Close()
//
//	conn.Do("FLUSHALL")
}
