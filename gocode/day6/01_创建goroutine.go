package main

import (
	"fmt"
	"time"
)

func test() {
	for {
		fmt.Println("this is test goroutine")
		time.Sleep(time.Second)
	}
}
func main() {
	go test() //由于后面是for死循环，放前面才会有机会执行
	for {
		fmt.Println("this is main goroutine")
		time.Sleep(time.Second)
	}
	//go test()放在后面，没有机会执行到这，协程启动不起来
}
