package main

import (
	"github.com/howeyc/fsnotify"
	"github.com/quexer/utee"
	"log"
	"os"
	"github.com/figoxu/Figo"
)

var (
	addCache = utee.NewTimerCache(1, func(k, v interface{}) {
		fileInfo,ok := v.(FileInfo)
		log.Println("@parse :",ok)
		filePath,err:= Figo.TpString(k)
		utee.Chk(err)
		if getFileSize(filePath) == fileInfo.fileSize {
			log.Println("Time To Do Business @file:", k)
		}

	})
	delCache = utee.NewTimerCache(1, func(k, v interface{}) {
		log.Println("@file:", k, "  was delete... do you want to clear any data ?")
	})
	renameCache = utee.NewTimerCache(1, func(k, v interface{}) {
		filePath,err:= Figo.TpString(k)
		utee.Chk(err)
		if getFileSize(filePath) == 0 {
			log.Println("@file:", k, " was delete by rename .")
			delCache.Put(k, v)
		} else {
			log.Println("@file:", k, " was add by rename .")
			addCache.Put(k, v)
		}
	})
)

type FileInfo struct {
	fileSize int64
	lastMt   int64
}

func main() {
	log.Println("Hello")
	watcher, err := fsnotify.NewWatcher()
	utee.Chk(err)
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				processImg(ev)

			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()
	dir4watch := "/home/figo/delete_it/dir4watch"
	err = watcher.Watch(dir4watch)
	utee.Chk(err)
	select {}
}

func getFileSize(fileName string) int64 {
	if fileInfo, err := os.Stat(fileName); err != nil {
		return 0
	} else {
		return fileInfo.Size()
	}
}

func processImg(event *fsnotify.FileEvent) {
	fileName := event.Name
	fileInfo := FileInfo{
		fileSize: getFileSize(fileName),
		lastMt:   utee.Tick(),
	}
	if event.IsCreate() || event.IsModify() {
		log.Println(fileName, " Was Create")
		addCache.Put(fileName, fileInfo)
	}
	if event.IsDelete() {
		log.Println(fileName, " Was Delete")
		delCache.Put(fileName, fileInfo)
	}
	if event.IsRename() {
		log.Println(fileName, " Was Rename")
		renameCache.Put(fileName, fileInfo)
	}
}
