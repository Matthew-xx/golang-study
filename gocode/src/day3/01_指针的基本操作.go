// 01_指针的基本操作
package main

import (
	"fmt"
)

func main() {
	var a int
	a = 10 //每一个变量都包含2个含义， 变量内存、变量地址（指针）
	fmt.Println(a, &a)
	//fmt.Println(&a)              //变量的内存
	fmt.Printf("%v,%v\n", a, &a) //变量的地址
	var p *int
	p = &a
	fmt.Printf("p是%v,a的地址是%v\n", p, &a)
	var q *int
	q = p
	*p = 77 //修改了该指针的内容，即便是后来修改的，*q也会变化。
	fmt.Printf("p是%v,a的地址是%v,p是%v\n", p, &a, q)
	fmt.Printf("p保存的内容是%v,q保存的内容是%v\n", *p, *q)
	//一个指针指向另一个指针（一个地址指向另一个地址）

	var m **int//指针的指针类型应该为**int
	m = &p
	fmt.Println(*(*m))
	fmt.Println(m)
}
