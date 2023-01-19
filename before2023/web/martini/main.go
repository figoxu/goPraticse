package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/quexer/utee"
	"log"
	"net/http"
	_ "expvar"
	_ "net/http/pprof"
	"runtime/debug"
	"time"
)

func main() {
	fmt.Println("hello")
	go freeAlloc()
	m := martini.Classic()
	m.Handlers(gzip.All(), martini.Recovery())
	Mount(m, 100)
	log.Printf("start gateway on %s\n", 5050)

	log.Fatal(http.ListenAndServe(":5050", nil))
//	m.RunOnAddr(":5050")
}

func Mount(m *martini.ClassicMartini, concurrent int) {


	fmt.Println("concurrent @v:", concurrent)
	m.Use(utee.MidConcurrent(concurrent))
	m.Use(render.Renderer())
	m.Use(utee.MidTextDefault)
	//	//map request web utilities
	m.Use(func(w http.ResponseWriter, c martini.Context) {
		web := &utee.Web{W: w}
		c.Map(web)
	})
	m.Use(func(c martini.Context){
		var msgs []string
		for i:=0;i<100000;i++ {
			msgs = append(msgs,fmt.Sprint("testMsg",i))
		}
		c.Map(msgs)
	})

	m.Group("/hello", func(r martini.Router) {
		r.Post("/world", test)
		r.Get("/world", test)
	})
	http.Handle("/", m)
//	http.HandleFunc("/test/a", testhandler)
}

func test(r *http.Request, web *utee.Web,msgs []string) (int, string) {
	for k,v := range msgs {
		log.Println("@k:",k," @v:",v)
	}
	return web.Json(200, "hello world")
}

func testhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("i am in")
}
var (
FREE_ALLOC = 1
)
func freeAlloc() {
	log.Println("free alloc at each", FREE_ALLOC, " Minutes")
	if FREE_ALLOC <= 0 {
		log.Fatal("exit with reason : bad param @free_alloc:", FREE_ALLOC)
	}
	for {
		debug.FreeOSMemory()
		time.Sleep(time.Minute * time.Duration(FREE_ALLOC))
	}
}
