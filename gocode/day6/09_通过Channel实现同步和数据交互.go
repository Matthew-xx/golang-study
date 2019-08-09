package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	defer fmt.Println("主协程调用完毕")
	go func() {
		defer fmt.Println("子协程调用完毕")
		for i := 0; i < 3; i++ {
			fmt.Println("i=", i)
			time.Sleep(time.Second)
		}
		ch <- "牛逼呀"
	}()
	str := <-ch
	fmt.Println(str)
}
