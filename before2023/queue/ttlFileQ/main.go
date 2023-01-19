package main

import (
	"encoding/json"
	"encoding/hex"
	"fmt"
	"github.com/alicebob/qr"
	"github.com/figoxu/utee"
	"github.com/pborman/uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"time"
	"sync"
)
//todo 每个文件目录不能超过1000个文件，目录摆放需要规范  分成多级的 1000整除的目录,一个文件夹，不能超过1000 个文件
func main() {
	log.Println("hello")
	ttlQ := NewFTtlQ("/home/figo/data", "sampleTtlQ")
	st := utee.TickSec()
	for i := 0; i <= 100*10000; i++ {
		dvid := fmt.Sprintf("sysDevice%v", i)
		ttlQ.Enq(dvid, []byte(dvid), 60*60*24*7)
		if i%10000 == 0 && i>0 {
			log.Println("10000 enq finish")
		}
	}
	log.Println("10 0000 device enqueue a message ,cost @t:", (utee.TickSec() - st))
	log.Println("sleep 1.2 minute,then retry with gorutine")
	ttlQ.hibernateAll()
	time.Sleep(time.Second * time.Duration(70))
	var wg sync.WaitGroup
	st = utee.TickSec()
	for i := 0; i <= 100*10000; i++ {
		dvid := fmt.Sprintf("sysDevice%v", i)
		wg.Add(1)
		exc := func() {
			defer wg.Done()
			ttlQ.Enq(dvid, []byte(dvid), 60*60*24*7)
		}
		go exc()
	}
	wg.Wait()
	log.Println("100 0000 device enqueue a message with gorutime,cost @t:", (utee.TickSec() - st))
	ttlQ.hibernateAll()
	log.Println("shut down all queue")
	time.Sleep(time.Second * time.Duration(10))
	for len(ttlQ.shut_q) > 0 {
		time.Sleep(time.Second * time.Duration(1))
	}
	log.Println("all shut down done")

}

func NewFTtlQ(basePath, qname string) *FileTtlQ {
	c := make(chan *qr.Qr, 1000000)
	timerCache := utee.NewTimerCache(60, func(k, v interface{}) {
		c <- v.(*qr.Qr)
	})
	d := fmt.Sprintf("%s/%s/%s", basePath, qname, "ldb")
	log.Println("start @dbpath:", d)
	err := os.MkdirAll(d, 0777)
	utee.Chk(err)
	db, err := leveldb.OpenFile(d, nil)
	utee.Chk(err)
	q := &FileTtlQ{
		Ldb:      db,
		timerCache:   timerCache,
		basePath: basePath,
		qname:    qname,
		shut_q:   c,
	}

	closeQ := func(fq *qr.Qr){
		defer func() {
			if err := recover(); err != nil {
				log.Println(err, " (recover) @fq:",fq)
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
	timerCache   *utee.TimerCache
	basePath string
	qname    string
	shut_q   chan *qr.Qr
}

func getOkDir(s string) string{
	v := hex.EncodeToString(utee.Md5([]byte(s)))
	v = v[8:24]  //16位md5
	d := fmt.Sprint(v[0:2],"/",v[2:4],"/",v[4:6],"/",v[6:8],"/",v[8:10],"/",v[10:12],"/",v[12:14])
	return d
}

func (p *FileTtlQ) getQ(uid interface{}) *qr.Qr {
	qn:=p.parseQName(uid)
	v := p.timerCache.Get(uid)
	if v == nil {
		d := fmt.Sprintf("%s/%s/q/%s", p.basePath, p.qname,getOkDir(qn))
		err := os.MkdirAll(d, 0777)
		utee.Chk(err)
		q, err := qr.New(
			d,
			qn,
			qr.OptionBuffer(1000),
		)
		utee.Chk(err)
		p.timerCache.Put(uid, q)
		return q
	}
	q := v.(*qr.Qr)
	return q
}
func (p *FileTtlQ) hibernateAll()  {
	for _,k := range p.timerCache.Keys() {
		v := p.timerCache.Get(k)
		if v == nil {
			log.Println("@k:", k, "  is not ok @v:",v)
			continue
		}
		p.shut_q <- v.(*qr.Qr)
		p.timerCache.Remove(k)
	}
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
