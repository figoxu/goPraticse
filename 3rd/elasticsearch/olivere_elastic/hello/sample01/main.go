package main

import (
	"context"
	"fmt"
	"github.com/figoxu/Figo"
	"github.com/icrowley/fake"
	"github.com/quexer/utee"
	"github.com/sirupsen/logrus"
	"gopkg.in/olivere/elastic.v5"
	"reflect"
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
	c, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host), elastic.SetBasicAuth(account, password))
	utee.Chk(err)
	result, _, err := c.Ping(host).Do(ctx)
	utee.Chk(err)
	fmt.Println(Figo.JsonString(result))
	client = c
}

func main() {
	tryInsert()
	trySearch()
}

func tryInsert() {
	indexService := client.Index().Index("test").Type("employee").Id("1").BodyJson(&Employee{
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
		Age:       18,
		About:     "Love",
		Interests: []string{"BasketBall", "PingPong", "FootBall"},
	})
	rsp, err := indexService.Do(context.Background())
	utee.Chk(err)
	logrus.WithField("rsp", rsp).WithField("err", err).Println("tryInsert")
}

func trySearch() {
	q := elastic.NewQueryStringQuery("football")
	rsp, err := client.Search().Index("test").Type("employee").Query(q).Size(10).Do(context.Background())
	utee.Chk(err)
	logrus.WithField("rsp", rsp).WithField("err", err).Println("trySearch")
	for _, item := range rsp.Each(reflect.TypeOf(Employee{})) {
		logrus.Println(Figo.JsonString(item))
	}
}
