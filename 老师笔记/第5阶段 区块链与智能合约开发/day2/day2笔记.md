# bolt数据库的操作

![image-20181015115621341](https://ws1.sinaimg.cn/large/006tNbRwly1fw8rrks54vj31ag0n4wgl.jpg)

```go
package main

import (
	"fmt"
	"go一期/bolt"
	"log"
)

func main() {
	//1. 打开数据库
	db, err := bolt.Open("test.db", 0600, nil)
	defer db.Close()

	if err != nil {
		log.Panic("打开数据库失败！")
	}

	//将要操作数据库（改写）
	db.Update(func(tx *bolt.Tx) error {
		//2. 找到抽屉bucket(如果没有，就创建）
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			//没有抽屉，我们需要创建
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建bucket(b1)失败")
			}
		}

		//3. 写数据
		bucket.Put([]byte("11111"), []byte("hello"))
		bucket.Put([]byte("22222"), []byte("world"))

		return nil
	})

	//4. 读数据
	db.View(func(tx *bolt.Tx) error {
		//1. 找到抽屉，没有的化直接报错退出
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			log.Panic("bucket b1 不应该为空，请检查!!!!")
		}

		//2. 直接读取数据
		v1 := bucket.Get([]byte("11111"))
		v2 := bucket.Get([]byte("22222"))

		fmt.Printf("v1 : %s\n", v1)
		fmt.Printf("v2 : %s\n", v2)

		return nil
	})

}

```

# 二、BlockChain改写

![image-20181015144902987](https://ws4.sinaimg.cn/large/006tNbRwly1fw8wr7uwf0j31gk0owwjm.jpg)



# 三、gob测试

```go
package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"fmt"
)

type Person struct {
	Name string
	Age  uint
}

func main() {
	//- 定义一个结构Person

	var xiaoMing Person
	xiaoMing.Name = "小明"
	xiaoMing.Age = 20

	//编码的数据放到buffer
	var buffer bytes.Buffer

	//- 使用gob进行序列化（编码）得到字节流
	//1. 定义一个编码器
	//2. 使用编码器进行编码
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&xiaoMing)
	if err != nil {
		log.Panic("编码出错，小明不知去向")
	}

	fmt.Printf("编码后的小明：%v\n", buffer.Bytes())

	//- 使用gob进行反序列化（解码）得到Person结构
	//1. 定义一个解码器
	//func NewDecoder(r io.Reader) *Decoder {
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))

	var daMing Person
	//2. 使用解码器进行解码
	err = decoder.Decode(&daMing)
	if err != nil {
		log.Panic("解码出错!")
	}

	fmt.Printf("解码后的小明：%v\n", daMing)

}

```

