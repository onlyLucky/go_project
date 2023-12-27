// Author: fg
// Date: 2023-12-02 14:51:14
// LastEditors: fg
// LastEditTime: 2023-12-02 14:51:14
// Description: content
package dataType

import "fmt"

func init() {
	fmt.Println("hello array!!! array init")
	var list = [...]string{"a", "b", "c"}
  slices := list[:] // 左一刀，右一刀  变成了切片
  fmt.Println(slices)
  fmt.Println(list[1:2]) // b

	var arr = [5]int{3,4,5,6,7}
	var slice = arr[:3]
	slice[0] = 300
	arr[1] = 400
	fmt.Println(slice,arr) // 引用地址类型

	var s []int
	fmt.Println("s == nil",s == nil)
}	

func MakeFun(){
	fmt.Println("MakeFun")
}