package my_error

import (
	"fmt"
	"runtime/debug"
)

func read() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)// 捕获异常，打印错误信息
			// 打印错误的堆栈信息
      s := string(debug.Stack())
      fmt.Println(s)
		}
	}()
	var list = []int{2,3}
	fmt.Println(list[2])
}

func init(){
	read()
}
