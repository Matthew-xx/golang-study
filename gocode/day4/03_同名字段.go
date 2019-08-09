package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Student struct {
	Person
	name string //同名字段，与继承过来的变量重复了
	id   int
	addr string
}

func main() {
	var s Student
	s.name = "我是谁的名字呢"
	fmt.Printf("s=%+v", s)
	//s={Person:{name: age:0} name:我是谁的名字呢 id:0 addr:}
	//优先本作用域字段
	//显式调用
	s.Person.name = "P"
	fmt.Printf("s=%+v", s)
}
