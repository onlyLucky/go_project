package _case

import "fmt"

func SimpleCase() {
	var a, b = 3, 4
	var c, d float64 = 5, 6
	fmt.Println("不使用泛型，数字比较：", getMaxNumInt(a, b))
	fmt.Println("不使用泛型，数字比较：", getMaxNumFloat(c, d))

	// 由编译器推断输入的类型
	fmt.Println("使用泛型，数字比较：", getMaxNum(a, b))
	fmt.Println("使用泛型，数字比较：", getMaxNum(c, d))
}

func getMaxNumInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxNumFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// func getMaxNum[T interface{int |float64}](a,b T) T
func getMaxNum[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type CusNum interface {
	// 支持uint8 int32 float64 int64及其衍生体
	// ~ 表示支持类型为衍生类型
	// | 表示取并集
	// 多行之间取交集 多行没有交集就不能约束
	uint8 | int32 | float64 | ~int64
	int32 | float64 | ~int64 | uint16
}

// MyInt64为int64的衍生类型 并具有基础类型int64的新类型，与int64不相等
type MyInt64 int64

// MyInt32为int32的别名，与int32是同一种类型
type MyInt32 = int32

func CusNumCase() {
	var a, b int32 = 3, 4
	var a1, b1 MyInt32 = a, b
	fmt.Println("自定义泛型，数字比较：", getMaxCusNumTCase(a, b))
	fmt.Println("自定义泛型，数字比较：", getMaxCusNumTCase(a1, b1))

	var c, d float64 = 5, 6
	// 由编译器推断输入的类型
	fmt.Println("自定义泛型，数字比较：", getMaxCusNumTCase(c, d))

	var e, f int64 = 7, 8
	var g, h MyInt64 = 7, 8
	fmt.Println("自定义泛型，数字比较：", getMaxCusNumTCase(e, f))
	fmt.Println("自定义泛型，数字比较：", getMaxCusNumTCase(g, h))
}

func getMaxCusNumTCase[T CusNum](a, b T) T {
	if a > b {
		return a
	}
	return b
}

/* 内置泛型类型 */
func BuiltInCase() {
	var a, b string = "abc", "efg"
	fmt.Println("内置 comparable 泛型类型约束", getBuiltInComparable(a, b))
	var c, d float64 = 100, 100
	fmt.Println("内置 comparable 泛型类型约束", getBuiltInComparable(c, d))

	var f = 100.123
	printBuiltInAny(f)
	printBuiltInAny(a)
}

func getBuiltInComparable[T comparable](a, b T) bool {
	// comparable 类型，只支持 ==  != 两个操作
	if a == b {
		return true
	}
	return false
}

func printBuiltInAny[T any](a T) {
	fmt.Println("内置 any 泛型类型约束", a)
}
