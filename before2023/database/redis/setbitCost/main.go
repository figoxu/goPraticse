package main
import (
	"log"
	"github.com/garyburd/redigo/redis"
	"github.com/quexer/utee"
)

var (
	rp *redis.Pool
)

func main(){
	rp = utee.CreateRedisPool(30, "127.0.0.1:6379", "")
	v :=1;
	for i:=0;i<32;i++{
		v *=2
	}
	c := rp.Get()
	defer c.Close()
//	c.Do("set" ,"figo","xu")
	log.Println(v)
	for i:=0;i<v;i++{
		if i%100000 == 0 {
			c.Flush()
			log.Println(i)
		}

		c.Send("SETBIT","BIT",i,1)
	}
	c.Flush()
	log.Println("over")
}
