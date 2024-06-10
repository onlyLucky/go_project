package _case

import "fmt"

type Dove struct {
	animal
}

func NewDove() AnimalI {
	// 类型*T实现接口，只有T类型的指针实现了该接口
	return &Dove{}
	// return Dove{} // 会报错 method Drink has pointer receiver  要是一个指针receiver
}

// 重写 receiver 为指针类型
func (d *Dove) Eat() {
	fmt.Println("鸽子吃虫子")
}
func (d *Dove) Drink() {
	fmt.Println("鸽子喝水")
}
func (d *Dove) Sleep() {
	fmt.Println("鸽子睡觉")
}
func (d *Dove) Run() {
	fmt.Println("鸽子助跑起飞")
}
