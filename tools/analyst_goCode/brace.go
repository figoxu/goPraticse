package main

type Code string

func (p Code) Len() int{
	return len(p)
}

func (p Code) Split(brace string) []Code {
	codes:=make([]Code,0)
	return codes
}