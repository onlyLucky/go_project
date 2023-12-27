package channel_base

import (
	"fmt"
	"sync"
	"time"
)

var (
	wait = sync.WaitGroup{}
)

func singWait(){
	fmt.Println("singWait~~~")
	time.Sleep(1*time.Second)
	fmt.Println("sing end~")
	wait.Done()
}

func init() {
	wait.Add(4)
	go singWait()
	go singWait()
	go singWait()
	go singWait()
	wait.Wait()
	fmt.Println("主线程结束")
}