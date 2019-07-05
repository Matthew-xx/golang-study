package main

import (
	"fmt"
)

func Producer(ch chan<- int) { //一定要记得加上管道的数据类型
	for i := 0; i < 10; i++ {
		ch <- i * i
	}
	close(ch) //别忘了关闭管道
	return
}
func Consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println("num=", num)
	}
	return
}

func main() {
	//创建一个双向channel
	ch := make(chan int)
	//生产者生产数字，写入chnanel
	go Producer(ch)
	//消费者消费数字，从channel读出
	Consumer(ch)
	// //go Consumer(ch)没必要单独起协程，放主协程
	// for {
	// }
}
