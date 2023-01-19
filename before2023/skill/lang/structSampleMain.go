package main
import (
	"log"
	"encoding/json"
)

var config struct {// 定义一个用于全局配置结构体
	APIKey   string
	APPSecrt string
}

func init() {
	log.Println("Hello")
	config.APIKey = "10101010"
	config.APPSecrt = "01010101"
}

func main(){
	log.Println("@APIKey:",config.APIKey," @AppSecrt:",config.APPSecrt)

	data := struct {//匿名结构体的定义
		Title string  `json:"title"`
		Items []string `json:"items`
	}{
		Title: "Hello",
		Items: []string{"How","Old","Are","You"},
	}
	b,_:=json.Marshal(data)
	log.Println("that is temp struct for json :",string(b))


//	{"data": {"children": [
//	{"data": {
//	"title": "The Go homepage",
//	"url": "http://golang.org/"
//	}},
//	...
//	]}}
//
//	type Item struct {
//		Title string
//		URL   string
//	}
//	type Response struct {
//		Data struct {
//				 Children []struct {
//					 Data Item
//				 }
//			 }
	}

}