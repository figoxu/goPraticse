package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
	"github.com/chromedp/cdproto/cdp"
	"io/ioutil"
	"github.com/quexer/utee"
)

func main() {
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}
	err = c.Run(ctxt, zolSearch())
	if err != nil {
		log.Fatal(err)
	}

	err = c.Shutdown(ctxt)
	utee.Chk(err)
	err = c.Wait()
	utee.Chk(err)
}

func zolSearch() chromedp.Tasks {
	var buf []byte
	return chromedp.Tasks{
		chromedp.Navigate(`http://desk.zol.com.cn/bizhi/7287_90150_2.html`),
		chromedp.WaitVisible(`#bigImg`, chromedp.ByID),
		chromedp.ScrollIntoView(`#bigImg`, chromedp.ByID),
		chromedp.Sleep(2 * time.Second), // wait for animation to finish
		chromedp.Screenshot(`#bigImg`, &buf, chromedp.ByID),
		chromedp.ActionFunc(func(context.Context, cdp.Executor) error {
			return ioutil.WriteFile("screenshot.png", buf, 0644)
		}),
	}
}
