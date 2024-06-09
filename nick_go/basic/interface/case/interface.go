package _case

import "fmt"

// 基本接口，可用于变量的定义
type ToString interface {
	String() string
}

func (u user) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d", u.ID, u.Name, u.Age)
}

func (u address) String() string {
	return fmt.Sprintf("ID: %d, Province: %s, City: %s", u.ID, u.Province, u.City)
}

// 泛型接口 address user 中的ID 是 int int64
type GetKey[T comparable] interface {
	any
	Get() T
}

// var s GetKey[int64]

func (u user) Get() int64 {
	return u.ID
}

func (addr address) Get() int {
	return addr.ID
}

/* list 转 map 列表转集合*/
func listToMap[K comparable, T GetKey[K]](list []T) map[K]T {
	mp := make(MapT[K, T], len(list))
	for _, data := range list {
		mp[data.Get()] = data
	}
	return mp
}

func InterfaceCase() {
	userList := []GetKey[int64]{
		user{ID: 1, Name: "nick", Age: 18},
		user{ID: 2, Name: "king", Age: 19},
	}

	addrList := []GetKey[int]{
		address{ID: 1, Province: "浙江", City: "杭州"},
		address{ID: 2, Province: "湖南", City: "长沙"},
	}

	userMp := listToMap[int64, GetKey[int64]](userList)
	fmt.Println(userMp)
	addrMp := listToMap[int, GetKey[int]](addrList)
	fmt.Println(addrMp)
}
