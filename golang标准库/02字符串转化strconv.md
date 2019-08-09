# 字符串转换

基本数据类型可以转化为字符串，字符串类型也可以转化为其他基本数据类型。（互相转化）

## FormatType系列函数

实现了将其他类型转换为字符串类型

### FormatBool()

    str := strconv.FormatBool(false)
    fmt.Println(str)
    //false

### FormatFloat()

f指打印格式以小数方式，-1小数位数为紧缩，以64处理。

    str = strconv.FormatFloat(3.14, 'f', -1, 64)
    fmt.Println(str)
    //false

## ParseType系列函数

字符串类型转化为其他基本类型数据

### ParseBool()

    flag, err := strconv.ParseBool("true")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("type is %T\n", flag) //%T代表类型 %V %t??
    }

## 常用的string与int互转

### Itoa()

    str = strconv.Itoa(34534)
    //=FormatInt(34534,10)
    fmt.Println(str+1)
    //234341

### Atoi()

    a, _ := strconv.Atoi(str) //=ParseInt(str,10,0)
    fmt.Println(a+a)
    //34535

## AppendType系列函数

### 回忆内置函数append()

在切片一节，介绍了一个内置函数（在main包中定义）-append().它可以为任意类型的切片，追加新的切片元素。

AppendType系列函数实现了将任意基本类型转化为字符串，在转化为字符切片，追加到**字符切片**。

注意他的追加目标切片只能是字符切片，而非字符串切片。

### AppendBool()

    b := []byte{'a', 'b'}
    b = strconv.AppendBool(b, true)
    fmt.Printf("b=%c\n", b)
    //b=[a b c d t r u e]

### AppendInt()

    b1 = strconv.AppendInt(b1, 123, 10)
    //第3个参数代表进制
    fmt.Printf("b1=%c\n", b1)
    //b1=[a b g h t r u e 1 2 4]

### AppendQuite()

    b1 = strconv.AppendQuote(b1, "aaaaaaaa") //Quote表示引号
    fmt.Printf("b1=%c\n", b1)
    //b1=[a b g h t r u e 1 2 4 " a a a a a a a a "]