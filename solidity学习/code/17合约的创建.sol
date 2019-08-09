pragma solidity ^0.5.10;
// 创建合约对象的方式：
// 注意旧版solidity通过new出来的是合约的地址，需要强转成成合约对象才可使用。
contract C1{
    function getValue() public pure returns(uint){
        return 100;
    }
}
contract C2{
    C1 public c1;
    function getValue() public returns(uint){
        c1 = new C1();
        return c1.getValue();
    }
}