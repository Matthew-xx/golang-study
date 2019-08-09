package main

import (
	"fmt"
)

func main() {
	//有多少[]就是多少维数组,遍历是就需要多少次循环
	a := [3][4]int{{1, 2, 3, 4}, {7, 8, 9, 0}, {5, 6, 7, 8}}
	fmt.Println(a)
	//[3][4]int,表示该2维数组有3个元素，每个元素都是1个1维数组[4]int
	//总长度为3，每个元素长度为4
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			a[i][j] = j
			fmt.Printf("a[%d][%d]=%d\n", i, j, a[i][j])
		}
	}
}
