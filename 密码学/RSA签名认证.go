package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"io"
	"os"
)

// RSA使用私钥文件，签名(哈希算法使用sha512)
func RSASignature(plainText []byte, filePath string) []byte {
	// 从文件获得私钥
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	b := make([]byte, fileInfo.Size())
	n, err := file.Read(b)
	if err != nil && err != io.EOF {
		panic(err)
	}
	file.Close()
	p, _ := pem.Decode(b[:n])
	privateKey, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		panic(err)
	}
	// 得到数据的哈希
	hash := sha512.New()
	hash.Write(plainText)
	hashText := hash.Sum(nil)
	// 使用私钥签名
	signText, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hashText)
	if err != nil {
		panic(err)
	}
	return signText
}

// RSA使用公钥文件进行签名校验（哈希使用对应的sha512）
func RSAVerify(plainText, signText []byte, filePath string) bool {
	// 获取文件句柄
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	// 获取文件大小信息
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	// 根据文件大小信息创建切片，用来存储读入的文件字节流
	buffer := make([]byte, fileInfo.Size())
	// 读文件
	n, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}
	// 使用pem解码，得到ASN.1DER格式编码的公钥序列
	block, _ := pem.Decode(buffer[:n])
	// 使用x509解码,得到公钥初始序列
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := pubKey.(*rsa.PublicKey)
	// 获取hash
	hash := sha512.New()
	hash.Write(plainText)
	hashText := hash.Sum(nil)
	// 执行验证
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hashText, signText)
	if err == nil {
		return true
	}
	return false
}
