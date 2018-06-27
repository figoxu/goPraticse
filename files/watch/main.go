package main

import (
	"log"
	"github.com/fsnotify/fsnotify"
	"path/filepath"
	"os"
	"github.com/quexer/utee"
	"github.com/figoxu/Figo"
	"time"
	"strings"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					info, err := os.Stat(event.Name)
					utee.Chk(err)
					if info.IsDir() {
						watchDir(watcher, event.Name)
					}
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					asyncChange(event.Name)
					//log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()
	watchDir(watcher, "/Users/xujianhui/develop/env/history/xujianhui")
	<-done
	select {}
}

func watchDir(watcher *fsnotify.Watcher, dir string) {
	err := watcher.Add(dir)
	utee.Chk(err)
	for _, dir := range listAllDir(dir) {
		err = watcher.Add(dir)
		utee.Chk(err)
	}
}

func listAllDir(dir string) []string {
	dirs := make([]string, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirs = append(dirs, path)
		}
		return nil
	})
	return dirs
}

var (
	fileLineMap = make(map[string]int)
)

func asyncChange(filename string) {
	fut := Figo.FileUtee{}
	vs, err := fut.ReadLinesSlice(filename)
	utee.Chk(err)
	idx := fileLineMap[filename]
	log.Println("@idx:", idx)

	type LogItem struct {
		Ct       time.Time
		Command  string
		FileName string
	}
	items := make([]LogItem, 0)
	for i := idx; i < len(vs); i++ {
		k := vs[i]
		k = strings.Replace(k, "#", "", -1)
		sec, err := Figo.TpInt64(k)
		utee.Chk(err)
		t := time.Unix(sec, 0)
		i++
		v := vs[i]
		items = append(items, LogItem{
			Ct:       t,
			Command:  v,
			FileName: filename,
		})
		fileLineMap[filename] = i + 1
	}
	for _, item := range items {
		log.Println("==>", Figo.JsonString(item))
	}
}
