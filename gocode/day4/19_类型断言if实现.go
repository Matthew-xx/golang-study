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
	//fmt.Println(i[0].(int))
	//类型查询，类型断言s
	for index, data := range i {
		if value, ok := data.(int); ok { //第一个返回断言出来的变量内容，第二个返回断言此类型结果真假
			fmt.Printf("i[%d]类型为int,内容为%d\n", index, value)
		} else if value, ok := data.(string); ok == true { //ok==true可简写为ok
			fmt.Printf("i[%d]类型为string,内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("i[%d]类型为Student结构体,内容为name=%s,id=%d\n", index, value.name, value.id)
		}
	}
}