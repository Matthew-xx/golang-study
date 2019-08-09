    const fs = require('fs');
    const path = require('path')
    let target = 'test1.txt'
    fs.stat(target,(err,stat) => {
        if(err){
            throw(err)
        }
        if(stat.isFile()){
            console.log(target+"是文件，大小为"+stat.size)
        }
        if(stat.isDirectory()){
            console.log(target+"是目录,开始搜索目录下的文件")
            fs.readdir(target,(err,files)=>{
                // 遍历files
                files.forEach(f=>{
                    console.log(f)
                    if(path.extname(f)=='.txt'){
                        fs.stat(path.join(target,f),(err,stat)=>{
                            if(err){
                                throw(err)
                            }
                            console.log(f+"的大小为"+stat.size)
                        })
                    } 
                })
            })
        }
    })