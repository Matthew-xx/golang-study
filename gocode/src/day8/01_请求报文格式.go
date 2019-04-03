package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("err1=", err1)
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
}
