package _case

import (
	"fmt"
	"strconv"
)

func ConvertCase() {
	/*数字类型转换*/
	var num1 int = 100
	// 低精度向高精度类型转换，不会丢失精度，数字过大反之则丢失
	fmt.Println(int64(num1))
	var num2 int64 = 100
	fmt.Println(int(num2))

	/*字符串和数字转换*/
	var num3 = 100
	fmt.Println(strconv.Itoa(num3) + "abc")
	var str1 = "100"

}
