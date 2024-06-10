package _case

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// 自定义错误
type cusError struct {
	Code string
	Msg  string
	Time time.Time
}

func (err cusError) Error() string {
	return fmt.Sprintf("Code: %s, Msg: %s, Time: %s", err.Code, err.Msg, err.Time)
}

func getCusError(code, msg string) error {
	return cusError{Code: code, Msg: msg, Time: time.Now()}
}

func ErrorCase() {
	err := errors.New("程序发生了错误！！！")
	fmt.Println(err) // 程序发生了错误！！！

	var a, b = -1, -2
	res, err := sum(a, b)
	fmt.Println(res, err)
	if err != nil {
		log.Println(err) // 0 Code: 500, Msg: 两数求和不能同时小于0, Time: 2024-06-09 22:04:50.7056973 +0800 CST m=+0.002558501
		cusError, ok := err.(cusError)
		if ok {
			fmt.Println("打印自定义错误信息：", cusError.Code, cusError.Msg, cusError.Time) // 打印自定义错误信息： 500 两数求和不能同时小于0 2024-06-09 22:04:50.7056973 +0800 CST m=+0.002558501
		}
	}
}

func sum(a, b int) (int, error) {
	if a <= 0 && b <= 0 {
		return 0, getCusError("500", "两数求和不能同时小于0")
	}
	return a + b, nil
}
