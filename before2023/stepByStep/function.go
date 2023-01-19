package main
import "fmt"

func plus(a int,b int)int{
	return a+b;
}

func plusPlus(a,b,c int)int{
	return a+b+c;
}

func plusArray(a...int)int{
	res := 0;
	for v:=range a{
		fmt.Println("....v:",v,"  value:",a[v])
		res = res  + a[v];
	}
	return res
}

func main(){

	res := plus(1,2)
	fmt.Println(res)

	res = plusPlus(1,2,3)
	fmt.Println(res)

	res = plusArray(1,2,3,4,5)
	fmt.Println(res)
	res = plusArray(3,4,5,6,7)
	fmt.Println(res)
}
