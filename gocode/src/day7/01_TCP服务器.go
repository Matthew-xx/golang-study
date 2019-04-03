package main

import (
	"fmt"
	"net"
)

func main() {
	//监听listen
	//listener, _ := net.Listen(tcp, "127.0.0.1:2343")可省略本地ip
	listener, err := net.Listen("tcp", ":2343")
	if err != nil {
		fmt.Printf("err=", err)
		return //出错就结束
	}
	defer listener.Close() //关监听
	//等待用户接入
	con, err := listener.Accept()
	if err != nil {
		fmt.Printf("err=", err)
		return
	}
	defer con.Close() //关闭用户链接
	//接收用户请求
	buf := make([]byte, 1024*4)
	n, err1 := con.Read(buf)
	if err1 != nil {
		fmt.Printf("err1=", err1)
		return
	}
	fmt.Println(string(buf[:n]))

}
