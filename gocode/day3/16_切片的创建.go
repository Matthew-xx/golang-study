package main

import (
	"fmt"
)

func main() {
	//a := [5]int{1, 2, 3, 4, 5}
	//s := make([]int, 3, 10)  分别时len和cap
	s := make([]int, 3) //不指定cap,则cap=len
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
}

//len是指切片的长度，即切片中元素的个数。
//cap是指切片的容量，即切片当前开辟的内存空间可以存储的元素个数。
//当cap不够用时，会自动申请一个翻倍大小的内存空间，并把原来空间的内容全部复制进去。但这样非常耗时
//故cap的常用来指定一个足够大的数字，达到空间换时间的效果。
