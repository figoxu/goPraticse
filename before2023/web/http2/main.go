package main

import (
	"fmt"
	"github.com/figoxu/utee"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	m := martini.Classic()
	m.Get("/helloGet", getHandler("Get"))
	m.Post("/helloPost", getHandler("Post"))
	http.Handle("/", m)
	log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")
	go uteeH2Test()
	err := http.ListenAndServeTLS(":10443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getHandler(val string) func() (int, string) {
	return func() (int, string) {
		return 200, fmt.Sprint("hello", val)
	}
}

func uteeH2Test() {
	time.Sleep(time.Second * time.Duration(2))
	log.Println("")
	log.Println("utee h2 invoke sample")
	if v, err := utee.Http2Get("https://127.0.0.1:10443/helloGet"); err != nil {
		log.Println("utee h2 get @err:", err)
	} else {
		log.Println("utee h2 get @rsp:", string(v))
	}

	if v, err := utee.Http2Post("https://127.0.0.1:10443/helloPost", url.Values{}); err != nil {
		log.Println("utee h2 post @err:", err)
	} else {
		log.Println("utee h2 post @rsp:", string(v))
	}

}
