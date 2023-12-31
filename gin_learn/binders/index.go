package binders

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// gin中的bind可以很方便的将 前端传递 来的数据与 结构体 进行 参数绑定 ，以及参数校验

// 对应的tag: postShouldBindJson json ;  postShouldBindQuery form;
type userType struct {
  Name string `json:"name" form:"name" uri:"name"`
  Age  int    `json:"age" form:"age" uri:"age"`
  Sex  string `json:"sex" form:"sex" uri:"sex"`
}

type dataType interface {
	string | userType | commonBindersType | customizeErrorType | customizeBinderType
}

type Msg[T dataType] struct {
	Data   T  `json:"data"`
	Message string `json:"msg"`
	Code  int `json:"code"`
}

func BindersRouter(router *gin.Engine) {
	// ShouldBindJSON校验JSON
	router.POST("/binders/shouldBindJson",postShouldBindJson)
	// ShouldBindQuery校验form
	router.POST("/binders/shouldBindQuery",postShouldBindQuery)
	// ShouldBindUri校验uri
	router.POST("/binders/shouldBindUri/:name/:age/:sex",postShouldBindUri)
	// ShouldBind
	router.POST("/binders/shouldBind",postShouldBind)
	// 常用验证器
	router.POST("/binders/common",postCommon)
	// 自定义验证的错误信息
	router.POST("/binders/validatorMsg",postValidatorMsg)
	// 自定义验证器
	router.POST("/binders/customizeBinder",postCustomizeBinder)
	// 注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign",signValid)
	}
}

// 可以绑定json，query，param，yaml，xml

func postShouldBindJson(c *gin.Context){
	var userInfo userType
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		// error -> string err.Error()
		msg:= Msg[string]{"postShouldBindJson解析数据失败", err.Error(), http.StatusOK}
		c.JSON(http.StatusOK, msg)
		return
	}
	msg:= Msg[userType]{userInfo, "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

// tag对应为form,不添加tag 数据无法赋值
func postShouldBindQuery(c *gin.Context){
	var userInfo userType
	err := c.ShouldBindQuery(&userInfo)
	if err != nil {
		msg:= Msg[string]{"postShouldBindQuery解析数据失败", err.Error(), http.StatusOK}
		c.JSON(http.StatusOK, msg)
		return
	}
	msg:= Msg[userType]{userInfo, "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

// tag对应为uri,不添加tag 数据无法赋值
func postShouldBindUri(c *gin.Context){
	var userInfo userType
	err := c.ShouldBindUri(&userInfo)
	if err != nil {
		msg:= Msg[string]{"postShouldBindUri解析数据失败", err.Error(), http.StatusOK}
		c.JSON(http.StatusOK, msg)
		return
	}
	msg:= Msg[userType]{userInfo, "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}	
 
/* 
ShouldBind
绑定form-data、x-www-form-urlencode 会根据请求头中的content-type去自动绑定
form-data的参数也用这个，tag用form
*/
func postShouldBind(c *gin.Context){
	var userInfo userType
	err := c.ShouldBind(&userInfo)
	if err != nil {
		msg:= Msg[string]{"postShouldBind解析数据失败", err.Error(), http.StatusOK}
		c.JSON(http.StatusOK, msg)
		return
	}
	msg:= Msg[userType]{userInfo, "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

// 常用验证器
/* 
// 不能为空，并且不能没有这个字段
required： 必填字段，如：binding:"required"  

// 针对字符串的长度
min 最小长度，如：binding:"min=5"
max 最大长度，如：binding:"max=10"
len 长度，如：binding:"len=6"

// 针对数字的大小
eq 等于，如：binding:"eq=3"
ne 不等于，如：binding:"ne=12"
gt 大于，如：binding:"gt=10"
gte 大于等于，如：binding:"gte=10"
lt 小于，如：binding:"lt=10"
lte 小于等于，如：binding:"lte=10"

// 针对同级字段的
eqfield 等于其他字段的值，如：PassWord string `binding:"eqfield=Password"`
nefield 不等于其他字段的值


- 忽略字段，如：binding:"-"
*/

// 内置验证器
/* 
// 枚举  只能是red 或green
oneof=red green 

// 字符串  
contains=fengfeng  // 包含fengfeng的字符串
excludes // 不包含
startswith  // 字符串前缀
endswith  // 字符串后缀

// 数组
dive  // dive后面的验证就是针对数组中的每一个元素

// 网络验证
ip
ipv4
ipv6
uri
url
// uri 在于I(Identifier)是统一资源标示符，可以唯一标识一个资源。
// url 在于Locater，是统一资源定位符，提供找到该资源的确切路径

// 日期验证  1月2号下午3点4分5秒在2006年
datetime=2006-01-02
*/
type commonBindersType struct {
	Name string `json:"name" binding:"required,min=1,max=10" `
	Age int `json:"age" binding:"gt=18"`
	Password string `json:"password"`
	ResetPassword string `json:"resetPassword" binding:"eqfield=Password"`
	Sex string `json:"sex" binding:"oneof=men women"`
	Ip string `json:"ip" binding:"required,ip"`
	TimeDate string `json:"timeDate" binding: "datetime=2006-01-02"`
	List []string `json:"list" binding:"required,dive,startswith=user"`
}
func postCommon(c *gin.Context){

	var userInfo commonBindersType
	err := c.ShouldBind(&userInfo)
	if err != nil {
		msg:= Msg[string]{"postCommon解析数据失败", err.Error(), http.StatusOK}
		c.JSON(http.StatusOK, msg)
		return
	}
	msg:= Msg[commonBindersType]{userInfo, "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

// 自定义验证的错误信息
// 当验证不通过时，会给出错误的信息，但是原始的错误信息不太友好，不利于用户查看只需要给结构体加一个msg 的tag

type customizeErrorType struct {
	Username string `json:"username" binding:"required" msg:"用户名不能为空"`
	Password string `json:"password" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
  Email    string `json:"email" binding:"email" msg:"邮箱地址格式不正确"`
}

/* 
err：这个参数为ShouldBindJSON返回的错误信息
obj：这个参数为绑定的结构体
还有一点要注意的是，validator这个包要引用v10这个版本的，否则会出错
*/

func postValidatorMsg(c *gin.Context){
	var userInfo customizeErrorType
	err := c.ShouldBind(&userInfo)
	if err != nil {
		// 自定义
		var errMsg string = getValidMsg(err, &userInfo)

		msg:= Msg[string]{"postValidatorMsg解析数据失败", errMsg, http.StatusOK}
		c.JSON(http.StatusOK, msg)
		return
	}
	msg:= Msg[customizeErrorType]{userInfo, "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}

// 返回结构体中的msg参数
func getValidMsg(err error, obj any) string {
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _,e := range errs{
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			if f,exits := getObj.Elem().FieldByName(e.Field()); exits{
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}

// 如果用户名不等于fengfeng就校验失败
func signValid(f validator.FieldLevel) bool {
	name := f.Field().Interface().(string)
	if name != "fengfeng" {
		return false
	}
	return true
}

type customizeBinderType struct {
	Name string `json:"name" binding:"sign" msg:"用户名错误"`
  Age  int    `json:"age" binding:""`
}

func postCustomizeBinder(c *gin.Context){
	var userInfo customizeBinderType
	err := c.ShouldBind(&userInfo)
	if err != nil {
		// 自定义
		var errMsg string = getValidMsg(err, &userInfo)

		msg:= Msg[string]{"postValidatorMsg解析数据失败解析数据失败", errMsg, http.StatusOK}
		c.JSON(http.StatusOK, msg)
		return
	}
	msg:= Msg[customizeBinderType]{userInfo, "success", http.StatusOK}
	c.JSON(http.StatusOK, msg)
}