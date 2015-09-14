package main
import "fmt"

func main(){
	lb:
	for i:=0;i<=100;i++{
		if i%2==0 {
			continue lb
		}
		fmt.Println(i)
	}
}
