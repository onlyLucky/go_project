package _case

import "fmt"

// 如果值类型，go会给相应初始化赋值为默认值  如果是指针类型，在声明的时候要进行变量初始化赋值
type user struct {
	Name string
	Age  uint
	Addr Address // 如果类型为指针类型*Address，这里就会报空指针错误。默认值是{}
}

type Address struct {
	Province string
	City     string
}

func StructCase() {
	//值类型
	u := user{
		Name: "nick",
		Age:  18,
	}
	f2(u)
	// 指针类型
	u1 := &user{
		Name: "nick",
		Age:  20,
	}
	// 指针类型
	u2 := new(user) // 默认值 &{ 0}
	u2.Name = "nick"
	u2.Age = 21
	// 结构体为值类型，定义变量后默认初始化
	var u3 user
	fmt.Println(u, u1, u2, u3) // {nick 18 { }} &{nick 20 { }} &{nick 21 { }} { 0 { }}
}

func f2(u user) {
	u.Age = 20
}
