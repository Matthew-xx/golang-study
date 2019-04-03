package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		//子协程执行前打印缓冲区的情况
		fmt.Printf("管道的剩余数据个数：%d,容量%d\n", len(ch), cap(ch))
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Printf("子协程放入数据[%d]——>管道，已有数据个数：%d,容量%d\n", i, len(ch), cap(ch))
		}
	}()
	time.Sleep(time.Second * 2)
	for i := 1; i <= 10; i++ {
		num := <-ch
		fmt.Printf("主协程取出数据[%d]管道<——，剩余数据个数：%d,容量%d\n", num, len(ch), cap(ch))

	}
}

//运行结果有若干种，总结来看。当管道满了之后，才会彻底阻塞写入方协程。
//管道空了之后，才会彻底阻塞读出方协程。
//不少情况是边写入边读出。也就是写满之后，立马又能写1个了。读出之后，立马又能读1个了。
//另外由于printf是在管道动作后发生的，可能会出现 貌似是没放的数据就显示被取出来了。实际上并没有错误
//因为确实是先放进去在取出的了，只是print(放入)本身被随机到后面执行的
