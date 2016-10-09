package main
import (
	"log"
	"fmt"
)

func main(){
	type Category string
	type Typo string
	type Ids []string
	idMaps := make(map[string]map[string]Ids)
	for _, key := range []string{"Hello","World"} {
		log.Println("@v:", key)
		for i:=0;i<100;i++ {
			if idMaps[key]==nil {
				idMaps[key] = make(map[string]Ids)
			}
			format :=fmt.Sprint(i%2)
			if idMaps[key][format]==nil {
				idMaps[key][format] = make(Ids,0)
			}
			idMaps[key][format]= append(idMaps[key][format] ,fmt.Sprint(key,"_Item_",i))
		}
	}
	for tid,fmtIdMap := range idMaps {
		log.Println("=====>tid:",tid)
		for tp,ids := range fmtIdMap {
			log.Println("=========>@tp:",tp)
			for _,id := range ids {
				log.Println("==================>@id:",id)
			}
		}
	}
}
