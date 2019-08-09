# 正则表达式regexp

## 使用步骤

1.解释规则.

使用以下方法定制一个解析器，并返回对它的引用，传递给变量reg

    reg := regexp.MustCompile(`a.b`)

其中函数结构如下

    func MustCompile(str string) *Regexp

2.使用规则

使用正则表达式对象的查找方法

    reg.FindAllStringSubmatch(buf,-1)

其中的查找方法定义如下：

    func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
