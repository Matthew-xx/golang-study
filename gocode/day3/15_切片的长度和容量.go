package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[2:3:5]
	//var b []int
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
	fmt.Println(s)
	// len: 1
	// cap: 3
	// [3] 注意，所有数字的统计以数组的起点为初始点
}
