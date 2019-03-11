package main

import (
	"context"
	"fmt"
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var client *elastic.Client

var host = "http://localhost:9200/"
var account = "admin"
var password = "123456"
//http://localhost:9200/_nodes/http?pretty

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

func init() {

	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host), elastic.SetBasicAuth(account, password))
	utee.Chk(err)
	result, _, err := client.Ping(host).Do(ctx)
	utee.Chk(err)
	fmt.Println(Figo.JsonString(result))
}

func main() {
	log.Println("Hello World")
}
