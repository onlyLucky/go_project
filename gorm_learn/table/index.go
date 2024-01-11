package table

import "gorm.io/gorm"

type Student struct {
	ID uint // 默认使用ID作为主键
	Name string
	Email *string	// 使用指针是为了存空值
}

func CreateTableFunc(DB *gorm.DB) {
	// 可以放多个 AutoMigrate的逻辑是只新增，不删除，不修改（大小会修改）
	DB.AutoMigrate(&Student{})
	// CREATE TABLE `t_student` (`id` bigint unsigned AUTO_INCREMENT,`name` longtext,`email` longtext,PRIMARY KEY (`id`))
	// bigint,longtext默认的类型太大了
}