package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "mike", 2: "Bob"}
	//var key,value ：=直接初始化两个变量
	//第一个为key 第二个为value
	for key, value := range m {
		fmt.Printf("key=%v,value=%v\n", key, value)
	}
	//总结range返回map可遍历键值对
	value, ok := m[5]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}
	//总结map元素本身可返回两个返回值，第一个为元素键对应的值，第二个为是否存在

	m[1] = m[2] //把2键对应的值，复制一份给1键
	m[2] = "hello"
	fmt.Println(m)
}
