package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) //无缓存

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("子协程放入数据[%d]——>管道，已有数据个数：%d,容量%d\n", i, len(ch), cap(ch))
		}
		//完成后关闭管道
		close(ch) //参数为可写入的单向管道，意味着只可读出的管道不能作为参数
	}()
	//更简单的写法：直接range ch 迭代（不是ch<-），可循环每次读出数据，并且管道被关闭后，自动识别

	// for num := range ch {
	// 	fmt.Printf("主协程取出数据[%d]管道<——，剩余数据个数：%d,容量%d\n", num, len(ch), cap(ch))
	// }

	for {
		//光关闭还不行，还要告诉主协程关闭了，从而让它执行别的动作
		if num, ok := <-ch; ok { //迭代管道，两个返回值：一个是取出的内容，另一个是关闭与否的判断
			fmt.Printf("主协程取出数据[%d]管道<——，剩余数据个数：%d,容量%d\n", num, len(ch), cap(ch))
		} else {
			a := <-ch
			fmt.Println("管道已关闭，已关闭的管道永远取出：", num, a)
			break
		}
	}
}

//管道关闭后，可以继续读
//子协程循环写入5次
//主协程无限次读取，当5次之后，无写无读，造成死锁。
//可以关闭管道，防止死锁
