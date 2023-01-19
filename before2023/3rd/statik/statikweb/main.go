package main

import (
	"log"
	"github.com/rakyll/statik/fs"
	"github.com/quexer/utee"
	"net/http"
	_ "./statik"
)

func main(){
	log.Println("Hello World")

	statikFS, err := fs.New()
	utee.Chk(err)
	http.ListenAndServe(":8080", http.FileServer(statikFS))
}


