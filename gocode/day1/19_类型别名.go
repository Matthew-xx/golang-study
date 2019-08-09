package main

import (
	"fmt"
)

func main() {
	type long int
	var a long
	var b int
	a = 10
	b = int(a)
	fmt.Println(b)
}
