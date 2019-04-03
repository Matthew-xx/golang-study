package main

import (
	"fmt"
)

func test02() func() int {
	x := 0
	fmt.Println("初始化后的：", x)
	return func() int {
		fmt.Println("匿名函数刚执行", x)
		x++
		return x * x
	}
}
func main() {
	f := test02()
	fmt.Println("匿名函数第1次执行完：", f())
	fmt.Println("匿名函数第2次执行完：", f())
	fmt.Println("匿名函数第3次执行完：", f())
	fmt.Println("匿名函数第4次执行完：", f())
}

//证明多次调用一个包含匿名函数的母函数，母函数内的语句只执行一次。匿名函数每次都执行，而且是在同一块内存上累积执行，在上次执行结果上继续执行
