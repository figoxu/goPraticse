package main
import (
	"sync"
	"fmt"
	"time"
)


var syncfailureIds = []string{}

var seed =7
var loopCount = 1

var execute = 1

func main(){

	syncfailureIds = append(syncfailureIds,"hello")
	syncfailureIds = append(syncfailureIds,"how")
	syncfailureIds = append(syncfailureIds,"are")
	syncfailureIds= append(syncfailureIds,"you")
	syncfailureIds = append(syncfailureIds,"i am")
	syncfailureIds = append(syncfailureIds,"fine")
	syncfailureIds = append(syncfailureIds,"thank")
	syncfailureIds = append(syncfailureIds,"you")
	var wg sync.WaitGroup


	for ;len(syncfailureIds)>0 ;  {
		fmt.Println("=======================retry ",loopCount," times===========================")
		for _, id := range syncfailureIds {
			wg.Add(1)
			fmt.Println("prepare to sync @app", id)
			go business(&wg)
		}
		loopCount++
		execute = 1
		syncfailureIds = []string{}
		wg.Wait()
	}
	fmt.Println("success")
	
}

func business(wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(time.Second*time.Duration(1))
	if seed > loopCount+execute {
		syncfailureIds = append(syncfailureIds,"hello")
		execute++
	}
}




