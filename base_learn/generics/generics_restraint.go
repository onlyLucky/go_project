package generics

import (
	"fmt"
	"strconv"
)

type NumStr interface {
	Num | Str
}

// ~的意思就是底层数据类型
type Num interface {
	~int | ~int32 | ~int64 | ~uint8
}

type Str interface {
	string
}

type Status uint8

type mySlice1[T NumStr] []T

// 约束方法
type Price int

func (p Price) String() string {
	// int 转 数字
	return strconv.Itoa(int(p))
}

type Price2 string

func (p Price2) String() string {
	// int 转数字
	return string(p)
}

type showPrice interface {
	~int | ~string
	String() string
}

func showPriceFunc[T showPrice](p T){
	fmt.Println(p.String())
}

func init() {
	fmt.Println("Generics Restraint: ==========")
	m1 := mySlice1[int]{1, 2, 3}
	fmt.Println(m1)
	m2 := mySlice1[int64]{1,2,3}
	fmt.Println(m2)
	m3 := mySlice1[string]{"hello"}
  fmt.Println(m3)
  m4 := mySlice1[Status]{1, 2, 3}
  fmt.Println(m4)

	var p1 Price = 12
	showPriceFunc(p1)
	var p2 Price2 = "56"
	showPriceFunc(p2)
}