package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Msg struct {
	Data string `json:"data"`
	Message string
	Code int
}

func HandleRouter(router *gin.Engine) {
	// 加载模板 
	// 不同文件夹下模板名字可以相同，此时需要 LoadHTMLGlob() 加载两层模板路径。templates/**/*
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	// 返回字符串
	router.GET("/txt", getTxt)
	// 返回json
	router.GET("/json", getJson)
	// 结构体转json
	router.GET("/moreJSON", getMoreJson)
	// 返回xml
	router.GET("/xml", getXml)
	// 返回yaml
	router.GET("/yaml", getYaml)
	// 加载html
	router.GET("/html", getHtml)
	// 文件静态资源
	// 在golang总，没有相对文件的路径，它只有相对项目的路径
	// 网页请求这个静态目录的前缀， 第二个参数是一个目录，注意，前缀不要重复
	router.StaticFS("/static", http.Dir("static"))
	// 配置单个文件， 网页请求的路由，文件的路径
	router.StaticFile("/logo.png", "temp/Gopher.png")

	// 重定向
	router.GET("/redirect", getRedirect)
	router.GET("/redirectStatic", getRedirectStatic)
}

func getTxt(c *gin.Context) {
	c.String(http.StatusOK, "这是一个txt")
}

func getJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg":"success", "status":http.StatusOK})
}

func getMoreJson(c *gin.Context) {
	msg := Msg{Data: "gin more json", Message:"success", Code: 200}
	// 注意 msg.Data 变成了 "data" 字段
	// 输出为 {data: "gin more json", Message:"success", Code: 200}
	c.JSON(http.StatusOK, msg)
}

func getXml(c *gin.Context){
	msg := Msg{Data: "gin get xml", Message:"success", Code: http.StatusOK}
	c.XML(http.StatusOK, msg)
}

func getYaml(c *gin.Context){
	msg := Msg{Data: "gin get yaml", Message:"success", Code: http.StatusOK}
	c.YAML(http.StatusOK, msg)
}

func getHtml(c *gin.Context){
	//根据完整文件名渲染模板，并传递参数
  c.HTML(http.StatusOK, "templates/index.html", gin.H{
    "title": "hello html templates",
  })
}

func getRedirect(c *gin.Context){
	//支持内部和外部的重定向
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
}

func getRedirectStatic(c *gin.Context){
	c.Redirect(http.StatusMovedPermanently, "/static/index.html")
}

