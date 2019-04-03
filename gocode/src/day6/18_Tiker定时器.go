package main

import (
	"fmt"
	"time"
)

func main() {
	//跟timer类似
	i := 0
	ticker := time.NewTicker(time.Second)
	for {
		i++
		<-ticker.C
		fmt.Println("i=", i)
		if i == 5 {
			ticker.Stop() //停止计时并跳出循环
			break
		}
	}
}
