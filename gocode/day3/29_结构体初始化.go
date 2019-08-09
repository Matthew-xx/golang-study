package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
}

func main() {
	//顺序初始化
	var s student = student{1, "小米", 12}
	fmt.Println(s)
	//指定成员初始化
	var s1 student = student{id: 1, name: "小明"}
	fmt.Print(s1)
}
