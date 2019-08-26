package main

import "fmt"

type QueueNode struct {
	data interface{}
	Next *QueueNode
}

// 创建队列
func (queue *QueueNode) Create(data ...interface{}) {
	// 数据校验
	if queue == nil {
		return
	}
	if len(data) == 0 {
		return
	}
	// 遍历数据组
	for _, v := range data {
		// 一个数据创建一个节点
		newNode := &QueueNode{v, nil}
		// newNode := new(QueueNode)
		// newNode.data = v
		// 将新节点加入到队列头之后
		queue.Next = newNode
		// 队列头指向新的节点，为插入下一个节点做准备，也就是插入下一个节点时，直接使用queue.next就是要插入的位置。
		queue = newNode
	}
}

// 打印
func (queue *QueueNode) Print() {
	if queue == nil {
		return
	}
	// 打印时从头节点（遍历）开始遍历，
	for queue != nil {
		// 注意此时的头是最初始位置，包含一个nil节点
		if queue.data != nil {
			fmt.Print(queue.data, " ")
		}
		queue = queue.Next
	}
	fmt.Println()
}

// 返回个数
func (queue *QueueNode) Length() int {
	if queue == nil {
		return -1
	}
	i := 0
	// 注意此时的头是最初始位置，包含一个nil节点
	for queue != nil {
		i++
		queue = queue.Next
	}
	// 需要减去nil头节点，数量减1
	return i - 1
}

// 入列
func (queue *QueueNode) Push(data interface{}) {
	if queue == nil || data == nil {
		return
	}
	// 还是注意初始指针在队头，要入列，应该找到队列末尾
	for queue.Next != nil {
		queue = queue.Next
	}
	// 创建新节点，加入到队列末尾
	queue.Next = &QueueNode{data, nil}
}

// 出列
func (queue *QueueNode) Pop() {
	if queue == nil {
		return
	}
	// 出列应该从队列头部出去
	queue.Next = queue.Next.Next
	// queue = queue.Next 这样做不行，因为并没有从内存中丢弃掉要出列的数据
}
func main() {
	// 创建队列头，注意每次调用队列头的方法时，队列头都会回归到最初始的位置。
	queue := new(QueueNode)
	queue.Create(1, 2, 3, 4, 5)
	queue.Print()
	fmt.Println("节点个数：", queue.Length())
	queue.Push(6)
	queue.Push(666)
	queue.Print()
	fmt.Println("节点个数：", queue.Length())
	queue.Pop()
	queue.Print()
	fmt.Println("节点个数：", queue.Length())
}
