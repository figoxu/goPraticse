package main

import (
	"strings"
)

type Code string

func (p Code) Len() int {
	return len(p)
}

func (p Code) Split(brace string, leftFlag bool) CodeStack {
	codeStack := make(CodeStack, 0)
	content := string(p)
	vs := strings.Split(content, brace)
	for index, code := range vs {
		if leftFlag && index > 0 && strings.Index(content, vs[index-1]+brace) != -1 {
			codeStack = AppendCode(codeStack, brace)
		}
		if code != "" && code !="\t" {
			codeStack = AppendCode(codeStack, code)
		}
		if !leftFlag && index < len(vs)-1 && strings.Index(content, vs[index+1]+code) != -1 {
			codeStack = AppendCode(codeStack, brace)
		}
	}
	return codeStack
}
