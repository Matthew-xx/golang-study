const http = require('http')
const fs = require('fs')
const path = require('path')
const util = require('util')
// 实现一个网页版的文件浏览器，浏览当前目录下的所有内容，同时浏览子文件夹内容
// 思考：
// 1. 需要大家一个服务器，实现基本路由
// 2. 需要获取当前文件夹的所有内容
// 读取当前项目文件夹html的目录
let target = 'html'
let psReaddir = util.promisify(fs.readdir)
let psStat = util.promisify(fs.stat)

// // 工具函数，读目录,经测试，不能把async函数当做工具函数，返回值永远是promis对象
// async function readDirFiles(path){
//     let returnFiles = new Array()
//     try{
//         let files = await psReaddir(path)
//         files.forEach(f => {
//             returnFiles.push(f)
//         });
//         return returnFiles
//     }catch(err){
//         throw(err)
//     }
// }

const server = http.createServer((req,res)=>{
    res.writeHead(200,{'content-type':'text/html;charset=utf-8'})
    console.log(req.url)
    if(req.url==='/favicon.ico'){
        res.end("")
        return
    }
    showDir(req,res)
    
})
server.listen(5000)

async function showDir(req,res){
    // 默认根路径，访问html
    // 若非根路径则访问html/
    if(req.url != '/'){
        console.log("当前请求路径："+req.url)
        target = req.url
    }
    let listr = '';
    try{
        // 读取targe目录
        let dirpath = path.join(__dirname,target)
        console.log("读取文件夹："+dirpath)
        let files = await psReaddir(dirpath)
        files.forEach(f => {
            // 获得target目录下的每一个文件
            let fpath = path.join(dirpath,f)
            let url = path.join(target,f)
            console.log("文件路径"+fpath)
            let stat = fs.statSync(fpath)
            if(stat.isDirectory()){
                listr += `<li><a href="${url}">${f}</a></li>`
            }else{
                listr += `<li>${f}</li>`
            }
        })
    }catch(err){
            throw(err)
    }
    res.end(makeHtml(listr))
    target = 'html'
}
function makeHtml(list){
     return`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <style type="text/css">
        *{padding:0;margin:0}
        ul>li{
            list-style:none;
        }
        li{
            padding: .6rem 1rem;
            background-color:#DDD;
            transition:all 1s;
        }
        li:not(:first-child){
            border-top:solid 1px #999;
        }
        li:hover{
            background-color:#AAA;
            
        }
    </style>
</head>
<body>
	<ul>${list}</ul>
</body>
</html>
<script type="text/javascript">
     `
}
// 总结：本项目遇到的一些坑：
// 最主要的是路径问题，一定要清楚，当前需要什么要的路径
// 如：读系统的文件夹，是要文件夹路径的，不能加盘符。如果是按照path工具组成的路径，会默认加盘符
// url是相对路径