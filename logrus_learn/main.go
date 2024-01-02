package main

import (
	"fmt"
	"logrusLearn/logSplit"
)

func main() {
	fmt.Println("logrus init")

	// 1.logrus常用方法
	// common.CommonFunc()
	// 2.hook
	// hook.HookFunc()
	/* 
	3.日志分割
	按时间分割 自定义write方法
	*/
	// logSplit.LogSplitByTimer()
	// 按时间分割 自定义hook
	// logSplit.LogSplitByTimerHook()
	// 按日志等级分割
	logSplit.LogSplitByLevel()
}