package query

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// 创建表结构
type User struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}

func QueryDataFunc(DB *gorm.DB) {
	setMoreData(DB)
	whereQuery(DB)
}

/* 插入一些数据 */
func setMoreData(DB *gorm.DB){
	var users []User
	// 删除所有数据
	DB.Find(&users).Delete(&users) 
	users = []User{
		{ID: 1, Name: "李元芳", Age: 32, Email: ptrString("lyf@yf.com"), Gender: true},
    {ID: 2, Name: "张武", Age: 18, Email: ptrString("zhangwu@lly.cn"), Gender: true},
    {ID: 3, Name: "枫枫", Age: 23, Email: ptrString("ff@yahoo.com"), Gender: true},
    {ID: 4, Name: "刘大", Age: 54, Email: ptrString("liuda@qq.com"), Gender: true},
    {ID: 5, Name: "李武", Age: 23, Email: ptrString("liwu@lly.cn"), Gender: true},
    {ID: 6, Name: "李琦", Age: 14, Email: ptrString("liqi@lly.cn"), Gender: false},
    {ID: 7, Name: "晓梅", Age: 25, Email: ptrString("xiaomeo@sl.com"), Gender: false},
    {ID: 8, Name: "如燕", Age: 26, Email: ptrString("ruyan@yf.com"), Gender: false},
    {ID: 9, Name: "魔灵", Age: 21, Email: ptrString("moling@sl.com"), Gender: true},
	}
	DB.Create(&users)
}

func ptrString(email string) *string {
  return &email
}

/* where  等价于sql语句中的where*/
func whereQuery(DB *gorm.DB){
	var users []User
	// 查询用户名是枫枫的
	DB.Where("name=?","枫枫").Find(&users)
	printQueryData(users)

}

func printQueryData(list []User){
	for _, item := range list{
		data,_ := json.Marshal(item)
		fmt.Println(string(data))
	}
}
