package channel_base

import (
	"fmt"
	"time"
)

func sing() {
	fmt.Println("sing~~~")
	time.Sleep(1*time.Second)
	fmt.Println("sing end~")
}

func init() {
	go sing()
	go sing()
	go sing()
	go sing()
	time.Sleep(2*time.Second)
}