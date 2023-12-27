package main

import (
	// _ "baseLearn/interface_base"
	// _ "baseLearn/generics"
	// _ "baseLearn/reflection"
	// _ "baseLearn/files"
	// _ "baseLearn/base"
	"baseLearn/client"
	"baseLearn/server"
)


func main() {
	server.TcpServer()
	client.TcpClient()
}