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

}
