pragma solidity ^0.5.10;

contract TestMapping{
    mapping(uint => string) public id_names;

    // 使用构造函数初始化
    constructor() public{
        id_names[1] = "Lily";
        id_names[2] = "Alice";
        id_names[3] = "Bob";
        // 覆盖
        id_names[3] = "Tom";
    }
    //通过key来查看对应的value
    function GetNameById(uint id)public view returns(string memory){
        return id_names[id];
    }
    // 通过key来修改对应的value
    function SetNameById(uint id)public {
        id_names[id] = "Jim";
    }
}