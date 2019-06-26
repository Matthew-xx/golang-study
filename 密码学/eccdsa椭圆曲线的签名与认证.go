package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"os"
)

// 生成秘钥对
func GenerateEccKey() {
	// 生产
	privateKey, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		panic(err)
	}
	// x509序列化
	derText, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	// pem序列化
	block := pem.Block{
		Type:  "ECDSA PRIVATE KEY",
		Bytes: derText,
	}
	// 创建文件对象
	file, err := os.Create("eccPrivate.pem")
	if err != nil {
		panic(err)
	}
	// pem写入文件
	pem.Encode(file, &block)
	file.Close()

	// 公钥
	publicKey := privateKey.PublicKey
	// x509序列化
	derText, err = x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	// pem序列化
	block = pem.Block{
		Type:  "ECDSA PUBLIC KEY",
		Bytes: derText,
	}
	// 创建文件对象
	file, err = os.Create("eccPublic.pem")
	if err != nil {
		panic(err)
	}
	// pem写入文件
	pem.Encode(file, &block)
	file.Close()
}

// eccdsa私钥签名
func EccSingature(plainText []byte, filePath string) (rText, sText []byte) {
	// 打开私钥文件
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	// 获取文件大小
	fileStatu, err := file.Stat()
	if err != nil {
		panic(err)
	}
	// 根据大小创建临时存储区并存入私钥序列
	buffer := make([]byte, fileStatu.Size())
	n, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}
	// pem反序列话
	block, _ := pem.Decode(buffer[:n])
	// x509反序列化得到私钥
	eccPrivateKey, err := x509.ParseECPrivateKey(block.Bytes)
	// 获取哈希
	hash := sha1.New()
	hash.Write(plainText)
	hashText := hash.Sum(nil)
	// 开始签名,得到签名结果r,s指针
	r, s, err := ecdsa.Sign(rand.Reader, eccPrivateKey, hashText)
	if err != nil {
		panic(err)
	}
	// 对指针r,s指向的bigint数据，序列化为[]byte
	rText, err = r.MarshalText()
	if err != nil {
		panic(err)
	}
	sText, err = s.MarshalText()
	if err != nil {
		panic(err)
	}
	return
}

// ecc公钥认证
func EccVerify(plainText, rText, sText []byte, filepath string) bool {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, fileInfo.Size())
	n, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(buffer[:n])
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	eccPublicKey := pubKey.(*ecdsa.PublicKey)
	// 获取hash
	hash := sha1.Sum(plainText)
	hashText := hash[:]
	// 反序列话指纹r为int
	var r, s big.Int
	err = r.UnmarshalText(rText)
	if err != nil {
		panic(err)
	}
	err = s.UnmarshalText(sText)
	if err != nil {
		panic(err)
	}
	// 开始验证
	return ecdsa.Verify(eccPublicKey, hashText, &r, &s)
}
