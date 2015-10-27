package main
import (
	"log"
	"github.com/gocql/gocql"
)


type Tweet struct {
	Timeline      string
	ID            gocql.UUID  `cql:"id"`
	Ingored       string      `cql:"-"`
	Text          string      `teXt`
	OriginalTweet *gocql.UUID `json:"origin"`
}

func main(){
	log.Println("hello")

	m, ok := StructToMap("str")
	if m != nil {
		log.Fatal("map is not nil when val is a string")
	}
	if ok {
		log.Fatal("ok result from StructToMap when the val is a string")

	}
}

