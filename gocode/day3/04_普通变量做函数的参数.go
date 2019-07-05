// 04_普通变量做函数的参数
package main

import (
	"fmt"
)

func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap 函数内的结果a=%v,b=%v", a, b)
}
func main() {
	a, b := 10, 20
	swap(a, b)
	fmt.Printf("main 函数内的结果a=%v,b=%v", a, b)
}
