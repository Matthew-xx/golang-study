package main

import "fmt"

func add(a, b int) int {
	return a + b
}

type long int

func (c long) add2(a, b long) long {
	c = a + b
	return c
}

// func (c long) add2(a, b int) long {
// 	c = long(a) + long(b)
// 	return c
// }注意类型检查很严格
func main() {
	c := add(1, 2)
	fmt.Println(c)
	var m long
	m = m.add2(3, 5)
	fmt.Println(m)

}
