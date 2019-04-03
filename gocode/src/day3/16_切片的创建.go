package main

import (
	"fmt"
)

func main() {
	//a := [5]int{1, 2, 3, 4, 5}
	//s := make([]int, 3, 10)  分别时len和cap
	s := make([]int, 3) //不指定cap,则cap=len
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
}
