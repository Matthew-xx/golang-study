package main

import (
	"fmt"
)

func test02() func() int {
	x := 0
	fmt.Println("初始化后的自由变量值：", x)
	return func() int {
		fmt.Println("匿名函数本次执行的自由变量值", x)
		x++
		return x * x
	}
}
func main() {
	//test02的返回值是一个匿名函数，匿名函数操作一个自由变量并返回
	f := test02()//接收匿名函数
	fmt.Println("开始执行匿名函数")
	fmt.Println("第1次执行匿名函数：", f())
	fmt.Println("第2次执行匿名函数：", f())
	fmt.Println("第3次执行匿名函数：", f())
	fmt.Println("第4次执行匿名函数：", f())
}

//证明多次调用一个包含匿名函数的母函数，母函数内的语句只执行一次。匿名函数每次都执行，而且是在同一块内存上累积执行，在上次执行结果上继续执行
