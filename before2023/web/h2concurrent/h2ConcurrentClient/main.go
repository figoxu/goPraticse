package main

import (
	"crypto/tls"
	"fmt"
	"github.com/quexer/utee"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	//	"time"
)

var (
	http2Client = &http.Client{
		Transport: &http2.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	resp_chan = make(chan *http.Response, 10000)

//	latch     = utee.NewThrottle(3)
)

func main() {
	log.Println("hello")
	go respProcess()
	t := 10000
	st := utee.Tick()
	for i := 0; i < t; i++ {
		//		latch.Acquire()
		//		go
		test()
	}
	cost := utee.Tick() - st
	log.Println("@times:", t, " invoke cost:", cost, " second")
	//	time.Sleep(time.Second*time.Duration(2))
	for {
		if len(resp_chan) == 0 {
			tcost := utee.Tick() - st
			log.Println("@times:", t, " invoke total cost:", tcost, " second")
			os.Exit(1)
		}
	}
}

func test() {
	//	defer latch.Release()
	if err := httpPost("https://127.0.0.1:10443/helloPost", url.Values{}); err != nil {
		log.Println("utee h2 post @err:", err)
	}
}

func httpPost(postUrl string, q url.Values, credential ...string) error {
	var resp *http.Response
	var err error
	req, err := http.NewRequest("POST", postUrl, strings.NewReader(q.Encode()))
	if err != nil {
		return fmt.Errorf("[http] err %s, %s\n", postUrl, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if len(credential) == 2 {
		req.SetBasicAuth(credential[0], credential[1])
	}
	resp, err = http2Client.Do(req)
	if err != nil {
		return fmt.Errorf("[http] err %s, %s\n", postUrl, err)
	}
	resp_chan <- resp
	return nil
}

func respProcess() {

	for {
		select {
		case resp := <-resp_chan:
			processing(resp)
		}
	}
}

func processing(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	postUrl := resp.Request.URL
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("[http] status err %s, %d\n", postUrl, resp.StatusCode)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[http] read err %s, %s\n", postUrl, err)
	}
	//	log.Println("h2 post @rsp:", string(b))
	return b, nil

}
