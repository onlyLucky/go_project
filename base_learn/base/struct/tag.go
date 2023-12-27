package struct_base

import (
	"encoding/json"
	"fmt"
)

var Tag = 123;

type Students struct{
	Name string `json:"name"`
	Age int `json:"age"` // `json:"age,omitempty"` 空值会被省略
}

func init() {
	fmt.Println("Struct_tag init-----")
	s:= Students{
		Name: "ff",
		Age: 18,
	}
	byteData,_ := json.Marshal(s)
	fmt.Println(string(byteData))
}