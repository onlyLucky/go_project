package base_func

import (
	"errors"
	"fmt"
)

// 使用func关键字定义一个函数
func sayHello() {
  fmt.Println("hello")
}

func add(n1 int, n2 int) {
  fmt.Println(n1, n2)
}

// 参数类型一样，可以合并在一起
func add1(n1, n2 int) {
  fmt.Println(n1, n2)
}

// 多个参数
func add2(numList ...int) {
  fmt.Println(numList)
}

// 无返回值
func fun1() {
  return // 也可以不写
}

// 单返回值
func fun2() int {
  return 1
}

// 多返回值
func fun3() (int, error) {
  return 0, errors.New("错误")
}

// 命名返回值
func fun4() (res string) {
  return // 相当于先定义再赋值
  //return "abc"
}



func init(){
	// 函数()调用函数
  sayHello()
	add(1, 2)
  add1(1, 2)
  add2(1, 2)
  add2(1, 2, 3, 4)

	fmt.Println("请输入要执行的操作：")
  fmt.Println(`1：登录
2：个人中心
3：注销`)
	var num int
  fmt.Scan(&num)
  var funcMap = map[int]func(){
    1: func() {
      fmt.Println("登录")
    },
    2: func() {
      fmt.Println("个人中心")
    },
    3: func() {
      fmt.Println("注销")
    },
  }
  funcMap[num]()
}
