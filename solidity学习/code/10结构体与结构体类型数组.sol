pragma solidity ^0.5.10;

contract TestStruct{
    // 定义与golang的结构体类似，注意1.类型在前，变量名在后2.注意是分号；分割。
    struct Student {
        string name;
        uint age;
        uint score;
        string sex;
    }
    // 定义结构体类型数组
    Student[] public Students;

    // 初始化stu对象,有两种方式，类似于go，但注意是()类似于函数，而不是{}只用于键值对。
    Student public stu1 = Student("Marry",12,99,"female");
    Student public stu2 = Student({name:"Bob",age:13,score:98,sex:"male"});

    function assgin()public{
        Students.push(stu1);
        Students.push(stu2);
        stu1.name = "Lily";
    }

    // solidity目前的ABI不支持直接返回自定义类型，使用测试版ABI，remixIDE又解析不了，很不方便开发
    // 解决方案：使用元组：tuple
    // 元组可以认为是多返回值。用（）包裹
    function returnStudent()public view returns(string memory,uint,uint,string memory){
        return(stu1.name,stu1.age,stu1.score,stu1.sex);
    }
}