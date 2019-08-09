pragma solidity ^0.5.10;
// 通过1个合约向另一个合约转钱
contract C1 {
    function payMoney()public payable returns(uint){
        return 22;
    }
    function getBalance() public view returns(uint){
        return address(this).balance;
    }
}

contract C2 {
    C1 public c1;
    function()external payable{

    }
    function setC1(address addr)public{
        c1 = C1(addr);
    }
    function callFeed()public{
        c1.payMoney.value(10).gas(300000)();
    }
    function getBalance() public view returns(uint){
        return address(this).balance;
    }
}