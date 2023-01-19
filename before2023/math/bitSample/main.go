package main

import (
	"fmt"
	"log"
)

func main() {
	a := 1 << 3
	log.Println("左移 1<<3==", a)

	a = 256 >> 8
	log.Println("右移 1<<3==", a)

	a = 10 ^ 2
	log.Println("异或 10 ^ 2==", a)
	//	结果是如果某位不同则该位为1, 否则该位为0

	a = 10 | 2
	log.Println("或 10|2==", a)
	//	两个相应的二进位中只要有一个为1, 该位的结果值为1

	a = 10 & 17
	log.Println("与 10&7==", a,"   10 bit:",bitStr(10),"  17 bit:",bitStr(17),"  10&17 bit:",bitStr(10&17))
	//	两个相应的二进位都为1, 该位的结果值才为1,否则为0

	a = ^125
	log.Println("取反 ^125==", a,"   125 bit:",bitStr(125),"    ^125: ",bitStr(^125))

	//	减1取反---补码
	a = ^(125 - 1)
	log.Println("补码 ^(125-1)==", a)

	a, b := 3, 7
	log.Println(a, "  ", b)

	a, b = swap(a, b)
	log.Println(a, "  ", b)

	log.Println("1 : ",bitStr(1))
	log.Println("2 : ",bitStr(2))
	log.Println("4 : ",bitStr(4))
	log.Println("8 : ",bitStr(8))
	log.Println("1|2|4|8 :" , bitStr(1 | 2 | 4| 8))
	log.Println("1&2&4&8 :" , bitStr(1 & 2 & 4& 8))

	log.Println(bitStr( bitAtPosition(3)+bitAtPosition(5)+bitAtPosition(1) ))

}

// 互换两个变量的值
// 不需要使用第三个变量做中间变量
func swap(a, b int) (int, int) {
	log.Println("[step1] a:", bitStr(a), " b:", bitStr(b))
	a ^= b // 异或等于运算
	log.Println("[step2] a:", bitStr(a), " b:", bitStr(b))
	b ^= a
	log.Println("[step3] a:", bitStr(a), " b:", bitStr(b))
	a ^= b
	log.Println("[step4] a:", bitStr(a), " b:", bitStr(b))
	return a, b
}

func bitStr(v int) string {
	return fmt.Sprintf("%b", v)
}

func bitAtPosition(position int) int{
	if(position-1<0){
		return 0
	}
	result := 1
	for i:=1;i<position;i++ {
		result *= 2
	}
	return result;
}
