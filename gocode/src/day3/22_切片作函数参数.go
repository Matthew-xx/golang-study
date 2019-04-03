package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}

}
func main() {
	n := 10
	s := make([]int, n)
	InitData(s)
	fmt.Print(s)
}
