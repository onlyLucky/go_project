package _case

import "fmt"

func VarDeclareCase() {
	// 通过var 声明变量
	var i, z, x int
	// 通过var 声明变量 并进行赋值
	var j int = 100
	var f float32 = 100.23

	//通过:= 推断的方式定义变量并赋值，此方式只能用于局部变量定义
	b := true

	//数组
	var arr = [5]int{1, 2, 3, 4, 5}
	arr1 := [...]int{2, 3, 4, 5, 6}
	var arr2 [5]int
	arr2[2] = 4
	arr2[3] = 5
	fmt.Println(i, z, x, j, f, b, arr, arr1, arr2)

	//指针类型 用来表示变量地址类型
	var intPtr *int
	var floatPtr *float32
	var i1 = 100
	f1(&i1)
	// 指针引用类型默认值为nil，值类型一定会有一个默认值。

	//接口类型 空接口类型可以接收任何类型
	var inter interface{}
	inter = i1
	fmt.Println(i1, intPtr, floatPtr, inter)
}

func f1(i *int) {
	//指针类型前面加* 表示取值
	*i = *i + 1
}
