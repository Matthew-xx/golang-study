package main

import "fmt"

// 快速排序，原理是将原来的数组分成两部分，然后取中间值为基准值，比他小放左边，比他大放右边。
// 然后分别对左边和右边重复这个过程
// 在实现中，需要两个指针分别从左和右同时移动，分别和基准值比较，决定是否交换

// 为了方便递归调用，多传入两个参数，表示排序范围
func QuickSort(arr []int, left int, right int) {
	//设置基准值为数组最左边的数，并记录基准值的下标
	temp := arr[left]
	index := left
	i := left
	j := right
	// 循环移动左右指针
	for i <= j {
		// 从右面找到比基准值小的数据
		for j >= index && arr[j] >= temp {
			j--
		}
		// 将基准值放在合适位置
		if j > index {
			// 原基准值的位置，存入找到的数据
			arr[index] = arr[j]
			// 找到的位置，作为基准值的新位置，基准值内容没变仍是temp
			index = j
		}
		// 从左面找到比基准值大的数据
		for i <= index && arr[i] <= temp {
			i++
		}
		// 获取基准值合适下标
		if i <= index {
			// 原基准值的位置，存入找到的数据
			arr[index] = arr[i]
			// 在找到数据的位置，作为基准值的新位置。基准值内容没变仍是temp，注意此时a[index]内的数据没变
			// 注意，基准值是个虚量，有名义上的位置和值，但当前基准值的位置指向的内存还没写入真正的基准值
			index = i
		}

	}
	// 将基准值写入基准值指向的内存中
	arr[index] = temp
	// 递归处理一分为二的数据，直到不可再分
	if index-left > 1 {
		QuickSort(arr, left, index-1)
	}
	if right-index > 1 {
		QuickSort(arr, index+1, right)
	}
}

func main() {
	arr := []int{3, 5, 2, 8, 1, 10, 45}
	fmt.Println(arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
