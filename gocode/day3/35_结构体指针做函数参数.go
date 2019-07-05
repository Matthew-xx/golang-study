package main

import "fmt"

type student struct {
	id   int
	name string
}

//通过传递结构体指针来修改内容
func test(s *student) {
	s.id = 3
	fmt.Println(*s)
}
func main() {
	p := &student{id: 2, name: "string"}
	fmt.Println("原内容", *p)
	test(p)
	fmt.Println("新内容", *p)
}
