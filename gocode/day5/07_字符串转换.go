package main

import (
	"fmt"
	"strconv"
)

func main() {

	//append系列函数，将其他类型转化为字符串后，再转为字节切片并追加到现有字节型切片，成为一个新的字节切片
	b1 := []byte{'a', 'b'}
	b1 = strconv.AppendBool(b1, true)
	fmt.Printf("b1=%c\n", b1)
	b1 = strconv.AppendInt(b1, 124, 10) //第3个参数代表进制
	fmt.Printf("b1=%c\n", b1)
	b1 = strconv.AppendQuote(b1, "aaaaaaaa") //Quote表示引号
	fmt.Printf("b1=%c\n", b1)

	//append系列函数，将其他类型转化为字符串后，再转为字节切片并追加到现有字节型切片，成为一个新的字节切片
	s := []string{"a", "b"}
	s = strconv.AppendBool(s, true)
	fmt.Printf("b1=%c\n", s)
	s = strconv.AppendInt(s, 124, 10) //第3个参数代表进制
	fmt.Printf("b1=%c\n", s)
	s = strconv.AppendQuote(s, "aaaaaaaa") //Quote表示引号
	fmt.Printf("b1=%c\n", s)

	//format系列函数，将其他类型转换为字符串类型
	str := strconv.FormatBool(false)
	fmt.Println(str)
	//f指打印格式以小数方式，-1小数位数为紧缩，以64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println(str)
	//字符串转其它类型
	flag, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("type is %T\n", flag) //%T代表类型 %V %t??
	}
	//常用的字符换与int互转,
	str = strconv.Itoa(34534) //=FormatInt(34534,10)
	fmt.Println(str + "1")
	a, _ := strconv.Atoi(str) //=ParseInt(str,10,0)
	fmt.Println(a + 1)
	//

}
