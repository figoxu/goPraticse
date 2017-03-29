package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"github.com/quexer/utee"
	"log"
	"tendcloud.com/push-gateway/vendor/github.com/figoxu/Figo"
	"time"
)

type HCWriter interface {
	write(interface{})
}

type RedisqlWriter struct {
	hotPrefix    string
	hotWriter    redis.Pool
	coldWriter   orm.Ormer
	coldChan     chan string
}

func NewRedisqlWriter(hotWriter redis.Pool, coldWriter orm.Ormer) RedisqlWriter {
	writer := RedisqlWriter{
		hotWriter:  hotWriter,
		coldWriter: coldWriter,
	}
	go writer.coldSchedule()
	go writer.hibernate()
	return writer
}

func (p *RedisqlWriter) score(t time.Time) string{
	return t.Format("20060102150405")
}

func (p *RedisqlWriter) coldSchedule() error {
	fetchColdKey := func() {
		endScore := p.score(time.Now().Add(time.Duration(-1) * time.Minute))
		r := p.hotWriter.Get()
		defer r.Close()
		keys, err := redis.Strings(r.Do("ZRANGE", p.zsetKey(), 0, endScore))
		defer Figo.Catch()
		utee.Chk(err)
		for _, key := range keys {
			p.coldChan <- key
		}
		r.Do("ZREM", p.zsetKey(), keys...)
	}
	heartbeat := time.Tick(1 * time.Minute)
	for {
		select {
		case <-heartbeat:
			fetchColdKey()
		}
	}
	return nil
}

func (p *RedisqlWriter) hashKey() string {
	return fmt.Sprint("hcw_", p.hotPrefix, "_hash")
}

func (p *RedisqlWriter) zsetKey() string {
	return fmt.Sprint("hcw_", p.hotPrefix, "_zset")
}

func (p *RedisqlWriter) hibernate() error {
	//todo scan sorted set  to cold schedule
	//距离现在超过1分钟的数据进行入库   <- coldChan

	fetchStruct := func(pk string) HotColdInfo {
		defer Figo.Catch()
		r := p.hotWriter.Get()
		defer r.Close()
		dataMap, err := redis.StringMap(r.Do("HSCAN", p.hashKey(), 0, "MATCH", fmt.Sprint(pk, "_*")))
		//todo dataMap for batch insert


		utee.Chk(err)
		bs, err := json.Marshal(dataMap)
		utee.Chk(err)
		info := HotColdInfo{}
		err = json.Unmarshal(bs, info)
		utee.Chk(err)
		return info
	}

	coldDataChan := make(chan interface{})
	for key := range p.coldChan {
		log.Println("@hibernate key :", key)
		info := fetchStruct(key)
		coldDataChan <- info
	}

//	batchSize := 10
//	index := 0
//	for data:= range coldDataChan {
//		//todo batch sql insert
//	}
	return nil
}

func (p *RedisqlWriter) write(data interface{}) error {
	info, ok := data.(HotColdInfo)
	if !ok {
		return errors.New("bad data format")
	}
	if info.IdentityNo == "" {
		return errors.New("info should have identity property")
	}

	r:=p.hotWriter.Get()
	defer r.Close()

	r.Send("ZADD",p.zsetKey(),info.IdentityNo, p.score(time.Now()))

	//reflect save
	r.Send("HMSET",p.hashKey(),)

	//	hashKey := fmt.Sprint("hcw_"+p.hotPrefix)
	//	http://www.runoob.com/redis/redis-hashes.html

	//	storage rank at sorted set
	//	http://www.runoob.com/redis/redis-sorted-sets.html

	return nil
}
