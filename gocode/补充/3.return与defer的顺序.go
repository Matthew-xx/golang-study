package main

import (
	"fmt"
)

// 多个defer的执行顺序为“后进先出”；
// defer、return、返回值三者的执行逻辑应该是：
// return最先执行，return负责将结果写入返回值中；
// 接着defer开始执行一些收尾工作；
// 最后函数携带当前返回值退出。

// 个人理解
// 可以将return理解为一个非原子操作。需要先获取返回值的value，再携带value退出函数
// 当函数结束后，defer执行前，return就已经拿到了返回值的副本，defer如果对返回值所在的变量修改，只会被defer自己保留下来
// 除非defer直接修改返回值。这种情况主要是提前在函数结构声明了返回值变量。
// 推论，类似函数参数的值传递，返回值也有这个过程。如果直接在函数体中，声明返回值变量，程序效率更高，减少了传递的过程。
func main01() {
	fmt.Println("return:", a()) // 打印结果为 return: 0
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}

// 打印结果：
// defer1: 1
// defer2: 2
// return: 0

func main02() {
	fmt.Println("return:", b()) // 打印结果为 return: 2
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i // 或者直接 return 效果相同
}

// 打印结果：
// defer1: 1
// defer2: 2
// return: 2

func main() {
	fmt.Println("c return:", *(c())) // 打印结果为 c return: 2
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("c defer2:", i) // 打印结果为 c defer: 2
	}()
	defer func() {
		i++
		fmt.Println("c defer1:", i) // 打印结果为 c defer: 1
	}()
	return &i
}

// 打印结果：
// defer1: 1
// defer2: 2
// return: 2
