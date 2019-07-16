# solidity高级语法

## 自动推导类型var（不推荐）

var类似于golang中的:=可以以var替代一些自定义类型名，由系统自动完成类型推导。

不推荐的原因：var也可以自动推导基本类型，但默认是字节占用最小的，如数字被自动推导为int8，这有时是不符合需求的，容易导致溢出。

## 全局变量：msg

solidity自带一些全局变量，专门指代以太坊的交互对象。

### msg.sender

msg.sender是一个address类型
每次与以太坊交互，都会产生一笔交易，sender指代当前交易的发送者就是sender

### msg.value

msg.value是一个uint类型
与以太坊交互包含转账时（发送eth），value指代当前的交易量。故**必须与payable修饰符配套使用**。

## 其他全局变量

### block系列

    block.number
    block.coinbase
    block.difficulty
    block.gaslimit
    ...

### msg系列

    msg.sender
    msg.value
    msg.sig
    msg.data
    ...

### 其他

## 错误处理

旧版本使用throw，可回退交易
新版本已经废弃，使用require,revert,assert来替代

实例请查看13msg.valuet测试.sol

require()

- 代替if专门进行错误判断。
- 括号里面是逻辑判断，若true则继续执行，false则抛出异常返回
- 允许自定义异常内容.

assert\revert与require用法类似。细微区别

## 修饰器modifier

modifier由一种特殊的函数定义而成，专门用来修饰函数，从而限制一个函数的行为。、
可以被继承，可以被派生合约重写。

本质上，修饰器是提取出来专门用来做某些判断的子函数，一个函数若被修饰器修饰了，先执行修饰器中的函数，在执行本函数内的代码。

优势就是节省了代码，提取常用判断代码做成一个公用的修饰器。

## 货币单位与时间单位

### 货币单位

    1 ether = 10^18 wei = 10^3 finney = 10^6 szabo

 ### 时间单位

    1 years = 365 days
    1 weeks = 7days
    1 days = 24 hours
    1 hours = 60 minutes
    1 minutes = 60 seconds
    1 = 1 seconds

## 事件event

事件可以理解为一种特殊的函数，只有函数定义和调用，没有函数体。
调用事件，会在交易体的log中，展示其中的参数。

    - 使用emit发射event

## 自动创建访问函数getter

任何一个public类型的状态变量，都会自动创建一个同名的访问函数，get()。
合约外（如另外一个合约中）可通过XX.data()方式访问。

## 合约间转账语法

    c1.payMoney.value(10).gas(300000)();

## 元组tuple

solidity无法返回自定义的数据结构，可以使用元组达到同样的目的。
元组，是（）包裹的一组基本数据类型的数据。可实现函数的多返回值。

 - 注意：元组的数据无法修改。

## 内置数学函数

常用的就1个sha3()哈希算法，但目前名称被废弃，改名为keccak256()

## 继承

- 继承语法 ： is
- 最远继承原则：若同时继承两个父类存在相同的方法，则使用最远继承的那个父类的方法。
- 可以使用".父类名.父类方法X"调用父类的指定方法。

## new、delete

new用来创建对象（开辟内存空间）

delete 用来删除任何变量，但不是真删除，是重置为默认值。

- delete不定长数组，删除所有元素，长度置为0
- delete定长数组，重置所有索引值
- delete mapping类型，什么都不发生
- delete mapping类型中的一个key，删除与该key相关联的值。
