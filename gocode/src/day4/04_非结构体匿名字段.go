package main

import "fmt"

type str string
type Person struct {
	name string
	age  int
}
type Student struct {
	Person //结构体类型匿名字段
	int    //基础类型匿名字段
	str    //自定义类型匿名字段
}

func main() {
	//顺序初始化
	s := Student{Person{"小明", 18}, 888, "beijing"}
	fmt.Printf("s=%+v\n", s)
	//指定类型初始化.没有名字字段，用类型名代替
	s1 := Student{Person: Person{name: "小红", age: 30}, int: 18, str: "shanghai"}
	fmt.Printf("s1=%+v\n", s1)
	s1.int = 20
	fmt.Printf("s1=%+v\n", s1)
}
