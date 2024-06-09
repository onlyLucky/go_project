package _case

import "fmt"

// 泛型结构体
type MyStruct[T interface{ *int | *string }] struct {
	Name string
	Data T
}

// 泛型 receiver 不支持泛型方法，只能通过receiver来实现方法的泛型处理
// func (myStruct MyStruct[T]) GetData[s int](t T) T  s类型声明报错
/* func (myStruct MyStruct[T]) GetData(t T) T {
	var i interface{} = 20
	a,ok := i.(int)
	b,ok := t.(int)  // 泛型传参不支持断言
	return myStruct.Data
} */
func (myStruct MyStruct[T]) GetData() T {
	return myStruct.Data
}

func ReceiverCase() {
	data := 18
	myStruct := MyStruct[*int]{
		Name: "nick",
		Data: &data,
	}
	data1 := myStruct.GetData()
	fmt.Println(*data1)

	str := "abcdefg"
	myStruct1 := MyStruct[*string]{
		Name: "nick",
		Data: &str,
	}
	str1 := myStruct1.GetData()
	fmt.Println(*str1)
}
