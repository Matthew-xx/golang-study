# git标签

git可以给历史中的某一个提交打上标签，标识其重要的里程碑式的阶段节点。

## 列出标签

    git tag

## 创建标签

两种类型的标签

### 附注标签

    git tag -a v1.1 -m 'my version 1.1'

为上次commit打上附注标签 V1.1,并记录本次操作日志。

### 轻量标签

git tag v1.1.1

## 查看标签

    git show v1.1
