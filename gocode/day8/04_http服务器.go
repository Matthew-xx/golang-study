package main

import (
	"fmt"
	"net/http"
)

//w给客户端回复数据
//r读取客户端发送的数据
func HandConn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello go"))
	fmt.Println(r.URL)
	fmt.Println(r.Header)
	fmt.Println(r.Body)

}
func main() {
	//注册处理函数,其实就是路由，但不是restful路由
	http.HandleFunc("/go", HandConn)
	//监听绑定
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
	// req, err := http.NewRequest("GET", "http://example.com", nil)
	// // ...
	// req.Header.Add("User-Agent", "Gobook Custom User-Agent")
	// // ...
	// client := &http.Client{ //... }
	// resp, err := client.Do(req)
}
