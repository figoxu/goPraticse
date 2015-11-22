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
	"time"
)

func main() {
	log.Println("hello")

	ttlQ := NewFTtlQ("./", "sampleTtlQ")
	ttlQ.Enq("figo", []byte("hello world"), 30)
	v, _ := ttlQ.Deq("figo")

	fmt.Println("@times1:",string(v))
	v, _ = ttlQ.Deq("figo")
	fmt.Println("@times2:",string(v))
	ttlQ.Enq("figo", []byte("hello world"), 30)
	v, _ = ttlQ.Deq("figo")
	fmt.Println("@times3:",string(v))
	ttlQ.Enq("figo", []byte("hello world"), 30)
	log.Println("sleep 31 second to use time out")
	time.Sleep(time.Second*time.Duration(31))
	v, _ = ttlQ.Deq("figo")
	fmt.Println("time out check @times3:",string(v))



	st := utee.TickSec()
	for i:=0 ;i<=10*10000;i++ {
		dvid := fmt.Sprintf("sysDevice",i)
		ttlQ.Enq(dvid,[]byte(dvid),60*60*24*7)
		if i%10000 ==0 {
			log.Println("10000 enq finish")
		}
	}
	log.Println("10 0000 device enqueue a message ,cost @t:",(utee.TickSec()-st) )

	latch := utee.NewThrottle(10000)
	st = utee.TickSec()
	for i:=0 ;i<=10*10000;i++ {
		dvid := fmt.Sprintf("sysDevice",i)
		latch.Acquire()
		exc := func(){
			defer latch.Release()
			ttlQ.Enq(dvid,[]byte(dvid),60*60*24*7)
		}
		go exc()
	}
	log.Println("10 0000 device enqueue a message with gorutime,cost @t:",(utee.TickSec()-st) )

	st = utee.TickSec()


}

func NewFTtlQ(basePath, qname string) *FileTtlQ {
	onEvicted := func(key interface{}, value interface{}) {
		qr := value.(*qr.Qr)
		qr.Close()
	}
	m, err := glru.NewWithEvict(1000000, onEvicted)
	utee.Chk(err)
	db, err := leveldb.OpenFile(fmt.Sprint(basePath, qname), nil)
	utee.Chk(err)
	q := &FileTtlQ{
		Ldb:    db,
		qCache: m,
	}
	return q
}

type FileTtlQ struct {
	Ldb    *leveldb.DB
	qCache *glru.Cache
}

func (p *FileTtlQ) getQ(uid interface{}) *qr.Qr {
	v, ok := p.qCache.Get(uid)
	if !ok {
		q, _ := qr.New(
			fmt.Sprint("./q/", uid),
			"qTest",
			qr.OptionBuffer(1000),
		)
		p.qCache.Add(uid, q)
		return q
	}
	q := v.(*qr.Qr)
	return q
}

//ttl unit is second
func (p *FileTtlQ) Enq(uid interface{}, data []byte, ttl ...uint32) error {
	fmt.Sprintf("q%v", uid)
	q := p.getQ(uid)
	k := string(uuid.NewUUID()) //16 byte
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
		case k :=<- p.getQ(uid).Dequeue() :
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
			time.Sleep(time.Duration(1)*time.Nanosecond)
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
