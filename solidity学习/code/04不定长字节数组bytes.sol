pragma solidity ^0.5.10;

contract TestBytes{
    bytes public b;

    function getLen()public view returns(uint){
        return b.length;
    }
    function setLen(uint i) public returns(uint){
        b.length = i;
        return b.length;
    }
    // 注意0.5.10的bytes类型只能接收0x开头的16进制数组，每两个当做一个元素。
    function setValue(bytes memory input)public {
        b = input;
    }
    function setValueByIndex(uint i)public {
        b[i] = 'h';
    }
    function getByIndex(uint i)public view returns(byte){
        return b[i];
    }
    function pushData()public {
        b.push('h');
    }
    
}
// 测试结果:
// 1. 不初始化，直接读取b的长度，未报错，结果为0
// 2. 不初始化，直接读b的元素，err
// 3. 不初始化，直接赋值，自动完成初始化，注意输入只能为16进制数字
// 4. 不初始化，直接设置长度，自动完成初始化，切内容以0填充
// 5. 不初始化，直接设置b的元素，err,初始化后可以。
// 6. 不初始化，直接push元素，可自动完成初始化，结果为尾部追加
