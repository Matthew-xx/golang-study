package main

import "fmt"

func main() {
	var day int
	fmt.Println("请输入天数：")
	fmt.Scanf("%d", &day)
	week := day / 7
	rest := day % 7
	fmt.Printf("这是%d周零%d天\n", week, rest)
}
