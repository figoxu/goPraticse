package main

import (
	"github.com/zhenjl/bloom/standard"
	"github.com/zhenjl/cityhash"
	"log"
	"fmt"
	_ "expvar"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
)

func main() {
	go generateGarbage();
	log.Fatal(http.ListenAndServe(":6666", nil))
}
//10亿设备量的过滤器，占用1.8G内存   存在碰撞的数量: 229660

func generateGarbage(){
	var capSize uint = 1000000000
	filter := standard.New(capSize)
	filter.SetHasher(cityhash.New64())
	v := []byte("Love")
	b := filter.Add(v).Check(v)
	log.Println("check @v:", b)
	bad := 0
	for i:=0;i<1000000000;i++ {
		if i%1000000 == 0 {
			log.Println(i)
			debug.FreeOSMemory()
		}
		d := []byte( fmt.Sprint("data",i) )
		if filter.Check(d) ==true{
			bad++
			//			panic(fmt.Sprint("should not exist @d:",string(d)))
		}
		if flag := filter.Add(d).Check(d) ;flag==false {
			panic(d)
		}
	}
	log.Println("====>>That is all @bad:",bad)
}
