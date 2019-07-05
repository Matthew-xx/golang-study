package main

import "fmt"

type Humaner interface {
	sayhi()
}
type Student struct {
	name string
	id   int
}

func (s *Student) sayhi() {
	fmt.Printf("student say hi:%s,%d\n", s.name, s.id)
}

type Teacher struct {
	addr  string
	group string
}

func (t *Teacher) sayhi() {
	fmt.Printf("teacher say hi:%s,%s\n", t.addr, t.group)
}

type MyStr string

func (s *MyStr) sayhi() {
	fmt.Printf("mystr say hi:%s\n", *s)
}

//定义一个普通函数，函数的参数为接口类型
//一个函数多种不同表现即为多态
func WhoSayHi(i Humaner) {
	i.sayhi()
}
func main() {
	s := &Student{"mike", 666}
	t := &Teacher{"beijing", "go"}
	var str MyStr = "hello mike"
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)
	//make为humaner接口类型的切片开辟内存
	x := make([]Humaner, 3)s
	x[0] = s
	x[1] = t
	x[2] = &str
	for _, i := range x {
		i.sayhi()
	}
}
