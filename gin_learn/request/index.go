package request

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type dataType interface {
	string | []string | dynamicType
}

// omitempty 忽略空值字段
// - 忽略某个字段
type dynamicType struct {
	UserId string `json:"user_id"`
	BookId string `json:"book_id,omitempty"`
}



type Msg[T dataType] struct {
	Data   T  `json:"data"`
	Message string `json:"msg"`
	Code  int `json:"code"`
}

func RequestRouter(router *gin.Engine) {
	// 查询参数
	router.GET("/query",getQuery)
	// 动态参数
	router.GET("/dynamicParam/:user_id",getDynamicParam)
	router.GET("/dynamicParam/:user_id/:book_id",getDynamicParam)
	// 表单提交
	router.POST("/postForm", postForm)
	// 原始参数
	router.POST("/getRawData", postGetRawData)
}

func getQuery(c *gin.Context){
	fmt.Println("getQuery:",c.Query("user"))
	user,ok := c.GetQuery("user")
  if(ok){
		fmt.Println("getQuery:",user)
	}
  fmt.Println("getQuery:",c.QueryArray("user")) // 拿到多个相同的查询参数
  fmt.Println("getQuery:",c.DefaultQuery("addr", "sichuan"))
	msg := Msg[[]string]{c.QueryArray("user"), "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

func getDynamicParam(c *gin.Context){

	fmt.Println(c.Param("user_id"))
  fmt.Println(c.Param("book_id"))
	// 强制类型转换
	/* var a interface{} =10
	t,ok:= a.(int)
	if ok{
			fmt.Println("int",t)
	}
	t2,ok:= a.(float32)
	if ok{
			fmt.Println("float32",t2)
	} */
	msg := Msg[dynamicType]{
		dynamicType{UserId: c.Param("user_id"), BookId: c.Param("book_id")},
		"success",
	 	http.StatusOK,
	}
	c.JSON(http.StatusOK, msg)
}

func postForm(c *gin.Context){
	fmt.Println("postForm: ",c.PostForm("name"))
	fmt.Println("postForm: ",c.PostFormArray("name"),reflect.TypeOf(c.PostFormArray("name")))
	fmt.Println("postForm: ",c.DefaultPostForm("addr", "sichuan")) // 如果用户没传，就使用默认值
	forms, err := c.MultipartForm()               // 接收所有的form参数，包括文件
  fmt.Println("postForm: ",forms, err)
	msg := Msg[string]{"postForm", "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

func postGetRawData(c *gin.Context){
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
}