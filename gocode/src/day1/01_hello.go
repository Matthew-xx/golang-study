// 01_hello.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	a := 3.14159
	fmt.Printf("%.30f\n", a) //???????
	fmt.Printf("==%6d==\n", 30)
	fmt.Printf("==%-6d==\n", 30)
	fmt.Printf("==%06d==\n", 30)
	// 不同数据类型之间无法直接运算，需进行转换

	// 置换的另一种方法
	// a=a+b
	// b=a-b
	// a=a-b
	// CPU运算器 控制器 寄存器

	//格式化输出的类型对应
	// %t 布尔类型
	// %s 字符串类型
	// %p 地址类型
	// %T 类型本身
	// %c字符类型

	// scan int类型 输入任何类型都接收首先能认识的那部分
	// 123.324 结果123
	// 213ndie结果：2123
	// 空格或回车作为接收结束
	// s1, s2 := "", ""
	// fmt.Scan(&s1, &s2)
	// fmt.Println(s1)
	// fmt.Printf(s2)

	var a1 int
	var b1 string
	fmt.Scanf("%3d", &a1)
	fmt.Scanf("%s", &b1)
	fmt.Println(a1)
	fmt.Println(b1)
	//输入4567890 xfcyuijk
	//输出456 7890

	//float32 7位   float64 15位
	//AsC11
	//码0到31都无法打印， 码32到126可以打印

	//0对应码48 A对应码65 a对应码97
	//大小写相差32
	return
}
