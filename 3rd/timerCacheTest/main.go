package main
import (
	"log"
	"github.com/quexer/utee"
	"github.com/robfig/cron"
	_ "net/http/pprof"
	_ "expvar"
	"net/http"
	"github.com/pborman/uuid"
	"fmt"
)

var cache = utee.NewTimerCache(1*30, func(k,v interface{}){
//	log.Println("time out @k:",k," @v:",v)
})

func main(){
	log.Println("cool")
	initCron()
}

func initCron() {
	c := cron.New()
	cron_10Minutes := "0 0/1 * * * *"
	c.AddFunc(cron_10Minutes, func() {
		log.Println(" 1 minute invoke begin")
		runRec(func(){
			log.Println("task invoked")
			k:= fmt.Sprintf("mk%v", uuid.NewUUID())
			for i:=0;i<10000*50;i++{
				cache.Put(fmt.Sprintf(k,i),i)
			}


		})
	})
	c.Start()
	log.Fatal(http.ListenAndServe(":6666", nil))
}


func runRec(exec func()) {
	rc := func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}
	defer rc()
	exec()
}
