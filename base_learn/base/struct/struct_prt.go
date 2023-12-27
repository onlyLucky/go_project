package struct_base

import "fmt"

type students struct {
	name string
	age int
}

func setAge(info students,age int){
	info.age = age
}
func setAge1(info *students,age int){
	info.age = age
}

func init() {
	fmt.Println("struct_prt init ======")
	s := students{
    name: "枫枫",
    age:  21,
  }
  fmt.Println(s.age)
  setAge(s, 18)
  fmt.Println(s.age)
  setAge1(&s, 17)
  fmt.Println(s.age)
}