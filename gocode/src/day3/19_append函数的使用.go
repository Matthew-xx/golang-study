package main

import (
	"fmt"
)

func main() {
	s := []int{1}
	fmt.Println(s, cap(s))
	s1 := append(s, 12) //创造新切片1
	fmt.Println("追加内容后的新s1", s1)
	fmt.Println("原s", s)
	s = s1 //s、s1指向新切片1
	fmt.Println("s1内容给s后新s", s)
	s2 := append(s1, 12) //创造新切片2，s2指向她
	fmt.Println("s1追加内容后的给新s2")
	fmt.Println("此时s1", s1)
	s = s2 //s、s2都指向新切片2
	fmt.Println("S2内容给s", s)
	s[0] = 8
	fmt.Println(s, s1, s2)
	//说明
	//1.append函数实际创建了全新切片（即本身，不是引用哦），需要返回值给原来切片才能实现给原切片追加
	//2.同一类型的切片之间可以随意赋值，不需要长度相同
	//3.再次证明切片的底层逻辑，赋值操作，在表面进行，但修改其中一个，一改俱改

}
