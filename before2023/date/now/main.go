package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"log"
	"time"
)

func main() {
	t := time.Now()
	n := now.New(t)
	log.Println(n.EndOfMonth())
	log.Println(midDay(t))
}

func midDay(t time.Time) time.Time {
	v := fmt.Sprint(t.Format("2006-01-02"), " 12:00:00")
	t2, _ := time.Parse("2006-01-02 15:04:05", v)
	return t2
}
