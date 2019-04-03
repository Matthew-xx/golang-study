package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("这是一个自己的错误")
	fmt.Println(e)
	//e2 := fmt.Errorf("%s", "这是两个自己的错误")
	e2 := fmt.Errorf("%s", e)
	fmt.Println(e2)
}
