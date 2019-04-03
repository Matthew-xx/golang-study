package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "mike"
	i[2] = Student{"mike", 666}
	//类型查询，类型断言
	for index, data := range i {
		switch value := data.(type) { //可能是一种特殊用法 value为断言出来的内容本身，type是断言类型
		case int:
			fmt.Printf("i[%d]类型为int,内容为%d\n", index, value)
		case string:
			fmt.Printf("i[%d]类型为string,内容为%s\n", index, value)
		case Student:
			fmt.Printf("i[%d]类型为Student结构体,内容为name=%s,id=%d\n", index, value.name, value.id)
		}
	}
}
