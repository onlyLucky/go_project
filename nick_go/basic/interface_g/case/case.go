package _case

func InterfaceCase() {

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
	// 多行之间取交集
	uint8 | int32 | float64 | ~int64
	int32 | float64 | ~int64 | uint16
}

// MyInt64为int64的衍生类型 并具有基础类型int64的新类型，与int64不相等
type MyInt64 int64

// MyInt32为int32的别名，与int32是同一种类型
type MyInt32 = int32

func getMax
func CusNumCase[T CusNum](a, b T) T {
	if a > b {
		return a
	}
	return b
}
