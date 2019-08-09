//浪费我一个小时的bug竟然是，测试文件出了问题，代码没有问题
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(conn net.Conn, path string) {
	//只读形式打开文件
	f, err := os.Open(path)

	if err != nil {
		fmt.Println("err open =", err)
		return
	}
	fmt.Println("文件读取成功，开始发送")
	defer f.Close()
	//读多少发送多少
	buf := make([]byte, 1024)
	var i int
	for {
		n, err := f.Read(buf)
		fmt.Println("发送进度", i, "长度：", len(buf[:n]))
		if err != nil {
			if err == io.EOF {
				fmt.Println("发送完毕")
				return
			} else {
				fmt.Println("err read=", err)
				return
			}

		}
		_, err1 := conn.Write(buf[:n])
		if err1 != nil {
			fmt.Println("err1 wri=", err1)
		}
		i++
	}

}

func main() {
	//提示输入文件
	fmt.Println("请输入要发送的文件名（包含路径）：")
	var path string
	fmt.Scan(&path)
	//获取文件信息
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	//主动接入服务器
	conn, err2 := net.Dial("tcp", ":2342")
	if err2 != nil {
		fmt.Println("err2=", err2)
		return
	}
	defer conn.Close()

	//先发送文件信息
	fmt.Println("我要开始发送了")
	_, err3 := conn.Write([]byte(info.Name()))
	if err3 != nil {
		fmt.Println("err3=", err3)
		return
	}

	//接收服务器返回的信息
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if string(buf[:n]) == "ok" {
		fmt.Println("文件名发送成功，开始发送内容")
		SendFile(conn, path)
	}

}
