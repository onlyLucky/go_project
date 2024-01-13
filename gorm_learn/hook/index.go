package hook

import (
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}

func (user *Student) BeforeCreate(tx *gorm.DB) (err error) {
	email := fmt.Sprintf("%s@qq.com", user.Name)
	user.Email = &email
	return nil
}

/* 
除了BeforeCreate外，Gorm还支持其他的钩子方法：

BeforeSave
AfterSave
BeforeUpdate
AfterUpdate
BeforeDelete
AfterDelete
BeforeFind
AfterFind
每个钩子都有对应的接收器函数签名，例如BeforeCreate(scope *gorm.Scope)和AfterCreate(scope *gorm.Scope)，通过scope参数可以访问并修改即将执行的操作相关的上下文信息。
*/