package main

import (
	"fmt"
	"reflect"
)

type BinaryNode struct {
	data   interface{}
	lChild *BinaryNode
	rChild *BinaryNode //右子树
}

// 创建二叉树，仅供演示，不使用
func (node *BinaryNode) Create() {
	node = new(BinaryNode)
}

// 先序遍历：先根，再左，再右
func (node *BinaryNode) PreOrder() {
	// 数据校验，也是递归的结束条件
	if node == nil {
		return
	}
	fmt.Print(node.data, " ")
	node.lChild.PreOrder()
	node.rChild.PreOrder()
}

// 中序遍历：先左，再根，再右
func (node *BinaryNode) MidOrder() {
	if node == nil {
		return
	}
	node.lChild.MidOrder()
	fmt.Print(node.data, " ")
	node.rChild.MidOrder()
}

// 后序遍历
func (node *BinaryNode) RearOrder() {
	if node == nil {
		return
	}
	node.lChild.RearOrder()
	node.rChild.RearOrder()
	fmt.Print(node.data, " ")
}

// 高度
func (node *BinaryNode) TreeHeight() int {
	// 叶子节点的高度为0，也是递归的结束条件
	if node == nil {
		return 0
	}
	// 计算左子树高度
	lh := node.lChild.TreeHeight()
	// 计算右子树高度
	rh := node.rChild.TreeHeight()
	// 哪棵树高，返回哪颗
	if lh > rh {
		// 子树高+1才是整棵树的高度
		lh++
		return lh
	} else {
		rh++
		return rh
	}
}

// 叶子节点的数量,为了用递归实现，计数器由外部传入
func (node *BinaryNode) LeafCount(count *int) {
	if node == nil {
		return
	}
	// 找到叶子节点，找到1个加1次
	if node.lChild == nil && node.rChild == nil {
		*count++
	}
	node.lChild.LeafCount(count)
	node.rChild.LeafCount(count)
}

// 二叉树的查找
func (node *BinaryNode) Search(data interface{}) {
	if node == nil {
		return
	}
	if reflect.TypeOf(node.data) == reflect.TypeOf(data) && data == node.data {
		fmt.Println("找到数据", node.data)
		return
	}
	node.lChild.Search(data)
	node.rChild.Search(data)
}

// 二叉树的销毁
func (node *BinaryNode) Destory() {
	if node == nil {
		return
	}
	node.lChild.Destory()
	node.lChild = nil
	node.rChild.Destory()
	node.rChild = nil
	node.data = nil
}

// 反转二叉树：前提是满二叉树
func (node *BinaryNode) Reverse() {
	if node == nil {
		return
	}
	// 交换节点：利用go的多重赋值
	node.lChild, node.rChild = node.rChild, node.lChild
	node.lChild.Reverse()
	node.rChild.Reverse()
}

// 二叉树拷贝
func (node *BinaryNode) Copy() *BinaryNode {
	// 当节点为空时，停止递归
	if node == nil {
		return nil
	}
	//对当前节点的左、右子树，进行拷贝
	lChild := node.lChild.Copy()
	rChild := node.rChild.Copy()
	// 创建新节点
	newNode := new(BinaryNode)
	// 将当前的拷贝结果，存入新节点的对应位置
	newNode.data = node.data
	newNode.lChild = lChild
	newNode.rChild = rChild
	// 返回新节点
	return newNode
}

// 主函数
func main() {
	// 创建
	tree := new(BinaryNode)
	// 创建节点
	node1 := BinaryNode{1, nil, nil}
	node2 := BinaryNode{2, nil, nil}
	node3 := BinaryNode{3, nil, nil}
	node4 := BinaryNode{4, nil, nil}
	node5 := BinaryNode{5, nil, nil}
	node6 := BinaryNode{6, nil, nil}
	// node7 := BinaryNode{7, nil, nil}
	// 创建结构
	tree.data = 0
	tree.lChild = &node1
	tree.rChild = &node2
	node1.lChild = &node3
	node1.rChild = &node4
	node2.lChild = &node5
	node2.rChild = &node6
	// node3.lChild = &node7
	// 先序遍历
	tree.PreOrder()
	fmt.Println()

	// 中序遍历
	tree.MidOrder()
	fmt.Println()

	// 后续遍历
	tree.RearOrder()
	fmt.Println()

	fmt.Println("高度为", tree.TreeHeight())
	count := 0
	tree.LeafCount(&count)
	fmt.Println("叶子数量", count)

	// 查找数据
	tree.Search(6)
	// 销毁二叉树
	// tree.Destory()
	fmt.Println("翻转前：")
	tree.MidOrder()
	fmt.Println()
	tree.Reverse()
	fmt.Println("翻转后：")
	tree.MidOrder()
	fmt.Println()

	newTree := tree.Copy()
	// 原始树
	fmt.Println("原始树 ：")
	tree.PreOrder()
	fmt.Println()
	fmt.Println("拷贝树：")
	newTree.PreOrder()
	fmt.Println()

}
