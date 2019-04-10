package main

import (
	"errors"
	"fmt"
)

// MyDiv 除法函数
func MyDiv(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
	} else {
		result = a / b
	}
	return
}
func main() {
	result, err := MyDiv(5, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
