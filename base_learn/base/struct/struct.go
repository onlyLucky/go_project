package struct_base

import "fmt"

type Student struct {
	name string
	age  int
}

func (s Student) printInfo() {
	fmt.Printf("name:%s age:%d\n", s.name, s.age)
}

func init() {
	s:=Student{
		name: "f",
		age: 18,
	}

	s.name = "ff"
	s.printInfo()

}