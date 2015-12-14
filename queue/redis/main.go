package main


import (
	as "github.com/aerospike/aerospike-client-go"
	"github.com/quexer/utee"
	"github.com/garyburd/redigo/redis"
//	"github.com/gocql/gocql"
//	mgo "gopkg.in/mgo.v2"
	"log"
	"time"
	"fmt"
)

const (
	DB     = "mpush"
	C_PUSH = "push"
)

const (
	NS_PUSH   = "push"
	SET_DV    = "dv"    //v2(app + id: dv2 obj), v1(id: dv data obj)
	SET_ALIAS = "alias" //app+alias : {id: "app:id", alias: ""}
	SET_USTAT = "ustat" //app(could be empty)+date: UserStat obj
	SET_CACHE = "cache" //key : {val: interface{}}
)

var i=0;
func main(){
	log.Println("hello")


	as:=connectAs("10.10.45.35:3000");

	r := createPool("10.10.23.154:6379","")

	syncAppAsDv(as,r,"5497c99059ba07085f000e7b")

	log.Println("success")
}


func connectAs(s string) *as.Client {
	h, port, err := utee.ParseUrl(s)
	utee.Chk(err)
	ac, err := as.NewClient(h, port)
	for err != nil {
		log.Println("as connect err:", err, " retry 2 sec later")

		time.Sleep(time.Second * time.Duration(2))
		ac, err = as.NewClient(h, port)
	}
	return ac
}



func syncAppAsDv(ac *as.Client, pool *redis.Pool, app string) error {
	stm := as.NewStatement(NS_PUSH, SET_DV, "id", "tp", "mt")
	stm.Addfilter(as.NewEqualFilter("app", app))
	p := as.NewQueryPolicy()
	p.RecordQueueSize = 200000
	rs, err := ac.Query(p, stm)
	if err != nil {
		fmt.Println("error happen When we sync @app:", app, " @err:", err)
		return err
	}
	latch := utee.NewThrottle(10000)
	for res := range rs.Results() {
		if res.Err != nil {
			return res.Err
		}
		bins := res.Record.Bins
		id, ok := bins["id"].(string)
		if !ok {
			//			log.Println("warn, id is empty", bins)
			continue
		}
		latch.Acquire()
		go Deq(pool,latch,id)
	}
	return nil
}


func qname(uid interface{}) string {
	return fmt.Sprintf("q%v", uid)
}


func  Deq(pool *redis.Pool,latch *utee.Throttle,uid interface{}) ([]byte, error) {
	c := pool.Get()
	defer c.Close()
	defer latch.Release()
	for {
		name := qname(uid)
		k, err := redis.String(c.Do("LPOP", name))
		if err != nil && err != redis.ErrNil {
			continue
		}

		if len(k) == 0 {
			break
		}
		b, err := redis.Bytes(c.Do("GET", k))
		if err != nil && err != redis.ErrNil {
			continue
		}
		if b != nil {
			c.Send("DEL", k)
			continue
		}
	}
	i++
	if(i%10000==0){
		log.Println("@success:",i)
	}
	return nil, nil
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