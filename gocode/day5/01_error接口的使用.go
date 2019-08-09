package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("errors:错误")
	fmt.Println(e)
	//e2 := fmt.Errorf("%s", "这是两个自己的错误")
	e2 := fmt.Errorf("test:%s", "fmt：错误")
	fmt.Println(e2) //ewrrrrrr
	var e3 error
	//e3 = fmt.Errorf("%s", e)
	fmt.Println(e3)
}
