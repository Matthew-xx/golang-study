package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//主动接入服务器
	con, err := net.Dial("tcp", ":2343") //dial函数连接
	if err != nil {
		fmt.Println("err=", err)
	}
	defer con.Close()
	//创建新协程来发送内容
	go func() {

		buf := make([]byte, 1024*4)
		//从键盘给服务器发送内容
		for {
			//循环输入
			//fmt.Scan(buf)
			n, err1 := os.Stdin.Read(buf)
			fmt.Println("我说：", string(buf[:n-1]))
			if err1 != nil {
				fmt.Println("err1:", err1)
			}
			//发送给服务器
			_, err2 := con.Write(buf[:n-1])
			if err2 != nil {
				fmt.Println("err2:", err2)
			}
		}
	}()
	//主协程从服务器接收返回来的数据,不断接收
	buf := make([]byte, 1024*4)
	for {
		n, err := con.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("您已退出")
				return
			} else {
				fmt.Println("err:", err)
				return
			}
		}
		fmt.Println("服务器回复我：", string(buf[:n]))

	}
	//关闭流程为：成功发送exit，服务端接收确认后，关闭发送，客户端接收到eof,客户端再关闭

}
