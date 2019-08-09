package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleCon(con net.Conn) { //传递参数con
	defer con.Close()
	//获取发送者地址
	addr := con.RemoteAddr().String() //需要转化为字符串
	fmt.Println(addr, "连接服务器成功！")
	buf := make([]byte, 1024*4)
	for { //客户端可以发送多次
		// buf := make([]byte, 1024*4)
		n, err1 := con.Read(buf)
		if err1 != nil {
			fmt.Printf("err1=", err1)
			return
		}
		if string(buf[:n]) == "exit" {
			fmt.Printf("用户退出")
			return
		}
		fmt.Println(addr, "说：", string(buf[:n]))
		con.Write([]byte(strings.ToUpper(string(buf[:n]))))
		//使用strings包中的大小写转换工具
	}

}
func main() {
	//监听listen
	listener, err := net.Listen("tcp", ":2343")
	if err != nil {
		fmt.Printf("err=", err)
		return //出错就结束
	}
	defer listener.Close() //关监听
	//等待用户接入

	for { //循环可接收多个用户连接
		con, err := listener.Accept()
		if err != nil {
			fmt.Printf("err=", err)
			return
		}
		//defer con.Close() //关闭用户链接
		//接收用户请求,需要开多任务，单独处理每个用户，否则会覆盖前面的用户
		go HandleCon(con)

	}

}
