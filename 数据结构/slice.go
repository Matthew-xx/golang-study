package main

/*
#include <stdlib.h>
*/
import "C"

// 以上格式必须严格遵守，不能有空行，最好单独impor，否则报错！
import (
	"fmt"
	"unsafe"
)

// 指针偏移量
const TAG = 8

type Slice struct {
	// Data *interface{} ,不使用interface,不方便计算，而且需要断言或者反射
	Data unsafe.Pointer //万能指针类型，对应void*(C语言)
	len  int            //有效长度
	cap  int            //有效容量
}

func (s *Slice) Create(l int, c int, Data ...int) {
	// 如果数据为空，则返回
	if len(Data) == 0 {
		return
	}
	// 如果长度或者容量小于零，或者数据个数超过了长度，则返回
	if l < 0 || c < 0 || l > c || len(Data) > l {
		return
	}
	// 通过C语言开辟空间，存储数据。注意将go语言的int变量转换成C语言的变量再使用
	// ulang:无符号长整型
	s.Data = C.malloc(C.ulong(c) * 8)
	s.cap = c
	s.len = l
	// 切片中的数据指针，转成可计算的指针类型，用来位移指针
	p := uintptr(s.Data)
	// 开始存储数据,注意这里的data是传递进来的一组数据，而非结构体中的data
	for _, v := range Data {
		// 给当前指针所指位置，存入数据。
		// 1.转回万能指针，
		// 2.强转为int指针
		// 3.找到int指针内容
		*(*int)(unsafe.Pointer(p)) = v
		// 指针偏移
		p += TAG
	}
}

// 打印切片
func (s *Slice) Print() {
	if s == nil {
		return
	}
	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		fmt.Print(*(*int)(unsafe.Pointer(p)), "")
		p += TAG
	}
}

// 切片追加
func (s *Slice) Append(data ...int) {
	// 校验数据
	if len(data) == 0 {
		return
	}
	// 若数据长度超出可存储范围，则开辟更多的内存,规则为2倍扩容，不够再次扩容
	// 1.确定扩容的倍数
	mul := 1
	for s.cap*mul < s.len+len(data) {
		mul *= 2
	}
	// 2.如果倍数>1则扩容
	if mul > 1 {
		fmt.Println("需要扩容，倍数：", mul)
		s.Data = C.realloc(s.Data, C.ulong(s.cap*8*mul))
		// 扩容完，别忘了更新cap值
		s.cap *= mul
	}
	// 开始写入数据
	// 1.先偏移指针到尾部
	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		p += TAG
	}
	// 2.插入数据
	for _, v := range data {
		*(*int)(unsafe.Pointer(p)) = v
		fmt.Println("正在插入数据：", v)
		// 偏移指针
		p += TAG
	}
	// 别忘了更新len
	s.len = s.len + len(data)
	fmt.Println("插入完成")
}

// 按照索引获取值
func (s *Slice) GetData(index int) int {
	// 校验数据index
	if index < 0 || index > s.len-1 {
		return -1
	}
	// 校验slice
	if s == nil || s.Data == nil {
		return -1
	}
	// 开始获取对应的数据
	// 1.偏移指针到对应位置
	p := uintptr(s.Data)
	for i := 0; i < index; i++ {
		p += TAG
	}
	// 2.得到元素值
	return *(*int)(unsafe.Pointer(p))

}

// 按照值获取索引
func (s *Slice) Search(data int) int {
	// 校验slice
	if s == nil || s.Data == nil {
		return -1
	}
	// 便利切片，通过指针位移实现遍历
	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		// 查找数据
		if *(*int)(unsafe.Pointer(p)) == data {
			return i
		}
		// 指针偏移，如果没有找到，需要找下一次
		p += TAG
	}
	return -1
}

// 按照索引位置删除元素
func (s *Slice) Delete(index int) {
	// 校验slice
	if s == nil || s.Data == nil {
		return
	}
	// 校验数据index
	if index < 0 || index > s.len-1 {
		return
	}
	// 指针偏移到要删除的位置
	p := uintptr(s.Data)
	for i := 0; i < index; i++ {
		p += TAG
	}
	// 删除的原理：所有后面的元素依次往前移一位，覆盖前者
	// 循环覆盖
	for i := index; i < s.len; i++ {
		// p现在指向被覆盖位置，创建一个临时指针，指向覆盖者的位置
		temp := p + TAG
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(temp))
		// 偏移p
		p += TAG
	}
	// 别忘了len-1
	s.len--
}

// 按照索引位置插入元素
func (s *Slice) Insert(index int, data int) {
	// 校验数据index
	if index < 0 || index > s.len-1 {
		return
	}
	/// 校验slice
	if s == nil || s.Data == nil {
		return
	}
	// 特殊处理在末尾插入的位置
	if index == s.len-1 {
		s.Append(data)
		return
	}
	// 指针偏移到要插入的位置
	p := uintptr(s.Data)
	p += uintptr(TAG * index)
	// 获取末尾的后面一个指针位置
	temp := uintptr(s.Data)
	temp += uintptr(TAG * s.len)
	// 从最后一个元素开始依次往后移动一位
	for i := s.len; i > index; i-- {
		*(*int)(unsafe.Pointer(temp)) = *(*int)(unsafe.Pointer(temp - TAG))
		temp -= TAG
	}
	// 将数据插入
	*(*int)(unsafe.Pointer(p)) = data
	// 增加长度
	s.len++

}

// 销毁切片
func (s *Slice) Destory() {
	C.free(s.Data)
	s.Data = nil
	s.len = 0
	s.cap = 0
	s = nil
}
func main() {
	var s Slice
	s.Create(5, 5, 1, 2, 3, 4, 5)
	fmt.Println(s)
	s.Print()
	s.Append(6, 7, 8, 9, 1, 1, 2, 3)
	s.Print()
	fmt.Println("获取第3个元素数据：")
	fmt.Println(s.GetData(3))
	fmt.Println("获取值为4的元素索引位置：")
	fmt.Println(s.Search(4))
	s.Delete(5)
	fmt.Println("删除第5个位置的元素")
	s.Print()
	fmt.Println("第12个位置插入666")
	s.Insert(11, 666)
	s.Print()
	s.Destory()
	s.Print()
}

// func main01() {
// 	var i int
// 	fmt.Println(unsafe.Sizeof(i))
// 	// int占8个字节
// 	var s []int
// 	fmt.Println(unsafe.Sizeof(s))
// 	// slice占24个字节，原因：需要保存3个信息：1.数据指针2.length3.容量cap
// }
