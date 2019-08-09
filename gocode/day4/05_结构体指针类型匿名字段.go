package main

import "fmt"

type str string
type Person struct {
	name string
	age  int
}
type Student struct {
	*Person //结构体指针类型匿名字段
	int
	str
}

func main() {
	//1.直接顺序初始化
	s := Student{&Person{"xiaomi", 29}, 12, "hahah"}
	fmt.Println(s.name, s.age, s.int, s.str)
	//2.声明变量，利用变量操作指针类型匿名字段，然后new分配空间
	var s1 Student
	s1.Person = new(Person) //分配空间
	s1.name = "huawei"
	fmt.Println(s1.name, s1.age, s1.int, s1.str)
	//3.声明变量,利用变量操作指针类型匿名字段直接赋值地址
	var s2 Student
	s2.Person = &Person{"oppo", 19}
	fmt.Println(s2.name, s2.age, s2.int, s2.str)

}
