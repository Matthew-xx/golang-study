# HTTP编程

HTTP（HyperText Transfer Protocol，超文本传输协议）是互联网上应用最为广泛的一种网络协议，定义了客户端和服务端之间请求与响应的传输标准。

Go 语言标准库内建提供了 net/http 包，涵盖了 HTTP 客户端和服务端的具体实现。使用net/http 包，我们可以很方便地编写 HTTP 客户端或服务端的程序。

## http默认客户端

http包提供了一个默认客户端供用户直接使用，该客户端包含以下方法

### http.Get()

    http.Get()用于请求一个资源，（等价于http.DefaultClient.Get()

即可，示例代码如下：

    resp, err := http.Get("http://example.com/")
    if err != nil { //
        处理错误 ...
        return
    }
    defer resp.Body.close()
    io.Copy(os.Stdout, resp.Body)

### http.Post()

以 POST 的方式发送数据

    http.Post(url string, contentType string, body io.Reader)

三个参数分别为：

1. 请求的目标 URL
2. 将要 POST 数据的资源类型（MIMEType）
3. 数据的比特流（[]byte 形式）

上传一张图片：

    resp, err := http.Post("http://example.com/upload", "image/jpeg", &imageDataBuf)
    if err != nil
    { // 处理错误
        return
    }
    if resp.StatusCode != http.StatusOK {
    // 处理错误
        return
    }

### http.PostForm()

提交表单。实现了标准编码格式为 application/x-www-form-urlencoded的表单提交。

示例：模拟 HTML 表单提交一篇新文章

    resp, err := http.PostForm("http://example.com/posts", url.Values{"title":
    {"article title"}, "content": {"article body"}})
    if err != nil
    { // 处理错误
        return
    }

### http.Head()

只请求目标URL的头部信息，即 HTTP Header 而不返回 HTTPBody。可以认为是简化版的http.Get()。

示例：请求一个网站首页的 HTTP Header 信息：

    resp, err := http.Head("http://example.com/")

### (*http.Client).Do()

在多数情况下，http.Get()和 http.PostForm() 就可以满足需求，但是如果我们发起的HTTP 请求需要更多的定制信息，我们希望设定一些自定义的 Http Header 字段，比如：
 设定自定义的"User-Agent"，而不是默认的 "Go http package"
 传递 Cookie
此时可以使用 net/http 包 http.Client 对象的 Do()方法来实现：

    req, err := http.NewRequest("GET", "http://example.com", nil)
    // ...
    req.Header.Add("User-Agent", "Gobook Custom User-Agent")
    // ...
    client := &http.Client{ //... }
    resp, err := client.Do(req)

## http客户端示例

    package main

    import (
        "fmt"
        "net/http"
    )

    func main() {

        r, err := http.Get("http://127.0.0.1:8000")
        if err != nil {
            fmt.Println("err:", err)
        }
        fmt.Println(r.Status)
        fmt.Println(r.Header)
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

## http 服务端

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
        //注册处理函数
        http.HandleFunc("/", HandConn)
        //监听绑定
        err := http.ListenAndServe(":8000", nil)
        if err != nil {
            return
        }
        req, err := http.NewRequest("GET", "http://example.com", nil)
        // ...
        req.Header.Add("User-Agent", "Gobook Custom User-Agent")
        // ...
        client := &http.Client{ //... }
        resp, err := client.Do(req)
    }
