pragma solidity ^0.5.10;
 contract CustTest{
     uint256[10] public numbers = [10,20,30,40,50,60,70,80,90,100];
     uint256 public sum;

     function total()public{
         for(uint256 i = 0;i < numbers.length;i++){
             sum += numbers[i];
         }
    }
    function changeValue(uint i,uint value)public{
        numbers[i] = value;
    }

    // 内置定长bytes数组
    bytes10 public fixedContent = 0x2ce68891e69da5e4ba86;
    // 自定义定长bytes数组
    byte[10] public custContent = [byte(0x2c),0xe6,0x88,0x91,0xe6,0x9d,0xa5,0xe4,0xba,0x86];
    // 不定长字节数组
    bytes public dynamicContent;

    function ToDynamic()public{
        for(uint256 i = 0;i < custContent.length;i++){
            // dynamicContent[i] = custContent[i];
            dynamicContent.push(custContent[i]);
        }
    }
 }