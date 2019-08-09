    const fs =  require('fs')
    const path = require('path')
    fs.readFile('test.txt',(err,data)=>{
        if(err){
            throw err;
        }
        console.log(data.length)
    })
    // readfile会异步执行，执行完的结果被callback,作为回调函数的参数继续执行，而整个回调函数也作为readfile的参数。
    // 这是回调函数的平常用法，当回调过多时会产生代码可读性的问题，成为回调地狱