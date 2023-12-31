package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type dataType interface {
	string
}

type Msg[T dataType] struct {
	Data   T  `json:"data"`
	Message string `json:"msg"`
	Code  int `json:"code"`
}

// 中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等
// Gin中的中间件必须是一个gin.HandlerFunc类型

func MiddlewareRouter(router *gin.Engine) {
	// 单独注册中间件
	// (m1处于handleSingleRegister函数的前面,请求来之后,先走m1,再走后面处理函数)
	router.GET("/middle/singleRegister",m1,handleSingleRegister)
}

func handleSingleRegister(c *gin.Context){
	fmt.Println("singleRegister")
	msg:= Msg[string]{"处理单独注册中间件", "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

// 定义一个中间件
func m1(c *gin.Context){
	fmt.Println("m1 in.........")
}