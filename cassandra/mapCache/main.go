package main
import (
	"fmt"
//	"log"
//	"time"
	"github.com/hailocab/gocassa"
	"github.com/quexer/utee"
	"log"
)

type MCache struct{
	Key string
	Value string
}

func main(){
	ks_mpush, err := gocassa.ConnectToKeySpace("test", []string{"127.0.0.1"}, "", "")


//	ks_mpush.Table()

	utee.Chk(err)
	ks_mpush.DebugMode(true)
	cache := ks_mpush.MapTable("cache","Key",MCache{})
	cache.Create()

	figoC := &MCache{
		Key:"figo",
		Value:"ok",
	}
	err =cache.Set(figoC).Run()
	if err != nil {
		log.Fatal(err)
	}

	tmp := &MCache{}
	err =cache.Read("figo",tmp).Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tmp.Key,"---",tmp.Value)
}
