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
	/* // 1.插入数据
	// setMoreData(DB)
	// 2.where 查询
	whereQuery(DB)
	// 3.结构体查询
	structQuery(DB)
	// 4.使用map查询 不会过滤零值
	mapQuery(DB)
	// 5.not条件 和where中的not等价
	notQuery(DB)
	// 6.Or条件
	orQuery(DB)
	// 7.Select 选择字段 
	selectQuery(DB)
	// 8.排序
	sortQuery(DB)
	// 9.分页查询
	pageQuery(DB)
	// 10.去重
	distinctQuery(DB)
	// 11.分组查询
	groupQuery(DB)
	// 12.执行原生sql
	sqlQuery(DB) */
}

/* 1.插入一些数据 */
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

/* 2.where  等价于sql语句中的where*/
func whereQuery(DB *gorm.DB){
	var users []User
	// 查询用户名是枫枫的
	DB.Where("name=?","枫枫").Find(&users)
	PrintQueryData(users,"查询用户名是枫枫的:")
	// 查询用户名不是枫枫
	DB.Where("name <> ?","枫枫").Find(&users)
	PrintQueryData(users,"查询用户名不是枫枫:")
	// 查询用户名包含 如燕，李元芳的
	DB.Where("name in ?",[]string{"如燕", "李元芳"}).Find(&users)
	PrintQueryData(users,"查询用户名包含 如燕，李元芳的:")
	// 查询姓李的
	DB.Where("name like ?","李%").Find(&users)
	PrintQueryData(users,"查询姓李的:")
	// 查询年龄大于23，是qq邮箱的
	DB.Where("age > ? and email like ?","23","%@qq.com").Find(&users)
	PrintQueryData(users,"查询年龄大于23，是qq邮箱的:")
	// 查询是qq邮箱的，或者是女的
	DB.Where("gender = ? or email like ?",false,"%@qq.com").Find(&users)
	PrintQueryData(users,"查询是qq邮箱的，或者是女的:")
}

/* 3.使用结构体查询  会过滤零值*/
func structQuery(DB *gorm.DB){
	var users []User
	DB.Where(&User{Name: "李元芳",Age: 0}).Find(&users)
	
	PrintQueryData(users,"使用结构体查询:")
}

/* 4.使用map查询 不会过滤零值 */
func mapQuery(DB *gorm.DB){
	var users []User
	DB.Where(map[string]any{"name":"李元芳","age":0}).Find(&users)
	PrintQueryData(users,"使用map查询 不会过滤零值:")
}

/* 5.not条件 和where中的not等价*/
func notQuery(DB *gorm.DB){
	// 排除年龄大于23的
	var users []User
	DB.Not("age>23").Find(&users)
	PrintQueryData(users,"not条件,排除年龄大于23的:")
}

/* 6.Or条件  和where中的or等价*/
func orQuery(DB *gorm.DB){
	var users []User
	DB.Or("gender = ?",false).Or(" email like ?","%@qq.com").Find(&users)
	PrintQueryData(users,"Or条件,查询是qq邮箱的，或者是女的:")
}

/* 7.Select 选择字段 */
func selectQuery(DB *gorm.DB){
	var users []User
	DB.Select("name", "age").Find(&users)
	PrintQueryData(users,"Select 选择字段,没有被选中，会被赋零值:")

	// 可以使用扫描Scan，将选择的字段存入另一个结构体中
	type selectUser struct {
		Name string
		Age  int
	}
	users = []User{}
	var sUsers []selectUser
	DB.Select("name", "age").Find(&users).Scan(&sUsers)
	fmt.Println(sUsers)
	// 上面写法写也是可以的，不过最终会查询两次，还是不这样写
	// 下面这样写就只查询一次了
	DB.Model(&User{}).Select("name", "age").Scan(&sUsers)
	fmt.Println(sUsers)
	
	// 还可以这样
	DB.Table("t_user").Select("name", "age").Scan(&sUsers)
	fmt.Println(sUsers)

	// Scan是根据column列名进行扫描的
	type scanUser struct {
		Name123 string `gorm:"column:name"`
		Age  int
	}
	var colUsers []scanUser
	DB.Table("t_user").Select("name", "age").Scan(&colUsers)
	fmt.Println(colUsers)
}
/* 8.排序 */
func sortQuery(DB *gorm.DB){
	// 根据年龄倒序 desc 降序   asc 升序
	var users []User
	DB.Order("name asc,age desc").Find(&users)
	PrintQueryData(users,"排序,根据年龄倒序:")
}
/* 9.分页查询 */
func pageQuery(DB *gorm.DB){
	var users []User
	// 一页两条，第1页
	DB.Limit(2).Offset(0).Find(&users)
	PrintQueryData(users,"分页查询,一页两条，第1页:")
	// 第2页
	DB.Limit(2).Offset(2).Find(&users)
	PrintQueryData(users,"分页查询,一页两条，第2页:")
	// 第3页
	DB.Limit(2).Offset(4).Find(&users)
	PrintQueryData(users,"分页查询,一页两条，第3页:")

	// 通用的写法
	users = []User{}
	// pageSize
	limit := 2
	// pageNum
	page := 1
	offset := (page - 1) * limit
	DB.Limit(limit).Offset(offset).Find(&users)
	PrintQueryData(users,"分页查询通用写法,一页两条，第1页:")
}

/* 10.去重 */
func distinctQuery(DB *gorm.DB){
	var ageList []int
	DB.Table("t_user").Select("age").Distinct("age").Scan(&ageList)
	fmt.Println(ageList)
	// 其他写法
	DB.Table("t_user").Select("distinct age").Scan(&ageList)
	fmt.Println(ageList)
}

/* 11.分组查询 */
func groupQuery(DB *gorm.DB){
	var ageList []int
	// 查询男生的个数和女生的个数
	DB.Table("t_user").Select("count(id)").Group("gender").Scan(&ageList)
	fmt.Println(ageList)
	// 精确哪一个是男生，哪一个是女生
	type AgeGroup struct {
		Gender int
		Count  int `gorm:"column:count(id)"`
	}
	var ageTotal []AgeGroup
	// 查询男生的个数和女生的个数
	DB.Table("t_user").Select("count(id)","gender").Group("gender").Scan(&ageTotal)
	fmt.Println(ageTotal)
	// 再精确点
	type AgeGroupMore struct{
		Gender int
		Count int `gorm:"column:count(id)"`
		Name string `gorm:"column:group_concat(name)"`
	}

	var ageMore []AgeGroupMore
	// 查询男生的个数和女生的个数
	DB.Table("t_user").Select("count(id)","gender","group_concat(name)").Group("gender").Scan(&ageMore)
	fmt.Println(ageMore)

	/* 使用gorm不会让你忘记原生sql的编写 */
}

/* 12.执行原生sql */
func sqlQuery(DB *gorm.DB){
	type AgeGroupMore struct{
		Gender int
		Count int `gorm:"column:count(id)"`
		Name string `gorm:"column:group_concat(name)"`
	}
	var ageMore []AgeGroupMore
	DB.Raw(`SELECT count(id), gender, group_concat(name) FROM students GROUP BY gender`).Scan(&ageMore)
	fmt.Println(ageMore)
}

func PrintQueryData(list []User,desc string){
	fmt.Println(desc+"=================")
	if(len(list)<=0){
		fmt.Println(list)
		return
	}
	for _, item := range list{
		data,_ := json.Marshal(item)
		fmt.Println(string(data))
	}
}

