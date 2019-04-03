package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{1: "mike", 2: "Bob"}
	//如果对应的键已经存在，则赋值就是修改其值
	fmt.Println(m1)
	m1[2] = "haha"
	fmt.Println(m1)
	//如果对应的键不存在，则赋值就是追加内容类似append
	m1[4] = "我是直接追加的"
	fmt.Println(m1)
	m1 = append()
	fmt.Println(m1)
}
