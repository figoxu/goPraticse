package main
import (
	"fmt"
	"net/url"
	"net/http"
)

func main(){
	fmt.Println("hello")
	anxinjieSms("1860XXXXX46","hello")
}

const (
	ANXINJIE_URL  = "XXX"
	ANXINJIE_NAME = "XXX"
	ANXINJIE_PASS = "XXX"
)

func anxinjieSms(mobile, content string) error {

	res_url := fmt.Sprintf("%s?name=%s&pass=%s&mobiles=%s&content=%s",
		ANXINJIE_URL,
		ANXINJIE_NAME,
		ANXINJIE_PASS,
		mobile,
		content,
	)

	l, err := url.Parse(res_url)
	if err != nil {
		return err
	}
	qe := l.Query().Encode()
	resp, err := http.Get(fmt.Sprintf("%s?%s", ANXINJIE_URL, qe))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//可选 查看返回结果, read resp.Body
	return nil

}
