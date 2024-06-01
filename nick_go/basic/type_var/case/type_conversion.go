package _case

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

func ConvertCase() {
	/*数字类型转换*/
	var num1 int = 100
	// 低精度向高精度类型转换，不会丢失精度，数字过大反之则丢失
	fmt.Println(int64(num1)) // 100
	var num2 int64 = 100
	fmt.Println(int(num2)) // 100

	/*字符串和数字转换*/
	var num3 = 100
	fmt.Println(strconv.Itoa(num3) + "abc") // 100abc
	var str1 = "100"
	fmt.Println(strconv.Atoi(str1)) // 100 <nil>

	var num4 int64 = 1010
	fmt.Println(strconv.FormatInt(num4, 10)) // 1010

	var str2 = "1010"
	fmt.Println(strconv.ParseInt(str2, 10, 64)) // 1010 <nil>

	/* 字符串与[]byte 转换 */
	var str3 = "今天天气很好"
	bytes1 := []byte(str3)
	fmt.Println(bytes1)         // [228 187 138 229 164 169 229 164 169 230 176 148 229 190 136 229 165 189]
	fmt.Println(string(bytes1)) // 今天天气很好

	/* 字符串与rune转换（int32） */
	// 将字符申转换为rune切片，实际上rune切片中存储了字符申的Unicode码点
	var rune1 = []rune(str3)
	fmt.Println(rune1)         // [20170 22825 22825 27668 24456 22909] 6个字符 码点
	fmt.Println(string(rune1)) // 今天天气很好
	fmt.Println(rune1[3])      // 气
	fmt.Println([]int32(str3)) // [20170 22825 22825 27668 24456 22909]

	/* 接口类型转其他类型 */
	var inf interface{} = 100
	var infStruct interface{} = user{}
	i, ok := inf.(int)
	fmt.Println(i, ok) // 100 true
	u, ok := infStruct.(user)
	fmt.Println(u, ok) // { 0 { }} true

	/* 时间类型转字符串 */
	var t time.Time
	t = time.Now()
	timeStr := t.Format("2006-01-02 15:04:05Z07:00") // Z07:00 表示时区
	fmt.Println(timeStr)                             // 2024-06-01 20:23:53+08:00 东八区

	/* 字符串转时间 */
	t2, _ := time.Parse("2006-01-02 15:04:05Z07:00", timeStr)
	fmt.Println(t2)

	/* uintptr */
	u1 := user{}
	// unsafe.Pointer 是一个通用的指针类型，不能用于计算
	uPtr := unsafe.Pointer(&u1)
	// 这里获取对象下面Name属性指针地址 一般都是 对象地址+对象偏移量属性指针地址
	// unsafe.Pointer不能用于计算，所以uPtr要进行转换 uintptr(uPtr)
	namePtr := unsafe.Pointer(uintptr(uPtr) + unsafe.Offsetof(u1.Name))
	// *string 将 namePtr 转化为字符串指针 再加一个* 转化为值类型 进行赋值
	*(*string)(namePtr) = "nick"
	fmt.Println(u1) // {nick 0 { }}
}
