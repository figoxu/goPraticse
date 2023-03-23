package main

import (
	"gopkg.in/russross/blackfriday.v2"
	"os"
)

func main() {
	v := md2html(`# Hello
* world
`)
	err := html2jpg(v, "/Users/xujianhui/mobvista/mtg/github/goPraticse/2023/doc/a.jpg")
	if err != nil {
		panic(err)
	}
}

func md2html(body string) string {
	return string(blackfriday.Run([]byte(body)))
}

func html2jpg(html, target string) error {
	c := ImageOptions{BinaryPath: "/usr/local/bin/wkhtmltoimage", Input: "-", HTML: html, Format: "jpg"}
	bs, err := GenerateImage(&c)
	if err != nil {
		return err
	}
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	_, err = f.Write(bs)
	return err
}
