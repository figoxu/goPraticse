package main

import (
	"crypto/tls"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"time"
)

func main() {
	m := martini.Classic()
	m.Get("/hello", helloHandler)
	http.Handle("/", m)
	log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")
	go h1Test()
	go h2Test()
	err := http.ListenAndServeTLS(":10443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler() (int, string) {
	log.Println("invoke")

	return 200, "hello"
}

func h1Test() {
	time.Sleep(time.Second * time.Duration(2))
	log.Println("h1 invoke sample")
	if rsp, err := http.Get("https://127.0.0.1:10443/hello"); err != nil {
		log.Println("@err:", err)
	} else {
		log.Println("@rsp:", rsp)
	}
}

func h2Test() {
	time.Sleep(time.Second * time.Duration(5))
	log.Println("h2 invoke sample")
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	if rsp, err := client.Get("https://127.0.0.1:10443/hello"); err != nil {
		log.Println("@err:", err)
	} else {
		log.Println("@rsp:", rsp)
	}
}
