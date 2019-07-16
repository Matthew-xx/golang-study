pragma solidity ^0.5.10;

contract TestSender{
    address public owner;

    // 谁调用这个函数 owner就是谁。
    constructor()public{
        // 构造函数中使用，可确定合约创建者，从而指定管理员
        owner = msg.sender;
    }

    function changeOwner() public{
        owner = msg.sender;
    }

    
}