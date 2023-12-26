package generics

import "fmt"

type myMap[K string | int, V any] map[K]V
type _User struct {
	Name string
}

func init() {
	fmt.Println("generics map: ================")
	m1 := myMap[string, string]{
		"key": "fengfeng",
		"value": "fValue",
	}
	fmt.Println(m1)

	m2 := myMap[int, _User]{
    0: _User{Name: "枫枫"},
  }
  fmt.Println(m2)
} 