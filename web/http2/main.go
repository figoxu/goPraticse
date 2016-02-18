package main

import (
	"crypto/tls"
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
	go h1Test()
	go h2Test()
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

func h1Test() {
	time.Sleep(time.Second * time.Duration(2))
	log.Println("")
	log.Println("h1 invoke sample")
	if rsp, err := http.Get("https://127.0.0.1:10443/helloGet"); err != nil {
		log.Println("h1 get @err:", err)
	} else {
		log.Println("h1 get @rsp:", rsp)
	}
	if rsp, err := http.PostForm("https://127.0.0.1:10443/helloPost", url.Values{}); err != nil {
		log.Println("h1 post @err:", err)
	} else {
		log.Println("h1 post @rsp:", rsp)
	}
}

func h2Test() {
	time.Sleep(time.Second * time.Duration(5))
	log.Println("")
	log.Println("h2 invoke sample")
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	if rsp, err := client.Get("https://127.0.0.1:10443/helloGet"); err != nil {
		log.Println("h2 get @err:", err)
	} else {
		log.Println("h2 get @rsp:", rsp)
	}
	if rsp, err := client.PostForm("https://127.0.0.1:10443/helloPost", url.Values{}); err != nil {
		log.Println("h2 post @err:", err)
	} else {
		log.Println("h2 post @rsp:", rsp)
	}

}

func uteeH2Test() {
	time.Sleep(time.Second * time.Duration(10))
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
