    const fs = require('fs');
    let promise = new Promise((resolve,reject)=>{
        fs.stat('test1.txt',(err,stat) => {
            if(err){
                reject(err)
            }else{
                resolve(stat)
            }          
        })
    })
    // 直接使用promise对象，仍然需要回调来使用，并没有方便多少。
    promise.then((stat)=>{
        if(stat.isFile()){
            console.log("是文件，大小为"+stat.size)
        }else{
            console.log("是目录")
        }
    }).catch(()=>{
        throw(err)
        console.log("成功")
    })

    