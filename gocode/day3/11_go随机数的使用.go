// 11_go随机数的使用
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //设置种子
	for i := 0; i < 5; i++ {
		//fmt.Println(rand.Int())
		fmt.Println(rand.Intn(100))
	}

}
