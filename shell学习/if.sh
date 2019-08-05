#/bin/bash
# 判断d传递到脚本的内容是文件还是文件夹
if [ -d $1 ];then
	echo "$1是个文件夹"
elif [ -s $1 ];then
	echo "$1是个文件，且文件不为空"
else
	echo "$1不存在，或者是大小为0的文件"
fi
