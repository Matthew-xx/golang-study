package main

//1定义一个map
//2.解析json到map
import (
	"encoding/json"
	"fmt"
)

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
	m := make(map[string]interface{})          //设置长度
	err := json.Unmarshal([]byte(jsonbuf), &m) //虽然map为引用类型，但是还是需要取指针的，可能是源代码函数定义时导致的
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", m)
	str := m["company"]

	fmt.Printf("str type is %T\n", str)
	fmt.Println(str)
}
