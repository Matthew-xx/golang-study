pragma solidity ^0.5.10;

contract test{
    address public addr1 = 0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c;

    // 匿名函数被用作后路执行（fallback），如调用不存在的函数时，即运行它，可用作低成本转账
    // 注意用作转账时需要设置value
    function () external payable{
    }

    function getAddrBalance() public view returns(uint){
        return addr1.balance;
    }
    function getContrBalance() public view returns(uint){
        return address(this).balance;
    }
    // 合约地址向别人转账8eth,**代表指数
    function contrTransfer() public payable{
        // transfer函数由被转账方调用,注意把函数设置为payable
        //最新版solidity把address分为了payable address和address
        address(uint160(addr1)).transfer(8*10**18);
    }
}
