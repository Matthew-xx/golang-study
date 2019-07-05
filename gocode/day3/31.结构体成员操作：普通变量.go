package main

import "fmt"

type student struct {
	id   int
	name string
}

func main() {
	var s student
	s.id = 1
	s.name = "小明"
	fmt.Println(s)
	//m := make(student, 3) //cannot make type student
	//fmt.Println(m)
}
