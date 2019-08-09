package main

import (
	"fmt"
	"runtime"
	//"runtime"
	//"time"
)

func test() {
	defer fmt.Println("cccc")
	//return //终止此函数，ddd无法打印，但bbb会继续打印
	runtime.Goexit() //终止所在协程，ddd\bbb均无法打印
	fmt.Println("ddddddd")
}
func main() {
	go func() {
		fmt.Println("aaaa")
		test()
		fmt.Println("bbbb")
	}()
	//不让主协程结束
	for {

	}
}
