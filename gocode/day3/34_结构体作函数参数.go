package main

import "fmt"

type student struct {
	id   int
	name string
}

func test(s student) {
	s.id = 1
	fmt.Println(s) //1
}
func main() {
	s := student{id: 2}
	test(s)
	fmt.Println(s) //2
}

//证明是值传递，想改内容需要地址传递
