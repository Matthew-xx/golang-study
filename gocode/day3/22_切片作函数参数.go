package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}

}
func main() {
	n := 10
	s := make([]int, n)
	InitData(s)
	fmt.Print(s)
}

//引用类型数据做函数参数时，传递进函数的引用的副本，即形参也指向实参的底层数组，这时修改形参的值，就会同步更新到实参。
