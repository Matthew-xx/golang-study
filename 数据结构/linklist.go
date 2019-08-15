package main

import (
	"fmt"
	"reflect"
)

// 链表的节点结构体，嵌套本身（结构体）的指针，实现链表
type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

func (node *LinkNode) Create(data ...interface{}) {
	// 校验数据
	if node == nil {
		return
	}
	if len(data) == 0 {
		return
	}
	// 将当前节点作为最后一个节点
	lastNode := node
	// 有多少个数据，就生成多少个新节点
	for i := 0; i < len(data); i++ {
		// 创建新节点
		newNode := new(LinkNode)
		// newNode := &LinkNode{
		// 	data[i],
		// 	nil,
		// }
		// 给新节点存入数据
		newNode.Data = data[i]
		// 把新节点的地址，存入上一节点中
		lastNode.Next = newNode
		// 新节点成为上一个节点
		lastNode = newNode
		// 开始下一次循环
	}
}

// 打印全部节点
func (node *LinkNode) Print() {
	// 校验数据,同时也是递归的结束条件
	if node == nil {
		return
	}
	// 打印数据
	if node.Data != nil {
		fmt.Println(node.Data)
	}
	node.Next.Print()
}

// 链表的长度
func (node *LinkNode) Length() int {
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

// 插入数据（头插）
func (node *LinkNode) InsertByHead(data interface{}) {
	if node == nil || data == nil {
		return
	}
	newNode := &LinkNode{data, nil}
	// 无需考虑后继为空的情况，此种情况node.Next==nil不影响
	newNode.Next = node.Next
	node.Next = newNode
}

// 插入数据（尾插）
func (node *LinkNode) InsertByTail(data interface{}) {
	if node == nil || data == nil {
		return
	}
	// 创建新节点
	newNode := &LinkNode{data, nil}
	// 找到尾巴,并插入
	for node.Next != nil {
		node = node.Next
	}
	// 插入新节点
	node.Next = newNode
}

// 插入数据，按照位置（中间插）
func (node *LinkNode) InsertByIndex(index int, data interface{}) {
	// 校验数据，不能为空，下标不能为负数，最大不能超过链表的长度-1
	if node == nil || data == nil || index < 0 || index > node.Length()-1 {
		return
	}
	// 将next指针指向对应的节点
	for i := 0; i < index; i++ {
		node = node.Next
	}
	// 创建新节点,新节点接受指针后续节点
	newNode := &LinkNode{data, node.Next}
	// 将新节点，成为原来指针的后继节点
	node.Next = newNode
}

// 删除节点，基于位置删除
func (node *LinkNode) DeleteByIndex(index int) {
	if node == nil || index < 0 || node.Length() < index+1 {
		return
	}
	// 生命一个变量，存储上个节点，初始内容为节点本身
	preNode := node
	// 将next指针指向对应的节点
	for i := 0; i < index; i++ {
		// 移动指针之前，先保存node，作为上一个节点
		preNode = node
		// 移动指针，node成为当前节点
		node = node.Next
	}
	// 开始删除：把下一个节点，直接成为自己上个节点的下一节点
	preNode.Next = node.Next
	// 释放内存
	node.Data = nil
	node.Next = nil
	node = nil
}

// 删除包含某数据的节点
func (node *LinkNode) DeleteByData(data interface{}) {
	if node == nil || data == nil {
		return
	}
	// 只要删除，就需要记录上一个节点
	preNode := node
	// 遍历节点数据，进行比对，原理：下一个节点不为空则比较下一个节点
	for node.Next != nil {
		// 把遍历过的当前节点保存为上一节点
		preNode = node
		// 把下一节点保存为当前节点，开始比较
		node = node.Next
		// 比对数据时，需要比对两点：类型和内容，必须完全相同
		if reflect.TypeOf(data) == reflect.TypeOf(node.Data) && node.Data == data {
			preNode.Next = node.Next
			node.Data = nil
			node.Next = nil
			node = nil
			// return
			// 因为已经把当前节点删除了，新的当前节点还没有比对过，需要重新进行本次遍历，需要上一个节点重新成为当前节点
			node = preNode
		}
		// 没有删除的话，则正常往下边走
	}
}

// 查找数据
func (node *LinkNode) Search(data interface{}) int {
	if node == nil || data == nil {
		return -1
	}
	i := 0
	for node.Next != nil {
		i++
		node = node.Next
		if reflect.TypeOf(data) == reflect.TypeOf(node.Data) && node.Data == data {
			return i
		}
	}
	return -1
}

// 销毁链表
func (node *LinkNode) Destory() {
	if node == nil {
		return
	}
	node.Next.Destory()
	node.Data = nil
	node.Next = nil
	node = nil
}

// 测试链表的数据
type Student struct {
	id   int
	name string
	sex  string
	age  int
}

func main() {
	s1 := Student{1001, "张三", "男", 18}
	s2 := Student{1007, "李四", "男", 20}
	s3 := Student{1034, "王花", "女", 21}
	s4 := Student{1055, "孙里", "男", 23}
	s5 := Student{1054, "张红", "女", 23}
	s6 := Student{1054, "赵刘", "男", 33}
	s7 := Student{1023, "李子", "男", 31}
	head := new(LinkNode)
	head.Create(s1, s2, s3, s4)
	head.Print()
	count := head.Length()
	fmt.Println(count)
	head.InsertByHead(s5)
	fmt.Println("头插：")
	head.Print()
	head.InsertByTail(s6)
	fmt.Println("尾插")
	head.Print()
	head.InsertByIndex(2, s7)
	fmt.Println("下标插")
	head.Print()
	fmt.Println("删除2位置的数据")
	head.DeleteByIndex(2)
	head.Print()
	fmt.Println("删除张红，王花，赵刘")
	head.DeleteByData(s5)
	head.DeleteByData(s3)
	head.DeleteByData(s6)
	head.Print()
	index := head.Search(s2)
	fmt.Println("李四的位置是：", index)
	head.Destory()
	fmt.Println("已销毁")
	head.Print()
}
