package main

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/quexer/utee"
	"github.com/robfig/cron"
	"log"
	"net"
	"net/http"
	"strings"
)

var (
	rp        *redis.Pool
	keyMe     string
	healthMap = make(map[string]*HealthApi)
)

const (
	cron_healthReport_task = "0/30 * * * * ?" //per 30 second
	cron_healthChk_task    = "0/10 * * * * ?" //per 30 second
)

func main() {
	rp = utee.CreateRedisPool(30, "106.75.27.144:6379", "baulk3?speed")
	addr := flag.String("p", ":6666", "address where the server listen on")
	flag.Parse()
	keyMe = healthKey(*addr)
	log.Println("@key:", keyMe)

	c := cron.New()
	c.AddFunc(cron_healthReport_task, func() {
		healthReport(keyMe)
	})
	c.AddFunc(cron_healthChk_task, healthCheck)
	c.Start()
	log.Printf("start server http 1 on %s \n", *addr)
	http.HandleFunc("/api/health", healthCheckHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))

}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprint(w, "ok")
	return
}

type HealthApi struct {
	key      string
	apis     []string
	api      string
	illCount int
}

func healthCheck() {
	c := rp.Get()
	defer c.Close()
	keys, err := redis.Strings(c.Do("HKEYS", "service_gateway"))
	utee.Chk(err)
	check := func(api string) bool {
		if b, err := utee.HttpGet(fmt.Sprint(api, "/api/health")); err == nil && string(b) == "ok" {
			return true
		}
		return false
	}
	for _, key := range keys {
		if key == keyMe {
			log.Println("is not need to check myself")
			continue
		}
		if healthMap[key] == nil {
			healthMap[key] = &HealthApi{
				key:  key,
				apis: strings.Split(key, ","),
				api:  "",
			}
		}
		if check(healthMap[key].api) {
			continue
		}
		healthFlag := false
		for _, api := range healthMap[key].apis {
			log.Println("@api:", api)
			if b, err := utee.HttpGet(fmt.Sprint(api, "/api/health")); err == nil && string(b) == "ok" {
				healthFlag = true
				healthMap[key].api = api
				log.Println("okay  @api:", api, "  ==[skip other api]==")
				break
			}
		}
		if !healthFlag {
			log.Println("[warn] @host:", key, " is not health . ")
			healthMap[key].illCount++
		}
		if healthMap[key].illCount > 3 {
			ill2die(key)
		}
	}

}

func ill2die(key string) {
	log.Println("[warn] kick not health @host:", key, " from conf server")
	delete(healthMap, key)
	c := rp.Get()
	defer c.Close()
	c.Do("HDEL", "service_gateway", key)
}

func healthReport(key string) {
	c := rp.Get()
	defer c.Close()
	c.Do("HSETNX", "service_gateway", key, 1)
}

func healthKey(port string) string {
	hk := ""
	ifaces, err := net.Interfaces()
	utee.Chk(err)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		utee.Chk(err)
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			ipStr := ip.String()
			if strings.Index(ipStr, "::") != -1 {
				continue
			}
			if ipStr == "127.0.0.1" {
				continue
			}
			if hk != "" {
				hk = fmt.Sprint(hk, ",")
			}
			hk = fmt.Sprint(hk, "http://", ipStr, port)
		}
	}
	return hk
}
