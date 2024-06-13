package main

import (
	_case "features/channelSelect/case"
	"os"
	"os/signal"
)

/*
子协程， main主协程运行很快，这里go定义的子协程执行不到  可以使用channel阻塞
*/
func main() {
	// _case.Communication()
	// _case.ConcurrentSync()
	_case.NoticeAndMultiplexing()

	// 使用channel阻塞 让子进程打印出来 进程不会退出
	ch := make(chan os.Signal, 0)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
