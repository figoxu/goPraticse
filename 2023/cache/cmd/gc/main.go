package main

import (
	"flag"
	"fmt"
	"github.com/golang/groupcache"
	"log"
	"net/http"
)

func main() {

	var seed string
	var port int

	flag.StringVar(&seed, "seed", "localhost:8888", "seedAddress")
	flag.IntVar(&port, "port", 8888, "port listen")

	flag.Parse()

	fmt.Println("seed is ", seed, " port is ", port)

	// step 2: 创建 groupcache 实例
	group := groupcache.NewGroup("example-group", 1024*1024, groupcache.GetterFunc(getFromStorage))

	// step 3: 使用 groupcache 获取数据
	var result []byte
	err := group.Get(nil, "key-1", groupcache.AllocatingByteSliceSink(&result))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(result))

	// 启动 HTTP 服务
	http.HandleFunc("/cache", func(w http.ResponseWriter, r *http.Request) {
		// 使用 HTTP GET 请求获取数据
		var result []byte
		err := group.Get(nil, "key-1", groupcache.AllocatingByteSliceSink(&result))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getFromStorage(ctx groupcache.Context, key string, dest groupcache.Sink) error {
	// 这里可以写具体的查询逻辑，从数据库或者文件系统中读取对应的数据
	data := fmt.Sprintf("This is value for key: %s", key)
	return dest.SetBytes([]byte(data))
}
