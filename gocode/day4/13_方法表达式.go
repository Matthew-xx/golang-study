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

func main() {
	p := Person{"mike", 'm', 18}
	f := Person.PrintInfo //注意格式，没有括号
	f(p)                  //显式把接受着传递过去
}
