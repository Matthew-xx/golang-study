# solidity环境搭建

## 在线IDE

    http://remix.ethereum.org/

## 本地环境

### 安装nodejs

建议安装lts版本，这里安装

    node --version
    v10.16.0

自带的npm版本

    npm -v
    6.9.0

### 更换node源为淘宝镜像

1. 打开.npmrc文件（nodejs\node_modules\npm\npmrc
2. 增加
 
         registry =https://registry.npm.taobao.org 

或者使用命令：

    npm config set registry https://registry.npm.taobao.org

检测是否成功

    // 配置后可通过下面方式来验证是否成功
    npm config get registry
    // 或
    npm info express

另外若包非常多，建议安装配套的cnpm

### 安装solidity和node的插件

需要根据自己的IDE，安装对应的插件包。
