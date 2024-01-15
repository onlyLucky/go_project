package query

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)


func MoreQueryFunc(DB *gorm.DB) {
	/* // 1.子查询
	childQuery(DB)
	// 2.命名参数
	nameQuery(DB)
	// 3.find到map
	findToMapQuery(DB)
	// 4.查询引用Scope
	var users []User
	DB.Scopes(scopeQuery).Find(&users)
	PrintQueryData(users,"查询引用Scope:") */
}

/* 1.子查询 */
func childQuery(DB *gorm.DB){
	// 原生sql: select * from students where age > (select avg(age) from students);
	var users []User
	DB.Model(User{}).Where("age > (?)", DB.Model(User{}).Select("avg(age)")).Find(&users)
	PrintQueryData(users,"子查询,查询大于平均年龄的用户:")
} 
/* 2.命名参数 */
func nameQuery(DB *gorm.DB){
	var users []User
	DB.Where("name = @name and age = @age", sql.Named("name","枫枫"), sql.Named("age",23)).Find(&users)
	// DB.Where("name = @name and age = @age", map[string]any{"name":"枫枫","age":23}).Find(&users)
	PrintQueryData(users,"命名参数:")
}

/* 3.find到map */
func findToMapQuery(DB *gorm.DB){
	var res []map[string]any
	DB.Table("t_user").Find(&res)
	fmt.Println(res)
}


/* 4.查询引入Scope */
func scopeQuery(DB *gorm.DB) *gorm.DB {
	// 可以再model层写一些通用的查询方式，这样外界就可以直接调用方法即可
	return DB.Where("age > ?", 23)
}