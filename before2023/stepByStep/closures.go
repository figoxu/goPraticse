package main

import "fmt"

func intSeq() func() int{
	i := 0
	return func() int{
		i+=1
		return i
	}
}

func main(){
	nextInt:=intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())


	fmt.Println("[let me see the memory address]  nextInt:",nextInt,"    newInts:",newInts)
}