package affairs

import (
	"fmt"

	"gorm.io/gorm"
)

func AffairsFunc(DB *gorm.DB) {
	// 1.建表
	createFunc(DB)
	addDataFunc(DB)
	// 2.普通事务
	generalTransactionFunc(DB)
	generalTransaction1Func(DB)
	// 3.手动事务
	manualTransactionFunc(DB)
}

type User struct {
  ID    uint   `json:"id"`
  Name  string `json:"name"`
  Money int    `json:"money"`
}

// InnoDB引擎才支持事务，MyISAM不支持事务
// DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

/* 
1. 创建表
*/
func createFunc(DB *gorm.DB){
	DB.Migrator().DropTable(&User{})
	DB.AutoMigrate(&User{})
}
func addDataFunc(DB *gorm.DB){
	DB.Create(&User{
		Name: "张三",
		Money: 1000,
	})
	DB.Create(&User{
		Name: "李四",
		Money: 1000,
	})
}
/* 
2.普通事务
以张三给李四转账为例，不使用事务的后果
*/

func generalTransactionFunc(DB *gorm.DB){
	var zhangsan,lisi User
	DB.Take(&zhangsan,"name = ?","张三")
	DB.Take(&lisi,"name = ?","李四")
	// 张三给李四转账100元
	// 先给张三-100
	zhangsan.Money -= 100
	DB.Model(&zhangsan).Update("money",zhangsan.Money)
	// 模拟失败的情况

	// 再给李四+100
	lisi.Money += 100
	DB.Model(&lisi).Update("money",lisi.Money)
}

func generalTransaction1Func(DB *gorm.DB){
	var zhangsan,lisi User
	DB.Take(&zhangsan,"name = ?","张三")
	DB.Take(&lisi,"name = ?","李四")
	// 张三给李四转账100元
	DB.Transaction(func(tx *gorm.DB) error{
		// 先给张三-100
		zhangsan.Money -= 100
		err := tx.Model(&zhangsan).Update("money",zhangsan.Money).Error
		if err != nil {
			fmt.Println(err)
			return err
		}

		// 再给李四 +100
		lisi.Money += 100
		err = tx.Model(&lisi).Update("money",lisi.Money).Error
		if err != nil{
			fmt.Println(err)
			return err
		}

		// 提交事务
		return nil
	})
}

// 使用事务之后，他们就是一体，一起成功，一起失败

/* 
3.手动事务
// 开始事务
tx := db.Begin()
// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
tx.Create(...)
// ...
// 遇到错误时回滚事务
tx.Rollback()
// 否则，提交事务
tx.Commit()
*/

func manualTransactionFunc(DB *gorm.DB){
	var zhangsan,lisi User
	DB.Take(&zhangsan,"name = ?","张三")
	DB.Take(&lisi,"name = ?","李四")
	// 张三给李四转账100元
	tx := DB.Begin()

	// 先给张三-100
	zhangsan.Money -= 100
	err := tx.Model(&zhangsan).Update("money",zhangsan.Money).Error
	if err != nil{
		tx.Rollback()
	}
	// 再给李四+100
	lisi.Money += 100
	err = tx.Model(&lisi).Update("money",lisi.Money).Error
	if err != nil{
		tx.Rollback()
	}
	// 提交事务
	tx.Commit() 
}