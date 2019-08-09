package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	info, err := os.Stat(list[1])
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("文件名称：", info.Name())
	fmt.Println("文件大小：", info.Size())
}
