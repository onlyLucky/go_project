package _case

import "fmt"

/*
声明AnimalI接口
定义AnimalI行为
*/
type AnimalI interface {
	// 吃
	Eat()
	// 喝
	Drink()
	// 睡觉
	Sleep()
	// 跑
	Run()
}

// animal 为 T
type animal struct{}

// receiver 为值类型
func (a animal) Eat() {
	fmt.Println("Animal Eat 接口默认实现")
}

// receiver 为指针类型
// func (a *animal) Drink() {}

func (a animal) Drink() {
	fmt.Println("Animal Drink 接口默认实现")
}
func (a animal) Sleep() {
	fmt.Println("Animal Sleep 接口默认实现")
}
func (a animal) Run() {
	fmt.Println("Animal Run 接口默认实现")
}

/* func init() {
	// 类型T实例value或pointer（值类型或指针类型）可以调用全部的方法，编译器会自动转换
	a := animal{}
	a.Eat()
	a.Drink()
} */
