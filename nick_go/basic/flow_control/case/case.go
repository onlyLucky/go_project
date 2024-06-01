package _case

import "fmt"

func FlowControlCase() {
	ifElseCase(0)
	ifElseCase(1)
	ifElseCase(2)
	forCase()
	switchCase("A", 1)
	switchCase("C", "")
}

func ifElseCase(a int) {
	if a == 0 {
		fmt.Println("执行if 语句块")
	} else if a == 1 {
		fmt.Println("执行 else if 语句块")
	} else {
		fmt.Println("执行else 语句块")
	}
}

func forCase() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			// 跳过本地循环
			continue
		}
		fmt.Println("情况1 执行 for 语句块 i:", i)
	}

	var i int
	var b = true
	for b { // 死循环 更改条件后结束
		fmt.Println("情况2 执行 for 语句块 i:", i)
		i++
		if i >= 10 {
			b = false
		}
	}

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index, data := range list {
		fmt.Printf("情况3 执行 for 语句块 便利切片 index: %d, data: %d \n", index, data)
	}

	mp := map[string]string{"A": "a", "B": "b", "C": "c", "D": "d"}
	for key, value := range mp {
		fmt.Printf("情况4 执行 for 语句块 便利集合 key: %v, value: %v \n", key, value)
	}

	str := "今天天气很好"
	for index, char := range str {
		fmt.Printf("情况5 执行 for 语句块 便利字符串 index: %v, char: %v \n", index, char) // index: 0, char: 20170
	}

	var j int
	for {
		fmt.Println("情况6 执行 for 语句块 j:", j)
		j++
		if j >= 10 {
			break
		}
	}

	// 嵌套循环
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("情况7 执行 for 语句块 i: %d, j: %d \n", i, j)
		}
	}

	// 跳出循环
lab1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("情况8 执行 for 语句块 i: %d, j: %d \n", i, j)
			break lab1
		}
	}
}

func switchCase(alpha string, in interface{}) {
	switch alpha {
	case "A":
		fmt.Println("执行 A 语句块")
	case "B":
		fmt.Println("执行 B 语句块")
	case "C", "D":
		fmt.Println("执行 C D 语句块")
		// 命中 C D  会直接执行后面语句
		fallthrough
	case "E":
		fmt.Println("执行 E 语句块")
	case "F":
		fmt.Println("执行 F 语句块")
	}

	switch in.(type) {
	case string:
		fmt.Println("in 输入为字符串")
	case int:
		fmt.Println("in 输入为int类型")
	default:
		fmt.Println("in 输入为其他类型")
	}
}

// goto容易造成死循环
func gotoCase() {
	var a = 0

lab1:
	fmt.Println("goto 位置1")
	for i := 0; i < 10; i++ {
		a += i
		if a == 0 {
			a += 1
			goto lab1
		}
	}
}
