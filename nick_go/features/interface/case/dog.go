package _case

import "fmt"

type Dog struct {
	animal
}

func NewDog() AnimalI {
	// 类型T实现接口，不管是T还是*T都实现了该接口
	return &Dog{}
	// return Dog{}
}

// 重写
func (d Dog) Eat() {
	fmt.Println("狗吃肉包子")
}
func (d Dog) Drink() {
	fmt.Println("狗喝水")
}
func (d Dog) Sleep() {
	fmt.Println("狗睡觉")
}
func (d Dog) Run() {
	fmt.Println("狗跑步")
}
