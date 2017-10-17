package main

import (
	"fmt"
	pinyingo "github.com/struCoder/Go-pinyin"
)

func main() {
	str := "中国"
	//str1 := "重阳"
	//py := pinyingo.NewPy(pinyingo.STYLE_TONE, pinyingo.NO_SEGMENT)       //string with tone        -> 中国: ["zhōng", "guó"]
	py := pinyingo.NewPy(pinyingo.STYLE_NORMAL, pinyingo.NO_SEGMENT) //string without tone     -> 中国: ["zhong", "guo"]
	//py := pinyingo.NewPy(pinyingo.STYLE_INITIALS, pinyingo.NO_SEGMENT) // get initials of string -> 中国: ["zh", "g"]

	//segment
	//py := pinyingo.NewPy(pinyingo.STYLE_TONE, pinyingo.USE_SEGMENT)       //string with tone        -> 重阳: ["chóng", "yáng"]

	fmt.Println(py.Convert(str))
}
