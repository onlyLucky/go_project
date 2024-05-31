package main

import (
	"flag"
	"fmt"
	"time"
)

var j = flag.Int("j", 0, "")

func main() {
	flag.Parse()
	fmt.Println(j)
	var i = 0
	for {
		fmt.Println("demo1 print:", i, *j)
		i++
		time.Sleep(time.Second)
	}
}
