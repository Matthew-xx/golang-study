package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//实现打印方法
func (p Person) PrintInfo() {
	fmt.Printf("Person:name=%v,sex=%c,age=%d", p.name, p.sex, p.age)
}

func main() {
	p := Person{"mike", 'm', 18}
	p.PrintInfo()                  //传统调用方法
	f := p.PrintInfo               //把方法入口地址赋值给一个变量，调用时隐藏了接收者，
	f()                            //利用变量去调用
	fmt.Printf("\ntype is %+T", f) //查看值类型type is func()

}
