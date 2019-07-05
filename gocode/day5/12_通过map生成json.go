package main

import (
	"encoding/json"

	"fmt"
)

func main() {
	// m := map[string]interface{}{
	// 	company: "黑马",
	// 	subject: []string{"GO", "C++", "Python", "C"},
	// 	isok:    true, Price: 324.34}
	//犯错：把键当变量了
	m := map[string]interface{}{
		"company": "黑马",
		"subject": []string{"GO", "C++", "Python", "C"},
		"isok":    true,
		"Price":   324.34}
	// fmt.Println(m)
	b, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println(string(b))
}
