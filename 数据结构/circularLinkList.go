package main

import (
	"fmt"
)

// 单向循环链表
type CircularLinkList struct {
	Data interface{}
	Next *CircularLinkList
}

// 创建循环链表
func (node *CircularLinkList) Create(data ...interface{}) {
	if node == nil || len(data) == 0 {
		return
	}
	// 记录头结点
	head := node
	for _, v := range data {
		newNode := &CircularLinkList{v, nil}
		node.Next = newNode
		node = node.Next
	}
	// 最后头尾相接,注意头结点不存储任何信息，故需要跳过头结点
	node.Next = head.Next
}

// 打印循环链表
func (node *CircularLinkList) Print() {
	if node == nil {
		return
	}
	// 记录循环的起始位置
	start := node.Next
	for {
		node = node.Next
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
		// 当下一个节点又是起始位置时，中断循环
		if start == node.Next {
			break
		}
	}
}

// 打印循环链表的长度
func (node *CircularLinkList) Length() int {
	if node == nil {
		return -1
	}
	// 记录循环的起始位置
	start := node.Next
	i := 0
	for {
		node = node.Next
		i++
		if start == node.Next {
			break
		}
	}
	return i
}

// 按照下标插入数据
func (node *CircularLinkList) Insert(index int, data interface{}) {
	if node == nil || index < 0 || index > node.Length() || data == nil {
		return
	}
	preNode := node
	for i := 0; i < index; i++ {
		preNode = node
		node = node.Next

	}
	// 在第0个位置插入时的特殊情况，需要处理
	if index == 0 {
		// 暂存头结点
		head := node
		// 新节点的后继节点为头结点之后的第一个节点
		newNode := &CircularLinkList{data, node.Next}
		// 找到最后一个节点
		for i := 0; i < node.Length(); i++ {
			node = node.Next
		}
		// 将最后一个节点的后继设置为新节点，完成插入
		node.Next = newNode
		// 因为是在头结点插入，头结点的位置也要换到新位置
		head.Next = newNode

	} else {
		// 其他情况，只需把当前节点设置为新节点的后继
		newNode := &CircularLinkList{data, node}
		// 把新节点设置为上一节点的后继
		preNode.Next = newNode
	}
}

// 按照下标删除数据
func (node *CircularLinkList) Delete(index int) {
	// 注意index不能为0,不能删除头结点
	if node == nil || index <= 0 || index > node.Length() {
		return
	}
	// 删除一个节点是需要将被删除的节点的前后节点重新缝合的，而第一个节点的比较特殊
	preNode := node
	// 删除第一个节点时，需要特殊处理,因为初始状态下，第一个节点的上一节点为头结点
	if index == 1 {
		head := node
		// 找到最后一个节点
		for i := 0; i < node.Length(); i++ {
			node = node.Next
		}
		node.Next = head.Next.Next
		// 调整头结点的位置
		head.Next = head.Next.Next
	} else {
		// 调整指针到所需位置
		for i := 2; i <= index; i++ {
			preNode = node
			node = node.Next
		}
		// 开始删除当前节点
		preNode.Next = node.Next
		node.Next = nil
		node.Data = nil
		node = nil
	}
}
func main() {
	list3 := new(CircularLinkList)
	list3.Create(1, 2, 3, 4)
	list3.Print()
	count := list3.Length()
	fmt.Println("节点个数：", count)
	list3.Insert(0, 6)
	fmt.Println("在第0个位置插入节点")
	list3.Print()
	fmt.Println("删除节点")
	list3.Delete(1)
	list3.Print()

	// 约瑟夫环的实现
	list3.Create("01", "02", "03", "04", "05", "06", "07", "08", "09", 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32)
	fmt.Println()
	list3.Print()
	fmt.Println()
	i := 1
	for list3.Length() > 2 {
		i += 3
		if i > list3.Length() {
			i = list3.Length() % 3
		}
		list3.Delete(i)
		list3.Print()
		fmt.Println()
	}
}
