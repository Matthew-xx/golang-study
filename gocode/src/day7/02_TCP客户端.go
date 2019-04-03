package main

import (
	"fmt"
	"net"
)

func main() {
	//主动接入服务器
	con, err := net.Dial("tcp", ":2343") //dial函数连接
	if err != nil {
		fmt.Println("err=", err)
	}
	defer con.Close()

	//发送数据
	buf := []byte("are u ok?")
	con.Write(buf)
}
