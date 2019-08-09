package main

import "fmt"

//求1～100之间奇数的和
func main() {
	num := 0
	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			num += i
		}
		continue
	}
	fmt.Println(num)
}
