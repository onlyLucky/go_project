package interface_base

import "fmt"

type data interface{}

type Dog struct {
	Name string
}

func Print(d data) {
	fmt.Println(d)
}

func init() {
	fmt.Println("space_interface ========")
	d := Dog{Name: "小黑"}

	Print(d)
	Print("123")
	Print(true)
	Print([]int{1,2,3})
	Print(make(map[string]string, 2))
}