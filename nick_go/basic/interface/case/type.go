package _case

import "fmt"

type user struct {
	ID   int64
	Name string
	Age  uint8
}

type address struct {
	ID       int
	Province string
	City     string
}

// 集合转列表
func mapToList[K comparable, T any](mp map[K]T) []T {
	list := make([]T, len(mp))
	var i int
	for _, data := range mp {
		list[i] = data
		i++
	}
	return list
}

func myPrintln[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}

func TTypeCase() {
	userMp := make(map[int64]user, 0)
	userMp[1] = user{ID: 1, Name: "nick", Age: 18}
	userMp[2] = user{ID: 2, Name: "king", Age: 19}
	userList := mapToList[int64, user](userMp)

	ch := make(chan user)
	go myPrintln(ch)
	for _, u := range userList {
		// 向管道中写入
		ch <- u
	}

	addrMp := make(map[int]address, 0)
	addrMp[1] = address{ID: 1, Province: "浙江", City: "杭州"}
	addrMp[2] = address{ID: 2, Province: "湖南", City: "长沙"}
	addrList := mapToList[int, address](addrMp)

	addrCh := make(chan address)
	go myPrintln(addrCh)
	for _, uAddr := range addrList {
		// 向管道中写入
		addrCh <- uAddr
	}
}

// 泛型切片的定义
type List[T any] []T

// 泛型集合的定义
// 声明两个泛型，分别为K,V
type MapT[K comparable, V any] map[K]V

// 泛型通道的定义
type Chan[T any] chan T

func TTypeCase1() {
	userMp := make(MapT[int64, user], 0)
	userMp[1] = user{ID: 1, Name: "nick", Age: 18}
	userMp[2] = user{ID: 2, Name: "king", Age: 19}
	var userList List[user]
	userList = mapToList[int64, user](userMp)

	ch := make(Chan[user])
	go myPrintln(ch)
	for _, u := range userList {
		// 向管道中写入
		ch <- u
	}

	addrMp := make(MapT[int, address], 0)
	addrMp[1] = address{ID: 1, Province: "浙江", City: "杭州"}
	addrMp[2] = address{ID: 2, Province: "湖南", City: "长沙"}
	var addrList List[address]
	addrList = mapToList[int, address](addrMp)

	addrCh := make(Chan[address])
	go myPrintln(addrCh)
	for _, uAddr := range addrList {
		// 向管道中写入
		addrCh <- uAddr
	}
}
