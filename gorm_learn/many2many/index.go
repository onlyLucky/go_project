package many2many

import "gorm.io/gorm"

/*
多对多关系，需要用第三张表存储两张表的关系
*/
func ManyToManyFunc(DB *gorm.DB) {

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