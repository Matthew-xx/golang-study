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
    // 使用async函数将promise对象的使用表现为同步写法，减少了回调函数的使用。但创建promise对象还是很麻烦
    async function asyncDemo(){
        try{
            let stat = await promise;
            if(stat.isFile()){
                console.log("是文件，大小为"+stat.size)
            }else{
                console.log("是目录")
            }
        }catch(err){
            throw(err)
        }
        
    }
    asyncDemo()
    