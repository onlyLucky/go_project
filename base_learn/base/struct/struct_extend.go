package struct_base

import "fmt"

type people struct {
	time string
}

func (p people) info() {
	fmt.Println("people", p.time)
}

type student struct{
	people
	name string
	age int
}

func (s student) printInfo(){
	fmt.Printf("name: %s age:%d\n", s.name,s.age)
}

func init() {
	p:=people{time:"2023-12-04 08:00"}
	
	s:=student{
		people: p,
		name: "f",
		age: 18,
	}

	s.name = "ff"
	s.printInfo()
	s.info()
	fmt.Println(s.time,s.people.time)
}