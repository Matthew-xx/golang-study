package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var C chan bool

func SpiderPage(i int, pagech chan<- int) {
	const baseurl string = "http://tieba.baidu.com/f?kw=java&ie=utf-8&pn="
	url := baseurl + strconv.Itoa((i-1)*50)
	fmt.Println(url)

	//爬取内容,自己封装一个函数
	str, err := HttpGet(url)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	//把内容写入到文件
	f, err1 := os.Create(strconv.Itoa(i) + ".html")
	if err != nil {
		fmt.Println("err1:", err1)
		return
	}
	f.WriteString(str)
	defer f.Close()
	pagech <- i

}

func HttpGet(url string) (result string, err error) {
	r, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer r.Body.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := r.Body.Read(buf)
		if n == 0 { //读取结束或出问题
			//err = err2
			break
		}
		result += string(buf[:n])
	}
	return

}

func DoWork(start, end int) {
	fmt.Printf("正在爬取%d到%d的页面\n", start, end)
	//确定爬取目标范围

	pagech := make(chan int)
	for i := start; i <= end; i++ {
		//再次封装一个
		go SpiderPage(i, pagech) //没有管道，直接加协程会导致被主协程关闭
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d页面爬取完成\n", <-pagech)
	}
}

func main() {
	var start, end int
	fmt.Println("请输入起始页：（>=1）")
	fmt.Scan(&start)
	fmt.Println("请输入结束页：（>=起始页）")
	fmt.Scan(&end)
	DoWork(start, end)
}
