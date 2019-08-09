// 07_数组的基本操作
package main

import (
	"fmt"
)

func main() {
	//定义数组，[10]int,[20]int 是两个类型数组，长度必须指明为常量。不可修改
	var a, b [10]int
	//var a [10]int
	var c [30]int
	for m, n := range a {
		fmt.Printf("m=%v,n=%v\n", m, n)
	}
	fmt.Println()
	for m, n := range b {
		fmt.Printf("m=%v,n=%v\n", m, n)
	}
	for m, n := range c {
		fmt.Printf("m=%v,n=%v\n", m, n)
	}
}
