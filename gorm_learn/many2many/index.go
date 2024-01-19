package many2many

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

/*
多对多关系，需要用第三张表存储两张表的关系
*/
func ManyToManyFunc(DB *gorm.DB) {
  /* // 1.创建数据
  createFunc(DB)
  // 2.添加数据
  addDataFunc(DB)
  // 3.查询
  queryFunc(DB)
  // 4.更新
  uploadFunc(DB)
  // 5.自定义连接表
  customizeFunc(DB)
  // 6.操作案例
  demoFunc(DB)
  // 7.自定义连接表主键 
  customizeTipFunc(DB) */
}

type Tag struct {
  ID       uint
  Name     string
  Articles []Article `gorm:"many2many:article_tag;"` // 用于反向引用
}

type Article struct {
  ID    uint
  Title string
  Tags  []Tag `gorm:"many2many:article_tag;"`
}

/* 1.表结构创建 */
func createFunc(DB *gorm.DB){
	DB.Migrator().DropTable(&Tag{},&Article{})
	DB.AutoMigrate(&Tag{},&Article{})
}

/* 2.多对多添加 */
func addDataFunc(DB *gorm.DB){
  // 添加文章，并创建标签
  DB.Create(&Article{
    Title:"python基础课程",
    Tags:[]Tag{
      {Name:"python"},
      {Name:"基础课程"},
    },
  })
  // 添加文章，选择标签
  var tags []Tag
  DB.Find(&tags,"name=?","基础课程")
  DB.Create(&Article{
    Title: "golang基础",
    Tags: tags,
  })
}

/* 3.多对多查询 */
func queryFunc(DB *gorm.DB){
  // 查询文章，显示文章的标签列表
  var article Article
  DB.Preload("Tags").Take(&article, 1)
  fmt.Println(article)
  // 查询标签，显示文章列表
  var tag Tag
  DB.Preload("Articles").Take(&tag,2)
  fmt.Println(tag)
}

/* 4.多对多更新 */
func uploadFunc(DB *gorm.DB){
  // 移除文章的标签
  var article Article
  DB.Preload("Tags").Take(&article,1)
  DB.Model(&article).Association("Tags").Delete()
  fmt.Println(article)
  // 更新文章的标签
  article = Article{}
  var tags []Tag
  DB.Find(&tags,[]int{1,2,6,7})
  DB.Preload("Tags").Take(&article,2)
  DB.Model(&article).Association("Tags").Replace(tags)
  fmt.Println(article)
}
type ArticleTag struct {
  ArticleID uint `gorm:"primaryKey"`
  TagID     uint `gorm:"primaryKey"`
  CreatedAt time.Time
}
/* 5.自定义连接表 */
func customizeFunc(DB *gorm.DB){
  type Article struct {
    ID    uint
    Title string
    Tags  []Tag `gorm:"many2many:article_tag"`
  }
  
  type Tag struct {
    ID   uint
    Name string
  }
  DB.Migrator().DropTable(&Tag{},&Article{},&ArticleTag{})
	DB.AutoMigrate(&Tag{},&Article{},&ArticleTag{})
}

/* 6.操作案例 */
func demoFunc(DB *gorm.DB){
  // 添加文章并添加标签，并自动关联
  DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})  // 要设置这个，才能走到我们自定义的连接表
  DB.Create(&Article{
    Title: "flask零基础入门",
    Tags: []Tag{
      {Name: "python"},
      {Name: "后端"}, 
      {Name: "web"},
    },
  })
  // 添加文章，关联已有标签
  var tags []Tag
  DB.Find(&tags,"name in ?",[]string{"python","web"})
  DB.Create(&Article{
    Title: "flash请求对象",
    Tags: tags,
  })
  // 给已有文章关联标签
  article := Article{
    Title: "django基础",
  }
  DB.Create(&article)
  var at Article
  tags = []Tag{}
  DB.Find(&tags,"name in ?",[]string{"python","web"})
  DB.Take(&at,article.ID).Association("Tags").Append(tags)
  // 替换已有文章的标签
  article = Article{}
  tags = []Tag{}
  DB.Find(&tags,"name in ?",[]string{"后端"})
  DB.Take(&article,"title = ?", "django基础")
  DB.Model(&article).Association("Tags").Replace(tags)
  // 
  articles := []Article{}
  DB.Preload("Tags").Find(&articles)
  fmt.Println(articles)
}

/* 
7.自定义连接表主键 
实用案例：那么按照gorm默认的主键名，那就分别是ArticleModelID，TagModelID，太长了，根本就不实用
joinForeignKey 连接的主键id
JoinReferences 关联的主键id
*/
type ArticleModel struct {
  ID    uint
  Title string
  Tags  []TagModel `gorm:"many2many:article_tags;joinForeignKey:ArticleID;JoinReferences:TagID"`
}

type TagModel struct {
  ID       uint
  Name     string
  Articles []ArticleModel `gorm:"many2many:article_tags;joinForeignKey:TagID;JoinReferences:ArticleID"`
}

type ArticleTagModel struct {
  ArticleID uint `gorm:"primaryKey"` // article_id
  TagID     uint `gorm:"primaryKey"` // tag_id
  CreatedAt time.Time
}

func customizeTipFunc(DB *gorm.DB){
  DB.SetupJoinTable(&ArticleModel{}, "Tags", &ArticleTagModel{})
  DB.SetupJoinTable(&TagModel{}, "Articles", &ArticleTagModel{})
  DB.Migrator().DropTable(&ArticleModel{},&TagModel{},&ArticleTagModel{})
	DB.AutoMigrate(&ArticleModel{},&TagModel{},&ArticleTagModel{})
}
