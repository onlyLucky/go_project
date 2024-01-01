package middleware

import (
	"fmt"
	"net/http"
	"time"

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
	/* 
	单独注册中间件
	(m1处于handleSingleRegister函数的前面,请求来之后,先走m1,再走后面处理函数) 
	m1 -> handleMiddle
	*/
	router.GET("/middle/singleRegister",m1,handleMiddle)
	/* 
	多个中间件 
	m1 -> handleMiddle -> m2
	*/
	router.GET("/middle/multiple",m1,handleMiddle,m2)
	/* 
	中间件拦截响应
	c.Abort()拦截，后续的HandlerFunc就不会执行了
	m3 ->  后续操作被拦截了 handleMiddle -> m2
	*/
	router.GET("/middle/interceptRes",m3,handleMiddle,m2)
	/* 
	中间件放行
	c.Next()，Next前后形成了其他语言中的请求中间件和响应中间件
	m4 in.........
	handleMiddleNext in ...
	m5 in.........
	m5 out.........
	handleMiddleNext out ...
	m4 out.........
	*/
	router.GET("/middle/middleNext",m4,handleMiddleNext,m5)
	/* 
	全局注册中间件
	使用Use去注册全局中间件，Use接收的参数也是多个HandlerFunc
	*/
	router.Use(mGlobal)
	router.GET("/middle/globalNext",handleMiddleNext)
	/* 
	中间件传递数据
	使用Set设置一个key-value, 在后续中间件中使用Get接收数据
	value的类型是any类型，所有我们可以用它传任意类型，在接收的时候做好断言即可
	*/
	router.Use(mTransform)
	router.GET("/middle/transform",handleMiddleTransform)
	/* 
	路由分组
	将一系列的路由放到一个组下，统一管理
	路由分组注册中间件
	*/
	apiR := router.Group("/api").Use(m1,m6())
	apiR.GET("/index",func(c *gin.Context){
		msg:= Msg[string]{"/api/index", "success", http.StatusOK}
		c.JSON(http.StatusOK, msg)
	})
	apiR.GET("/home",func(c *gin.Context){
		msg:= Msg[string]{"/api/home", "success", http.StatusOK}
		c.JSON(http.StatusOK, msg)
	})
	/* 
	gin.Default
	gin.Default()默认使用了Logger和Recovery中间件，其中：
	Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。 Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
	*/

	/* 
	中间件案例
	权限验证 验证用户是否登录了
	耗时统计 统计每一个视图函数的执行时间
	*/
	api2R := router.Group("/api/v2")
	apiUser := api2R.Group("")
	{
		apiUser.POST("login",func(c *gin.Context){
			msg:= Msg[string]{"登录成功", "success", http.StatusOK}
			c.JSON(http.StatusOK, msg)
		})
	}
	apiHome := api2R.Group("system").Use(JwtTokenMiddleware)
	{
		apiHome.GET("/index",func(c *gin.Context){
			msg:= Msg[string]{"/api/v2/system/index", "success", http.StatusOK}
			c.JSON(http.StatusOK, msg)
		})
		apiHome.GET("/home",func(c *gin.Context){
			msg:= Msg[string]{"/api/v2/system/home", "success", http.StatusOK}
			c.JSON(http.StatusOK, msg)
		})
	}
}

func handleMiddle(c *gin.Context){
	fmt.Println("handleMiddle")
	msg:= Msg[string]{"处理路由", "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

func handleMiddleNext(c *gin.Context){
	fmt.Println("handleMiddleNext in ...")
	msg:= Msg[string]{"处理路由", "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
	c.Next()
	fmt.Println("handleMiddleNext out ...")
}

// 定义一个中间件
func m1(c *gin.Context){
	fmt.Println("m1 in.........")
}
func m2(c *gin.Context){
	fmt.Println("m2 in.........")
}

func m3(c *gin.Context){
	fmt.Println("m3 ...in")
  msg:= Msg[string]{"第一个中间件拦截了", "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
  c.Abort()
}

func m4(c *gin.Context){
	fmt.Println("m4 in.........")
	c.Next()
	fmt.Println("m4 out.........")
}

func m5(c *gin.Context){
	fmt.Println("m5 in.........")
	c.Next()
	fmt.Println("m5 out.........")
}

func mGlobal(c *gin.Context){
	fmt.Println("mGlobal in.........")
	c.Next()
	fmt.Println("mGlobal out.........")
}

func mTransform(c *gin.Context){
	fmt.Println("mTransform in.........")
	c.Set("name","fengfeng")
}

func handleMiddleTransform(c *gin.Context){
	fmt.Println("handleMiddleTransform in.........")
	name, _ := c.Get("name")
	fmt.Println(name)
	// value的类型是any类型，所有我们可以用它传任意类型，在接收的时候做好断言即可
	msg:= Msg[string]{name.(string), "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

// 中间件其他方式使用
func m6() gin.HandlerFunc {
	// 这里的代码是程序一开始就会执行
	return func(c *gin.Context){
		// 这里是请求来了就会执行
		fmt.Println("m6 in ....")
	}
}

func JwtTokenMiddleware(c *gin.Context){
	// 获取请求头的token
	token := c.GetHeader("token")
	// 调用jwt的验证函数
	if token == "1234"{
		// 验证通过
		c.Next()
		return
	}
	// 验证不通过
	msg:= Msg[string]{"权限验证失败", "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
	c.Abort()
}

func TimeMiddleware(c *gin.Context) {
  startTime := time.Now()
  c.Next()
  since := time.Since(startTime)
  // 获取当前请求所对应的函数
  f := c.HandlerName()
  fmt.Printf("函数 %s 耗时 %d\n", f, since)
}