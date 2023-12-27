package my_error

import "os"

func init() {
	// 读取配置文件中
	_, err := os.ReadFile("xxx")
	if err != nil {
		panic(err.Error())
	}
}