package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type dataType interface {
	string | []string | dynamicType | getRawDataUserType | articleType | []articleType
}

// omitempty 忽略空值字段
// - 忽略某个字段
type dynamicType struct {
	UserId string `json:"user_id"`
	BookId string `json:"book_id,omitempty"`
}

// 原始参数类型
type getRawDataUserType struct{
	Name string `json:"name,omitempty"`
	Age int `json:"age,omitempty"`
	Func string `json:"func,omitempty"`
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
	// 四大请求方式
	/* Restful风格指的是网络应用中就是资源定位和资源操作的风格。不是标准也不是协议。
	GET：从服务器取出资源（一项或多项）
	POST：在服务器新建一个资源
	PUT：在服务器更新资源（客户端提供完整资源数据）
	PATCH：在服务器更新资源（客户端提供需要修改的资源数据）
	DELETE：从服务器删除资源 */
	router.GET("/articles",getArticlesList) // 文章列表
	router.GET("/articles/:id",getArticlesDetail) // 文章详情
	router.POST("/articles",createArticles) // 添加文章
	router.PUT("/articles/:id",updateArticles) // 修改某一篇文章
	router.DELETE("/articles/:id",deleteArticles) // 删除某一篇文章
	// 请求头参数获取
	router.GET("/getReqHeader",getReqHeader)
	// 设置响应头
	router.GET("/setResHeader",setResHeader)
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
	/* 
	form-data
	postGetRawData:  ----------------------------461199465389766013934536
	Content-Disposition: form-data; name="name"
	枫枫
	----------------------------461199465389766013934536
	Content-Disposition: form-data; name="age"
	18
	----------------------------461199465389766013934536--
	*/

	/* 
	x-www-form-urlencoded
	postGetRawData:  name=%E6%9E%AB%E6%9E%AB&age=18
	*/

	/* 
	json
	postGetRawData:  {
    "name": "枫枫",
    "age": 21
	}
	*/
	fmt.Println("postGetRawData: ",string(body))

	msg := Msg[getRawDataUserType]{getRawDataUserType{Func: "postGetRawData:"+contentType}, "success", http.StatusOK}
	switch contentType {
	case "application/json":
		
		var user getRawDataUserType
		err := json.Unmarshal(body, &user)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(user)
		msg = Msg[getRawDataUserType]{user, "success", http.StatusOK}
	}
	c.JSON(http.StatusOK, msg)
}

type articleType struct {
  Title   string `json:"title"`
  Content string `json:"content"`
}

func getArticlesList(c *gin.Context){
	list := []articleType{
		{"Go语言入门", "这篇文章是《Go语言入门》"},
    {"python语言入门", "这篇文章是《python语言入门》"},
    {"JavaScript语言入门", "这篇文章是《JavaScript语言入门》"},
	}
	msg := Msg[[]articleType]{list, "get article list success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

func getArticlesDetail(c *gin.Context){
	// 获取param中的id
  fmt.Println("getArticlesDetail: ",c.Param("id"))
	detail := articleType{"Go语言入门", "这篇文章是《Go语言入门》"}
	msg := Msg[articleType]{detail, "get article detail success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

func _bindJson(c *gin.Context, obj any)(err error){
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
  switch contentType {
  case "application/json":
    err = json.Unmarshal(body, &obj)
    if err != nil {
      fmt.Println(err.Error())
      return err
    }
  }
  return nil
}

func createArticles(c *gin.Context){
	// 接收前端传递来的json数据
	var article articleType

	err:=_bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		msg := Msg[string]{"", "create fail", http.StatusOK}
		c.JSON(http.StatusOK, msg)
		return
	}else{
		msg := Msg[articleType]{article, "create success", http.StatusOK}
		c.JSON(http.StatusOK, msg)
	}
}

func updateArticles(c *gin.Context){
	fmt.Println("updateArticles: ",c.Param("id"))
  var article articleType
  err := _bindJson(c, &article)
  if err != nil {
    fmt.Println(err)
		msg := Msg[string]{"", "update fail", http.StatusOK}
		c.JSON(http.StatusOK, msg)
    return
  }else{
		msg := Msg[articleType]{article, "update success", http.StatusOK}
		c.JSON(http.StatusOK, msg)
	}
}
func deleteArticles(c *gin.Context){
	fmt.Println("deleteArticles: ",c.Param("id"))
	msg := Msg[string]{"", "delete success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

func getReqHeader(c *gin.Context){
	// 首字母大小写不区分  单词与单词之间用 - 连接
  // 用于获取一个请求头
	fmt.Println(c.GetHeader("User-Agent"))
  // fmt.Println(c.GetHeader("user-agent"))
  // fmt.Println(c.GetHeader("user-Agent"))
  // fmt.Println(c.GetHeader("user-AGent"))

	// Header 是一个普通的 map[string][]string
	fmt.Println(c.Request.Header)
	// 如果是使用 Get方法或者是 .GetHeader,那么可以不用区分大小写，并且返回第一个value
	fmt.Println(c.Request.Header.Get("User-Agent"))
  fmt.Println(c.Request.Header["User-Agent"])
	// 如果是用map的取值方式，请注意大小写问题
	fmt.Println(c.Request.Header["user-agent"])

	// 自定义的请求头，用Get方法也是免大小写
  fmt.Println(c.Request.Header.Get("Token"))
  fmt.Println(c.Request.Header.Get("token"))
	msg := Msg[string]{"", "获取请求头成功！！！", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

func setResHeader(c *gin.Context){
	c.Header("Token", "e85d2352-ea9a-44b2-9e62-dce726128f0c")
	c.Header("Content-Type", "application/text; charset=utf-8")
	msg := Msg[string]{"", "设置响应头成功！！！", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}