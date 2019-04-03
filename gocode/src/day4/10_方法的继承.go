package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//实现打印方法
func (p Person) PrintInfo() {
	fmt.Printf("name=%v,sex=%c,age=%d", p.name, p.sex, p.age)
}

type Student struct {
	Person
	id int
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 123}
	s.PrintInfo()
}
