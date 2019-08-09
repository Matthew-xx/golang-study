package main

import "fmt"

func main() {
	var p *int
	fmt.Println(p)
	//*p = 3343,没有内存空间就操作，导致报错
}
