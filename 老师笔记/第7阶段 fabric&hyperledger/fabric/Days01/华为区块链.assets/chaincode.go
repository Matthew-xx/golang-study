package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Test struct {
	// 测试区块链数据的读和写
}

func (this *Test) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (this *Test) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// chaincode 入口,更新,添加,查询数据 都可以走这个方法
	// 一句传递的数据来区分到底调用哪个方法
	// 传递:调用方法的名称 和 依据不同的方法名所依赖的参数
	// 使用ChaincodeStubInterface 下的 GetFunctionAndParameters
	// 约定:Parameters 如果是get放法index为0中存储可以,如果是set 方法,index为0中存储可以,index为1中存储value

	function, parameters := stub.GetFunctionAndParameters()
	if function == "get" {
		this.get(stub, parameters[0])
	} else if function == "set" {
		return this.set(stub, parameters[0], []byte(parameters[1]))
	}

	return shim.Error("Invalid Smart Contract function name.")

}

func (this *Test) set(stub shim.ChaincodeStubInterface, key string, value []byte) peer.Response {
	err := stub.PutState(key, value)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (this *Test) get(stub shim.ChaincodeStubInterface, key string) peer.Response {

	data, err := stub.GetState(key)

	// 异常处理
	if err != nil {
		return shim.Error(err.Error())
	}

	// 出 data == nil
	if data == nil {
		return shim.Error("Data not Availate.")
	}

	return shim.Success(nil)
}

func main() {
	shim.Start(new(Test))
}
