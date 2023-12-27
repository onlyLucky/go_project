package files

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 常规写
func writeFile() {
	file, err := os.OpenFile("./base/files/write.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR,0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteSlice := []byte("hello world! writeFile")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", bytesWritten)
}

// 快速写
func writeFast(){
	err := ioutil.WriteFile("./base/files/write.txt", []byte("add a new line writeFast"), 0644)
	if err != nil {
		panic(err)
	}
}

// 缓冲写
func writeBuffer(){
	file, err := os.OpenFile("./base/files/write.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR,0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	msg := "Hello World!\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.Write([]byte(msg))
	}
	writer.Flush()
}

// 复制文件
func writeCopy(){
	read, _ := os.Open("./base/files/write.txt")
	write, _ := os.Create("./base/files/write_copy.txt")
	n, err := io.Copy(write,read)
	fmt.Println(n, err)
}

// 目录操作
func readDir(path string){
	dir, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range dir {
		name := entry.Name()
		nPath := fmt.Sprintf("%s/%s", path, name)
		if entry.IsDir(){
			readDir(nPath)
			continue
		}
		fmt.Println(nPath)
	}
}

func init() {
	fmt.Println("file write : ========")
	writeFile()
	writeFast()
	writeBuffer()
	writeCopy()
	readDir("./base")
}