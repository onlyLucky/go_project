package _case

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
defer关键词用来声明一个延迟调用函数
该函数可以是匿名函数也可以是具名函数
多个延时函数执行顺序为后进先出
*/
func DeferCase1() {
	fmt.Println("开始执行DeferCase1")
	defer func() {
		fmt.Println("调用匿名函数1")
	}()
	defer f1()
	defer func() {
		fmt.Println("调用匿名函数2")
	}()
	fmt.Println("DeferCase1执行结束")
}

/* 参数预计算 */
func DeferCase2() {
	i := 1
	// 传参
	defer func(j int) {
		fmt.Println("defer j: ", j)
	}(i + 1)
	// 闭包
	defer func() {
		fmt.Println("defer j: ", i)
	}()
	i = 99
	fmt.Println("i: ", i)
}

/*
返回值
defer 函数执行在return之后
*/
func DeferCase3() {
	i, j := f2()
	fmt.Printf("i: %d, j: %d, g: %d \n", i, *j, g)
}

/*
异常处理
*/
func ExceptionCase() {
	defer func() {
		// 捕获异常,恢复协程
		err := recover()
		// 异常处理
		if err != nil {
			fmt.Println("异常信息 defer recover:", err)
		}
	}()
	fmt.Println("开始执行ExceptionCase方法")
	panic("ExceptionCase方法执行结束")
	fmt.Println("ExceptionCase方法执行结束") // 这里不执行
}

/*
资源处理
*/
func FileReadCase() {
	file, err := os.Open("README.md") // ../README.md 上一级文件夹
	if err != nil {
		log.Fatal(err)
	}
	// 通过defer调用资源释放的方法
	defer func() {
		file.Close()
		fmt.Println("资源释放结束")
	}()
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		fmt.Println(buf[:n])
	}
}

func f1() {
	fmt.Println("调用具名函数1")
}

var g = 100

func f2() (int, *int) {
	defer func() {
		g = 200
	}()
	fmt.Println("f2 g:", g)
	return g, &g
}
