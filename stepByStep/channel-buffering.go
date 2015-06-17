package main
import "fmt"

func main(){
	//创建长度为2的字符串 channel
	messages := make(chan string,2)

	//往channel里面写入2个字符串
	messages <- "buffered"
	messages <- "channel"

	//从channel里面读取字符串
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
