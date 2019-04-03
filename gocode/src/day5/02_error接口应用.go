package main

import (
	"errors"
	"fmt"
)

func MyDiv(a, b int) (err error, r int) {
	if b == 0 {
		return errors.New("除数不能为0"), 0
	} else {
		return nil, a / b
	}
	// if b == 0 {
	// 	err = errors.New("除数不能为0")
	// } else {
	// 	result = a / b
	// }
	// return

}
func main() {
	err, result := MyDiv(5, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
