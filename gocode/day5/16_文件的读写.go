package main

//打开文件 新建文件
//读写文件
//关闭文件
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	f, err := os.Create(path)
	//Creat()根据文件名新建文件=openFile(name,读写|创建|截断，0666 )
	//NewFile()根据文件描述符创建文件，返回一个文件对象
	//Open()根据文件名以只读形式打开文件 =openFile(name,只读，0| )
	//OpenFile()打开或创建文件，最灵活的函数，设置3个参数，文件名、打开方式、权限
	if err != nil {
		fmt.Println("err=", err)
	}
	defer f.Close()
	var buf string
	for i := 0; i < 10; i++ { //循环条件之间用分号；
		buf = fmt.Sprintf("i=%d\n", i)
		_, err1 := f.WriteString(buf)
		if err1 != nil { //不应放在循环外面
			fmt.Println("err1=", err1)
		}
	}

}

//正常无缓冲一个字节一个字节读取
func ReadFile(path string) {
	f, err := os.Open(path) //只读形式打开文件
	defer f.Close()
	if err != nil {
		fmt.Println("err=", err)
	}
	//var b []byte 错误，创建一个字节切片存储内容不是这么创建的
	b := make([]byte, 5) //正确方法。同时指定合适的大小
	for {
		n, err := f.Read(b)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				return
			} else {
				fmt.Println("err=", err)
				return
			}
		}
		fmt.Println("本次读取字符数：", n)
		fmt.Println(string(b[:n])) //注意用了多少，取多少。否则会有一堆空
	}
}

//一行一行读取，需用到带缓冲的IO 即bufio
func ReadFileLine(path string) {
	f, err1 := os.Open(path) //只读形式打开文件
	defer f.Close()
	if err1 != nil {
		fmt.Println("err1=", err1)
	}
	//调用bufio的newreader函数
	//参数io.reader,为IO包中的reader接口，找一个实现read()方法的接收者类型即可，很明显是os.File类型
	//返回一个bufreader接收者类型指针
	br := bufio.NewReader(f)
	b := make([]byte, 1024*4) //正确方法。同时指定合适的大小
	var err error
	//b, err = br.ReadBytes("\n")不要用双引号
	for {
		b, err = br.ReadBytes('\n')
		//参数为分割符，在这里就是结束符，结束读取，故需要for循环循环读取
		//遇到'\n'结束读取，但是把它也读进去了
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				//return读完了不要结束，否则后面不执行了
				break //终端循环即可
			} else {
				fmt.Println("err=", err)
			}

		}
		fmt.Print(string(b)) //读取一行，使用一行
	}

}
func main() {

	//WriteFile("./demo.txt")
	// ReadFile("./demo.txt")
	ReadFileLine("./demo.txt")
}
