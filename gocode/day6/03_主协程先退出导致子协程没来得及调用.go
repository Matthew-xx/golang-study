package main

import (
	"fmt"
	"time"
)

func main() {
	//主函数执行 过程中，遇到有go修饰方法，会把这个方法扔出去，去独立执行，而main继续执行自己后面的语句
	//可以想象 go修饰的函数都长大了，该自己去照顾自己了。main()就不在你身上费心思了。
	go func() {
		for {
			fmt.Println("this is son goroutine")
			time.Sleep(time.Second)
		}

	}() //定义并执行匿名函数

}
