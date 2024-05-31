package _case

import "fmt"

// iota 连续的 从0递增下去 枚举不想要连续的时候 可以使用_ 不接受这个变量， 不会影响后面的值
/*
const (
	B  = 1 << (10 * 0)
	KB = 1 << (10 * 1)
	MB = 1 << (10 * 2)
	GB = 1 << (10 * 3)
	TB = 1 << (10 * 4)
)
*/
const (
	B = 1 << (10 * iota)
	KB
	MB
	_
	TB
)

type Gender uint

const (
	FEMALE Gender = iota
	MALE
	THIRD
)

func ConstAndEnumCase() {
	const (
		A = 10
		B = 20
	)
	size := MB
	var gender Gender = MALE
	fmt.Println(A, B, gender, size)

}
