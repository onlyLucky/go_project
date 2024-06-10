package _case

import (
	"errors"
	"fmt"
	"reflect"
)

func ReflectCase() {
	type user struct {
		ID    int64
		Name  string
		Hobby []string
	}
	type outUser struct {
		ID    int64
		Name  string
		Hobby []string
	}

	u := user{ID: 1, Name: "nick", Hobby: []string{"唱", "跳", "rap", "篮球"}}
	out := outUser{}

	res := copy(&out, u)
	fmt.Println(res, out)

	listUser := []user{
		{ID: 1, Name: "nick", Hobby: []string{"唱", "跳", "rap", "篮球"}},
		{ID: 2, Name: "nick1", Hobby: []string{"唱1", "跳1", "rap1", "篮球1"}},
		{ID: 3, Name: "nick2", Hobby: []string{"唱2", "跳2", "rap2", "篮球2"}},
	}
	list := sliceColumn(listUser, "Hobby")
	fmt.Println(list)

}

func copy(dest interface{}, source interface{}) error {
	sType := reflect.TypeOf(source)
	sValue := reflect.ValueOf(source)
	// 如果为指针类型，获取他的值
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
		sValue = sValue.Elem()
	}
	dType := reflect.TypeOf(dest)
	dValue := reflect.ValueOf(dest)
	// 如果赋值对象不为指针
	if dType.Kind() != reflect.Ptr {
		return errors.New("目标对象必须为struct指针类型")
	}
	dType = dType.Elem()
	dValue = dValue.Elem()
	if sValue.Kind() != reflect.Struct {
		return errors.New("复制的源对象必须为struct或struct的指针")
	}
	if dValue.Kind() != reflect.Struct {
		return errors.New("目标对象必须为struct指针类型")
	}
	destObj := reflect.New(dType)
	for i := 0; i < dType.NumField(); i++ {
		destField := dType.Field(i)
		if sourceField, ok := sType.FieldByName(destField.Name); ok {
			if destField.Type != sourceField.Type {
				continue
			}
			value := sValue.FieldByName(destField.Name)
			destObj.Elem().FieldByName(destField.Name).Set(value)
		}
	}
	dValue.Set(destObj.Elem())
	return nil
}

func sliceColumn(slice interface{}, column string) interface{} {
	t := reflect.TypeOf(slice)
	v := reflect.ValueOf(slice)
	// 如果是指针
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	// 如果是结构体
	if t.Kind() == reflect.Struct {
		val := v.FieldByName(column)
		return val.Interface()
	}
	// 如果不是切片
	if t.Kind() != reflect.Slice {
		return nil
	}

	t = t.Elem()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	f, _ := t.FieldByName(column)
	sliceType := reflect.SliceOf(f.Type)
	s := reflect.MakeSlice(sliceType, 0, 0)

	for i := 0; i < v.Len(); i++ {
		o := v.Index(i)
		if o.Kind() == reflect.Struct {
			val := o.FieldByName(column)
			s = reflect.Append(s, val)
		}
		if o.Kind() == reflect.Ptr {
			v1 := o.Elem()
			val := v1.FieldByName(column)
			s = reflect.Append(s, val)
		}
	}

	return s.Interface()

}
