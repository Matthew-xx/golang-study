package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc fewbhn frew afc auc acc"
	//1.解释规则
	//2.使用规则

	//解析正则表达式
	reg := regexp.MustCompile(`a.b`) //使用反引号。
	reg := regexp.MustCompile(`a[0-9]b`)
	reg := regexp.MustCompile(`a\db`)
	//使用正则表达式对象的查找方法
	result := reg.FindAllStringSubmatch(buf, -1) //-1是所有
	fmt.Println(result)
}
