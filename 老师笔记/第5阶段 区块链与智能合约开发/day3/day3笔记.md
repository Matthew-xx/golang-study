挖矿交易即coinbase并非没有输入，而是输入并用引用其他的输出

特点：这条挖矿交易只有一个输入，一个输出，并且输入引用id为-1，索引为-1，签名字段可以随便写数据。



转账记住两点：

1. 每一笔能够支配的钱来源于上一个交易输出（普通交易）
2. 每一个花费的输出都要一次性花完，如果有剩余，转给自己





# 步骤

## 1. 定义交易结构

```go
gob: type main.Output has no exported field
```

==结构里面的数据要大写==(大坑！！三个小时)

## 2. setTXID

## 3. 创建CoinbaseTX

## 4. 使用Transaction改写代码

block.go, 

pow.go,   //引出挖矿真相，TODO

blockchain.go, //处理创世块交易，屏蔽addBlock

cli.go, //屏蔽addBlock，因为参数已经完全变了

main.go //地址先写成中本聪



编译通过

### - 挖矿的真相

挖矿其实是对区块头做哈希，而不是对整个区块，区块体是通过梅克尔根影响到区块头的。

![image-20181016170425026](/Users/duke/Library/Application Support/typora-user-images/image-20181016170425026.png)

z





梅克尔根是一个二叉树，我们这里就做一个哈希的拼接即可。

### - MakeMerkelRoot

```go

func (block *Block) MakeMerkelRoot() []byte {
	//这是模拟梅克尔树，目前只是做哈希的拼接，后续再完成二叉树的构造

	txs := block.Transactions
	var hashInfo []byte
	for _, tx := range txs {
		hashInfo = append(hashInfo, tx.TXID...)
	}

	return hashInfo
}
```

注意在block中引用MakeMerkelRoot函数即可，在pow已经调用MerkelRoot了



## 5. 分析图

![1.模拟转账演示](https://ws3.sinaimg.cn/large/006tNbRwly1fwbdnqy6qzj31kw0wy1dp.jpg)



## 6.查找余额，实现FindUTXOs



## 7. 创建普通交易

- FindUTXOs

  先实现空函数

- output转成input
- 创建新的output

- 找零

## 

## 8. 转账send

## 9. 获取余额

## 10. 拆分函数

每次不用把所有的钱都统计出来，够用就行





