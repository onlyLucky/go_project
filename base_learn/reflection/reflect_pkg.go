package reflection

import (
	"fmt"
	"reflect"
)

type Users struct {
	Name string `json:"name" feng:"name_xxx"`
	Age  int    `json:"age" feng:"age"`
}

func FPrint(inter interface{}) {
	t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	// fmt.Println(t.Kind()) // 获取这个接口的底层类型
  // fmt.Println(t.Elem()) // 变量的原始类型

	for i:=0; i<t.NumField();i++{
		//fmt.Println()
    // 字段的类型
    // 字段名
    // 字段的值
    // 字段的tag
    fmt.Println(
      t.Field(i).Type,
      t.Field(i).Name,
      v.Field(i),
      t.Field(i).Tag.Get("feng"),
    )
	}
}

func reviseFPrint(inter interface{}){
	v := reflect.ValueOf(inter)
  e := v.Elem()  // 必须用这个
  e.FieldByName("Name").SetString("枫枫知道")
}

func init() {
	fmt.Println("reflect_pkg: ======")
	user := Users{"枫枫", 21}
  FPrint(user)
	fmt.Println("reflect_revise: ======")
	reviseFPrint(&user)  // 必须传指针
  fmt.Println(user)
}