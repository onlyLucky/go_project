package singleQuery

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SingleQueryFunc(DB *gorm.DB) {
	
	// 1.建表
	// createTableFunc(DB)
	// 2.单表插入
	// addDateFunc(DB)
	// 3.批量插入
	// multiAddDateFunc(DB) 
	
	// 4.查询单条记录
	// querySingleFunc(DB)
	// 5.查询多条记录
	// multiQuery(DB)
	// 6.查询like条件的数据
	// likeQueryFunc(DB)
	// 7.更新
	// uploadFunc(DB)
	// 8.删除
	// delFunc(DB)
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
	// Take的第二个参数，默认会根据主键查询，可以是字符串，可以是数字

	/* 根据其他条件查询 */
	var otherUser User
	DB.Take(&otherUser,"name=?","机器人27号")
	fmt.Println(otherUser,*otherUser.Email)
	// 使用？作为占位符，将查询的内容放入?从而自带防止sql注入
	otherUser = User{}
	DB.Take(&otherUser,fmt.Sprintf("name='%s'","机器人27号' or 1=1#"))
	fmt.Println(otherUser,*otherUser.Email) 

	/* 根据struct查询 */
	var structUser User
	// 只有一个主要值
	structUser.ID = 2
	structUser.Name = "机器人1号"
	DB.Take(&structUser)
	fmt.Println(structUser,*structUser.Email)

	/* 获取查询结果 */
	// 获取查询的记录数
	var userList []User
	count := DB.Find(&userList).RowsAffected
	fmt.Println("count: ",count)
	// 查询是否失败
	err := DB.Find(&userList).Error
	fmt.Println("err: ",err)

	var conditionUser User
	e := DB.Take(&conditionUser,"0").Error
	switch e{
	case gorm.ErrRecordNotFound:
		fmt.Println("not found")
	default:
		fmt.Println(e)
	}
}

/* 查询多条记录 */
func multiQuery(DB *gorm.DB){
	var users []User
	DB.Find(&users)
	for _, user := range users{
		fmt.Println(user,*user.Email)
	}

	// 由于email是指针类型，所以看不到实际的内容(序列化之后，会转换为我们可以看得懂的方式)
	var userList []User
	DB.Find(&userList)
	for _, user := range users{
		data,_ := json.Marshal(user)
		fmt.Println(string(data))
	}

	// 根据主键列表查询
	userList = []User{}
	DB.Find(&userList,[]int{1,3,5,7,9})
	// DB.Find(&studentList, 1, 3, 5, 7, 9)  同上效果
	fmt.Println(userList)

	/* 根据其他条件查询 */
	userList = []User{}
	DB.Find(&userList,"name in ?",[]string{"枫枫","zhangsan"})
	fmt.Println(userList)
}

// LIKE条件查询
func likeQueryFunc(DB *gorm.DB){
	// 单条
	var user User
	DB.Where("name LIKE ?","%机器人9%").Find(&user)
	fmt.Println(user)
	var users []User
	DB.Where("name LIKE ?","%机器人9%").Find(&users)
	for _, user := range users{
		data,_ := json.Marshal(user)
		fmt.Println(string(data))
	}
	
}

/* 更新 */
func uploadFunc(DB *gorm.DB){
	var user User
	DB.Take(&user)
	user.Age = 15  // 0 零值也会更新
	// 全字段更新
	DB.Save(&user)


	/* 更新指定字段 */
	user = User{}
	DB.Take(&user)
	user.Age = 20
	DB.Select("age").Save(&user)

	/* 批量更新 */
	var userList []User
	DB.Find(&userList,"age = ?",20).Update("email","dingding@hz.com")
	// UPDATE `t_user` SET `email`='dingding@hz.com' WHERE age = 20 AND `id` IN (1,73,100)
	DB.Model(&User{}).Where("age = ?", 23).Update("email","dingding@hz23.com")
	// UPDATE `t_user` SET `email`='dingding@hz23.com' WHERE age = 23

	/* 更新多列 */
	// 如果是结构体，它默认不会更新零值
	email := "dingding@hz24.com"
	DB.Model(&User{}).Where("age=?",24).Updates(User{
		Email: &email,
		Gender: false, // 这个不会更新
	})

	// 如果想让他更新零值，用select就好
	email = "dingding@hz25.com"
	DB.Model(&User{}).Where("age = ?", 25).Select("gender", "email").Updates(User{
		Email:  &email,
		Gender: false,
	})

	// 如果不想多写几行代码，则推荐使用map
	email = "dingding@hz26.com"
	DB.Model(&User{}).Where("age = ?", 26).Updates(map[string]any{
		"email":  &email,
		"gender": false,
	})

	/* 更新选定字段 */
	// Select选定字段

	// Omit忽略字段
}

// 删除
func delFunc(DB *gorm.DB){
	var user User = User{ID: 10}
	DB.Delete(&user)
	// DELETE FROM `t_user` WHERE `t_user`.`id` = 10

	/* 删除多个 */
	DB.Delete(&User{}, []int{1,2,3})
	// DELETE FROM `t_user` WHERE `t_user`.`id` IN (1,2,3)

	// 查询到的切片列表
	var users []User
	DB.Where("name LIKE ?","%机器人9%").Find(&users)
	DB.Delete(&users)
	// DELETE FROM `t_user` WHERE `t_user`.`id` IN (91,92,93,94,95,96,97,98,99,100)
}