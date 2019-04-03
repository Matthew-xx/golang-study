package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("时间到，子协程打印")
	}()
	timer.Stop() //被停止
	for {
	}
}
