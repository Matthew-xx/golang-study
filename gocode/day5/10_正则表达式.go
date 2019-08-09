package main

import (
	"fmt"
	"regexp"
)

func main() {
	//``原生字符串，代表字符串最原始的样子
	buf := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<title>Go语言标准库文档中文版 | Go语言中文网 | Golang中文社区 | Golang中国</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
	<meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
	<meta charset="utf-8">
	<link rel="shortcut icon" href="/static/img/go.ico">
	<link rel="apple-touch-icon" type="image/png" href="/static/img/logo2.png">
	<meta name="author" content="polaris <polaris@studygolang.com>">
	<meta name="keywords" content="中文, 文档, 标准库, Go语言,Golang,Go社区,Go中文社区,Golang中文社区,Go语言社区,Go语言学习,学习Go语言,Go语言学习园地,Golang 中国,Golang中国,Golang China, Go语言论坛, Go语言中文网">
	<meta name="description" content="Go语言文档中文版，Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
</head>
	<div>和蔼
	好</div>
	<div>哈哈</div>
	<div>试</div>
	<div>你过啦</div>
	<div>安抚</div>





<frameset cols="15,85">
	<frame src="/static/pkgdoc/i.html">
	<frame name="main" src="/static/pkgdoc/main.html" tppabs="main.html" >
	<noframes>
	</noframes>
</frameset>
</html>`
	// reg := regexp.MustCompile(`<div>.*</div>`) //如果有换行，会被排除在外
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	result := reg.FindAllStringSubmatch(buf, -1)
	//fmt.Println(result)
	//过滤<></>
	for _, r := range result {
		//fmt.Println(r[0]) //带标签的内容
		fmt.Println(r[1]) //去掉包裹后的内容
		fmt.Printf("r Type is %T", r)
		fmt.Printf("result Type is %T", result)
	}
	//总结：解析器出的结果是个字符串二维切片，使用range对其遍历，获得的每个元素分别为一个字符串切片
	//每隔字符串切片由带标签和不带标签两个元素构成。
}
