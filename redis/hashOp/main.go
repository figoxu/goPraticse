package main

import (
	"github.com/figoxu/utee"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

func main() {
	log.Println("hello")
	REDIS_SERVER := "127.0.0.1:6379"
	REDIS_PASSWD := ""
	pool := createPool(REDIS_SERVER, REDIS_PASSWD)
	k := "testP"
	saveHash(pool, k)
	incHash(pool, k, 100, 22, 44)
	log.Println(getHash(pool, k))
	incHash(pool, k, 100, 22, 44)
	log.Println(getHash(pool, k))
	incHash(pool, k, 100, 22, 44)
	saveHash(pool, k)
	log.Println(getHash(pool, k))
	incHash(pool, k, 100, 22, 44)
	log.Println(getHash(pool, k))
}

func saveHash(pool *redis.Pool, id string) {
	c := pool.Get()
	defer c.Close()

	v, err := redis.Int(c.Do("EXISTS", id))
	utee.Chk(err)
	if v > 0 {
		log.Println("already exist")
		return
	}
	_, err = c.Do("HMSET", id, "total", 0, "read", 0, "send", 0)
	utee.Chk(err)
	c.Do("TTL", id, 60*60*2)
}

func incHash(pool *redis.Pool, id string, total, read, sent int) {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("HINCRBY", id, "total", total)
	utee.Chk(err)
	_, err = c.Do("HINCRBY", id, "read", read)
	utee.Chk(err)
	_, err = c.Do("HINCRBY", id, "sent", sent)
	utee.Chk(err)
}

func getHash(pool *redis.Pool, id string) map[string]int {
	c := pool.Get()
	defer c.Close()
	m, err := redis.IntMap(c.Do("HGETALL", id))
	utee.Chk(err)
	return m
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
