package main

import (
	"fmt"
)

func main() {
	srcslice := []int{6, 6}
	dstslice := []int{1, 2, 3, 4}
	copy(dstslice, srcslice) //对应位置复制
	fmt.Println(dstslice)    //6，6，3，4
}
