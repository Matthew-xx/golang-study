package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

// 使用sha256进行哈希
func myHash(src []byte) {
	// 方法一使用哈写接口
	myHash := sha256.New()
	myHash.Write(src)
	myHash.Write(src) //写1次以上会连续写入，相当于字符串拼接
	res := myHash.Sum(nil)
	log.Printf("%x\n", res)
	// 格式化为16进制形式
	myStr := hex.EncodeToString(res)
	log.Printf("%s\n", myStr)

	// 方法二 使用包装好的方法
	src = append(src, src...)
	res2 := sha256.Sum256(src)
	log.Printf("%x\n", res2)
}
