package main

import (
	"log"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/figoxu/utee"
)

func main() {
	logger := &lumberjack.Logger{
		Filename:   "./logs/sample.log",
		MaxSize:    500, // megabytes
		MaxBackups: 30,
		MaxAge:     28, //days
	}
	log.SetOutput(logger)
	for i := 0; i < 2000000; i++ {
		log.Println("Hello", i)
		if i%500000 == 0 && i!=0{
			err := logger.Rotate();
			utee.Chk(err)
		}
	}
}
