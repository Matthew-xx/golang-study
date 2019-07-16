pragma solidity ^0.5.10;

contract TestGetter{
    uint public data = 200;
    // 本函数其实不用写，与自动创建的getter功能重复
    function getData()public view returns(uint){
        return data;
    }
}

contract Test2{
    
    function getValue()public returns(uint){
        TestGetter t1 = new TestGetter();
        return t1.data();
    }
    
}