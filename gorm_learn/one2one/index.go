package one2one

import (
	"fmt"

	"gorm.io/gorm"
)

/*
一对一关系比较少，一般用于表的扩展
例如一张用户表，有很多字段
那么就可以把它拆分为两张表，常用的字段放主表，不常用的字段放详情表
*/

func OneToOneFunc(DB *gorm.DB) {
	// 1.表结构搭建
	createFunc(DB)
	// 2.添加记录
	addDataFunc(DB)
	// 3.查询
	queryFunc(DB)
}

type User struct {
  ID       uint
  Name     string
  Age      int
  Gender   bool
  UserInfo UserInfo // 通过UserInfo可以拿到用户详情信息
}
type UserInfo struct {
	User   *User  // 要改成指针，不然就嵌套引用了
  UserID uint // 外键
  ID     uint
  Addr   string
  Like   string
}

/* 1.表结构搭建 */
func createFunc(DB *gorm.DB){
	DB.Migrator().DropTable(&User{},&UserInfo{})
	DB.AutoMigrate(&User{},&UserInfo{})
}

/* 2.添加记录 添加用户，自动添加用户详情 */
func addDataFunc(DB *gorm.DB){
	DB.Create(&User{
		Name:   "枫枫",
		Age:    21,
		Gender: true,
		UserInfo: UserInfo{
			Addr: "湖南省",
			Like: "写代码",
		},
	})

	// 添加用户详情，关联已有用户
	DB.Create(&User{
		Name:   "峰峰",
		Age:    15,
		Gender: true,
	})
	var user User
	DB.Take(&user,2)
	DB.Create(&UserInfo{
		User: &user,
		UserID: 2,
		Addr:"南京市",
		Like: "唱跳rap",
	})
}
/* 3. 查询 */
func queryFunc(DB *gorm.DB){
	var users []User
	DB.Preload("UserInfo").Find(&users)
	fmt.Println(users)
}