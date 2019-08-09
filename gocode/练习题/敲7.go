package main

import "fmt"

// 1~100，7 7的倍数，带7的都不要打印
func main() {
	for i := 0; i <= 100; i++ {
		if i%7 == 0 {
			continue
		}
		if i >= 70 && i < 80 {
			continue
		}
		if i%10 == 7 {
			continue
		}
		fmt.Println(i)
	}
}
