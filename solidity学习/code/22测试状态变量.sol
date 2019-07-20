pragma solidity ^0.5.10;

contract A{
    uint public data = 200;
    function setValue(uint _data)public{
        data = _data;
    }
}

contract B{
    A a;
    constructor() public{
        a = new A();
    }
    function getValue()public view returns(uint){
        return a.data();
    }
    function setValue()public returns(uint){
        a.setValue(300);
    }
}
contract C{
    A a;
    constructor() public{
        a = new A();
    }
    function getValue()public view returns(uint){
        return a.data();
    }
    function setValue()public returns(uint){
        a.setValue(500);
    }
}