package main

import (
	"fmt"
	"net"
)

//设置一个结构体变量client，用来生成一个用户对象

type Client struct {
	C    chan string //包含一个string类型的管道字段，用来给用户子线程通信
	name string
	addr string
}

//设置一个全局map，用来保存当前在线用户列表
//其中 键为id 值为用户对象
//暂不初始化，用多少初始多少
var OnlinMap map[string]Client

//设置一个全局变量 channel，用来给各用户协程之间直接通信
//直接初始化，因为就这一个
var message chan string = make(chan string)

func main() {
	//设置监听器
	listener, _ := net.Listen("tcp", ":3422")
	defer listener.Close()

	//管理各用户协程，只需启用一个，不需要时因通信会自动阻塞
	go Mannage()

	//不断接受用户的链接，每接受一个，单独起一个用户协程为其服务
	for {
		//链接一个运行一次这个循环??????????
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err=", err)
		}
		// defer conn.Close()最好不要在循环中defer

		//每接受一个，单独起一个协程为其服务
		go HandleConn(conn)
	}

}

//处理链接
func HandleConn(conn net.Conn) {

	//主协程把链接给了子协程，最好让子协程结束时关闭链接
	defer conn.Close()

	//获取该用户地址
	addr := conn.RemoteAddr().String()

	//把该用户存入，map
	cli := Client{make(chan string), addr, addr}
	OnlinMap[addr] = cli
	//给公用管道发送数据，
	message <- addr + ": login"

	//为本用户创建一个转发协程，专门为本用户转发消息
	go WritetoClient(conn, cli)
	//处理链接期间不能断开链接需保持本协程运行

	for {

	}
}

//管理用户协程之间通信，依赖公共管道
func Mannage() {
	//初始化一个map
	OnlinMap = make(map[string]Client)

	//每当公用管道有数据时，故需要循环遍历等待数据，而不是一次性执行
	//for {
	for msg := range message {
		//msg := <-message
		//给所有用户发送这条数据,需range
		for _, cli := range OnlinMap {
			cli.C <- msg
		}
	}

}

//用户协程与转发协程间通信，依赖每个用户对象自己的管道
func WritetoClient(conn net.Conn, cli Client) {
	//用户结构体内的管道有数据时，需使用range,确保每当用户需要转发时，都能处理，而非处理一次
	//msg := <-cli.C

	//此管道在接收用户协程的消息，用户协程可能多次向此管道发送数据
	for msg := range cli.C {
		//转发给用户对应的客户端
		conn.Write([]byte(msg + "\n"))
	}

}
