package main
import (
	"log"
	"fmt"
	"encoding/json"
	"github.com/figoxu/utee"
)


func main(){
	log.Println("boot grid mix json by interface")


	items := []interface{}{}
	for i := 0; i < 3; i++ {
		items = append(items, DevItem{
			Daytime: fmt.Sprint("2016-04-", i),
			Alive:   1310 + i,
			Inc:     300 + i,
			Total:   63000 + i,
		})
	}

	for i := 0; i < 3; i++ {
		items = append(items, SampleItem{
			Key:fmt.Sprint("result",i),
			Value:fmt.Sprint("value",i),
		})
	}


	dataTable := BootGrid{
		Current:  3,
		RowCount: 10,
		Total:    1000,
		Rows:     items,
	}
	v,e:=json.Marshal(dataTable)
	utee.Chk(e)
	log.Println("@v:",string(v))
}



type BootGrid struct {
	Current  int       `json:"current"`
	RowCount int       `json:"rowCount"`
	Total    int       `json:"total"`
	Rows     []interface{} `json:"rows"`
}

type DevItem struct {
	Daytime string `json:"daytime"`
	Alive   int    `json:"alive"`
	Inc     int    `json:"inc"`
	Total   int    `json:"total"`
}

type SampleItem struct {
	Key string `json:"key"`
	Value string `json:"value"`
}