package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreatNum(num *int) {
	rand.Seed(time.Now().UnixNano())
	for {
		if rand.Intn(10000) > 999 {
			*num = rand.Intn(10000)
			break
		}
	}
	//fmt.Println(*num)临时显示
}
func GetNum(s []int, num int) {
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10
}
func OnGame(s []int) {
	var a int
	s2 := make([]int, 4)
	for {
		for {
			fmt.Println("请输入一个4位数：")
			//fmt.Scanf("%d", &a)循环里使用scanf时会出现未输入就被视为输入完成的情况
			fmt.Scan(&a)
			if a > 999 && a < 10000 {
				break
			} //else {
			fmt.Println("输入错误")
			//}
		}
		GetNum(s2, a)
		n := 0
		for i := 0; i < 4; i++ {
			if s[i] < s2[i] {
				fmt.Printf("第%d个数大了\n", i)
			} else if s[i] > s2[i] {
				fmt.Printf("第%d个数小了\n", i)
			} else {
				fmt.Printf("第%d个数对了\n", i)
				n++
			}
		}

		if n == 4 {
			fmt.Printf("都猜对了，游戏结束")
			break
		} //无else即不符合时什么都不做
	}
}
func main() {
	var num int
	CreatNum(&num)
	s := make([]int, 4)
	GetNum(s, num)
	//fmt.Println(s)
	OnGame(s)
}
