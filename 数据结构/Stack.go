package main

import (
	"fmt"
)

type StackNode struct {
	Data interface{}
	Next *StackNode
}

func CreateStack(data ...interface{}) *StackNode {
	if len(data) == 0 {
		return nil
	}
	// 专门存储栈顶的指针,默认为空节点
	s := new(StackNode)
	// 当前节点的下一个节点指针，指代那些已经入栈的节点,默认为空
	var nextNode *StackNode = nil
	// 遍历需要存储的数据组
	for _, v := range data {
		// 创建存储数据的节点，有几个数据，创建几个
		newNode := new(StackNode)
		newNode.Data = v
		// 将新的节点，放在栈顶位置，成为新的栈顶
		s = newNode
		// 判断当前节点是不是入栈的第一个节点，只需判断当前节点的下一节点是不是空
		// 如果不是，则把已经入栈的节点则放到新栈顶的下一个位置。
		// 如果是，则无需处理
		if nextNode != nil {
			s.Next = nextNode
		}
		// 当前节点入栈完成，会有新的节点过来，当前节点提前设置成未来节点的下一节点，方便后续操作
		nextNode = s
	}
	return s
}
func PrintStack(s *StackNode) {
	if s == nil {
		return
	}
	for s != nil {
		fmt.Print(s.Data, " ")
		s = s.Next
	}
}

// 链栈的个数
func LengthStack(s *StackNode) int {
	if s == nil {
		return -1
	}
	i := 0
	for s != nil {
		i++
		s = s.Next
	}
	return i
}

// 入栈
func Push(s *StackNode, data interface{}) *StackNode {
	// 数据校验
	if s == nil || data == nil {
		return s
	}
	// 创建新节点，存储要入栈的数据，原栈顶成为他的下一个节点
	newNode := &StackNode{data, s}
	return newNode
}

// 出栈
func Pop(s *StackNode) *StackNode {
	if s == nil {
		return s
	}
	s = s.Next
	return s
}

// 清空链栈
func Clear(s *StackNode) *StackNode {
	return nil
}
func main() {
	stack := CreateStack("a", "b", "c")
	PrintStack(stack)
	fmt.Println("链栈的长度为", LengthStack(stack))
	stack = Push(stack, "d")
	PrintStack(stack)
	fmt.Println("链栈的长度为", LengthStack(stack))
	stack = Pop(stack)
	fmt.Println("出栈1次：")
	PrintStack(stack)
	fmt.Println("清除链栈")
	stack = Clear(stack)
	PrintStack(stack)
}
