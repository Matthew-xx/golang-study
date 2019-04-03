package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[2:3:5]
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
}
