package main

import (
	"fmt"
)

func main() {
	s := []int{1}
	fmt.Println("初始内容及容量", s, cap(s))
	s1 := append(s, 12) //利用append 创造新切片s1
	fmt.Println("追加内容后的新s1", s1,cap(s1))
	//修改s1,看s
	s1[0] = 8
	fmt.Println("修改s1看s", s)
	fmt.Println("修改s1看s", s1)

	s2 := append(s1, 18) //创造新切片2，s2指向她
	s=s2
	fmt.Println("s1追加内容后的给新s2", s2)
	//s3 := append(s, s1, s2)
	s3 := append(append(s, s1...), s2...)

	fmt.Println(s3)
	//说明
	//1.append函数实际根据被追加的切片指向的底层数组，复制了全新的底层数组，返回值为指向新底层数组的新切片引用。
	//2.同一类型的切片之间可以随意赋值，不需要长度相同
	//3.再次证明切片是引用类型，通过赋值指向同一底层数组的切片，一改俱改
	//4.append函数可不仅可以追加多个元素，也可以追加一个同类型切片，但需要加上...将元素打散

}
