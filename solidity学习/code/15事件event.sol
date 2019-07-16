pragma solidity ^0.5.10;
 contract TestEvent{
    // 定义一个事件,注意是个语句
    event payEvent(address,uint,uint);

    function payMoney()public payable{
        emit payEvent(msg.sender,msg.value,block.difficulty);
    }
 }