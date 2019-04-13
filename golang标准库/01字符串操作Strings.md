# 字符串操作包：Strings

## Contains()

判断一个字符串A是否包含另一个字符串B，两个字符串作为参数，包含返回true,

    s := "hello kitty"
    fmt.Println(strings.Contains(s, "hello"))
    //true

## Join()

将一个字符串切片的元素，拼接成一个字符串，可以指定用某个字符来拼接。

    slice := []string{"a", "b", "c", "d"}
    fmt.Println(strings.Join(slice, "#"))
    //a#b#c#d

## Index()

一个字符串在另一个字符串中的位置(首次出现)，没有则返回-1

    fmt.Println(strings.Index(s, "l"))
    //2

## repeat（）

重复一个字符串X次

    fmt.Println(strings.Repeat("go", 3))
    //gogogo

## split

切割字符串。把一个字符串，根据指定字符串拆开成字符串切片的一组元素。

    slice1 := strings.Split("185734549@qq.com@myhome", "@")
    fmt.Println(slice1[0])
    fmt.Println(slice1)
    //185734549
    //[185734549 qq.com myhome]

## trim()

把一个字符串的两头的指定字符串全部去掉。常用来删除两头空格

    s2 := "   a   b    c  d    "
    fmt.Println(strings.Trim(s2, " "))
    //a   b    c  d

## field()

把一个字符串按照其中的空格拆成一个切片, 相当于Split指定的参数为空格

    s3 := strings.Fields(s2)
    for _, str := range s3 {
        fmt.Println(str)
    }
    //a
    //b
    //c
    //d


