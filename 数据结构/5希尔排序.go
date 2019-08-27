package main

import "fmt"

// 缩小增量排序:
// 类似于冒泡排序，只是并非是相邻比较，而是增量比较，增量设定为数据数量的一半，并且不断减半，
// 每次得到新的增量，都要进行一次循环比较，直到增量为1
func ShellSort(arr []int) {
	// 根据数据量，确定增量变化次数，从而确定需要循环比较的次数
	for inc := len(arr) / 2; inc > 0; inc /= 2 {
		// 每次循环比较，都是从第1个数据与增量处数据比较，并向后拓展，所以次数为len(arr)-inc次，或者inc-0次，
		// 这里使用inc为循环起点，len(arr)为循环终点
		for i := inc; i < len(arr); i++ {
			// 定位要比较的数，初始位置为j,
			// 注意，不单单是比当前的，每次比完了，都要把之前比过的重新比一次
			// 故在这里，随着i的增长，j还是要做区间自减，循环比较。
			// 这实际上就是，按照增量将数据分组，同1步长触及到的所有端点，作为一个数据集合，再不断进行冒泡排序
			for j := i - inc; j >= 0; j -= inc {
				if arr[j+inc] < arr[j] {
					arr[j], arr[j+inc] = arr[j+inc], arr[j]
				} else {
					break
				}
			}
		}
	}
}

func main() {
	arr := []int{2, 6, 9, 1, 0, 4, 3, 8}
	ShellSort(arr)
	fmt.Println(arr)
}
