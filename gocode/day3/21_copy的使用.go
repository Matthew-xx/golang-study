package main

import (
	"fmt"
)

func main() {
	slice1 := []int{6, 6}
	slice2 := []int{1, 2, 3, 4}
	copy(slice1, slice2) //对应位置复制
	fmt.Println(slice1)  //1,2
	slice1 = []int{6, 6}
	slice2 = []int{1, 2, 3, 4}
	copy(slice2, slice1) //对应位置复制
	fmt.Println(slice2)  //6，6，3，4

}
