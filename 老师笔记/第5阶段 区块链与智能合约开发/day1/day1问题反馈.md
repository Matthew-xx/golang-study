

# 一、节奏快

问题:

讲课节奏稍快，思路不容易跟

有些重点，讲慢点，我大脑还没处理过来就过去了



答：

放慢节奏，适当停留5分钟给大家消化



# 二、没讲清楚

问题：

老师上课可以拖，晚上也可以讲，别急，希望可以把过程讲的细点，就像今天join函数那样讲，之前有个binary的函数没见过也没讲明白



答：

go[语言几种序列化方式对比](https://blog.csdn.net/fengfengdiandia/article/details/79986237)



## 1. binary序列化

### - binary.Write

```
err := binary.Write(&buffer, binary.BigEndian, num)
```

这个函数的目的是将任意的数据转换成byte字节流，这个过程叫做序列化

### - binary.Read

同样，可以通过binary.Read方式进行反序列化，从字节流转回原始结构。

```go
binary.Read(buf, binary.LittleEndian, &num)
```

### - 特点

特点：高效。

缺点：如果在编码的结构中有不确定长度的类型，那么会报错。这时候可以使用gob来编码。



## 2. 对其方式 

为了提高传输效率，要涉及到字节对齐的问题，也就是大小端对齐。

### - 大端对齐

### - 小段对齐

```go
package main

import (
   "fmt"
   "unsafe" //go语言的sizeof
)
// 低 --------》 高
// 12 34  -> 大端 -> 高尾端
// 34 12  -> 小端 -> 低尾端

func main() {
   s := int16(0x1234)
   b := int8(s)
   fmt.Println("int16字节大小为", unsafe.Sizeof(s)) //结果为2
   if 0x34 == b {
      fmt.Println("little endian")
   } else {
      fmt.Println("big endian")
   }
}
```



# 三、疑问

问题：

为什么要用公钥生成地址，而不直接使用公钥做接受地址



答：

1. 公钥太长
2. 公钥包含信息太少，地址可以包含版本号（不同网络的地址不通用）
3. 地址可以做校验（最后4字节的校验码）

