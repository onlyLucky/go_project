package main

import (
	"flag"
	"fmt"
	"time"
)

var j = flag.Int("j", 0, "")

func main() {
	flag.Parse()
	var i = 0
	for {
		fmt.Println("demo2 print:", i, *j)
		i++
		time.Sleep(time.Second)
	}
}
