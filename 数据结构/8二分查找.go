package main

import "fmt"

// 前提是有序数据，可使用二分查找,返回下标
func BinarySort(arr []int, num int) int {
	// 起始点
	start := 0
	// 结束下标
	end := len(arr) - 1
	// 中间下标
	mid := (start + end) / 2
	for i := 0; i < len(arr); i++ {
		if num == arr[mid] {
			return mid
		} else if num > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
		mid = (start + end) / 2
	}
	return -1
}
func main() {
	arr := []int{3, 6, 7, 9, 11, 23, 56, 87, 99, 344, 349, 787}
	num := 23
	fmt.Println("返回的下标：", BinarySort(arr, num))
}
