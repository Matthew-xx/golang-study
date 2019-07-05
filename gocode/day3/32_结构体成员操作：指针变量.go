package main

import "fmt"

func main() {
	//1.1先声明普通变量，利用变量地址生成指针
	var s student   //不可省略，必须先有合法地址，才能有指针
	var p *students //可通过自动推倒类型省略了这一步
	p = &s
	p.id = 1
	p.name = "xiaomi"
	fmt.Println(s)
	//错误写法
	// var p1 *student
	// p1.id = 1
	// p1.name = "xiaomi"
	var p1 *student
	p1 = &s
	p1.id = 1
	p1.name = "xiaomi"
	fmt.Println(p1)
	//2直接利用new开辟内存，生成指针
	p2 := new(student)
	p2.id = 2         //可以直接用指针操作成员
	(*p2).name = "每组" //可以用变量操作成员

	fmt.Println(*p2) //取出内容
}

type student struct {
	id   int
	name string
}
