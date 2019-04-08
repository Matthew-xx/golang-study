# map

## 变量声明

    var mymap map[int]string
    fmt.Println(m)
结果

    map[]
此时map没有指向内存，打印出来是空的，直接赋值会panic

## 使用make创建map

仅仅声明并没有开辟内存空间，使用make可以创建有内存指向的map，还可以指定存储能力数量。

    mymap := make（map[int]string,3）

### 存储能力

map在使用make创建时可指定**存储能力**(可选参数)，但map不存在cap的概念，但有**len**的概念。

    m2 := make(map[int]string, 1) 
    //1是存储能力，内容长度为0
    fmt.Println(m2)//内容为空
    fmt.Println(len(m2)) //len为0
    m2[1] = "mike"
    m2[0] = "game"
    fmt.Println(m2)      //无序打印
    fmt.Println(len(m2)) //自动扩充为2，并不被存储能力1限制
结果：

    map[]
    0
    map[1:mike 0:game]
    2

## 声明时初始化创建map

    m3 := map[int]string{1: "mike", 2: "Bob"}

1. 注意格式","用来分割元素，":"用来分割键与值
2. 直接初始化不需要用make
3. 直接打印map，键值对无序显示

## 赋值

    m1 := map[int]string{1: "mike", 2: "Bob"}
    //如果对应的键已经存在，则赋值就是修改其值
    fmt.Println(m1)
    m1[2] = "haha"
    fmt.Println(m1)
    //如果对应的键不存在，则赋值就是追加内容类似append
    m1[4] = "I'm new!"
    fmt.Println(m1)
    m2 := m1
    m2[2] =  "lala"
    fmt.Println(m1)
    fmt.Println(m2)

结果

    map[1:mike 2:Bob]
    map[1:mike 2:haha]
    map[1:mike 2:haha 4:I'm new!]
    map[2:lala 4:I'm new! 1:mike]
    map[1:mike 2:lala 4:I'm new!]

1. 如果对应的键已经存在，则给对应键赋值就是修改其值
2. 如果对应的键不存在，则给对应键赋值就是追加内容，类似切片append函数，但不会开辟新内存。
3. map是引用类型，指向同一内存的map引用，所包含的键值对，一改俱改。

## map遍历

### 遍历map引用

    for key, value := range m {
        fmt.Printf("key=%v,value=%v\n", key, value)
    }

总结，使用range遍历一个map引用，会产生两个返回值，第1个为本次循环到的key，第2个本次循环对应的value。

### 键的返回值

    value, ok := m[5]
    if ok {
        fmt.Println(value)
    } else {
        fmt.Println("不存在")
    }

总结：map元素本身可返回两个返回值，第一个为元素键对应的值，第二个为是否存在

## delete函数

delete函数同样为go的内置函数。

    m := map[int]string{1: "mike", 2: "Bob"}
    m1:=m
    fmt.Println(m)
    delete(m, 2) //在m中删除键为2的键值对
    delete(m, 3)
    fmt.Println(m)
    fmt.Println(m1)

结果：

    map[1:mike 2:Bob]
    map[1:mike]
    map[1:mike]

1. 在m中删除键为3的键值对，不存在某个键为3，则什么也不发生。
2. delete函数时内存级别的删除，故同引用的map对应键值对也会被删除。