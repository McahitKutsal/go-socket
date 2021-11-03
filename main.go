package main

import (
	"fmt"
	"main/socket"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:8098")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server initialized at :8098")
	socket.Start(server)
}
