// 生成私钥并按规范写入文件
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"log"
	"os"
)

// 生成RSA的秘钥对
func GenerateRsaKey(keySize int) {
	// 使用rsa包中的方法生成rsa私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}
	// 使用x509标准将得到的私钥序列化为ASN.1的DER编码字符串
	derText := x509.MarshalPKCS1PrivateKey(privateKey)

	// 准备编码成pem格式
	block := pem.Block{
		Type:  "RSA Private Key",
		Bytes: derText,
	}
	// 开始编码
	file, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(file, &block)
	if err != nil {
		panic(err)
	}
	file.Close()
	// ===========公钥===========
	publicKey := privateKey.PublicKey
	derStream, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	block = pem.Block{
		Type:  "RSA Public Key",
		Bytes: derStream,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(file, &block)
	if err != nil {
		panic(err)
	}
	file.Close()
}

// RSA加解密
// 使用公钥文件加密
func RSAEncrypt(plainText []byte, filePath string) []byte {
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
	public, err := x509.ParsePKIXPublicKey(p.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey, isSuccess := public.(*rsa.PublicKey)
	if !isSuccess {
		panic(errors.New("断言错误"))
	}
	// 使用公钥加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	return cipherText
}

// /使用私钥文件解密
func RSADecrypt(crypherText []byte, filePath string) []byte {
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
	// 使用私钥解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, crypherText)
	if err != nil {
		panic(err)
	}
	return plainText
}

// RSA使用私钥文件，签名
func RSASign(filePath string) {
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
	log.Println(privateKey)
	// 使用私钥签名
	// cipherText,err:=rsa.SignPKCS1v15(rand.Reader,privateKey,)
}
