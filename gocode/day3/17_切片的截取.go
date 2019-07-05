package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s := a[:] //切取全部，不指定长度和容量的写法，默认全部
	fmt.Println(s)

	s2 := s[1:3:5] //切取a[1]到a[2],即2，3，容量为4
	fmt.Println(s2)

	s3 := s[:3] //切取头部到a[3]不含a[3] =s[0:3:3]
	fmt.Println(s3)
	s4 := s[5:] //切取a[5]到末尾,包含a[5]=s[5:10:10]
	fmt.Println(s4)
}
