package main

import "fmt"

func main() {
	s := "abcdef"
	// s[1] = 'm'
	fmt.Println(s, s[3])
	// golang中的string类型存储的字符串是只读的， 如果要修改string内容需要将string转换为[]byte或[]rune，并且修改后的string内容是重新分配的
	r := []rune(s)
	r[1] = 'm'
	s = string(r)
	fmt.Println(s)

}
