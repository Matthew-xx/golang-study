package main

import "fmt"

// 冒泡排序口诀：
/*
外层控制行：
内层控制列：
相邻要比较
满足则交换
*/

// 总结；
// 1. 要多次从头开始比较相邻元素，但比较一次，就得到一个最大的，下次只需找出剩下所有元素中最大的，
//     就不需要再比最大的那个了，随着次数增多，需要比较的元素越来越少
//   规律是：
// 当次外循环的内层循环执行次数 = 数据个数 - 当次外层执行编号 - 1
// 外层总需执行次数=数据个数-1
// 推论：
// 当次外循环的内层循环执行次数 = 数据个外层总需执行次数数 - 当次外层执行编号

func BubbleSort(arr []int) int {
	// 定义个计数器，统计比较次数
	count := 0
	// 外层循环控制行
	for i := 0; i < len(arr)-1; i++ {
		// 内层控制列
		for j := 0; j < len(arr)-1-i; j++ {
			count++
			// 相邻要比较，满足条件否
			if arr[j] > arr[j+1] {
				// 满足则交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return count
}

// 对冒泡排序进行优化，如果原始数据较有序，可减少比较次数:
// 原理：若比较后没有对数据进行交换操作，说明数据有序，只需要记录下这个状态，下次无需进行比较了，
func OptimizedBubbleSort(arr []int) int {
	// 定义个计数器，统计比较次数
	count := 0
	// 定义个标识位，存储是否发生过数据交换,默认为发生过交换，需要比较
	isSwap := true
	// 外层循环控制行，只有发生过数据交换，才有必要继续遍历外层循环，没有发生，则认为已经有序
	for i := 0; i < len(arr)-1 && isSwap; i++ {
		// 开始遍历过程中，每次都默认没有发生数据交换，如果遍历完，真没有发生，则会中断循环
		isSwap = false
		// 内层控制列,
		for j := 0; j < len(arr)-1-i; j++ {
			count++
			// 相邻要比较，满足条件否
			if arr[j] > arr[j+1] {
				// 满足则交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// 发生数据交换了，保存状态，以便继续遍历
				isSwap = true
			}
		}
	}
	return count
}
func main() {
	arr := []int{1, 2, 346, 7, 3, 8, 4, 6}
	count := BubbleSort(arr)
	fmt.Println(arr)
	fmt.Println("比较次数：", count)

	count = OptimizedBubbleSort(arr)
	fmt.Println(arr)
	fmt.Println("优化后比较次数：", count)

}
