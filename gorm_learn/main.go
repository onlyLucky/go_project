package main

import (
	"fmt"
	"gormLearn/connect"
	"gormLearn/singleQuery"
	"gormLearn/table"
)

func main() {
	// 1.连接
	db := connect.ConnectFunc()
	fmt.Println(db)
	// 2.模型定义
	table.CreateTableFunc(db)
	// 3.单表查询
	singleQuery.SingleQueryFunc(db)
}