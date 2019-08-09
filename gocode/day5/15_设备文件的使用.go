package main

import (
	"fmt"
	"os"
)

func main() {

	//os.Stdout.Close()默认是 打开的，通常也不关
	fmt.Println("a you ok")
	//往标准输出设备写内容
	os.Stdout.WriteString("are you ok?\n")
	var a int
	fmt.Println("输入数字：")
	os.Stdin.Close() //
	fmt.Scan(&a)
	fmt.Println("输入的数字：", a)

}
