# 文件读写io与bufio

## 文件读写步骤

1. 打开文件，或新建文件
2. 读写文件
3. 关闭文件

## 打开文件

使用os包内的函数,返回一个文件对象的引用。

- OpenFile()
    打开或创建文件，最灵活的函数，设置3个参数，文件名、打开方式、权限
- Creat()
    根据文件名新建文件=openFile(name,读写|创建|截断，0666 )
- Open()
    根据文件名以只读形式打开文件 =openFile(name,只读，0| )
- NewFile()
    根据文件描述符创建文件，返回一个文件对象。

示例：

    f, err := os.Create(path)
    if err != nil {
        fmt.Println("err=", err)
    }
    defer f.Close()

## 写入文件

文件对象的引用是在os包中创建，使用该引用的方法集中的写入方法实现写入内容

    func (f *File) WriteString(s string) (n int, err error)

示例：
    var buf string
    for i := 0; i < 10; i++ { //循环条件之间用分号；
        buf = fmt.Sprintf("i=%d\n", i)
        _, err1 := f.WriteString(buf)
        if err1 != nil { //不应放在循环外面
            fmt.Println("err1=", err1)
        }
    }

## 读文件

首先同样需要打开文件才能读取，不赘述

### 无缓存字节流读取

    func (f *File) Read(b []byte) (n int, err error)

注意：

传入的字符切片len()有多长就会读多长，返回读到的字符数和可能遇到的error：

1. 当文件首次就读取就读取结束了，则返回读到的字符数和error=nill
2. 多次调用Read()才结束，结束后返回0和 io.EOF.

读取文件示例：

b := make([]byte, 1024*4) //正确方法。同时指定合适的大小

    n, err := f.Read(b)
    if err != nil {
        if err == io.EOF {
            fmt.Println("文件读取完毕")
            return
        } else {
            fmt.Println("err=", err)
        }

    }
    fmt.Println("n=", n)
    fmt.Println(string(b[:n]))
    //注意用了多少，取多少。否则会有一堆空

### 带缓存区的读取bufio

不带缓存的Read()是读到一个字节就输出一个字节，带缓存的NewReader()是先一个字节一个字节的读入缓存区，然后，在从缓存区输出多个字节。默认缓存区大小为4096的Reader
函数定义如下：

    func NewReader(rd io.Reader) *Reader

参数io.reader,为IO包中的reader接口，找一个实现read()方法的接收者类型即可，很明显是os.File类型

返回值是一个新的带缓存的Reader对象的引用。

示例：一行一行的读取

    func ReadFileLine(path string) {

        //只读形式打开文件
        f, err1 := os.Open(path)
        defer f.Close()
        if err1 != nil {
            fmt.Println("err1=", err1)
        }

        //调用bufio的newreader函数
        br := bufio.NewReader(f)

        //正确方法。同时指定合适的大小
        b := make([]byte, 1024*4)
        var err error
        for {
            b, err = br.ReadBytes('\n')
            //参数为分割符，在这里就是结束符，结束读取，故需要for循环循环读取
            //遇到'\n'结束读取，但是把它也读进去了
            if err != nil {
                if err == io.EOF {
                    fmt.Println("文件读取完毕")
                    //读完了不要return，否则后面不执行了,break即可
                    break
                } else {
                    fmt.Println("err=", err)
                }
            }
            fmt.Print(string(b)) //读取一行，使用一行
        }
    }

输出：

    i=0
    i=1
    i=2
    i=3
    i=4
    i=5
    i=6
    i=7
    i=8
    文件读取完毕