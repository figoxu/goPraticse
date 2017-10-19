package main

import (
	"log"
	"github.com/figoxu/utee"
	"fmt"
)

type BlockingTask struct {
	TaskContent string
	DoneChan    chan bool
	TryTimes int
}

var(
	b_q utee.MemQueue
	t_c *utee.TimerCache
)

func retryMethod(k,v interface{}){
	task:=v.(*BlockingTask)
	task.TryTimes = task.TryTimes+1
	if task.TryTimes >= 3 {
		task.DoneChan <- false
	}else{
		businessExecute(task.TaskContent)
		t_c.Put(k,v)
	}
}


func main() {
	b_q=utee.NewLeakMemQueue(100, 1, blockingExecute)
	t_c=utee.NewTimerCache(3,retryMethod)
	mockInput()
	c :=make(chan bool)
	<-c
}

func mockInput(){
	newTask := func(v string)*BlockingTask{
		return &BlockingTask{
			TaskContent :v,
			DoneChan :make(chan bool),
			TryTimes :0,
		}
	}

	for i:=0;i<30;i++ {
		go b_q.Enq(newTask(fmt.Sprint("testData ",i)))
	}
}

func blockingExecute(v interface{}) {
	task := v.(*BlockingTask)
	t_c.Put(task.TaskContent,task)
	businessExecute(task.TaskContent)
	<-task.DoneChan
}

func businessExecute(content string){

	log.Println("Doing At:",content)
}
