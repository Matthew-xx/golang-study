pragma solidity ^0.5.10;

contract Base1 {
    function data() public pure returns(uint){
        return 1;
    }
}

contract Base2 {
    function data() public pure returns(uint){
        return 2;
    }
}

contract Son1 is Base1,Base2{

}

contract Son2 is Base2,Base1{
    
}