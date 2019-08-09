pragma solidity ^0.5.10;
contract TestModifier{
    // 需求：限定只有合约创建者，才可以调用一个函数
    // 实现方法1：使用错误处理require
    // 实现方法2：使用modifier限定
    address public owner;
    uint256 public value;
    constructor()public{
        owner = msg.sender;
    }
    modifier onlyOwner(){
        require(msg.sender == owner, "无调用权限");
        _;
    }
    function changeValue(uint _value)public onlyOwner{
        value = _value;
    }


}