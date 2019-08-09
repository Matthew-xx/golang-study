# 数据格式json

## 编码json数据

编码成json数据方法

    func Marshal(v interface{}) ([]byte, error)
编码带缩进的json数据方法：

    func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
//参数2为前缀符号（行首字符），参数3为缩进符号.均可指定为任意字符
输出

### 通过结构体生成json

构建一个结构体

成员变量名为满足包外访问必须大写，然而json格式中对应的键为小写会导致不一致。可使用struct tag起别名。

    type IT struct {
        Company string   `json:"company"` //此字段在json中的键用新名字
        Subject []string `json:"-"`       //此字段被忽略
        Isok    bool     `json:",string"` //此字段，视为string
        Price   float32
    }

声明并初始化一个该结构体类型的变量

    it := IT{"heima", []string{"Go", "C++", "Python", "Test"}, true, 5465.45}

将此变量转换为json格式

    byte, err := json.Marshal(it)
    fmt.Println(string(byte))
输出

    {"company":"heima","Isok":"true","Price":5465.45}

转化为带有缩进的json格式：

    byte, err := json.MarshalIndent(it, "", "  ")
输出

    {
    "company": "heima",
    "Isok": "true",
    "Price": 5465.45
    }

### 通过map生成json

声明并舒适化一个map类型变量

    m := map[string]interface{}{
        "company": "黑马",
        "subject": []string{"GO", "C++", "Python", "C"},
        "isok":    true,
        "Price": 324.34
        }

生成json

    b, err := json.MarshalIndent(m, "", "  ")
        if err != nil {
            fmt.Println("err=", err)
        }
    fmt.Println(string(b))

      "Price": 5465.45
}
输出

    {
            "Price": 324.34,
            "company": "黑马",
            "isok": true,
            "subject": [
                    "GO",
                    "C++",
                    "Python",
                    "C"
            ]
    }

## 解码json数据

    使用如下函数解码

    func Unmarshal(data []byte, v interface{}) error
此函数

    此函数将jsons格式数据解码并存储到变量V所指向的内存中，如果V不是一个指针，或者是nil，此函数返回InvalidUnmarshalError.

    解码是编码的反面

## 解码json到结构体

仍使用上述定义的结构体，将输出的json数据，存储到一个**原生字符串**中。

    jsonbuf := `{
        "price": 324.34,
        "company": "黑马",
        "isok": true,
        "subject": [
                "GO",
                "C++",
                "Python",
                "C"
        ]
    }`
    //原生字符串使用反引号包裹，允许换行，不识别转义。

开始解码到这个变量

    var it IT
    err := json.Unmarshal([]byte(jsonbuf), &it)
    //注意取地址
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%+v", it)

输出

    {Company:黑马 Subject:[GO C++ Python C] Isok:true Price:324.34}

### 解码json到map

继续利用原生字符串，直接解码

    m := make(map[string]interface{}, 4)
    err := json.Unmarshal([]byte(jsonbuf), &m)
    //虽然map为引用类型，但是还是需要取指针的，由函数内部实现是必须是指针类型，实现接管整块内存。
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%+v\n", m)