package main

import (
	"github.com/quexer/utee"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	REG_WASH = `//.*`
)

type Parser struct {
	prepareReg []string
	processReg []string
}

func (p *Parser) exe(content string) []string {
	prep := func(reg string, contents ...string) []string {
		var result []string
		for _, content := range contents {
			rs := regexp.MustCompile(reg).FindAllString(content, -1)
			result = append(result, rs...)
		}
		return result
	}
	proc := func(reg string, contents ...string) []string {
		var result []string
		for _, content := range contents {
			rs := regexp.MustCompile(reg).ReplaceAllString(content, "")
			result = append(result, rs)
		}
		return result
	}
	result := []string{content}
	for _, reg := range p.prepareReg {
		result = prep(reg, result...)
	}
	for _, reg := range p.processReg {
		result = proc(reg, result...)
	}
	return TrimAndClear(result...)
}

func TrimAndClear(strs ...string) []string {
	result := []string{}
	for _, v := range strs {
		v = strings.TrimSpace(v)
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	codes := ReadAllFile("/home/figo/develop/env/GOPATH/src/tendcloud.com/push-gateway/api-admin.go")
	codes = regexp.MustCompile(REG_WASH).ReplaceAllString(codes, "")
	varParse := &Parser{
		prepareReg: []string{`var.*\([^$]*?\)`, `.+ `},
		processReg: []string{"var", " "},
	}
	for _, v := range varParse.exe(codes) {
		log.Println(v)
	}

	constParse := &Parser{
		prepareReg: []string{`const.*\([^$]*?\)`, `.+=`},
		processReg: []string{"const", "=", "\\(", "\\)"},
	}
	for _, v := range constParse.exe(codes) {
		log.Println(v)
	}

	funcNameParse := &Parser{
		prepareReg: []string{`func.*\{[^$]*?\}`, `func.*\(`},
		processReg: []string{`\(.*\)`, `\(`, `\)`, `func`},
	}
	for _, v := range funcNameParse.exe(codes) {
		log.Println(v)
	}
}

func ReadAllFile(filePth string) string {
	f, err := os.Open(filePth)
	utee.Chk(err)
	b, err := ioutil.ReadAll(f)
	utee.Chk(err)
	return string(b)
}
