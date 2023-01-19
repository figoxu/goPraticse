package main

import (
	"log"
	"regexp"
	"fmt"
)

var (
	content = "\u003ci\u003e\u003cu\u003e\u003cb\u003e\u003cfont color=\"#00ff00\"\u003e大事发生的发生的分\u003c/font\u003e\u003c/b\u003e\u003c/u\u003e\u003c/i\u003e<ul><li> Hello </li><li> World </li></ul>"
)

func main() {
	log.Println("%s",content)
	content := fmt.Sprint(content)
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	content = re.ReplaceAllString(content, "")
	log.Println(content)
}
