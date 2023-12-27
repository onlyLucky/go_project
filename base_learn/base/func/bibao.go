package base_func

import (
	"fmt"
	"time"
)

func awaitAdd(t int) func(...int) int {
	time.Sleep(time.Duration(t) * time.Second)
	return func(numList ...int) int {
		var sum int
		for _,i2 := range numList{
			sum += i2
		}
		return sum
	}
}

func init() {
	fmt.Println(awaitAdd(2)(1,2,3))
}