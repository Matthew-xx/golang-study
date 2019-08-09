package main

import (
	"log"
)

// 测试文件
func main() {
	// 测试des加密：模式CBC
	log.Println("des 加解密")
	key := []byte("1234abcd")
	src := []byte("我是大佬")
	cipherText := desEncrypt(src, key)
	plainText := desDecrypt(cipherText, key)
	log.Println("密文是：", string(cipherText))
	log.Println("解密后：", string(plainText))
	// 测试des加密，模式:CRT
	log.Println("des 加解密:crt模式")
	newKey := []byte("12345678abcdefgh")
	cipherText = aesEncrypt(src, newKey)
	plainText = aesDecrypt(cipherText, newKey)
	log.Println("密文是：", string(cipherText))
	log.Println("解密后：", string(plainText))

	// 测试生成rsa秘钥文件
	GenerateRsaKey(1024)

	// 使用公钥文件加密
	cypherText := RSAEncrypt([]byte("go语言牛逼！"), "public.pem")
	log.Println("公钥加密后的密文：", string(cypherText))
	// 使用私钥解密
	plainText = RSADecrypt(cypherText, "private.pem")
	log.Println("私钥解密后的明文：", string(plainText))

	// 计算sha256哈希
	myHash(plainText)

	// 消息验证码
	hmac := GenerateHmac([]byte("我们结婚吧"), []byte("heliu"))
	// 校验消息验证码
	isSuccess := VerifyHmac([]byte("我们结婚吧"), []byte("heliu"), hmac)
	log.Println(isSuccess)

	// RSA签名
	RSASign := RSASignature([]byte("我们分手吧"), "private.pem")
	// 校验签名
	isSuccess = RSAVerify([]byte("我们分手吧"), RSASign, "public.pem")
	log.Println(isSuccess)

	// ecdsa密钥对生成
	GenerateEccKey()
	// ecc签名
	r, s := EccSingature([]byte("ecc签名"), "eccPrivate.pem")
	// ecc 认证
	isSuccess = EccVerify([]byte("ecc签名"), r, s, "eccPublic.pem")
	log.Println(isSuccess)
}
