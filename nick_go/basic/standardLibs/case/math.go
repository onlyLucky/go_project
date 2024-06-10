package _case

import (
	"fmt"
	"math"
)

func MathCase() {
	fmt.Println("2的10次方：", math.Pow(2, 10))
	fmt.Println("返回以2为底，1024的对数: ", math.Log2(1024))
	fmt.Println("返回两数中较大的值：", math.Max(2, 10))
	fmt.Println("向上取整：", math.Ceil(2.49))
	fmt.Println("向下取整：", math.Floor(2.89))
	fmt.Println("90角的正弦值：", math.Sin(math.Pi/2))
	fmt.Println("1的反正弦值：", math.Asin(1))
}
