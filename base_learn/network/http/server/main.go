package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

/* // handler
func handler(res http.ResponseWriter, req *http.Request) {
  res.Write([]byte("hello go"))
}

func main() {
  // 回调函数
  http.HandleFunc("/index", handler)
  // 绑定服务  http://localhost:8080/index
  http.ListenAndServe(":8080", nil)
} */

func IndexHandler(res http.ResponseWriter, req *http.Request) {
  switch req.Method {
  case "GET":
    data, err := os.ReadFile("E:/Code/go_project/base_learn/network/http/server/index.html")
    if err != nil {
      fmt.Println("GET",data)
    }
    res.Write(data)
    //res.Write([]byte("<h1>hello 枫枫 GET</h1>"))
  case "POST":
    // 获取body数据
    data, err := io.ReadAll(req.Body)
    // 拿请求头
    contentType := req.Header.Get("Content-Type")
    fmt.Println(contentType)
    //switch contentType {
    //case "application/json":
    //
    //}

    if err != nil {
      fmt.Println(data)
    }
    ma := make(map[string]string)
    json.Unmarshal(data, &ma)
    fmt.Println(ma["username"])

    type User struct {
      Username string `json:"username"`
    }
    var user User
    json.Unmarshal(data, &user)
    fmt.Println(user)
    // 设置响应头
    header := res.Header()
    header["token"] = []string{"y1gyf156sdgT%d44hjgj"}
    res.Write([]byte("hello go POST"))
  }
}

func main() {
  // 1. 绑定回调
  http.HandleFunc("/index", IndexHandler)
  // 2. 注册服务
  fmt.Println("listen server on http://localhost:8080")
  http.ListenAndServe(":8080", nil)
}