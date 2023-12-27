package channel_base

import "fmt"

func init() {
	var c chan int // 声明一个传递整形的通道
	// 初始化通道
	c = make(chan int, 1)
	c <- 1
	//c <- 2 // 会报错 deadlock
	fmt.Println(<-c)
	//fmt.Println(<-c) // 再取也会报错  deadlock

	c <-2
	n,ok := <-c
	fmt.Println(n,ok)
	defer close(c) // 关闭协程
	c <-3 // 关闭之后就不能再写或读了  send on closed channel
	fmt.Println(c)
}
