package main

import (
	"fmt"
	"ginLearn/binders"
	"ginLearn/download"
	"ginLearn/middleware"
	"ginLearn/request"
	"ginLearn/response"
	"ginLearn/upload"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context){
	context.String(http.StatusOK,"hello go!!!")
}

func main() {

	/* 
	gin 环境切换，有debug模式 release模式 
	如果不想看到启动gin，显示所有的路由debug日志，那么我们可以改为release模式,只会打印后续请求的日志
	*/
	gin.SetMode(gin.ReleaseMode) // 切换为release模式
	/* 
	定义路由格式
	启动gin，它会显示所有的路由，默认输出为：
	[GIN-debug] GET    /middle/multiple          --> ginLearn/middleware.m2 (5 handlers)
	[GIN-debug] GET    /middle/interceptRes      --> ginLearn/middleware.m2 (5 handlers)
	[GIN-debug] GET    /middle/middleNext        --> ginLearn/middleware.m5 (5 handlers)
	[GIN-debug] GET    /middle/globalNext        --> ginLearn/middleware.handleMiddleNext (4 handlers)

	更改输出格式可以使用 gin.DebugPrintRouteFunc
	输出格式如下：
	2024/01/01 19:25:04 [ fengfeng ] GET /middle/multiple ginLearn/middleware.m2 5
	2024/01/01 19:25:04 [ fengfeng ] GET /middle/interceptRes ginLearn/middleware.m2 5
	2024/01/01 19:25:04 [ fengfeng ] GET /middle/middleNext ginLearn/middleware.m5 5
	2024/01/01 19:25:04 [ fengfeng ] GET /middle/globalNext ginLearn/middleware.handleMiddleNext 4
	*/
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int){
		log.Printf("[ fengfeng ] %v %v %v %v\n",
		httpMethod,
		absolutePath,
		handlerName,
		nuHandlers,
		)
	}
	/* 
	日志处理
	输出到文件
	*/
	f, _ := os.Create("gin.log")
	// 只输入到日志文件中
	//gin.DefaultWriter = io.MultiWriter(f)
  // 如果需要同时将日志写入文件和控制台，请使用以下代码。（os.Stdout 控制台）
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	
	// 创建一个默认路由
	// router := gin.Default()
	/* 
	创建一个自定义格式log显示路由 
	默认的是这样的
	[GIN] 2024/01/01 - 19:37:23 | 200 |       582.8µs |       127.0.0.1 | POST     "/binders/common"
	
	修改log输出格式
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string{}))

	输出格式如下：
	[ API ] 1007-01-01 16:00:00 |  200  | 	 127.0.0.1 | 513.9µs |  POST     	  /binders/common

	添加log 颜色字符串
	
	*/
	router := gin.New()
	router.Use(gin.LoggerWithFormatter( LoggerWithFormatter ))
	// 上面修改log的格式也可以使用 gin.LoggerWithConfig
	/* router.Use(
    gin.LoggerWithConfig(
      gin.LoggerConfig{ Formatter: LoggerWithFormatter},
    ),
  ) */

	fmt.Printf("----\033[97;44m 打印颜色处理 \033[0m ----\n") // 打印颜色

	

	// 绑定路由规则和路由函数，访问/index的路由，将由对应的函数处理掉
	
	router.GET("/index", Index)
	// 1.响应 路由
	response.HandleRouter(router)
	// 2.请求 路由
	request.RequestRouter(router)
	// 3.绑定器 路由
	binders.BindersRouter(router)
	// 4.文件上传
	upload.UploadRouter(router)
	// 5.文件下载
	download.DownloadRouter(router)
	// 6.中间件
	middleware.MiddlewareRouter(router)
	// 7.日志的处理
	// fmt.Println(router.Routes())  // 它会返回已注册的路由列表

	// 启动监听，gin会把web服务运行在本机的0.0.0.0:8080端口上
	router.Run("0.0.0.0:8080")
	// 用原生http服务的方式，router.Run本质就是http.ListenAndServe的进一步封装
  http.ListenAndServe(":8080", router)
}

// 设置打印格式
func LoggerWithFormatter(params gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	statusColor = params.StatusCodeColor()
	methodColor = params.MethodColor()
	resetColor = params.ResetColor()
	return fmt.Sprintf(
		"[ API ] %s | %s %d %s | \t %s | %s | %s %-7s %s \t  %s\n",
		params.TimeStamp.Format("2007-01-01 16:04:00"),
		statusColor, params.StatusCode, resetColor, // 状态码
		params.ClientIP, // 客户端ip
		params.Latency, // 请求耗时
		methodColor, params.Method, resetColor, // 请求方法
		params.Path, // 路径
	)
}
