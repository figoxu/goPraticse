package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/figoxu/Figo"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/quexer/utee"
	"log"
	"strings"
	"time"
)

type HCWriter interface {
	write(interface{})
}

type RedisqlWriter struct {
	hotPrefix  string
	hotWriter  *redis.Pool
	coldWriter *sql.DB
	coldChan   chan string
}

func NewRedisqlWriter(hotWriter *redis.Pool, coldWriter *sql.DB) RedisqlWriter {
	writer := RedisqlWriter{
		hotWriter:  hotWriter,
		coldWriter: coldWriter,
		coldChan:make(chan string,1000),
	}
	go writer.coldSchedule()
	go writer.hibernate()
	return writer
}

func (p *RedisqlWriter) score(t time.Time) string {
	return t.Format("20060102150405")
}

func (p *RedisqlWriter) coldSchedule() error {
	fetchColdKey := func() {
		endScore := p.score(time.Now().Add(time.Duration(-1) * time.Minute))
		r := p.hotWriter.Get()
		defer r.Close()
		log.Println("ZRANGE", p.zsetKey(), 0, endScore)
		keys, err := redis.Strings(r.Do("ZRANGE", p.zsetKey(), 0, endScore))
		log.Println(">>>001")
		defer Figo.Catch()
		log.Println(">>>002")
		utee.Chk(err)
		log.Println(">>>003")
//		redisParam := []interface{}{p.zsetKey()}
		for _, key := range keys {
//			redisParam = append(redisParam,key)
			p.coldChan <- key
		}
//		log.Println(">>>004")
//		r.Do("ZREM", redisParam...)
		log.Println(">>>008")
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
	fetchStruct := func(pk string) HotColdInfo {
		defer Figo.Catch()
		r := p.hotWriter.Get()
		defer r.Close()
		log.Println("HSCAN", p.hashKey(), 0, "MATCH", fmt.Sprint(pk, "_*"))
		dataMap, err := redis.StringMap(r.Do("HSCAN", p.hashKey(), 0, "MATCH", fmt.Sprint(pk, "_*")))
		//how to mapping redis string map to golang struct
		utee.Chk(err)
		info := HotColdInfo{}
		for key, value := range dataMap {
			propName := strings.Replace(key, fmt.Sprint(pk, "_"), "", -1)
			isProp := func(name string) bool {
				return strings.ToLower(propName) == strings.ToLower(name)
			}
			if isProp("Id") {
				v, _ := Figo.TpInt(value)
				info.Id = v
			} else if isProp("IdentityNo") {
				info.IdentityNo = value
			} else if isProp("Name") {
				info.Name = value
			} else if isProp("Age") {
				v, _ := Figo.TpInt(value)
				info.Age = v
			} else if isProp("Salary") {
				v, _ := Figo.TpInt(value)
				info.Salary = v
			} else if isProp("Description") {
				info.Description = value
			}
		}
		return info
	}

	storage := func(){
		log.Println("try to storage")
		coldDataChan := make(chan HotColdInfo,100000)
		for key := range p.coldChan {
			log.Println("@hibernate key :", key)
			info := fetchStruct(key)
			coldDataChan <- info
		}

		tx, err := p.coldWriter.Begin()
		utee.Chk(err)
		stmt, err := tx.Prepare("INSERT INTO HotColdInfo(Id,IdentityNo,Name,Age,Salary,Description) values(?,?,?,?,?,?) ON DUPLICATE UPDATE HotColdInfo SET  Id=?,IdentityNo=?,Name=?,Age=?,Salary=?,Description=? WHERE Id=? ")
		//todo should add batch here
		for data := range coldDataChan {
			_, err := stmt.Exec(data.Id, data.IdentityNo, data.Name, data.Age, data.Salary, data.Description, data.Id, data.IdentityNo, data.Name, data.Age, data.Salary, data.Description, data.Id)
			if err != nil {
				log.Println("@stmt exec error :", err)
			}
		}
		tx.Commit()

	}
	for{
		storage();
	}
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

	r := p.hotWriter.Get()
	defer r.Close()
	idkey := func(prop string) string {
		return fmt.Sprint(info.IdentityNo, "_", prop)
	}
//	log.Println("ZADD", p.zsetKey(), p.score(time.Now()), info.IdentityNo)
	r.Send("ZADD", p.zsetKey(), p.score(time.Now()), info.IdentityNo)
//	log.Println("HMSET", p.hashKey(), idkey("Id"), info.Id, idkey("IdentityNo"), info.IdentityNo, idkey("Name"), info.Name, idkey("Age"), info.Age, idkey("Salary"), info.Salary, idkey("Description"), info.Description)
	r.Send("HMSET", p.hashKey(), idkey("Id"), info.Id, idkey("IdentityNo"), info.IdentityNo, idkey("Name"), info.Name, idkey("Age"), info.Age, idkey("Salary"), info.Salary, idkey("Description"), info.Description)
	r.Flush()
	return nil
}
