package main

import (
	"fmt"
)

func main() {
	a := []int{1, 6, 34, 87, 3, 8, 23, 8, 4, 67}
	fmt.Println("原a：", a)

	s := a[2:5] //34 87 3注意不要设置容量，否则影响二次切片长度，
	s[2] = 88   //3被改为88,注意不要越界，长度为3容量为4，但s[3]并不存在
	fmt.Println("第1次切片并改值", a)

	s2 := s[2:5] //3 8 23
	s2[2] = 99   //23被改为99
	fmt.Println("第2次切片并改值", a)

	// 两个不同切片一旦互相赋值，总指向同一底层数组，元素一改俱改

	as := []int{1, 2}
	bs := []int{2, 3}
	as = bs
	bs[0] = 8 //元素一改俱改
	fmt.Println(as, bs)
	bs = []int{6} //整个切片类型赋值，指向的底层数组也就变了
	fmt.Println(as, bs)

	// as := []int{1, 2}
	// bs := []int{2, 3}
	// as = bs
	// bs = []int{6}
	// fmt.Println(as, bs)
}
