package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//封装一个爬取函数，输入网址，返回内容
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 { //读完或者出错
			//fmt.Println("resp err:", err2)
			//这样写永远报错,永远有读完的时候啊
			break //结束循环
		}
		result += string(buf[:n])
	}
	return //结束函数
}

//封装一个执行函数，输入页码，获取页码对应的主网页中我们需要的内容
func SpiderPage(i int, page chan<- int) {
	filetile := make([]string, 0)
	filecontent := make([]string, 0)
	//根据页码生成url
	url := "https://www.pengfue.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d个网页:%s\n", i, url)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("err httpget:", err)
		return
	}
	//fmt.Println(result)
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	//你解析不出来报人家的错干嘛
	//冤枉你了呢,你只是空而已
	if re == nil {
		fmt.Println("regexp err")
		return
	}
	//返回笑话详情的链接

	joyurls := re.FindAllStringSubmatch(result, -1)
	// fmt.Println(joyurls)
	//第1个是二维数组下标，第二个是二维数组下标对应的一个元素
	for _, data := range joyurls {
		//每个元素都有两个子元素，data[1]为我们需要的元素
		//开始爬取每一个段子
		title, content, err := SpiderOneJoy(data[1])
		if err != nil {
			fmt.Println("spideronejoy err:", err)
			continue
		}
		// fmt.Println(title)
		// fmt.Println(content)
		filetile = append(filetile, title)
		filecontent = append(filecontent, content)
	}
	StoreJoyToFile(i, filetile, filecontent)
	page <- i
}
func StoreJoyToFile(i int, filetile, filecontent []string) {
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	defer f.Close()
	if err != nil {
		fmt.Println("file err:", err)
		return
	}
	for i := 0; i < len(filetile); i++ {
		f.WriteString(filetile[i] + "\n")
		f.WriteString(filecontent[i])
		f.WriteString("\n###############\n")
	}
}

//封装一个函数，输入地址，返回标题和内容
func SpiderOneJoy(url string) (title string, content string, err error) {

	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}

	//取标题
	re1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re1 == nil {
		err2 := fmt.Errorf("regexp err")
		err = err2
		return
	}
	//只取第一个
	tmptile := re1.FindAllStringSubmatch(result, 1)

	//第1个是二维数组下标，第二个是二维数组下标对应的一个元素
	for _, data := range tmptile {
		title = data[1]
		//去掉标题的tab,另外为啥你老把反斜杠弄混
		title = strings.Replace(title, "\t", "", -1)
		title = strings.Replace(title, "\n", "", -1)
		title = strings.Replace(title, " ", "", -1)
		break
	}

	//取内容
	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if re2 == nil {
		err2 := fmt.Errorf("regexp err")
		err = err2
		return
	}
	//只取第一个
	tmpcontent := re2.FindAllStringSubmatch(result, 1)

	//第1个是二维数组下标，第二个是二维数组下标对应的一个元素
	for _, data := range tmpcontent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "nbsp;", "", -1)
		content = strings.Replace(content, "&<br />", "", -1)
		content = strings.Replace(content, "<br />", "", -1)
		break
	}
	return
}

//工作函数，执行爬取任务
func DoWork(start int, end int) {
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个链接已经爬取完成\n", <-page)
	}
}
func main() {
	//获取目标范围
	var start, end int
	fmt.Println("请输入起始页（>1）:")
	fmt.Scan(&start)
	fmt.Println("请输入结束页（>起始页）:")
	fmt.Scan(&end)
	fmt.Printf("准备获取从%d,到%d的网页\n", start, end)
	DoWork(start, end)
}
