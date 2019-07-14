pragma solidity ^0.5.10;
contract testStroMemo{
    string public name = 'Lily';
    //注意，高版本的solidity已经不支持缺省的传递方式，若需要接收外部传递的复杂类型，必须显式声明传递方式。

    // 函数参数接收状态变量，显式传递方式：storage
    function setValue1(string storage _name)private{
        bytes(_name)[0] = "N";
    }
    function call1()public{
        setValue1(name);
    }
    // 状函数参数接收状态变量，显式声明传递方式：memory
    function setValue2(string memory _name)private pure{
        bytes(_name)[0] = "M";
    }
    function call2()public view{
        setValue2(name);
    }

    // 局部变量接受状态变量，显式传递方式：storage
    function localTest1()public{
        string storage tmp = name;
        bytes(tmp)[0] = "P";
    }
    // 局部变量接受状态变量，显式传递方式：memory
    function localTest2()public view{
        string memory tmp = name;
        bytes(tmp)[0] = "P";
    }
}