package generics

import "fmt"

// 遍历init 切片
func PrintIntSlice(slice []int){
	for _, v := range slice {
		fmt.Printf("PrintIntSlice: %T %v\n", v, v)
	}
}

// 遍历int64切片
func PrintInt64Slice(slice []int64){
	for _, v := range slice {
		fmt.Printf("PrintInt64Slice: %T %v\n", v, v)
	}
}

// int64 切片 转 int 切片
func Int64SliceToIntSlice(Int64Slice []int64) (IntSlice []int){
	for _, v := range Int64Slice {
		IntSlice = append(IntSlice, int(v))
	}
	return
}

// 泛型函数
func PrintSliceTypeSlice[T int | int64 | string](slice []T) {
  fmt.Printf("%T\n", slice)
  for _, v := range slice {
    fmt.Printf("%T  %v\n", v, v)
  }
}

func init() {
	fmt.Println("Generics func init: ===========")

	PrintIntSlice([]int{1, 2, 3, 4, 5})

	var int64Slice []int64 = []int64{4, 5, 7}
  PrintInt64Slice(int64Slice)

	var intSlice []int
  for _, v := range int64Slice {
    intSlice = append(intSlice, int(v))
  }
  PrintIntSlice(intSlice)

  PrintIntSlice(Int64SliceToIntSlice(int64Slice))

	fmt.Println("Generics type")
	PrintSliceTypeSlice([]int{1, 2, 3, 4, 5})
  PrintSliceTypeSlice([]int64{1, 2, 3, 4, 5})
  PrintSliceTypeSlice([]string{"hello"})
  
  // 标准写法
  PrintSliceTypeSlice[int]([]int{1, 2, 3, 4, 5})
  PrintSliceTypeSlice[int64]([]int64{1, 2, 3, 4, 5})
  PrintSliceTypeSlice[string]([]string{"hello"})
}