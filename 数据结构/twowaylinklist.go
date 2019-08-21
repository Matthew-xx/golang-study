package main

import "fmt"

type TwoWayLinkNode struct {
	Prev *TwoWayLinkNode
	Next *TwoWayLinkNode
	Data interface{}
}

func (node *TwoWayLinkNode) Create(data ...interface{}) {
	if node == nil || len(data) == 0 {
		return
	}

	for _, v := range data {
		// 创建新的节点，创建时就写入前驱和内容，就差后继了
		newNode := &TwoWayLinkNode{node, nil, v}
		// 当前节点的后继写入新节点
		node.Next = newNode
		// 将新节点变为当前节点
		node = newNode
		// 继续遍历当前节点
	}
}

// 打印双向链表
func (node *TwoWayLinkNode) Print() {
	if node == nil {
		return
	}
	if node.Data != nil {
		fmt.Println(node.Data)
	}
	node.Next.Print()
}

// 打印双向链表：反向
func (node *TwoWayLinkNode) ReverPrint() {
	if node == nil {
		return
	}
	// 指向链表末尾
	for node.Next != nil {
		node = node.Next
	}
	for node.Prev != nil {
		if node.Data != nil {
			fmt.Println(node.Data)
		}
		node = node.Prev
	}
}

// 返回俩表长度
func (node *TwoWayLinkNode) Length() int {
	if node == nil {
		return -1
	}
	i := 0
	for node.Next != nil {
		i++
		node = node.Next
	}
	return i
}

// 插入数据，按照位置插入
func (node *TwoWayLinkNode) InsertByIndex(index int, data interface{}) {
	if node == nil || index < 0 || data == nil || index > node.Length() {
		return
	}
	// 先生命一个变量用来保存当前节点的上一节点
	preNode := node
	// 把指针指向指定位置的节点
	for i := 0; i < index; i++ {
		// 随着下标增长，当前节点，上一节点的信息也要更新
		preNode = node
		node = node.Next
	}
	// 创建新节点
	newNode := &TwoWayLinkNode{preNode, node, data}
	// 把新节点与前后节点进行缝合
	// 成为上一节点的下一节点
	preNode.Next = newNode
	// 成为当前节点的上一节点
	node.Prev = newNode
	// 缝合结束
}

// 删除数据，按照位置删除
func (node *TwoWayLinkNode) DeleteByIndex(index int) {
	if node == nil || index < 0 || index > node.Length() {
		return
	}
	// 先生命一个变量用来保存当前节点的上一节点
	preNode := node
	// 把指针指向指定位置的节点
	for i := 0; i < index; i++ {
		// 随着下标增长，当前节点，上一节点的信息也要更新
		preNode = node
		node = node.Next
	}
	// 把当前节点的下一节点，作为上一节点的下一节点
	preNode.Next = node.Next
	// 把当前节点的上一节点，作为下一节点的上一节点
	node.Next.Prev = node.Prev
	// 释放内存
	node.Prev = nil
	node.Next = nil
	node.Data = nil
	node = nil

}

// 双向链表的销毁

func (node *TwoWayLinkNode) Destory() {
	if node == nil {
		return
	}
	node.Next.Destory()
	node.Prev = nil
	node.Next = nil
	node.Data = nil
	node = nil
}
func main() {
	list2 := new(TwoWayLinkNode)
	list2.Create(1, 2, 3, 4)
	list2.Print()
	list2.ReverPrint()
	count := list2.Length()
	fmt.Println("长度为", count)
	list2.InsertByIndex(2, 0)
	list2.Print()
	// list2.ReverPrint()
	fmt.Println("删除第3个节点")
	list2.DeleteByIndex(3)
	list2.Print()
	list2.Destory()
	fmt.Println("已销毁")
	list2.Print()
}
