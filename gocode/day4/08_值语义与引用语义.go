// 08_值语义与引用语义
package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

//值语义
func (p Person) setPerson1(name string, age int) {
	p.name = name
	p.age = age
	fmt.Println(p)
}

//引用语义
func (p *Person) setPerson2(name string, age int) {
	p.name = name
	p.age = age
	fmt.Println(*p)
}
func main() {
	var p1 Person
	p1.setPerson1("老王", 78)
	fmt.Println(p1)
	var p2 *Person   //或简写p2:= new(Person)
	p2 = new(Person) //分配空间
	p2.setPerson2("老张", 60)//等于(*p2).setPerson2("老张", 60)
	fmt.Println(*p2)
}
