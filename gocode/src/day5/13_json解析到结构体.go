package main

//1定义一个结构体
//2.解析json到结构体
import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string   `json:"company"` //结构体必须大写，但又对不上，可使用 struct-tag
	Subject []string `json:"subject"`
	Isok    bool     `json:"isok"`
	Price   float32  `json:"price"`
}

func main() {
	jsonbuf := `{
        "price": 324.34,
        "company": "黑马",
        "isok": true,
        "subject": [
                "GO",
                "C++",
                "Python",
                "C"
        ]
	}`
	//var it IT
	it := new(IT)
	fmt.Printf("it type is %T\n", it)
	//结构体不是引用类型，其他函数修改其内容时，需要取指针
	//err := json.Unmarshal(byte(jsonbuf), it)
	//并不是将字符串转换为字节，而是字节切片
	err := json.Unmarshal([]byte(jsonbuf), it) //传递指针进去，而且必须是指针
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", *it) //取指针内容
}
