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
  // 1.创建数据
  createFunc(DB)
  // 2.添加数据
  addDataFunc(DB)
  // 3.查询
  queryFunc(DB)
  // 4.更新
  uploadFunc(DB)
  // 5.customizeFunc
}

type Tag struct {
  ID       uint
  Name     string
  Articles []Article `gorm:"many2many:article_tags;"` // 用于反向引用
}

type Article struct {
  ID    uint
  Title string
  Tags  []Tag `gorm:"many2many:article_tags;"`
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
/* 5.自定义连接表 */
func customizeFunc(DB *gorm.DB){
  type Article struct {
    ID    uint
    Title string
    Tags  []Tag `gorm:"many2many:article_tags"`
  }
  
  type Tag struct {
    ID   uint
    Name string
  }
  
  type ArticleTag struct {
    ArticleID uint `gorm:"primaryKey"`
    TagID     uint `gorm:"primaryKey"`
    CreatedAt time.Time
  }
}



