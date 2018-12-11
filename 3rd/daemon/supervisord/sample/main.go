package main

import (
	"log"
	"time"
	"path/filepath"
	"os"
	"strings"
)

func main() {
	for {
		log.Println("Hello Figo.xu & Mr egg @Dir:", getCurrentDirectory(),"  @sample:",os.Args[0])
		time.Sleep(time.Second * time.Duration(5))
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
