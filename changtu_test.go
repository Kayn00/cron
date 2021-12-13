package cron

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

const spec = "0 * * * * *"

var numLastStart, numFistStart, numMinStart struct {
	num int
	sync.Mutex
}

func TestLastStart(t *testing.T) {
	cron := New()
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			cron.AddFunc(spec, handLastStart)
			wg.Done()
		}()
	}

	wg.Wait()
	cron.Start()
	time.Sleep(70 * time.Second)
	fmt.Println("numLastStart is : ", numLastStart.num)

	time.Sleep(time.Hour)
}

func TestFirstStart(t *testing.T) {
	cron := New()
	cron.Start()
	for i := 0; i < 10000; i++ {
		go cron.AddFunc(spec, handFistStart)
	}

	time.Sleep(70 * time.Second)
	fmt.Println("numFistStart is : ", numFistStart.num)

	time.Sleep(time.Hour)
}

func TestMinStart(t *testing.T) {
	cron := New()
	for i := 0; i < 10000; i++ {
		go cron.AddFunc(spec, handMinStart)
	}

	cron.Start()
	time.Sleep(70 * time.Second)
	fmt.Println("numMinStart is : ", numMinStart.num)

	time.Sleep(time.Hour)
}

func handLastStart() {
	numLastStart.Lock()
	numLastStart.num++
	numFistStart.Unlock()
}

func handFistStart() {
	numFistStart.Lock()
	numFistStart.num++
	numFistStart.Unlock()
}

func handMinStart() {
	numMinStart.Lock()
	numMinStart.num++
	numFistStart.Unlock()
}
