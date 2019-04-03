package main

import (
	"fmt"
	"time"
)

func Printer(str string) {

	for _, s := range str {
		fmt.Print(string(s))
		time.Sleep(time.Second / 7)
	}
	fmt.Print("\n")
}
func main() {
	//单任务执行时按顺序执行
	// Printer("hello world")
	// Printer("my name is mike")
	go Printer("hello world")
	Printer("my name is mike")
	//多任务执行时，资源被竞争
}
