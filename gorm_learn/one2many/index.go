package one2many

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

/*
一对多关系 表结构建立

在gorm中，官方文档是把一对多关系分为了两类，Belongs To 属于谁、Has Many 我拥有的
*/

func OneToManyFunc(DB *gorm.DB) {
	/* // 1.建表
	createTable(DB)
	// 2.一对多的添加
	addFuc(DB)
	// 3.外键添加
	addForeignKeyFunc(DB)
	// 4.查询
	queryFunc(DB)
	// 5.预加载
	reloadFunc(DB)
	// 6.删除
	delFunc(DB) */
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

	a1 = Article{Title: "ABC"}
	a2 = Article{Title: "数据与结构"}
	user = User{Name: "峰峰",Articles: []Article{a1,a2}}
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
	// "id=?",1   1  Take 参数
	/* DB.Take(&user,2)
	var article Article
	DB.Take(&article, 5)
	user.Name="峰峰"
	user.Articles = []Article{article}
	
	DB.Save(&user) */
	DB.Take(&user,1)
	a1 := Article{Title: "java零基础入门",UserID: 2}
	// Association 给User 某个对象中Articles添加数据
	DB.Model(&user).Association("Articles").Append(&a1) // 这里添加塞入前面的user数据要存在才可以

	/* 给现有的文章关联用户  类代码参考，不过要先添加文章，新用户 */
	/* 
	var article Article
	DB.Take(&article, 5)

	article.UserID = 2
	DB.Save(&article)
	*/
	/* 
	// Append方法
	var user User
	DB.Take(&user, 2)

	var article Article
	DB.Take(&article, 5)

	DB.Model(&article).Association("User").Append(&user)
	*/
}

/* 4.查询 */
func queryFunc(DB *gorm.DB){
	// 查询用户，显示用户的文章列表
	var user User
	DB.Take(&user,1)
	fmt.Println(user)
	// 直接这样，是显示不出文章列表
}

/* 5.预加载 */
func reloadFunc(DB *gorm.DB){
	// 我们必须要使用预加载来加载文章列表
	var user User
	DB.Preload("Articles").Take(&user,1)

	data,_ := json.Marshal(user)
	fmt.Println(string(data))

	// 预加载的名字就是外键关联的属性名

	// 查询文章，显示文章用户的信息
	var article Article
	DB.Preload("User").Take(&article,1)
	aData,_ := json.Marshal(article)
	fmt.Println(string(aData))

	/* 嵌套预加载 */
	// 查询文章，显示用户，并且显示用户关联的所有文章，这就得用到嵌套预加载了
	article = Article{}
	user = User{}
	DB.Preload("User.Articles").Take(&article,1)
	fmt.Println(article)

	/* 带条件的预加载 */
	// 查询用户下的所有文章列表，过滤某些文章
	article = Article{}
	user = User{}
	DB.Preload("Articles","id = ?", 1).Take(&user, 1)
	fmt.Println(user)

	/* 自定义预加载 */
	article = Article{}
	user = User{}
	DB.Preload("Articles", func(DB *gorm.DB) *gorm.DB {
		return DB.Where("id in ?", []int{1,2})
	}).Take(&user, 1)
	fmt.Println(user)
}

/* 6.删除 */
func delFunc(DB *gorm.DB){
	// 级联删除 删除用户，与用户关联的文章也会删除
	var user User
	var users []User
	/* DB.Take(&user,2)
	DB.Select("Articles").Delete(&user) */
	
	/* 清除外键关系 删除用户，与将与用户关联的文章，外键设置为null*/
	DB.Preload("Articles").Take(&user,2)
	DB.Model(&user).Association("Articles").Delete(&user.Articles)
	// 查看用户列表
	users = []User{}
	DB.Preload("Articles").Find(&users)
	fmt.Println(users)
}

