package main

import "fmt"

func main(){
	jobs := make(chan int,5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			}else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	//发送3个数据
	for j:=1 ;j<3;j++{
		jobs<-j
		fmt.Println("sent job",j)
	}
	//关闭channel
	close(jobs)
	fmt.Println("sent all jobs")


	 doneResult := <-done

	fmt.Println(doneResult)
}