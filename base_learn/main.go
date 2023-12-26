package main

import (
	// _ "baseLearn/interface_base"
	// _ "baseLearn/generics"
	// _ "baseLearn/reflection"
	// _ "baseLearn/files"
	"baseLearn/client"
	"baseLearn/server"
)


func main() {
	server.TcpServer()
	client.TcpClient()
}