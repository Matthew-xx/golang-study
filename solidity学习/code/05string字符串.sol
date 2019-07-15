pragma solidity ^0.5.10;

contract testString{
    
    string public str = "HelloWorld,我来了";

    // 可强制转换为bytes
    function ToBytes() public view returns(bytes memory){
        return bytes(str);
    }
    // 强转换为bytes后，可以获取长度，修改内容
    // 字符串由utf编码存储，一个字母1字节，一个汉字3字节，共20字节。
    function getLength() public view returns(uint){
        return bytes(str).length;
    }
    function changeValue() public {
        bytes(str)[0] = "h";
    }
    // setlength不合适的话，会出现一个汉字不够拆的情况，虽然会set成功，但也导致字节序列不再符合utf8而读取报错。
    function setLength(uint i) public{
    bytes(str).length = i;
    }
}