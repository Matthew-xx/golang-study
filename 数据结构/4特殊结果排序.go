package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 在一些特定场景，可使用此排序方法，得到排序结果
// 原理：如果要排序的数字，在一定的范围，而且有大量重复，可通过统计个数，进行排序
// 由原理可知，其实此法并没有真正进行排序，算是投机取巧，因为已经知道数据值域在一个有限的范围，
// 又利用map统计了不同数据出现的次数，直接通过在值域顺序穷举的方式，将排序结果打印了出来。而原数据在内存中并未被排序

func main() {
	// 生成大量0~999的数字（10000个）
	rand.Seed(time.Now().UnixNano())
	s := make([]int, 0)
	for i := 0; i < 10000; i++ {
		s = append(s, rand.Intn(1000))
	}
	fmt.Println(s)

	// 利用map统计各数据出现的次数
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	fmt.Println(m)

	// 遍历数据的所有可能性，根据不同数据出现的次数，打印结果，map中不存在key的即出现次数为0
	for i := 0; i < 1000; i++ {
		// 根据当前出现的次数，决定打印次数
		for j := 0; j < m[i]; j++ {
			fmt.Print(i, " ")
		}
	}
}
