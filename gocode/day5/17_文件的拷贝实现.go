package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

//获取文件名
//只读打开文件
//读取内容
//新建文件并写入，边读边写
//读完写完
//关闭文件
func GetArgs() (sstr string, tstr string, err error) {
	strs := os.Args
	if len(strs) != 3 { //go run xxx.go xxx1 xxx2  run之后的都算
		//err = error.Error("请输入正确格式:XXX XXX")
		//生成一个错误类型,需要调用errors包，而非使用类型名调用方法
		err = errors.New("请输入正确格式:XXX XXX")
		return
	} else if strs[1] == strs[2] {

		err = errors.New("源文件名与目标文件名不能相同")
		return
	} else {
		sstr = strs[1]
		tstr = strs[2]
		return
	}
}
func CopyFile(sstr string, tstr string) {
	//打开文件
	f1, err1 := os.Open(sstr)
	if err1 != nil {
		fmt.Println("err=", err1)
		return
	}
	defer f1.Close()
	//创建目标文件
	f2, err2 := os.Create(tstr)
	if err2 != nil {
		fmt.Println("err=", err2)
		return
	}
	defer f2.Close()

	//核心内容处理，读多少写多少

	byteslice := make([]byte, 1024*4) //创建一个4K大小的临时缓冲区
	for {
		n, err := f1.Read(byteslice)
		if err != nil {
			if err == io.EOF {
				break //跳出循环
				fmt.Println("读取完毕")
			} else {
				fmt.Println("err=", err)
				return
			}
		}
		f2.Write(byteslice[:n])
		//不用处理err等返回值.必须有：n否则会写入很多.....................
	}

}
func main() {
	s, t, err := GetArgs()
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	CopyFile(s, t)

}
