package main

import (
	"encoding/json"
	"fmt"
	"github.com/alicebob/qr"
	"github.com/figoxu/utee"
	glru "github.com/hashicorp/golang-lru"
	"github.com/pborman/uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"time"
)
//todo 替换LRU Cache为，Timer Cache
func main() {
	log.Println("hello")

	q2, _ := qr.New(
		"/home/figo/data/sampleTtlQ/q/",
		"sample2",
		qr.OptionBuffer(1000),
	)
	q2.Enqueue("hello")
	q2.Close()

	q3, _ := qr.New(
		"/home/figo/data/sampleTtlQ/q/",
		"sample3",
		qr.OptionBuffer(1000),
	)
	q3.Enqueue("hello")

	ttlQ := NewFTtlQ("/home/figo/data", "sampleTtlQ")

	ttlQ.Enq("test", []byte("hello world"), 30)
	ttlQ.Enq("test", []byte("hello world"), 30)
	log.Println("get queue @q:", ttlQ.getQ("test"))
	ttlQ.getQ("test").Close()

	ttlQ.Enq("figo", []byte("hello world"), 30)
	v, _ := ttlQ.Deq("figo")

	fmt.Println("@times1:", string(v))
	v, _ = ttlQ.Deq("figo")
	fmt.Println("@times2:", string(v))
	ttlQ.Enq("figo", []byte("hello world"), 30)
	v, _ = ttlQ.Deq("figo")
	fmt.Println("@times3:", string(v))
	ttlQ.Enq("figo", []byte("hello world"), 30)
//	log.Println("sleep 31 second to use time out")
//	time.Sleep(time.Second * time.Duration(31))
	v, _ = ttlQ.Deq("figo")
	fmt.Println("time out check @times3:", string(v))
	if fq, ok := ttlQ.qCache.Get("figo"); !ok {
		log.Println("figo queue is not ok")
	} else {
		ttlQ.Enq("figo", []byte("hello world"), 30)
		fq.(*qr.Qr).Close()
		log.Println("figo queue is close ok")
	}

//	st := utee.TickSec()
//	for i := 0; i <= 10*10000; i++ {
//		dvid := fmt.Sprintf("sysDevice%v", i)
//		ttlQ.Enq(dvid, []byte(dvid), 60*60*24*7)
//		if i%10000 == 0 {
//			log.Println("10000 enq finish")
//		}
//	}
//	log.Println("10 0000 device enqueue a message ,cost @t:", (utee.TickSec() - st))

	latch := utee.NewThrottle(10000)
	st := utee.TickSec()
	for i := 0; i <= 10*10000; i++ {
		dvid := fmt.Sprintf("sysDevice %v", i)
		latch.Acquire()
		exc := func() {
			defer latch.Release()
			ttlQ.Enq(dvid, []byte(dvid), 60*60*24*7)
		}
		go exc()
	}
	log.Println("10 0000 device enqueue a message with gorutime,cost @t:", (utee.TickSec() - st))

	for _,k := range ttlQ.qCache.Keys() {
//		dvid := fmt.Sprintf("sysDevice", k)
		v, ok := ttlQ.qCache.Get(k)
		if !ok {
			log.Println("@k:", k, "  is not ok @v:",v)
			continue
		}
		if v == nil {
			continue
		}
		ttlQ.shut_q <- v.(*qr.Qr)
	}
	log.Println("shut down all queue")
	time.Sleep(time.Second * time.Duration(5))
	for len(ttlQ.shut_q) > 0 {
		time.Sleep(time.Second * time.Duration(1))
	}
	log.Println("all shut down done")

}

func NewFTtlQ(basePath, qname string) *FileTtlQ {
	c := make(chan *qr.Qr, 1000000)
	onEvicted := func(key interface{}, value interface{}) {
		c <- value.(*qr.Qr)
	}
	m, err := glru.NewWithEvict(1000000, onEvicted)
	utee.Chk(err)
	d := fmt.Sprintf("%s/%s/%s", basePath, qname, "ldb")
	log.Println("dbpath @d:", d)
	err = os.MkdirAll(d, 0777)
	utee.Chk(err)
	db, err := leveldb.OpenFile(d, nil)
	utee.Chk(err)
	q := &FileTtlQ{
		Ldb:      db,
		qCache:   m,
		basePath: basePath,
		qname:    qname,
		shut_q:   c,
	}

	closeQ := func(fq *qr.Qr){
		defer func() {
			if err := recover(); err != nil {
				log.Println(err, " (recover)")
			}
		}()
		fq.Close()
	}
	clean := func() {
		for fq := range c {
			closeQ(fq)
		}
	}
	go clean()
	return q
}

type FileTtlQ struct {
	Ldb      *leveldb.DB
	qCache   *glru.Cache
	basePath string
	qname    string
	shut_q   chan *qr.Qr
}

//g_as_basic_q            = make(chan *OpAsBasic, MAX_OPQ_AS_BASIC)

func (p *FileTtlQ) getQ(uid interface{}) *qr.Qr {
	v, ok := p.qCache.Get(uid)
	if !ok {
		d := fmt.Sprintf("%s/%s/q/", p.basePath, p.qname)
		err := os.MkdirAll(d, 0777)
		utee.Chk(err)
		q, err := qr.New(
			d,
			p.parseQName(uid),
			qr.OptionBuffer(1000),
		)
		utee.Chk(err)
		p.qCache.Add(uid, q)
		return q
	}
	q := v.(*qr.Qr)
	return q
}
func (p *FileTtlQ) parseQName(uid interface{}) string {
	return fmt.Sprintf("q_%v", uid)
}

//ttl unit is second
func (p *FileTtlQ) Enq(uid interface{}, data []byte, ttl ...uint32) error {
	q := p.getQ(uid)
	k := string(uuid.NewUUID().String()) //16 byte
	q.Enqueue(k)
	t := int64(-1) //never ood (out of day)
	if len(ttl) > 0 {
		t = utee.TickSec() + int64(ttl[0])
	}
	qv := QValue{
		Data: data,
		Dod:  t,
	}
	b, err := json.Marshal(qv)
	utee.Chk(err)
	p.Ldb.Put([]byte(k), b, nil)
	return nil
}

func (p *FileTtlQ) Deq(uid interface{}) ([]byte, error) {
	retry := false
	for {
		select {
		case k := <-p.getQ(uid).Dequeue():
			key, _ := k.(string)
			b, err := p.Ldb.Get([]byte(key), nil)
			if err != nil {
				return nil, err
			}
			if b != nil {
				v := &QValue{}
				json.Unmarshal(b, v)
				p.Ldb.Delete([]byte(key), nil)
				if v.Dod > utee.TickSec() || v.Dod == -1 {
					return v.Data, nil
				}
			}
		default:
			if retry {
				return nil, nil
			}
			time.Sleep(time.Duration(1) * time.Nanosecond)
			retry = true
		}
	}
}

func (p *FileTtlQ) Len(uid interface{}) (int, error) {
	return -1, nil
}

type QValue struct {
	Data []byte `json:"v"`
	Dod  int64  `json:"d"` //date of death
}

type Queue interface {
	Enq(uid interface{}, data []byte, ttl ...uint32) error
	Deq(uid interface{}) ([]byte, error)
	Len(uid interface{}) (int, error)
}
