# Node常用模块

## 理解Node编程与浏览器编写js的区别

浏览器写js,实际上是使用浏览器引擎提供的功能和方法写代码。
其中的js库，如jquery也不过是对这些接口的二次封装。

Node端写js，脱离了浏览器，是使用node封装好的一系列功能模块写代码，这造成了node端经常调用的方法，有很大不同。加之ES6语法的不同，让初学者感觉像是写完全不同的东西。

## NPM介绍

NPM是Node自带的包管理器，与node捆绑安装。用来拓展node的功能模块。

npm是世界上生态最丰富的社区，没有之一。基本上要啥有啥

### npm配置文件

在当前项目文件夹下使用以下命令

    node init
可初始化一个node项目，生产package.json文件


    npm install --save XX
安装一个包，并保存到package.json文件中

## Node自带全局变量

- __dirname: 当前文件目录
- __filename: 当前文件命+绝对路径
- console:控制台对象，用来输出信息
- process:进程对象，获取进程信息，环境变量
- setTimeout/clearTimeout:延时执行
- setInterval/clearInterval:定时器

## path模块

提供一些工具函数，用来处理文件与目录的路径相关

- path.basename:返回一个路径的最后一部分
- path.dirname:返回一个路径的目录名
- path.extname:返回一个路径的拓展名
- path.join:拼接路径片段
- path.normalize:将一个路径正常化

## fs模块

文件操作相关模块

- fs.stat/fs.statSync:访问文件的元数据，如文件大小，修改时间。
- fs.readFile/fs.readFileSync:异步/同步读取文件
- fs.writeFile/fs.writeFileSync:异步/同步写入文件
- fs.readdir/fs.readdirSync:读取文件夹内容
- fs.unlink/fs.unlinkSync:删除文件
- fs.rmdir/fs.rmdirSync：删除空文件夹，删除非空文件夹可使用fs-extra第三方模块实现删除
- fs.watchFile:监视文件变化

## 操作大文件:stream流式处理

fs.readFile是整个读入内存，不适用于大文件，可使用流对象处理

fs.createReadStream可创建一个流对象

## 异步IO的回调地狱

异步IO使用回调过度造成callback hell，为了解决这个问题，诞生了promise和async/await

### promise

promise可以对异步回调代码进行包装，把原来的一个回调拆成2个回调，以提高可读性。

    注意：resolve和reject只能调用1次，调此则彼失效。

创建promise对象

    let prms = new Promise((reslove,reject)=>{
        //开始异步操作
        ioOperaction(param,(err,data)=>{
            if(err){
                rejecct(err)
            }else{
                reslove(data.toString())
            }
        })
    })


使用promise对象

    prms.then((text)=>{
        //需要resolve的内容的具体操作
    }).catch((err)=>{
        //需要reject的err具体操作
    })

#### 使用util.promisefy快速生产promise对象

    const {promisify} = require('util');
    const 新异步操作函数名 = promisify(异步操作函数名);
    执行新异步操作函数(参数同原函数).then((text)=>{

    }).catch((err)=>{
        
    })

### async/await

将异步的Promise代码变为同步的写法，只是写法同步，实际仍然是异步

处理promise对象的新机制:定义async类型方法，并执行

    async function asyncDemo(){
        try{
            let text = await prms
            //需要resolve的内容的具体操作

            let text1 = await prms1
            ......
        }catch(err){
            //统一处理所有promise的错误
        }
    }
    asyncDemo()

### 终极使用方法

    1. 先使用promise包装，可使用util.promisify包
    2. 使用async/await编写同步风格的异步代码

## http模块

本模块包含了httpserver和httpclient

    http.createServer 创建server对象
    http.get:执行httpget请求

## 模块导出

当用户自己写了一个模块（XX.js）后，如果需要在其他文件中使用，并不能直接require到
需要用户自定义的的模块下加入以下代码

    modle.exports = funcName//只导出一个方法，注意是方法名
或

    modle.exports = {
        name : varName,
        add:funcName,
        multi //==multi:multi
    }
    //同时导出多个时，使用对象包裹起来，使用时通过点调用名字

另外注意require是使用相对路径,如./xxx/xx

## npm安装与卸载包

安装与卸载分为全局性操作和项目局部操作。

    npm i -g XXX 全局安装
    npm i --save XXX 局部安装

全局安装的是全系统生效，项目局部安装，只在当前项目内生效。

### 另外一大区别

    全局安装后，可使用对应软件包的命令行工具，如果是局部安装，则无法使用，只能由代码调用

如全局安装的ganaceh-cli，使用命令 ganache-cli  即可启动一个本地服务器，共本机软件调用
而拒不安装的gannache-cli，只能有代码rrequire，而且运行一次启动一次，结束调用后自动关闭，不能