package _case

import "errors"

// 形参
// 局部变量
// 全局变量

type User struct {
	Name string
	Age  uint
}

// user *User默认值为 nil nil 进行属性赋值会报错 runtime error: invalid memory address or nil pointer dereference
// func SumCase(a, b int) (user *User, sum int, err error)
func SumCase(a, b int) (sum int, err error) {
	// user.Name = "nick"
	// user.Age = 18
	if a <= 0 && b <= 0 {
		err = errors.New("两数相加不能同时小于0")
		// return 0, err
		return
	}
	sum = a + b
	// return sum,nil
	return
}

// 值传递，引用传递
func ReferenceCase(a int, b *int) {
	a += 1
	*b += 1
}

/* 变量作用域 */
// 全局变量
var g int
var G int

func ScopeCase(a, b int) {
	c := 100
	g = a + b + c
	G = g
}

// 获取user name属性值 传user指针类型
func (u *User) GetName() string {
	return u.Name
}

// 获取user age属性值 传user指针类型
func (u *User) GetAge() uint {
	return u.Age
}

// 返回user 指针类型
func NewUser(name string, age uint) *User {
	return &User{Name: name, Age: age}
}
