package main

import (
	"fmt"
	"gormLearn/affairs"
	"gormLearn/connect"
	"gormLearn/dataType"
	"gormLearn/many2many"
	"gormLearn/one2many"
	"gormLearn/one2one"
	"gormLearn/query"
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
	// 4.建表Hook  /hook/index.go
	// 5.高级查询
	query.QueryDataFunc(db)
	query.MoreQueryFunc(db)
	// 6.一对多关系
	one2many.OneToManyFunc(db)
	// 7.一对一关系
	one2one.OneToOneFunc(db)
	// 8.多对多关系
	many2many.ManyToManyFunc(db)
	// 9.自定义数据类型
	dataType.DataTypeFunc(db)
	// 10.事务
	affairs.AffairsFunc(db)
}