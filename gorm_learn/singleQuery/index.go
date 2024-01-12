package singleQuery

import (
	"fmt"
	"math/rand"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SingleQueryFunc(DB *gorm.DB) {
	/* 
	// 1.建表
	createTableFunc(DB)
	// 2.单表插入
	addDateFunc(DB)
	// 3.批量插入
	multiAddDateFunc(DB) 
	 */
	// 4.查询单条记录
	querySingleFunc(DB)
}

// 创建表结构
type User struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}
func createTableFunc(DB *gorm.DB){
	DB.AutoMigrate(&User{})
}
// 添加记录
func addDateFunc(DB *gorm.DB){
	email := "xxx@qq.com"
	// 创建记录
	user := User{
		Name: "枫枫",
		Age: 21,
		Gender: true,
		Email: &email,
	}
	DB.Create(&user)
	/* 
	有两个地方需要注意
		1.指针类型是为了更好的存null类型，但是传值的时候，也记得传指针
		2.Create接收的是一个指针，而不是值
	*/
}

// 批量插入
func multiAddDateFunc(DB *gorm.DB){
	var userList []User
	for i:=0; i<100; i++{
		email:=generateRandomEmail()
		userList = append(userList, User{
			Name: fmt.Sprintf("机器人%d号",i+1),
			Age: rand.Intn(39)+1, // 随机生成1-40
			Gender: rand.Intn(2)==0, // 随机true false [0,2)  0,1
			Email: &email,
		})
	}
	DB.Create(&userList)
}

func generateRandomEmail()string{
	// 随机用户部分（使用字母、数字混合）6-10 
	username := strings.ToLower(randomString(rand.Intn(4)+6))
	// 域名后缀
	domainSuffixes := []string{"gmail.com","yahoo.com","hotmail.com", "outlook.com","icloud.com"}
	domain := domainSuffixes[rand.Intn(len(domainSuffixes))]
	return fmt.Sprintf("%s@%s", username, domain)
}

// 生成指定长度的随机字符串
func randomString(length int) string{
	lettersAndDigits := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = lettersAndDigits[rand.Intn(len(lettersAndDigits))]
	}
	return string(bytes)
}

// 查询单条记录
func querySingleFunc(DB *gorm.DB){
	var user User
	DB.Take(&user)
	fmt.Println(user,*user.Email)
	// 获取单条记录的方法很多，我们对比sql就很直观了
	mysqlLogger := logger.Default.LogMode(logger.Info)
	DB = DB.Session(&gorm.Session{Logger: mysqlLogger})
	var firstUser User
	DB.First(&firstUser)
	fmt.Println(firstUser,*firstUser.Email)
	var lastUser User
	DB.Last(&lastUser)
	fmt.Println(lastUser,*lastUser.Email)
	/* 根据主键查询 */
	var keyUser User
	DB.Take(&keyUser,2)
	fmt.Println(keyUser,*keyUser.Email)
	keyUser=User{}  // 重新赋值
	DB.Take(&keyUser,"4")
	fmt.Println(keyUser,*keyUser.Email)
}