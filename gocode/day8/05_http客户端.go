package main

import (
	"fmt"
	"net/http"
)

func main() {
	//r, err := http.Get("http://www.baidu.com")
	r, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(r.Status)
	fmt.Println(r.Header)
	//fmt.Println(r.Body)不能直接打印
	buf := make([]byte, 1024)
	defer r.Body.Close()
	var temp string
	for {
		n, err1 := r.Body.Read(buf)
		if n == 0 {
			fmt.Println("err1:", err1)
			break
		}
		temp += string(buf[:n])
	}

	fmt.Println(temp)
}
