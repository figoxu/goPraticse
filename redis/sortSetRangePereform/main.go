package main
import (
	"time"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"github.com/quexer/utee"
)



const (
	DATE_FORMAT_YmD   = "20060102"
	DATE_FORMAT_YmDH   = "2006010215"
	DATE_FORMAT_YmDHMS = "2006 01 02 15 04 05"
)

func getDay(day int, format string) int {
	t := time.Now()
	t = t.Add(time.Hour * 24 * time.Duration(day))
	i, _ := strconv.Atoi(t.Format(format))
	return i
}

func main(){
	REDIS_SERVER := "10.10.81.163:6379"
	REDIS_PASSWD := ""
	//	REDIS_SERVER := "127.0.0.1:6379"
	//	REDIS_PASSWD := ""
	pool := createPool(REDIS_SERVER, REDIS_PASSWD)

	setName := "SetTest"
	stF := utee.Tick()
	for i:=0;i<60;i++ {
		dayValMax := getDay(-1*i,DATE_FORMAT_YmDH)
		dayValMin := getDay(-1*i-1,DATE_FORMAT_YmDH)
		st := utee.Tick()
		i,err:=zcount(pool,setName,dayValMin,dayValMax);
		if err!= nil {
			fmt.Println("err:",err)
		}
		fmt.Println(" get val range @min:",dayValMin," @max:",dayValMax," @count:",i," @spent:",(utee.Tick()-st))
	}
	fmt.Println(" get val range 60 times eachDayRange @spent:",(utee.Tick()-stF))

	dayValMax := getDay(0,DATE_FORMAT_YmDH)
	dayValMin := getDay(-30,DATE_FORMAT_YmDH)
	stF = utee.Tick()
	count,_:=zcount(pool,setName,dayValMin,dayValMax);
	fmt.Println(" get val range 30 Days @min:",dayValMin," @max:",dayValMax," @count:",count,"  @spent:",(utee.Tick()-stF))
	dayValMax = getDay(-9,DATE_FORMAT_YmDH)
	dayValMin = getDay(-16,DATE_FORMAT_YmDH)
	stF = utee.Tick()
	count,_=zcount(pool,setName,dayValMin,dayValMax);
	fmt.Println(" get val range 7 Days @min:",dayValMin," @max:",dayValMax," @count:",count,"  @spent:",(utee.Tick()-stF))
}


func zcount(pool *redis.Pool,uid interface{}, min, max int )(int, error){
	c := pool.Get()
	defer c.Close()
	return redis.Int(c.Do("ZCOUNT", uid, min, max))
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