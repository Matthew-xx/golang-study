# 关键字介绍

## 修饰符类

### private、public、external、internal

用来修饰函数和状态变量，决定他们的外部访问权限。

智能合约部署后，状态变量和函数是要上链被访问的，如果需要控制访问权限，可以用这些关键词修饰。

1. private:完全私有
2. public:完全公开
3. external仅合约外部可以调用，内部也可以使用this调用
4. internal:私有，但可以被继承。

### view、constant

用来标记函数

代表这个函数很安全，只会读取，而不修改任何状态变量。

另，remix会把view变量显示到IDE前端，方便用户查看

### pure

用来标记函数

代表这个函数很安全，没有使用任何状态变量

### payable

用来标记函数

代表这个函数重点留意，调用该函数可以付钱，也可以不付钱，注意不是汽油，而是转账，钱会给到智能合约账户。

### returns

用来标记函数的返回值

### storage与memory

storage与memory专门用来修饰引用类型的变量：string,struct,mapping.不能用来修饰其他类型。
