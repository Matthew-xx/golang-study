package main

import "fmt"

// 外层控制行
// 内层控制列
// 找到最大值
// 最大值与最后一个元素交换
func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		// 用来存储所得到的最大值的下标
		index := 0
		// 遍历数据，查找最大值
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] > arr[index] {
				index = j
			}
			// 此循环结束后，获得所有数据中的最大值的下标
		}
		arr[index], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[index]
	}
}
func main() {
	arr := []int{3, 6, 2, 1, 7, 5, 9, 0}
	fmt.Println("排序前：", arr)
	SelectSort(arr)
	fmt.Println("排序后：", arr)
}
