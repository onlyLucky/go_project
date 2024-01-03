package main

import (
	"fmt"
	"logrusLearn/integrate"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	// logSplit.LogSplitByLevel()

	/* 
	4. gin集成logrus
	*/
	integrate.InitFile("logs","logrusLearn")
	router := gin.New()
	router.Use(integrate.LogMiddleware())
	router.GET("/", func(c *gin.Context) {
    logrus.Info("来了")
    c.JSON(200, gin.H{"msg": "hello"})
  })
  router.Run(":8080")
}