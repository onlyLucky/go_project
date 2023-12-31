package download

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 文件下载
func DownloadRouter(router *gin.Engine){
	router.GET("/download",downloadFile)
}


// 有些响应，比如图片，浏览器就会显示这个图片，而不是下载，所以我们需要使浏览器唤起下载行为
func downloadFile(c *gin.Context){
	fmt.Println("downloadFile")
	c.Header("Content-Type", "application/octet-stream") // 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
	c.Header("Content-Disposition", "attachment; filename="+"file.jpg") // 用来指定下载下来的文件名
	c.Header("Content-Transfer-Encoding", "binary")// 表示传输过程中的编码形式，乱码问题可能就是因为它
	c.File("static/gin.png")
}
// 注意，文件下载浏览器可能会有缓存，这个要注意一下 解决办法就是加查询参数