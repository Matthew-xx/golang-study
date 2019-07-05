package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
}

func main() {
	//var s *student =(&student){1,"小米"，18}无需括号
	var s *student = &student{1, "小米", 18}
	fmt.Println(*s)
	//指定成员初始化
	p := &student{id: 1, name: "华为"}
	fmt.Println(*p)
	fmt.Println(p)
}
