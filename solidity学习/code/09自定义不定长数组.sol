pragma solidity ^0.5.10;

contract CustDynamic{
    uint[] public c = [1,2,3,4];

    function pushCon()public{
        c.push(5);
    }
    function getNumsC()public view returns(uint[] memory){
        return c;
    }

    uint[] public b;
    function pushCon2()public{
        b = new uint[](7);
        b.push(10);
    }

    function getNumsB()public view returns(uint[] memory){
        return b;
    }

}