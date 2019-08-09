pragma solidity ^0.5.10;

contract TestValue{
    uint public money;
    function payMoney()public payable{
        money = msg.value;
    }

    // 获取合约地址内的余额
    function getBalance()public view returns(uint){
        return address(this).balance;
    }

    //需求:限定转账金额，并记录对应转账人
    mapping(address => uint) public person_money;
    function payMoneyLimit()public payable{
        require(msg.value <= 100,"转账数不符合要求！");
        person_money[msg.sender] = msg.value;
    }
}