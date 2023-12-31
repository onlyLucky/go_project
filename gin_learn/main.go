package main

import (
	"ginLearn/binders"
	"ginLearn/request"
	"ginLearn/response"
	"ginLearn/upload"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context){
	context.String(http.StatusOK,"hello go!!!")
}

func main() {
	// 创建一个默认路由
	router := gin.Default()
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
	
	// 启动监听，gin会把web服务运行在本机的0.0.0.0:8080端口上
	router.Run("0.0.0.0:8080")
	// 用原生http服务的方式，router.Run本质就是http.ListenAndServe的进一步封装
  http.ListenAndServe(":8080", router)
}

