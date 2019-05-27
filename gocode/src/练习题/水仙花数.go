package main

import "fmt"

//水仙花数：一个三位数，各位数的立方和恰好等于它本身
func main() {
	for i := 100; i < 1000; i++ {
		baiwei := i / 100
		shiwei := (i % 100) / 10
		gewei := i % 10
		if baiwei*baiwei*baiwei+shiwei*shiwei*shiwei+gewei*gewei*gewei == i {
			fmt.Println(i)
		}
	}
}
