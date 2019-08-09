package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Student struct {
	Person //只有类型没有名字的字段，即把整个结构体类型名字放进来，集成其全部成员
	id     int
	addr   string
}

func main() {
	var s1 Student
	s1.name = "小明"
	s1.age = 19
	s1.id = 000
	fmt.Printf("s1=%+v\n", s1)
	s1.Person = Person{"小红", 18}
	fmt.Printf("s1=%+v\n", s1)

	//验证匿名字段被实例操作也不受方法集约束
	(&s1).Person = Person{"小蓝", 18}
	fmt.Printf("s1=%+v\n", s1)
	s2 := new(Student)
	s2.name = "小妞" //(*s2).name = "小牛"
	fmt.Printf("s1=%+v\n", *s2)

}
