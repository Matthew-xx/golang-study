package main

import (
	"fmt"
	"regexp"
)

func main() {
	//提取有效小数
	buf := "3.1415926 asfds 4.45646 9.  .7  93.3 fgrewof fewf.4"
	////1.解释规则
	reg := regexp.MustCompile(`\d+.\d+`)
	s := reg.FindAllString(buf, -1)
	fmt.Println(s)
}
