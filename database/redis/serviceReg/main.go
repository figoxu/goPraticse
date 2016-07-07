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
	"net/url"
	"strings"
)

var (
	rp *redis.Pool
)

const (
	cron_healthReport_task = "0/30 * * * * ?" //per 30 second
	cron_healthChk_task    = "0/10 * * * * ?" //per 30 second
)

func main() {
	rp = utee.CreateRedisPool(30, "106.75.27.144:6379", "baulk3?speed")
	addr := flag.String("p", ":6666", "address where the server listen on")
	flag.Parse()
	log.Printf("start server http 1 on %s \n", *addr)
	guards := NewHealthGuards(*addr)
	for _, v := range guards.m {
		healthApi := v.(*HealthApi)
		if v.key == guards.keyMe {
			continue
		}
		uri := fmt.Sprint(healthApi.api, "/cache/clear/:id")
		utee.HttpPost(uri, &url.Values{})
	}
	http.HandleFunc("/api/health", guards.healthCheckHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))

}

func NewHealthGuards(port string) *HealthGuards {
	keyMe := ""
	ifaces, err := net.Interfaces()
	utee.Chk(err)
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		utee.Chk(err)
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ipStr := ip.String()
			if strings.Index(ipStr, "::") != -1 {
				continue
			}
			if ipStr == "127.0.0.1" {
				continue
			}
			if keyMe != "" {
				keyMe = fmt.Sprint(keyMe, ",")
			}
			keyMe = fmt.Sprint(keyMe, "http://", ipStr, port)
		}
	}
	healthGuard := &HealthGuards{
		m:     make(map[string]*HealthApi),
		keyMe: keyMe,
	}

	c := cron.New()
	c.AddFunc(cron_healthReport_task, healthGuard.healthReport)
	c.AddFunc(cron_healthChk_task, healthGuard.healthCheck)
	c.Start()
	return healthGuard
}

type HealthGuards struct {
	m     map[string]*HealthApi
	keyMe string
}

func (p *HealthGuards) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
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

func (p *HealthGuards) healthCheck() {
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
		if key == p.keyMe {
			continue
		}
		if p.m[key] == nil {
			p.m[key] = &HealthApi{
				key:  key,
				apis: strings.Split(key, ","),
				api:  "",
			}
		}
		if check(p.m[key].api) {
			continue
		}
		healthFlag := false
		for _, api := range p.m[key].apis {
			log.Println("@api:", api)
			if b, err := utee.HttpGet(fmt.Sprint(api, "/api/health")); err == nil && string(b) == "ok" {
				healthFlag = true
				p.m[key].api = api
				break
			}
		}
		if !healthFlag {
			log.Println("[warn] @host:", key, " is not health . ")
			p.m[key].illCount++
		}
		if p.m[key].illCount > 3 {
			p.ill2die(key)
		}
	}

}

func (p *HealthGuards) ill2die(key string) {
	log.Println("[warn] kick not health @host:", key, " from conf server")
	delete(p.m, key)
	c := rp.Get()
	defer c.Close()
	c.Do("HDEL", "service_gateway", key)
}

func (p *HealthGuards) healthReport() {
	c := rp.Get()
	defer c.Close()
	c.Do("HSETNX", "service_gateway", p.keyMe, 1)
}
