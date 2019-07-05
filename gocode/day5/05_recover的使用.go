package main

import (
	//"errors"
	"fmt"
)

func testA() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}
func testB(x int) {
	//defer recover()
	//defer fmt.Println(recover())
	//recover 只能放到 defer 函数里面，不能放到子函数。
	//实测直接 defer recover() 也不行。
	// defer func() {
	// 	recover()
	// }() //调用一下
	defer func() {
		fmt.Println(recover())
	}() //这样不好，没panic也执行输出nil
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println(recover())经典错误：这样相当于重新调用一次recover()
			fmt.Println(err)
		}
	}() //调用一下

	var b [10]int
	b[x] = 8888
	fmt.Println("bbbbbbbbbbbbbbbbb")

}
func testC() {
	fmt.Println("ccccccccccccccccc")
}
func main() {
	testA()
	testB(0)
	testC()
	// defer testA()
	// defer testB()
	// defer testC()
}
