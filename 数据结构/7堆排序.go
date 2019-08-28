package main

import "fmt"

func HeapInit(arr []int) {
	length := len(arr)
	// 将切片转为二叉树模型，实现大根堆
	for i := length/2 - 1; i >= 0; i-- {
		HeapSort(arr, i, length-1)
	}
	// 根节点存储最大值，将根节点跟叶子节点数据交换
	for i := length - 1; i > 0; i-- {
		// 如果只剩下根节点和根节点下的左子节点
		if i == 1 && arr[0] <= arr[i] {
			break
		}
		arr[0], arr[i] = arr[i], arr[0]
		HeapSort(arr, 0, i-1)
	}

}
func HeapSort(arr []int, startNode int, maxNode int) {
	var max int
	// 定义左子树节点和右子树节点
	lChild := startNode*2 + 1
	rChild := lChild + 1
	// 如果子节点超出范围，跳出递归
	if lChild >= maxNode {
		return
	}
	if rChild <= maxNode && arr[rChild] > arr[lChild] {
		max = rChild
	} else {
		max = lChild
	}
	// 和根节点比较
	if arr[max] <= arr[startNode] {
		return
	}
	// 交换数据
	arr[startNode], arr[max] = arr[max], arr[startNode]
	// 递归进行下次比较
	HeapSort(arr, max, maxNode)
}
func main() {
	arr := []int{4, 6, 3, 88, 22, 345, 45}
	HeapInit(arr)
	fmt.Println(arr)
}
