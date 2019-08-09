package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

//值语义
func (p Person) setPersonValue(name string, age int) {
	p.name = name
	p.age = age
	fmt.Println(p)
}

//引用语义
func (p *Person) setPersonPointer(name string, age int) {
	p.name = name
	p.age = age
	fmt.Println(p)
}
func main() {
	p := &Person{"mike", 12}
	//类型的方法集是指可以被该类型的值调用的所有方法的集合
	//用实例value和pointer调用方法（含匿名字段）不受方法集的约束，总能找到全部方法，自动转换类型
	p.setPersonPointer("老张",60)
	(*p).setPersonPointer("老张",60)
	p.setPersonValue("老张",60)
	(*p).setPersonValue("老张",60)
}
