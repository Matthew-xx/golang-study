package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//实现打印方法
// func (p Person) PrintInfo() {
// 	fmt.Printf("name=%v,sex=%c,age=%d\n", p.name, p.sex, p.age)
// }
func (p *Person) PrintInfo2() {
	fmt.Printf("name=%v,sex=%c,age=%d\n", p.name, p.sex, p.age)
}

type Student struct {
	Person
	id int
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 123}
	// s.PrintInfo()
	// s.Person.PrintInfo()
	s.PrintInfo2()

	s.Person.PrintInfo2()

	// (&s).PrintInfo()
	// (&s).Person.PrintInfo()
	(&s).PrintInfo2()
	// (&s).Person.PrintInfo2()

}
