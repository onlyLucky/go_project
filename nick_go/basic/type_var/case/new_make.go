package _case

import "fmt"

func NewCase() {
	//	通过new函数，可以创建任意类型，并返回指针
	// new 所做的使用是在内存中开辟一个地址空间，里面放的空的默认值， 所以他的地址不为nil，值是为空的
	mpPtr := new(map[string]*user)
	if *mpPtr == nil {
		fmt.Println("map 为空")
	}
	//mpPtr为指针 进行属性赋值的时候，*获取值
	//(*mpPtr)["A"] = &user{}
	fmt.Println(mpPtr) // &map[]
	slicePtr := new([]user)
	if *slicePtr == nil {
		fmt.Println("切片值为空", *slicePtr)
	}
	// 切片切的是底层的数组
	// 切片长度在声明的时候已经确定了他的长度，append触发切片扩容机制
	*slicePtr = append(*slicePtr, user{Name: "nick"})

	userPtr := new(user)
	strPtr := new(string)
	fmt.Println(slicePtr, userPtr, strPtr) // &[{nick 0 { }}] &{ 0 { }} 0xc00008a090
}

// make仅用于切片、集合、通道的初始化
func MakeCase() {
	//	初始化切片，并设置长度和容量
	slice := make([]int, 10, 20)
	slice[0] = 10
	//	初始化集合，并设置集合的初始大小
	mp := make(map[string]string, 10)
	mp["A"] = "a"
	//	初始化通道，设置通道的读写方向和缓冲大小
	ch := make(chan int, 10)
	ch1 := make(chan<- int, 10) // 只写
	ch2 := make(<-chan int)     // 只读 没有缓冲区
	fmt.Println(slice, mp, ch, ch1, ch2)
	// [10 0 0 0 0 0 0 0 0 0] map[A:a] 0xc0000b6000 0xc0000b60b0 0xc000086120
}

func SliceAndMapCase() {
	//切片的定义
	var slice []int
	slice = []int{1, 2, 3, 4, 5, 6}
	slice1 := make([]int, 10)
	slice1[1] = 10
	fmt.Println(slice, slice1) // [1 2 3 4 5 6] [0 10 0 0 0 0 0 0 0 0]
	/*切片的截取*/
	slice2 := make([]int, 5, 10)
	fmt.Println(len(slice2), cap(slice2)) // 5 10
	slice2[0] = 0
	slice2[1] = 1
	slice2[2] = 2
	slice2[3] = 3
	slice2[4] = 4
	slice3 := slice2[0:5]
	fmt.Println(len(slice3), cap(slice3), slice3) // 5 10 [0 1 2 3 4]
	slice4 := slice2[2:5]
	fmt.Println(len(slice4), cap(slice4), slice4) // 3 8 [2 3 4]
	// 超过数组长度，但在容量之内
	slice5 := slice2[1:8]
	fmt.Println(len(slice5), cap(slice5), slice5) // 7 9 [1 2 3 4 0 0 0]
	//	超过切片容量大小之后会报错
	/*切片的附加*/
	slice6 := slice2[1:10]
	fmt.Println(len(slice6), cap(slice6), slice6) //9 9 [1 2 3 4 0 0 0 0 0]
	//进行附加之后，切片的容量改变了，本质上是新创建了一个切片，不再是原来的切片了。
	slice6 = append(slice6, 1, 2, 3, 4, 5)
	fmt.Println(len(slice6), cap(slice6), slice6) //14 18 [1 2 3 4 0 0 0 0 0 1 2 3 4 5]

	/*集合 无序  在其他类型的语言，集合的顺序是添加进入的顺序*/
	mp := make(map[string]string, 10)
	mp["A"] = "a"
	mp["B"] = "b"
	mp["C"] = "c"
	mp["D"] = "d"
	fmt.Println(mp) // map[A:a B:b C:c D:d]
	for k, v := range mp {
		fmt.Println(k, v)
	}
	/* 每次便利的顺序都是不一样的
	C c
	A a
	B b
	D d
	*/
}
