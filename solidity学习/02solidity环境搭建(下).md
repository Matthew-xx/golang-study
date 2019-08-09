# solidity环境搭建 （下）之VSCode

## 安装solidity和node的插件

需要根据自己的IDE，安装对应的插件包。VScode推荐solidity\Solidity Extend插件

注意:其实solidity插件安装好后，打开sol文件，直接F5就可实现编译，但是为了了解原理，建议自己写编译脚本
而且这里有个大坑(4小时)：插件编译后默认输出文件到VScode的工作区的第一个文件夹，巨坑无比。由于我的文件夹比较多，尝试编译后，没有反应，花费了太多时间。

## 安装配套环境软件

初始化项目

    node init

安装solc

    npm i --save solc

安装ganache

    npm i --save ganache

安装mocha

    npm i --save mocha

安装web3

    npm i --save web3

安装完以上软件，就算完成环境安装

### VScode对web3的代码自动补全

实际上，安装web3就已经附带了对应的typescript,只是vscode没有识别加载。需要在源文件所在目录新建jsconfig.json文件，写入一下内容

    {
        "compilerOptions": {
            "target": "es6"
        },
        "exclude": [
            "node_modules"
        ]
    }

至此，环境完美搭建完成！
