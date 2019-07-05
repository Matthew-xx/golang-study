package main

import (
	"fmt"
	"time"
)

//ch:=make(chan int) //全局变量不要自动推导类型
var ch = make(chan int) //为啥不能用ch:=make(chan int)?
func Printer(str string) {

	for _, s := range str {
		fmt.Print(string(s))
		time.Sleep(time.Second / 7)
	}
	fmt.Print("\n")
}
func Person1() {
	Printer("hello world")
	//给管道写数据，执行完毕动作后，写入数据，另一个任务才能解除阻塞
	ch <- 666
}
func Person2() {
	//直接从管道读取数据，显然一开始是没有数据的，阻塞状态
	<-ch
	Printer("my name is mike")
}

func main() {

	go Person1()
	go Person2()
	//多任务执行时，资源被竞争
	for {
	}
}
