package main

import (
	"sync"

	"figoxu.me/redis/pkg/ds"
	"figoxu.me/redis/pkg/ut"
	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	keyName := "hello" + uuid.New()
	var wg sync.WaitGroup
	for i := 2; i > 0; i-- {
		wg.Add(1)
		go run(&wg, keyName)
	}
	wg.Wait()
	logrus.Println(`Done`)
}

func run(wg *sync.WaitGroup, keyName string) {
	defer wg.Done()
	b := &TimeCountBarrier{
		Redis:    ds.DefaultRedis(),
		Bucket:   keyName,
		Duration: 60,
		Slot:     10,
	}
	for i := 0; i < 20; i++ {
		acquire, err := b.Acquire()
		if err == LockIsBusy {
			logrus.Println("稍后重试")
			continue
		}
		ut.Chk(err)
		leftCount, err := b.SlotLeft()
		ut.Chk(err)
		logrus.WithField("acquire", acquire).
			WithField("leftCount", leftCount).
			Println("process at ", i)
	}
}
