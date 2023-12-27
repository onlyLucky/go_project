package server

import (
	"fmt"
	"io"
	"net"
)

func init() {
	fmt.Println("server TCP: =======")
	// 创建tcp 监听地址
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	// tcp 监听
	listen,_ := net.ListenTCP("tcp", tcpAddr)
	for {
		// 等待连接
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		// 获取客户端的地址
		fmt.Println(conn.RemoteAddr().String() + " 进来了")
		// 读取客户端传来的数据
		for {
			var buf []byte = make([]byte, 1024)
			n, err := conn.Read(buf)
			// 客户端退出
			if err == io.EOF {
				fmt.Println(conn.RemoteAddr().String() + " 出去了")
				break
			}
			fmt.Println(string(buf[0:n]))
		}
	}
}