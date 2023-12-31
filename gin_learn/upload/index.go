package upload

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

// 文件上传
func UploadRouter(router *  gin.Engine){
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
  // 单位是字节， << 是左移预算符号，等价于 8 * 2^20
  // gin对文件上传大小的默认值是32MB
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// 单文件
	router.POST("/upload/singleFile",uploadSingleFile)
	// 保存只读文件 创建复制
	router.POST("/upload/createCopy",uploadCreateCopy)
	// 多文件上传
	router.POST("/upload/multiFile",uploadMultiFile)
}

func uploadSingleFile(c *gin.Context){
	// 单文件
	file, err := c.FormFile("file")
	if err != nil {
		msg:= Msg[string]{"上传失败", "fail", http.StatusOK}
		c.JSON(http.StatusOK, msg)
	}else{
		log.Println(file.Filename)

		dst := "./Uploads/"+file.Filename
		// 上传文件到指定的完整文件路径
		c.SaveUploadedFile(file,dst) // 文件对象  文件路径，注意要从项目根路径开始写

		msg:= Msg[string]{"上传成功", "success", http.StatusOK}
		c.JSON(http.StatusOK, msg)
	}
	
}

func uploadCreateCopy(c *gin.Context){
	file, err := c.FormFile("file")
	if err != nil {
		msg:= Msg[string]{"上传失败", "fail", http.StatusOK}
		c.JSON(http.StatusOK, msg)
	}else{
		log.Println(file.Filename)
		// 读取文件中的数据，返回文件对象
		fileRead, _ := file.Open()
		// 读取打印上传文件
		data, _ := io.ReadAll(fileRead)
		fmt.Println("uploadCreateCopy：",string(data))
		dst := "./Uploads/"+file.Filename
		// 创建一个文件
		out, e := os.Create(dst)
		if e != nil {
			fmt.Println(e)
		}
		defer out.Close()
		// 拷贝文件对象到out中
		io.Copy(out, fileRead)

		msg:= Msg[string]{"上传成功", "success", http.StatusOK}
		c.JSON(http.StatusOK, msg)
	}
}

func uploadMultiFile(c *gin.Context){
	// Multipart form
	form, _ := c.MultipartForm()
	files, err := form.File["upload"]  // 注意这里名字不要对不上了
	if !err {
		msg:= Msg[string]{"上传多个文件失败", "fail", http.StatusOK}
		c.JSON(http.StatusOK, msg)
	}else{
		for _, file := range files {
			log.Println(file.Filename)
			// 上传文件至指定目录
			c.SaveUploadedFile(file, "./Uploads/"+file.Filename)
		}
		msg:= Msg[string]{"上传多个文件成功", "success", http.StatusOK}
		c.JSON(http.StatusOK, msg)
	}
}