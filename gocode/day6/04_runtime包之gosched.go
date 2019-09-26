package main

import (
	"fmt"
	"runtime"
	//"time"
)

func main() {
	go func() {
		runtime.Gosched()
		for i := 0; i < 5; i++ {
			fmt.Println("this is goroutine 1")
		}
	}()
	go func() {
		runtime.Gosched()
		for i := 0; i < 5; i++ {
			fmt.Println("this is goroutine 2")
		}
	}()
	for i := 0; i < 5; i++ {
		// runtime.Gosched()
		fmt.Println("this is main  goroutine")
		//time.Sleep(time.Second)通过减慢主协程，发现可以执行子协程
		//无减慢过程则，子协程无法执行
	}
	runtime.Gosched()
	for i := 0; i < 5; i++ {
		// runtime.Gosched()
		fmt.Println("this is also main  goroutine")
		//time.Sleep(time.Second)通过减慢主协程，发现可以执行子协程
		//无减慢过程则，子协程无法执行
	}
	for {
	}
}
