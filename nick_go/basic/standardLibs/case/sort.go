package _case

import (
	"fmt"
	"sort"
)

type sortUser struct {
	ID   int64
	Name string
	Age  uint8
}

type ByID []sortUser

// 获取长度
func (a ByID) Len() int { return len(a) }

// 交换位置
func (a ByID) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// 返回i的ID是否小于j位置的ID
func (a ByID) Less(i, j int) bool {
	return a[i].ID < a[j].ID
}

func SortCase() {
	list := []sortUser{
		{ID: 1, Name: "nick", Age: 12},
		{ID: 5, Name: "nick", Age: 18},
		{ID: 2, Name: "nick", Age: 23},
		{ID: 3, Name: "nick", Age: 15},
		{ID: 8, Name: "nick", Age: 22},
		{ID: 12, Name: "nick", Age: 20},
		{ID: 7, Name: "nick", Age: 24},
		{ID: 9, Name: "nick", Age: 25},
		{ID: 25, Name: "nick", Age: 10},
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Age < list[j].Age
	})
	fmt.Println(list)

	// 实现sort interface 自定义排序
	sort.Sort(ByID(list))
	fmt.Println(list)
}
