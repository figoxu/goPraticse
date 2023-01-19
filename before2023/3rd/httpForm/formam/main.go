package main

import (
	"github.com/go-martini/martini"
	formam "github.com/monoculum/formam"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	log.Println("hello")
	m := martini.Classic()
	m.Handlers(martini.Recovery())
	m.Post("/test/post", Handler)
	http.Handle("/", m)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type InterfaceStruct struct {
	ID   int
	Name string
}

type Company struct {
	Public     bool      `formam:"public"`
	Website    url.URL   `formam:"website"`
	Foundation time.Time `formam:"foundation"`
	Name       string
	Location   struct {
		Country string
		City    string
	}
	Products []struct {
		Name string
		Type string
	}
	Founders  []string
	Employees int64

	Interface interface{}
}

func Handler(w http.ResponseWriter, r *http.Request) error {
	m := Company{
		Interface: &InterfaceStruct{}, // its is possible to access to the fields although it's an interface field!
	}
	r.ParseForm()
	decoder := formam.NewDecoder(&formam.DecoderOptions{TagName: "formam"})
	decoder.RegisterCustomType(func(vals []string) (interface{}, error) {
		return time.Parse("2006-01-02", vals[0])
	}, []interface{}{time.Time{}}, nil)
	if err := decoder.Decode(r.Form, &m); err != nil {
		return err
	}
	log.Println("@company:", m)
	return nil
}
