package _case

import "fmt"

type Cat struct {
	animal
}

func NewCat() AnimalI {
	return &Cat{}
}

// 重写
func (c *Cat) Eat() {
	fmt.Println("猫吃老鼠")
}
