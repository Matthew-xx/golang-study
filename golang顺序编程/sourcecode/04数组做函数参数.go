package main

import (
	"fmt"
)

func test(a [5]int) {
	a[2] = 32
	fmt.Println("test:", a)
}
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	test(a)
	fmt.Println("main:", a)
}

//数组做函数参数时值传递
