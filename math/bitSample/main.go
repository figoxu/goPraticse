package main
import "log"

func main(){
	a := 1<<3
	log.Println("左移 1<<3==",a)

	a = 256>>8
	log.Println("右移 1<<3==",a)

	a = 10 ^ 2
	log.Println("异或 10 ^ 2==",a)
//	结果是如果某位不同则该位为1, 否则该位为0

	a = 10 | 2
	log.Println("或 10|2==",a)
//	两个相应的二进位中只要有一个为1, 该位的结果值为1

	a = 10 & 2
	log.Println("与 10&2==",a)
//	两个相应的二进位都为1, 该位的结果值才为1,否则为0

	a = ^125
	log.Println("取反 ^125==",a)


//	减1取反---补码
	a = ^(125-1)
	log.Println("补码 ^(125-1)==",a)

	a,b := 11111,222222
	log.Println(a,"  ",b)

	a,b = swap(a,b)
	log.Println(a,"  ",b)

}


// 互换两个变量的值
// 不需要使用第三个变量做中间变量
func swap(a, b int) (int, int) {
	a ^= b // 异或等于运算
	b ^= a
	a ^= b
	return a, b
}