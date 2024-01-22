package main

import (
	"fmt"
	"redisLearn/basic"
	"redisLearn/connect"
)

func main() {
	fmt.Println("hello redis!!!")
	// 1.连接
	connect.ConnectFunc()
	// 2.基础操作
	basic.StringFuc(connect.DB)
}