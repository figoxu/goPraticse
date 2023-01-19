package main

import (
	"fmt"
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"math"
	"time"
)

func main() {
	rp := Figo.RedisPool("127.0.0.1:6379", "")
	rq := Figo.NewRedisZQueue(rp, "test", 20, func(v string, err error) {
		defer Figo.Catch()
		utee.Chk(err)
		fmt.Println(v)
		time.Sleep(time.Millisecond)
	})
	for i := 0; i < 100; i++ {
		if i == 10 {
			err := rq.Enq(fmt.Sprint("Hello ", i), math.MaxInt32-1)
			utee.Chk(err)
		} else {
			err := rq.Enq(fmt.Sprint("Hello ", i), i)
			utee.Chk(err)
		}
	}
	time.Sleep(time.Second * 30)
}
