pragma solidity ^0.5.10;

contract TestEnum{
    // 声明一个枚举类型的对象
    enum weekDays{
        Monday,Tuesday,Wendesday,Thursday,Friday,Saturday,Sunday
    }
    // 设置枚举对象类型的变量currentDay
    weekDays currentDay;
    //设置枚举对象的默认值
    weekDays defaultDay = weekDays.Sunday;
    
    function setDay(weekDays _day)public{
        currentDay = _day;
    }
    function getDay()public view returns(uint){
        return uint(currentDay);
    }
    function getDefaultDay() public view returns(uint){
        return uint(defaultDay);
    }

}
// 总结，枚举实际上就是把一个变量的值固定到一个范围，该变量只能在特定的几个值之中进行选择，并且，会自动给这几个值添加序号，用数字指代。
