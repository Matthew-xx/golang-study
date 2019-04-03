package main

import (
	"encoding/json"
	"fmt"
)

// type IT struct {
// 	compant string
// 	subject []string
// 	isok    bool
// 	price   float32
// }
//成员变量名必须大写，否则其他包的函数无法操作他，继而无法实现转为json格式
type IT struct {
	Company string   `json:"company"` //此字段在json中的键用新名字
	Subject []string `json:"-"`       //此字段被忽略
	Isok    bool     `json:",string"` //次字段，视为string
	Price   float32
}

func main() {
	it := IT{"heima", []string{"Go", "C++", "Python", "Test"}, true, 5465.45}
	byte, err := json.Marshal(it)
	//byte, err := json.MarshalIndent(it, "￥", "##")
	//有缩进格式化编码,参数2为前缀符号（行首字符），参数3为切割符号(行首，个数代表层级)
	if err != nil {
		fmt.Println("err=", err)
	}
	//fmt.Println(byte)直接打印字节是数组，需要转化一下
	fmt.Println(string(byte))
}
