package main

import (
	"fmt"
	"net/http"
	"os"

	"strconv"
)

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
	const baseurl string = "http://tieba.baidu.com/f?kw=java&ie=utf-8&pn="

	for i := start; i <= end; i++ {
		url := baseurl + strconv.Itoa((i-1)*50)
		fmt.Println(url)
		//爬取内容,自己封装一个函数
		str, err := HttpGet(url)
		if err != nil {
			fmt.Println("err", err)
			continue
		}
		//把内容写入到文件

		f, err1 := os.Create(strconv.Itoa(i) + ".html")
		if err != nil {
			fmt.Println("err1:", err1)
			continue
		}
		f.WriteString(str)
		defer f.Close()
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
