package main

import (
	"fmt"
)

func test(a *[5]int) {
	(*a)[2] = 32 //使用指针取单个数组元素的内容，注意加括号
	fmt.Println("test:", *a)
}
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	test(&a)
	fmt.Println("main:", a)
}

//数组指针做函数参数时传递地址，即引用传递
