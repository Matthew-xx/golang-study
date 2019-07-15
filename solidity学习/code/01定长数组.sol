pragma solidity ^0.5.10;

contract test{
    bytes1 b1 = 'h';
    bytes20 b20 = 'helloworld';
    //获取数组的长度
    function getLen() public view returns(int256){
        return b20.length;
    }
    //尝试修改某个数组元素,报错
    function setValue() public{
        // b20[0] = 'H';
        b20 = 'HelloWorld!!!!';
    }
    // 尝试访问数组内容，传递1越界会报错
    function getValue(uint i) public view returns(byte){
        return b20[i];
    }
}
