package main

import (
	"fmt"
	"redisLearn/connect"
)

func main() {
	fmt.Println("hello redis!!!")
	connect.ConnectFunc()
}