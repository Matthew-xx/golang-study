package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//实现打印方法
func (p Person) PrintInfo() {
	fmt.Printf("Person:name=%v,sex=%c,age=%d", p.name, p.sex, p.age)
}

type Student struct {
	Person
	id int
}

func (s Student) PrintInfo() {
	fmt.Printf("Student:name=%v,sex=%c,age=%d", s.name, s.sex, s.age)
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 123}
	s.PrintInfo()
	//依旧是就近原则调用
	fmt.Println()
	s.Person.PrintInfo()
	//显式调用继承的方法
}
