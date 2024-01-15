package one2many

import "gorm.io/gorm"

/*
一对多关系 表结构建立

在gorm中，官方文档是把一对多关系分为了两类，Belongs To 属于谁、Has Many 我拥有的
*/

func OneToManyFunc(DB *gorm.DB) {
	// 1.建表
	createTable(DB)
}

// 默认外键关联
type User struct {
	ID       uint      
	Name     string    `gorm:"size:8"`
	Articles []Article // 用户拥有的文章列表
}
// id title user_id
type Article struct {
	ID     uint   
	Title  string `gorm:"size:16"`
	UserID uint   // 属于   这里的类型要和引用的外键类型一致，包括大小
	User   User   // 属于
}

// 关于外键命名，外键名称就是关联表名+ID，类型是uint  UserID  User 表 + ID 默认外键
/* 1.重写外键关联 */

/* type User struct {
  ID       uint      
  Name     string    `gorm:"size:8"`
  Articles []Article `gorm:"foreignKey:UID"` // 用户拥有的文章列表
}
// id title uid
type Article struct {
  ID    uint   
  Title string `gorm:"size:16"`
  UID   uint   // 属于
  User  User   `gorm:"foreignKey:UID"` // 属于
} */
// 我改了Article 的外键，将UID作为了外键，那么User这个外键关系就要指向UID

func createTable(DB *gorm.DB){
	DB.Migrator().DropTable(&User{},&Article{})
	DB.AutoMigrate(&User{},&Article{})
}