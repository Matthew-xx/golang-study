package main

import (
	"fmt"
)

func main() {
	//创建一个双向channel
	ch := make(chan int)
	// 双向能隐式转换为单向
	var writeCh chan<- int
	var readCh <-chan int
	writeCh = ch //只能接收数据,不能发送，即 writeCh<-666
	readCh = ch  //只能发送数据,不能接收，即 <-writeCh

	// 双向能隐式转换为单向
	var writeCh chan<- int //单向不能转为双向
}
