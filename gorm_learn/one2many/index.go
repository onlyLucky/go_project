package one2many

import (
	"gorm.io/gorm"
)

/*
一对多关系 表结构建立

在gorm中，官方文档是把一对多关系分为了两类，Belongs To 属于谁、Has Many 我拥有的
*/

func OneToManyFunc(DB *gorm.DB) {
	// 1.建表
	createTable(DB)
	// 2.一对多的添加
	addFuc(DB)
	// 3.外键添加
	addForeignKeyFunc(DB)
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

/* 2.一对多的添加 */

func addFuc(DB *gorm.DB){
	// 创建用户，并且创建文章
	a1 := Article{Title: "python"}
	a2 := Article{Title: "golang"}
	user := User{Name: "枫枫",Articles: []Article{a1,a2}}
	DB.Create(&user)

	// 创建文章，关联已有用户
	a1 = Article{Title: "golang零基础入门",UserID: 1}
	DB.Create(&a1)
	
	user = User{}
	DB.Take(&user,1)
	DB.Create(&Article{Title: "python零基础入门",User: user})
}

/* 3.外键添加 */
func addForeignKeyFunc(DB *gorm.DB){
	var user User
	DB.Take(&user,2)
	var article Article
	DB.Take(&article, 5)
	user.Articles = []Article{article}
	DB.Save(&user)
}
