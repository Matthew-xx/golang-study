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
	//顺序初始化，注意加上类型名
	var s1 Student = Student{Person{"xiaoming", 18}, 001, "beijing"}
	fmt.Println(s1)
	fmt.Printf("s1=%+v\n", s1) //%+v显示更详细的
	//置顶成员初始化，未指定为0或空，注意加上类型名和格式
	s2 := Student{Person: Person{age: 20}, id: 002}
	fmt.Printf("s2=%+v", s2)
}
