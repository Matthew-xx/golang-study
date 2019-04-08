package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "mike", 2: "Bob"}
	m1:=m
	fmt.Println(m)
	delete(m, 2) //在m中删除键为2的键值对
	fmt.Println(m)
	fmt.Println(m1)
}
