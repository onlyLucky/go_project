package reflection

import "fmt"

type User struct {
	Name string
	Age  int
}

func Print(inter interface{}) {
	switch x := inter.(type) {
	case User:
		x.Name = "张三"
		fmt.Println(x.Name, x.Age)
	}
}

func PPrint(user *User){
	user.Name = "王五"
}

func init() {
	fmt.Println("reflection: ======")
	user := User{Name: "枫枫", Age: 21}
	Print(user)
	fmt.Println(&user)
}