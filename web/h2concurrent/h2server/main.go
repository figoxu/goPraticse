package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"time"
)

func main() {
	m := martini.Classic()
	m.Get("/helloGet", getHandler("Get"))
	m.Post("/helloPost", getHandler("Post"))
	http.Handle("/", m)
	log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")
	err := http.ListenAndServeTLS(":10443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getHandler(val string) func() (int, string) {
	return func() (int, string) {
		time.Sleep(time.Millisecond * time.Duration(10))
		return 200, fmt.Sprint("hello", val)
	}
}
