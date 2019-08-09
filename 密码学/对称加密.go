package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
)

// des 的cbc加密（密码分组链接模式）
// 1 填充 补位
// 2 刚好合适，再加一个分组
// 3 填充的字节的值=缺少的字节数（缺几个填几个）

func paddingLastGroup(plainText []byte, blockSize int) []byte {
	// 求分割后的余数
	lastNum := len(plainText) % blockSize
	// 要填充的字节数
	padNum := blockSize - lastNum
	// 要填充的字符，把数量当做内容
	char := []byte{byte(padNum)}
	// 准备，要填充的内容
	preparePlain := bytes.Repeat(char, padNum)
	// 组合
	newPlainText := append(plainText, preparePlain...)
	return newPlainText
}

// 去掉填充的数据
func unPaddingLastGroup(plainText []byte) []byte {
	// 1.读出切片的最后一个字节
	lastChar := plainText[len(plainText)-1]
	// 2.转为整型
	number := int(lastChar)
	// 切去切片
	return plainText[:len(plainText)-number]
}

// des加密
func desEncrypt(plainText, key []byte) []byte {
	// 1.创建一个底层使用des的密码接口的对象
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2.明文填充
	newText := paddingLastGroup(plainText, block.BlockSize())
	// 3.创建一个使用cbc分组的加密接口对象
	iv := []byte("12345678") //初始化向量
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 4.加密
	cipherText := make([]byte, len(newText))
	blockMode.CryptBlocks(cipherText, newText)

	return cipherText
}

// des解密
func desDecrypt(cipherText, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	newText := unPaddingLastGroup(plainText)
	return newText
}

// aes加密+crt
func aesEncrypt(plainText, key []byte) []byte {
	// 1.创建一个底层使用des的密码接口的对象
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 3.创建一个使用cbc分组的加密接口对象
	iv := []byte("12345678abcdefgh") //初始化向量
	stream := cipher.NewCTR(block, iv)
	// 4.加密
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)
	return cipherText
}

// aes解密
func aesDecrypt(cipherText, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)
	return plainText
}
