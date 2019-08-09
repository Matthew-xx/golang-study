package main

import "fmt"

type Humaner interface {
	sayhi()
}
type Student struct {
	name string
	id   int
}

func (s Student) sayhi() {
	fmt.Printf("student say hi:%s,%d\n", s.name, s.id)
}

type Teacher struct {
	addr  string
	group string
}

func (t Teacher) sayhi() {
	fmt.Printf("teacher say hi:%s,%s\n", t.addr, t.group)
}

type MyStr string

func (s MyStr) sayhi() {
	fmt.Printf("mystr say hi:%s\n", s)
}
func main() {
	//定义接口类型变量
	var i Humaner
	//只要实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以给i赋值。
	s := Student{"mike", 666}
	i = s
	//s.sayhi()
	i.sayhi()
	t := Teacher{"beijing", "go"}
	i = t
	i.sayhi()
}
