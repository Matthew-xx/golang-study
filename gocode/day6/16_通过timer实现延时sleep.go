package main

import (
	"fmt"
	"time"
)

func main() {
	//第一种方法：生产timer对象，返回值为结构体变量，需要调用结构体内channel发送出去
	fmt.Println("开始计时")
	timer := time.NewTimer(2 * time.Second)
	// t :=
	fmt.Println("两秒时间到", <-timer.C)
	//第2种方法，直接使用sleep()
	time.Sleep(2 * time.Second)
	fmt.Println("又两秒时间到")
	//第3种方法，直接利用time.After方法，返回值本身就是channel,故只需把本身发送出去
	<-time.After(2 * time.Second)
	fmt.Println("又又两秒时间到")
}
