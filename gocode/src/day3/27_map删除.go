package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "mike", 2: "Bob"}
	fmt.Println(m)
	delete(m, 2) //在m中删除键为2的键值对
	fmt.Println(m)
}
