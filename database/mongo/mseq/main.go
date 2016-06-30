package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/quexer/utee"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"sync"
)

var (
	rp *redis.Pool
	se *MgoSe
)

const (
	DB    = "mpush"
	C_SEQ = "seq"
)

type Seq struct {
	Id  string `bson:"_id"`
	Seq int64  `bson:"seq"`
}

func main() {
	rp = utee.CreateRedisPool(30, "106.75.27.144:6379", "baulk3?speed")
	con, err := mgo.Dial("mpush:talkingdata@106.75.27.144/mpush?maxPoolSize=50")
	utee.Chk(err)
	se = &MgoSe{
		referSession: con,
	}

	knameA := "figo_test"
	initSeq(knameA)
	nextSeq(knameA)
	initSeq(knameA)
	nextSeq(knameA)
	initSeq(knameA)
	nextSeq(knameA)
	nextSeq(knameA)
	nextSeq(knameA)

	knameB := "figo_testB"
	initSeq(knameB)
	nextSeq(knameB)
	initSeq(knameB)
	nextSeq(knameB)
	initSeq(knameB)
	nextSeq(knameB)
	nextSeq(knameB)
	nextSeq(knameB)
	nextSeq(knameB)

}

func seqKey(name string) string {
	return fmt.Sprint("SEQ_", name)
}

func initSeq(name string) {
	seq := &Seq{}
	if err := se.DB(DB).C(C_SEQ).FindId(name).One(seq); err == mgo.ErrNotFound {
		seq = &Seq{
			Id:  name,
			Seq: 0,
		}
		se.DB(DB).C(C_SEQ).Insert(seq)
	}
	log.Println(seq)
	c := rp.Get()
	defer c.Close()
	c.Do("SET", seqKey(name), seq.Seq)
}

func nextSeq(name string) int {
	c := rp.Get()
	defer c.Close()
	v, _ := redis.Int(c.Do("INCR", seqKey(name)))
	se.DB(DB).C(C_SEQ).UpdateId(name, bson.M{"seq": v})
	return v
}

type MgoSe struct {
	sync.Mutex
	realSession  *mgo.Session
	referSession *mgo.Session
}

func (p *MgoSe) SetMode(consistency mgo.Mode, refresh bool) {
	p.session().SetMode(consistency, refresh)
}

func (p *MgoSe) DB(name string) *mgo.Database {
	return p.session().DB(name)
}

func (p *MgoSe) IsOpen() bool {
	if p.realSession == nil {
		return false
	}
	return true
}

func (p *MgoSe) session() *mgo.Session {
	p.Lock()
	defer p.Unlock()
	if !p.IsOpen() {
		if p.referSession == nil {
			log.Panicln("error : referSession is not configure")
		}
		p.realSession = p.referSession.Copy()
	}
	return p.realSession
}

func (p *MgoSe) Close() {
	p.Lock()
	defer p.Unlock()
	if p.IsOpen() {
		p.realSession.Close()
		p.realSession = nil
	}
}
