# Socket编程

go语言可以基于网络协议TCP/IP等网络协议进行网络编程，所用的编程接口叫做socket接口。所以socket本身不是一种协议，而是书写TCP或者UTP协议的用户接口。

## 传统Socket编程步骤

1. 建立Socket：使用socket()函数
2. 绑定Socket:使用bind()
3. 监听：listen(),链接：connet()
4. 接受链接：accept()
5. 接收数据：receive();发送数据：send()

## Dial()实现GO的Socket

go语言对传统Socket编程步骤进行了抽象和封装，可使用net包中的dial函数

    func Dial(network, address string) (Conn, error)
其中参数network可以为：

    "tcp",  "tcp4" (IPv4-only), "tcp6" (IPv6-only)
    "udp",  "udp4" (IPv4-only), "udp6" (IPv6-only)
    "ip",   "ip4"  (IPv4-only), "ip6"  (IPv6-only)
    "unix", "unixgram"        ，"unixpacket".

对于TCP/UDP,地址参数格式为host:port。host可以为IP地址或域名。

对于其他类型协议，不做解释

## TCP客户端

    package main

    import (
        "fmt"
        "net"
    )

    func main() {
        //主动接入服务器
        con, err := net.Dial("tcp", ":2343") //dial函数连接
        if err != nil {
            fmt.Println("err=", err)
        }
        defer con.Close()

        //发送数据
        buf := []byte("are u ok?")
        con.Write(buf)
    }


## TCP服务器

    package main

    import (
        "fmt"
        "net"
    )

    func main() {

        //监听listen
        listener, err := net.Listen("tcp", ":2343")
        if err != nil {
            fmt.Printf("err=", err)
            return //出错就结束
        }

        //关监听
        defer listener.Close()

        //等待用户接入
        con, err := listener.Accept()
        if err != nil {
            fmt.Printf("err=", err)
            return
        }

        //关闭用户链接
        defer con.Close()

        //接收用户请求
        buf := make([]byte, 1024*4)
        n, err1 := con.Read(buf)
        if err1 != nil {
            fmt.Printf("err1=", err1)
            return
        }
        fmt.Println(string(buf[:n]))

    }
