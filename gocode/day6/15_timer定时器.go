package main

import (
	"fmt"
	"time"
)

//Timer是个结构体类型，包含了以个只可发送时间的channel类型字段C
func main01() {
	//1.创建一个该结构体类型的变量.该变量X时间后，向变量.C里写入那时的时间。不到时间，不写入。
	timer := time.NewTimer(time.Second * 2)
	//当前时间：
	fmt.Println("time now:", time.Now())
	t := <-timer.C //不到时间，立马没东西，阻塞。
	fmt.Println("time then:", t)
}

//验证只会写入一次，就关闭channel
func main() {

	timer := time.NewTimer(time.Second * 2)
	for {
		t := <-timer.C
		fmt.Println("time then:", t)
	}

}

//执行结果：
//写入一次后死锁
