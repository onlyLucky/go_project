package my_error

import (
	"errors"
	"fmt"
)

func Parent() error{
	err:= method()
	return err
}

func method() error{
	return errors.New("出错了")
}

func init() {
	fmt.Println(Parent())
}