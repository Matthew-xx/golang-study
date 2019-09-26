package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//r, err := http.Get("http://www.baidu.com")
	r, err := http.Get("http://127.0.0.1:8000/go")
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
		n, err := r.Body.Read(buf)
		temp += string(buf[:n])
		if err != nil {
			if err == io.EOF {
				fmt.Println("finished", n)
				break
			} else {
				fmt.Println("err=", err)
				break
			}
		}

	}

	fmt.Println(temp)
}
