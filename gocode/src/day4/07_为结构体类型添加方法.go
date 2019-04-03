package main

import "fmt"

type Person struct {
	name string
	age  int
}

//打印函数
func (p Person) printinfo() {
	fmt.Println(p)
}

//只能是自定义类型，不能是指针类型，如果要用基类型，需要先转换为自定义类型
//接收者类型不一样，就算不同的方法（即便同名），不算重载方法
//不支持重载方法，即接受者类型、方法名相同，但参数列表不同的方法才算重载，这里不支持

//type personpointer  *Person
//func (p perisonpointer) setPerson(name string, age int)
//上面是错误的

//赋值函数
func (p *Person) setPerson(name string, age int) {
	p.age = age
	p.name = name
}
func main() {
	p := Person{"老李", 50}
	p.printinfo()
	p.setPerson("老王", 60) //无需取地址
	p.printinfo()
}
