package main

import "fmt"

func myfunc(m map[int]string) {
	m[1] = "小米"
	m[3] = "苹果"
	// delete(m, 2) //删除某个key
}
func main() {
	m := make(map[int]string)
	m[2] = "华为"
	fmt.Println(m)
	myfunc(m)
	fmt.Println(m)
}

// map为引用传递
