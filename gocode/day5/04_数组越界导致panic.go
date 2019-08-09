package main

import (
	//"errors"
	"fmt"
)

func testA() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}
func testB(x int) {
	var b [10]int
	b[x] = 8888
	fmt.Println("bbbbbbbbbbbbbbbbb")

}
func testC() {
	fmt.Println("ccccccccccccccccc")
}
func main() {
	testA()
	testB(30)
	testC()
	// defer testA()
	// defer testB()
	// defer testC()
}
