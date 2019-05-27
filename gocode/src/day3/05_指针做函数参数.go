package main

import "fmt"

func swap(a, b *int) {
	a, b = b, a
	// *a, *b = *b, *a
	fmt.Printf("swap: a=%d,b=%d\n", *a, *b)
}
func main() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Printf("main：a=%d,b=%d\n", a, b)
	//总结：指针操作可以操作传过来的内存地址，通过修改内容地址内保持的内容，从而是作用域外部的变量值发生变化。
	//即：传递参数给形参，只能修改形参内容（值传递），无法影响传给形参的原始变量。而传递指针（地址传递），可以通过操作指针，把原始变量的内容进行修改。
	//如果swap()内的语句是a,b=b,a,则不会有影响外部变量值。因为此时函数参数是指针，函数操作的也是指针。函数运行完毕形参（指针）也就被清空了，不会影响外部

	//传指针的本质是，利用指针把存储实参的地址的变量复制一份传到函数内部，此时利用地址可以找到存储实参值的内存空间，从而修改实参的值。
	//如果传递指针，直接操作指针而非操作地址值，这时操作的其实是存储地址的指针副本（形参），效果相当于值传递，不会改变实参。

	//二次总结：说明参数的传递本质就是复制一份内存数据传递进去，值传递是复制一份值传递进入函数操作，顾不影响外面的值。
	//指针传递，是复制一份存储值的地址进入函数，操作地址内容，是在操作副本指向的内容，故可影响外部实参值。操作地址本身，是在操作副本本身。故不影响外部实参值。
}
