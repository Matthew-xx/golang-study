pragma solidity ^0.5.10;

contract testConv{
    // 初始化一个定长数组
    bytes10 public fixedBytes = 0x2ce68891e69da5e4ba86;
    // 初始化一个可变数组
    bytes public unfixedBytes = new bytes(fixedBytes.length);
    string public str;

    // 将定长数组内容赋值给不定长数组
    function fixedToUnfixed()public{
        // unfixedBytes = bytes(fixedBytes); err
        for(uint256 i = 0;i < fixedBytes.length;i++){
            unfixedBytes[i] = fixedBytes[i];
        }
    }
    // 将定长数组内容转换为string,需要先转换为unfixedbytes
    function fixedToString()public{
        fixedToUnfixed();
        str = string(unfixedBytes);
    }
}