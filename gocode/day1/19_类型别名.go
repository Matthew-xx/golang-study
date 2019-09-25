package main

import (
	"fmt"
)

func main() {
	type long int     //相当基于旧类型，创建新类型
	type str = string //相当于起别名
	var a long
	var b int
	a = 10
	b = int(a)
	var s str = "2343"
	fmt.Printf("%T,%T,%T", a, b, s)
}
