pragma solidity ^0.5.10;

contract TestDelete{
    // test string：变为空字符串
    string public str1 = "hello";
    function setStr(string memory input) public{
        str1 = input;
    }
    function deleteStr()public {
        delete str1;
    }
    // test fixed arry:定长数组长度不变，全部用0填充
     uint256[10] public fixedArry = [1,2,3,4,5];
     function deleteFixedArry() public {
         delete fixedArry;
     }
    // test dynamic arry：动态数组内存全部清空，长度置为0
    uint256[] public dynamicArry = [1,2,3,4,5];
    function deleteDynamicArry() public {
         delete dynamicArry;
     }
    //  test mapping : 不能删除map本身，但可以根据某个键，删除某个键值对，删除效果是值置为默认值。
    mapping(uint256 => string) public map;
    function setMapping()public{
        map[110] = "匪警";
        map[119] = "火警";
    }
    function deletemapping() public{
        // delete map;
        delete map[110];
    }
}