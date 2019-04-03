package main

import . "fmt"

func swap(a, b *int) {
	a, b = b, a
	Printf("swap:a=%d,b=%d", *a, *b)
}
func main() {
	a, b := 10, 20
	swap(&a, &b)
	Printf("main：a=%d,b=%d", a, b)
	//总结：指针操作可以操作传过来的内存地址，通过修改内容地址内保持的内容，从而是作用域外部的变量值发生变化。
	//即：传递参数给形参，只能修改形参内容（值传递），无法影响传给形参的原始变量。而传递指针（地址传递），可以通过操作指针，把原始变量的内容进行修改。
}
