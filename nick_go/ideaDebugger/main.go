package main

import (
	"flag"
	"fmt"
	"time"
)

var j = flag.Int("j", 0, "")

// 初始化函数，golang每个包的引用会优先调用此函数
func init() {
	fmt.Println("init")
}

// 函数程序入口
func main() {
	flag.Parse()
	fmt.Println(j)
	var i = 0
	fmt.Println(fmt.Sprintf("变量打印： %d", i))
	for i := 0; i < 100; i++ {
		fmt.Println("demo print:", i, *j)
		i++
		time.Sleep(time.Second)
	}
	/**/
}
