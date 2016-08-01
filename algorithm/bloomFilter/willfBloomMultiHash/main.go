package main

import (
	"crypto/md5"
	_ "expvar"
	"fmt"
	"github.com/zhenjl/bloom"
	"github.com/zhenjl/bloom/standard"
	"github.com/zhenjl/cityhash"
	"hash"
	"hash/fnv"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
)

func main() {
	go generateGarbage()
	log.Fatal(http.ListenAndServe(":6666", nil))
}

//10亿设备量的过滤器，占用1.8G*1 = 1.8G 内存   存在碰撞的数量: 236*1000 (cityhash) [大致值]
//10亿设备量的过滤器，占用1.8G*2 = 3.6G 内存   存在碰撞的数量: 236 (cityhash+md5)
//10亿设备量的过滤器，占用1.8G*3 = 5.4G 内存   存在碰撞的数量: 0 (cityhash+md5+fnv.New64())
func generateGarbage() {
	var capSize uint = 1000000000

	//
	//	fvv := []string{}
	//	fvv = append(fvv,"hello")

	filter := NewBloomWrap(capSize)
	v := []byte("Love")
	b := filter.Add(v).Check(v)
	log.Println("check @v:", b)
	bad := 0
	for i := 0; i < 1000000000; i++ {
		if i%1000000 == 0 {
			log.Println(i)
			debug.FreeOSMemory()
		}
		d := []byte(fmt.Sprint("data", i))
		if filter.Check(d) == true {
			bad++
			//			panic(fmt.Sprint("should not exist @d:",string(d)))
		}
		if flag := filter.Add(d).Check(d); flag == false {
			panic(d)
		}
	}
	log.Println("====>>That is all @bad:", bad)
}

type BloomWrap []bloom.Bloom

func NewBloomWrap(capSize uint) BloomWrap {
	wrap := []bloom.Bloom{}
	newFilter := func(h hash.Hash) {
		filter := standard.New(capSize)
		filter.SetHasher(h)
		wrap = append(wrap, filter)
	}
	newFilter(cityhash.New64())
	newFilter(fnv.New64())
	newFilter(md5.New())
	log.Println("bloomLen: ", len(wrap))
	return wrap
}

func (p *BloomWrap) Add(d []byte) *BloomWrap {
	for _, b := range *p {
		b.Add(d)
	}
	return p
}

func (p *BloomWrap) Check(d []byte) bool {
	for _, b := range *p {
		if !b.Check(d) {
			return false
		}
	}
	return true
}
