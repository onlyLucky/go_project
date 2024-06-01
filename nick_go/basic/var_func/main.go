package main

import (
	_case "basic/var_func/case"
	"fmt"
)

func main() {
	a := 10
	b := 20
	// 函数传参为变量的值 将变量的值进行深拷贝 并不会影响变量的值
	fmt.Println(_case.SumCase(a, b)) // 30 <nil>
	fmt.Println(a, b)                // 10 20
	_case.ReferenceCase(a, &b)
	fmt.Println(a, b) // 10 21

	fmt.Println(_case.G) // 0
	_case.ScopeCase(a, b)
	fmt.Println(_case.G) // 131

	user := _case.NewUser("nick", 18)
	fmt.Println(user.GetName(), user.GetAge()) // nick 18
}
