package main
import (
	"net/url"
	"net/http"
	"fmt"
	"bytes"
	"strconv"
)

func main(){
	apiUrl := "http://127.0.0.1:3000"
	resource := "/admin/reload/dispatcher"
	data := url.Values{}
	data.Set("flag", "true")
	data.Add("conf", "")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/"

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload

	r.SetBasicAuth("xxxxx", "xxxxx")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
}