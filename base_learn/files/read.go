package files

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readOnce() {
	bs, err := os.ReadFile("./files/abc.txt")
	if err != nil {
		fmt.Printf("Read file error: %v\n", err)
		return
	}
	fmt.Printf("%s\n",bs)
}

func readSlice(){
	file, err := os.Open("./files/abc.txt")
	if err != nil {
		panic(err)                                                            
	}
	defer file.Close()
	for {
		buf := make([]byte, 32)
		_,err = file.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", buf)
	}
}

func readBuffer(){
	file, err := os.Open("./files/abc.txt")
	if err != nil {
		fmt.Printf("Open file error: %v\n",err) 
		return                                                         
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// 按行读取
		line,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
}

func readCursor(){
	file, _ := os.Open("./files/abc.txt")
	defer file.Close()
	// 开始位置前进5个字节
	var whence = 0
	var offset int64 = 5
	pos, _ := file.Seek(offset, whence)
	fmt.Println("Jump forward 5 bytes from start position:", pos)

	// 当前位置回退2个字节
	whence = 1
	offset = -2
	pos, _ = file.Seek(offset, whence)
	fmt.Println("Jump back 2 bytes from current position:", pos)
}


func init() {
	readOnce()
	readSlice()
	readBuffer()
	readCursor()
}