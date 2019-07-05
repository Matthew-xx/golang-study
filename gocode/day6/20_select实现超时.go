package main

import (
	"fmt"
	"time"
)

//select
func main() {
	ch := make(chan int)
	quite := make(chan bool)

	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("我在工作")
			case <-time.After(3 * time.Second):
				fmt.Println("超时了")
				quite <- true
			}
		}
	}()

	for i := 0; i < 3; i++ {
		ch <- i
	}
	<-quite //注意是从管道接收数据
}
