package main

import (
	"fmt"
)

func main() {
	var m map[int]string
	//定义一个变量，类型为map[int]string
	//m[1] = "Bob"
	fmt.Println(m)
	fmt.Println(len(m) /*cap(m)*/) //没有cap
	//invalid argument m (type map[int]string) for cap
	//使用make创建，可指定长度（可选）
	m2 := make(map[int]string, 1) //1是存储能力，内容长度为0
	fmt.Println(m2)
	fmt.Println(len(m2)) //依旧为0
	//使用make创建，
	m2[1] = "mike"
	m2[0] = "game"
	fmt.Println(m2)      //无序打印
	fmt.Println(len(m2)) //自动扩充为2，并不被存储能力1限制
	//初始化
	//m3 := make(map[int]string){1,"mike",2,"Bob"}
	//1.注意格式","用来分割元素，":"用来分割键与值
	//2.直接初始化不需要用make,make创建规避空指针的空map
	m3 := map[int]string{1: "mike", 2: "Bob"}

	fmt.Println(m3) //无序打印
	fmt.Println(len(m3))
	// 总结：
	// make可指定开辟的内存个数，但不会被这个数字限制
	// map是有长度的，长度就是键值对的个数。
}
