package main

import (
	"github.com/donnie4w/go-logger/logger"
	"github.com/quexer/utee"
	"log"
	"os"
)

func main() {
	rollCrcLog("./", "crc.log")
	for i := 0; i < 100; i++ {
		logger.Warn("This is Warning Message")
	}
	logger.Error("This is ", "Error Message")
	log.Println("End")
}

func rollCrcLog(fileDir, fileName string) {
	err := os.MkdirAll(fileDir, 0777)
	utee.Chk(err)
	logger.SetConsole(true)
	logger.SetRollingDaily(fileDir, fileName)
	logger.SetLevel(logger.WARN)
}
