package main
import (
	"github.com/peterbourgon/diskv"
	"log"
	"encoding/json"
	"fmt"
	"github.com/figoxu/utee"
//	"strings"
)

type Test struct {
	Name  string	 `json:"name"`
	Tp    string	 `json:"tp"`
	Count int	 `json:"count"`
}
func main(){
	c := diskv.New(diskv.Options{
		BasePath:     "./my-diskv-data-directory",
		Transform:    func(s string) []string { return []string{} },
		CacheSizeMax: 1024 * 1024, // 1MB
	})
	b, _ := json.Marshal(Test{Name:"figo",Tp:"android",Count:1024})
	c.Write("test",b)
	for k:= range c.Keys(nil) {
		log.Println("@k:",k)
		d,_:=c.Read(k)
		v:=&Test{}
		json.Unmarshal(d,v)
		log.Println("@name:",v.Name,"@count:",v.Count,"@tp:",v.Tp)
	}
	d,_:=c.Read("test")
	v:=&Test{}
	json.Unmarshal(d,v)
	log.Println("@name:",v.Name,"@count:",v.Count,"@tp:",v.Tp)


	st := utee.Tick()
//	1000*
	for i:=0;i<100*10000;i++ {
		b, _ := json.Marshal(Test{Name:"figo",Tp:"android",Count:1024})
		c.Write(fmt.Sprint("test",i),b)
//		c.WriteStream(fmt.Sprint("test",i),strings.NewReader( string(b)),false)
	}
	log.Println("100,0000 write cost ",(utee.Tick()-st),"m second")
	st = utee.Tick()
	for k := range c.Keys(nil) {
		b,e:=c.Read(k)
		if e!=nil {
			log.Println("@error:",e)
		}
		v:=&Test{}
		json.Unmarshal(b,v)
		log.Println("@name:",v.Name,"@count:",v.Count,"@tp:",v.Tp)

	}
	log.Println("100,0000  read cost ",(utee.Tick()-st),"m second")
}


