package main

import "fmt"

func main() {
	//空接口是万能类型，可以保存任何类型的值
	var i interface{} = 1
	fmt.Println("i=", i)
}
