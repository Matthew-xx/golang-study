pragma solidity ^0.5.10;

contract testKeccak{
    function getHash()public pure returns(bytes32){
        return keccak256("helloworld");
    }
}