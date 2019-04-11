//fibonacci数列：1,1,2,3,5,8,13
package main

import (
	"fmt"
	"time"
)

//注意管道方向，前者为只可往管道发送数据，后者为只可从管道接收数据
func fibonacci(numch chan<- int64, ch <-chan bool) {
	x, y := int64(1), int64(1)
	for {
		select {
		case numch <- x: //注意冒号位置在最后面
			x, y = y, x+y			close(numch)
			//	close(ch) err:cannot close receive-only channel,不能关闭只可接收数据的管道
			return
		}
	}

}
func main() {
	//创建数字通道
	numch := make(chan int64)

	//创建传输命令的通道
	ch := make(chan bool)

	//生成数字的函数必须放前面，因为此程序模式为子协程生成数字，主协程打印。而非主协程生成，子协程打印
	go fibonacci(numch, ch)

	//主协程负责打印，有数字就接收，不想接收了就发送结束信号
	func() {
		for i := 0; i < 10; i++ { //循环10次，即打印到第10个
			fmt.Println(<-numch)
			time.Sleep(time.Second / 5)
		}
		//发送结束信号
		ch <- true //不一定非得是true只要写进去内容即可

	}()

}
