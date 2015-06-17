package main

import "fmt"

func main(){

	var a[5] int
	fmt.Println("emp:",a)
	//数组的默认值是 0

	//通过数组角标进行赋值、取值
	a[4] = 100
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])
	//数组的长度
	fmt.Println("len:",len(a))

	//通过语法来声明和初始化一个数组
	b := [5]int{1,2,3,4,5}
	fmt.Println("dc1:",b)


	//多维数组创建
	var twoD[2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j]=i+j;
		}
	}


	fmt.Println("2d:",twoD)

}
