//一只牛寿命6年，第3岁和第5岁剩下小牛，小牛同理。初始一只牛，传入参数N年返回牛的数量
// 分析这是一道递归题目，模拟牛生牛即可

package main

func main() {

}

var count int = 1

func NewCattle(year int) (count int) {

	age := 1
	for {
		age++
		year--
		if age == 3 {
			count++
			NewCattle(year - 3)
		}
		if age == 5 {
			count++
			NewCattle(year - 5)
		}
		if age == 6 {
			count--
		}

	}

}
