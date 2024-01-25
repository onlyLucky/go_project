package main

import (
	"redisLearn/basic"
	"redisLearn/connect"
)

func main() {
	// 1.连接
	connect.ConnectFunc()
	// 2.基础操作
	// 字符串操作
	basic.StringFuc(connect.DB)
	// 列表操作
	basic.ListFunc(connect.DB)
	// 集合
	basic.SetFunc(connect.DB)
	// 哈希
	basic.HashFunc(connect.DB)
	// 有序集合
	basic.SortedSet(connect.DB)
}