package main

import (
	"fmt"
)

func main() {
	//声明时初始化

	//1.全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	for m, n := range a {
		fmt.Printf("m=%v,n=%v\n", m, n)
	}
	//2.自动推导类型
	b := [5]int{1, 2, 3, 4, 5}
	for m, n := range b {
		fmt.Printf("m=%v,n=%v\n", m, n)
	}
	//3.部分初始化
	c := [5]int{1, 2, 3}
	for m, n := range c {
		fmt.Printf("m=%v,n=%v\n", m, n)
	}
	//3.指定位置初始化
	//d := [5]int{0: 1, 2: 2, 3: 3}
	d := [5]int{4: 1, 3: 2, 1: 3}
	//":"分割序号和值，“，”分割数组元素
	for m, n := range d {
		fmt.Printf("m=%v,n=%v\n", m, n)
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d) //直接打印数组

}
