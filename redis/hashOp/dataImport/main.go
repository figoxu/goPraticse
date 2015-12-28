package main

import (
	"bufio"
	"github.com/garyburd/redigo/redis"
	"github.com/quexer/utee"
	"log"
	"os"
	"strings"
	"time"
	"fmt"
)

func main() {

	REDIS_SERVER := "127.0.0.1:6379"
	REDIS_PASSWD := ""
	pool := createPool(REDIS_SERVER, REDIS_PASSWD)
	file, err := os.Open("/home/figo/develop/env/GOPATH/src/github.com/figoxu/goPraticse/redis/hashOp/dataImport/data.txt")
	if err != nil {
		log.Println("@err:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		put(scanner.Text(), pool)
	}
	log.Println("hello")
}

func put(txt string, pool *redis.Pool) {
	c := pool.Get()
	defer c.Close()
	v := strings.Split(txt, " # ")
	id := v[0]
	pwd := v[1]
	mail := v[2]
	_, err := c.Do("HMSET",fmt.Sprint("account", id), "pwd", pwd, "mail", mail)
	utee.Chk(err)

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
