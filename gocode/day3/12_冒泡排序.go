package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println("排序前：\n", a)
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println("排序后：\n", a)

}
