    const fs = require('fs');
    const path = require('path')
    let target = 'test1.txt'
    const util = require('util');
    // 将fs.stat()转为新方法promiseStat()，新方法直接返回promise对象
    let promiseStat = util.promisify(fs.stat)
    // 将fs.readir()转为新方法promiseReadDir()，新方法直接返回promise对象
    let promiseReadDir = util.promisify(fs.readdir)


    // 使用async函数将promise对象的使用表现为同步写法，减少了回调函数的使用。而且使用工具创建promise对象。
    async function asyncDemo(){
        try{
            // 使用新方法，获得promise对象
            let stat = await promiseStat(target);
            if(stat.isFile()){
                console.log("是文件，大小为"+stat.size)
            }if(stat.isDirectory){
                console.log("是目录")
                let files = await promiseReadDir(target)
                files.forEach(async f => { 
                    console.log(f)
                    if(path.extname(f)==".txt"){
                        // 在方法f里面，使用await，需要使用async
                        let stat = await promiseStat(path.join(target,f))
                        console.log(f+"大小为"+stat.size)
                    }         
                });
            }
        }catch(err){
            throw(err)
        }
        
    }
    asyncDemo()
    