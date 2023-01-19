package main
import (
	"log"
	"encoding/json"
)

func main(){
	type Pa struct {
		name string  `json:"name,omitempty"`
		badge int    `json:"badge,omitempty"`
	}

	type Pb struct{
		name string  `json:"name,omitempty"`
	}

	a := Pa{
		name :"figo",
	}
	ma   :=  map[string]interface{}{}
	ma["pp"] = a
	ja,_ := json.Marshal(&ma)
	b := Pb{
		name :"figo",
	}
	mb   :=  map[string]interface{}{}
	mb["pp"] = b
	jb,_ := json.Marshal(&mb)


	log.Println(uint16(len(ja)))
	log.Println(uint16(len(jb)))
}
