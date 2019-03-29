# git标签

git可以给历史中的某一个提交打上标签，标识其重要的里程碑式的阶段节点。

## 列出标签

    git tag

## 查看标签

    git show v1.1

## 创建标签

两种类型的标签

### 创建附注标签

    git tag -a v1.1 -m 'my version 1.1'

为上次commit打上附注标签 V1.1,并记录本次操作日志。

### 创建轻量标签

    git tag v1.1.1

### 为之前的提交打标签

1. 查看提交日志，找到对应的校验号

        git log --pretty=oneline

    输出

        41d3033fe4724f299904e21a8c58b8f1900a2e85 修改格式错误
        f85e4ef504a2eba53b8679cc2f71647077f95802 (lxgo/master) 移动文件

2. 利用校验号创建标签

    git tag -a V1.0 -m '1.0版本' 41d3033fe472

## 共享标签

### 共享单个

push到远程仓库时，默认不会传送标签，可以手动传送

    git push origin v1.1.1
注意这个操作不仅会传送标签，也会把上次commit的所有内容一同push。

### 共享全部

    git push origin --tags

## 删除标签

删除本地标签与删除远程标签不同

### 删除本地标签

    git tag -d v1.1.1

### 删除远程标签

    git push origin :refs/tags/v1.1.1

## 检出标签

操作危险，咱不考虑
