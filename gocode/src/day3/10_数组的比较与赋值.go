// 10_数组的比较与赋值
package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{8, 7, 6}

	fmt.Println("a==b?", a == b)
	fmt.Println("a==c?", a == c)
	//c := [3]int{1, 2, 3}err 不同类型的数组不能复制
	a = c
	fmt.Println(a)
}
