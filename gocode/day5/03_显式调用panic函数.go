package main

import (
	//"errors"
	"fmt"
)

func testA() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}
func testB() {
	fmt.Println("bbbbbbbbbbbbbbbbb")
	panic("panic test")
}
func testC() {
	fmt.Println("ccccccccccccccccc")
}
func main() {
	// testA()
	// testB()
	// testC()
	defer testA()
	defer testB()
	defer testC()
	testA()
	testB()
	testC()
	defer testC()
}
