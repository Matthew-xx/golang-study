package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello kitty"
	//一个字符串是否包含另一个字符串，两个字符串作为参数，包含返回true,
	fmt.Println(strings.Contains(s, "hello"))
	//Join,将一个字符串切片的内容，用X字符串拼接成一个字符串
	slice := []string{"a", "b", "c", "d"}
	fmt.Println(strings.Join(slice, "#"))
	//Index,一个字符串在另一个字符串中的位置(第一次出现)，没有则返回-1
	fmt.Println(strings.Index(s, "l"))
	//repeat,重复一个字符串X次
	fmt.Println(strings.Repeat("go", 3))
	//split,把一个字符串，根据指定字符串拆开成一个切片。
	slice1 := strings.Split("185734549@qq.com@myhome", "@")
	fmt.Println(slice1[0])
	fmt.Println(slice1)
	//trim,把一个字符串的两头的指定字符串全部去掉。常用来删除两头空格
	s2 := "   a   b    c  d    "
	fmt.Println(strings.Trim(s2, "a"))
	fmt.Println(strings.Trim(s2, " "))
	//field,把一个字符串按照其中的空格拆成一个切片
	s3 := strings.Fields(s2)
	for _, s4 := range s3 {
		fmt.Println(s4)
	}

}
