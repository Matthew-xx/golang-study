package main

import (
	"fmt"
)

func myfunc(a int, b ...int) int {
	for _, d := range b {
		a = a + d
	}
	return a
}

func main() {
	fmt.Println(myfunc(1, 2, 3, 4, 9, 5))
}
