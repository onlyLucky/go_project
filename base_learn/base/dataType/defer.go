// Author: fg
// Date: 2023-12-04 09:47:55
// LastEditors: fg
// LastEditTime: 2023-12-04 09:48:02
// Description: content

// defer这种机制可以用于资源的释放、错误处理、性能优化等多种场景
package dataType

import "fmt"

func Func(){
	defer fmt.Println("defer2")
	fmt.Println("func")
	defer fmt.Println("defer1")
}

func init(){
	fmt.Println("defer-init")
	defer fmt.Println("defer4")
	Func()
	defer fmt.Println("defer3")
}


