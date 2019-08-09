package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func ReceiveFile(name string, conn net.Conn) {
	f, err := os.Create(name + "2")
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("创建文件成功")
	//边读边写
	buf := make([]byte, 1024)
	var i int
	for {

		n, err1 := conn.Read(buf)
		fmt.Println("发送进度", i, "长度：", len(buf[:n]))
		fmt.Println(len(buf[:n]))
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("接收完毕")
				return
			} else {
				fmt.Println("err1 reci=", err1)
				return
			}

		}

		_, err2 := f.Write(buf[:n])
		if err2 != nil {
			fmt.Println("err2=", err2)
		} else {
			fmt.Println("文件接收已接收完成")
		}
		i++
	}
}
func main() {
	//设置监听器
	lisener, err := net.Listen("tcp", ":2342") //tcp必须小写
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer lisener.Close()
	//监听客户端连接
	conn, err1 := lisener.Accept()
	if err1 != nil {
		fmt.Println("err1=", err1)
		return
	}
	fmt.Println("检测到用户要发送文件")
	defer conn.Close()
	//读取客户端发来信息
	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("err2=", err2)
		return
	}
	//告诉客户端接收成功，可发送文件
	_, err3 := conn.Write([]byte("ok"))
	if err3 != nil {
		fmt.Println("err3=", err3)
		return
	}
	fmt.Println("收到文件名，请开始发送内容")
	ReceiveFile(string(buf[:n]), conn)
}
