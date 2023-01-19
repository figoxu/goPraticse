package main
import (
	"log"
	"sync"
	"github.com/figoxu/utee"
	"net/url"
)

func main(){
	log.Println("hello")

	for k:=0 ;k<100;k++ {
		var wg sync.WaitGroup
		f := func(){
			defer wg.Done()
			utee.HttpPost("http://localhost:5050/hello/world",url.Values{})
		}
		for i:=0;i<1000;i++ {
			wg.Add(1)
			go f()
		}
		wg.Wait()
		log.Println("finish 100 invoke")

	}
	log.Println("finish total invoke")
}


