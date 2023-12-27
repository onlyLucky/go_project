// 异步处理
package channel_base

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int)

func pay(name string, money int, wait *sync.WaitGroup){
	fmt.Printf("%s 开始购物\n", name)
  time.Sleep(1 * time.Second)
  fmt.Printf("%s 购物结束\n", name)

	moneyChan <-money
	wait.Done()
}

func init() {
	var wait sync.WaitGroup
	startTime:=time.Now()
	wait.Add(3)
	// 主线程结束，协程函数跟着结束
  go pay("张三", 2, &wait)
  go pay("王五", 3, &wait)
  go pay("李四", 5, &wait)

	go func(){
		defer close(moneyChan)
		// 在协程函数里面等待上面三个协程函数结束
    wait.Wait()
	}()

	var moneyList []int
	for money := range moneyChan {
    moneyList = append(moneyList, money)
  }
	fmt.Println("购买完成", time.Since(startTime))
  fmt.Println("moneyList", moneyList)
}