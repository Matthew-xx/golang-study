package main

import "fmt"

func main() {
	a := '中'
	// var b byte = '国'
	// constant 22269 overflows byte,byte装不下代表 国这个字符的int编码
	var b rune = '华'

	c := `人`
	d := "民"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(string(b))
	fmt.Println(c, d)
	fmt.Println(c + d)
}
