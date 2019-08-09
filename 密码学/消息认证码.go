package main

import (
	"crypto/hmac"
	"crypto/sha1"
)

// 生成消息验证码
func GenerateHmac(plainText, key []byte) []byte {
	// 创建hmac对象，并指定哈希算法和秘钥
	hash := hmac.New(sha1.New, key)
	// 给哈希对象添加数据
	hash.Write(plainText)
	// 计算散列值
	hashText := hash.Sum(nil)
	return hashText
}

// 校验消息验证码
func VerifyHmac(plainText, key, hashText []byte) bool {
	hash := hmac.New(sha1.New, key)
	hash.Write(plainText)
	myhashText := hash.Sum(nil)
	return hmac.Equal(hashText, myhashText)
}
