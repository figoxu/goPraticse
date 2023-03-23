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

	rowspanHtml := `
<html lang="en">
<head>
    <meta charSet="utf-8"/>
    <style >
        table{
            border: 1px solid #000;

        }
        td{
            border: 1px solid #000;

        }
        tr{
            border: 1px solid #000;

        }

    </style>
</head>
<body>
<table >
    <tr>
        <td>类别</td>
        <td>名称</td>
    </tr>
    <tr>
        <td rowspan="2">颜色</td>
        <td>红色</td>
    </tr>
    <tr>
        <td>黄色</td>
    </tr>
    <tr>
        <td colspan="2">姓氏</td>
    </tr>
    <tr>
        <td>王</td>
        <td>张</td>
    </tr>
</table></body></html>`
	err = html2jpg(rowspanHtml, "/Users/xujianhui/mobvista/mtg/github/goPraticse/2023/doc/b.jpg")
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
