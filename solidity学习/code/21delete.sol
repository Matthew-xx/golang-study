pragma solidity ^0.5.10;

contract TestDelete{
    // test string
    string public str1 = "hello";
    function setStr(string memory input) public{
        str1 = input;
    }
    function deleteStr()public {
        delete str1;
    }
    // test arry
     
}